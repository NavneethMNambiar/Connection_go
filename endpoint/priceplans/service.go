package priceplans

import (
	"sort"
	"time"

	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

type Service interface {
	CompareAllPricePlans(smartMeterId string) (domain.PricePlanComparisons, error)
	RecommendPricePlans(smartMeterId string, limit uint64) (domain.PricePlanRecommendation, error)
}

type service struct {
	logger     *logrus.Entry
	pricePlans repository.PriceRepository
	accounts   *repository.Accounts
}

func NewService(
	logger *logrus.Entry,
	pricePlans repository.PriceRepository,
	accounts *repository.Accounts,
) Service {
	return &service{
		logger:     logger,
		pricePlans: pricePlans,
		accounts:   accounts,
	}
}

func (s *service) CompareAllPricePlans(smartMeterId string) (domain.PricePlanComparisons, error) {
	pricePlanId := s.accounts.PricePlanIdForSmartMeterId(smartMeterId)
	costs := s.CalculateCost(smartMeterId)
	// consumptionsForPricePlans := s.pricePlans.ConsumptionCostOfElectricityReadingsForEachPricePlan(smartMeterId)
	if len(costs) == 0 {
		return domain.PricePlanComparisons{}, domain.ErrNotFound
	}
	return domain.PricePlanComparisons{
		PricePlanId:          pricePlanId,
		PricePlanComparisons: costs,
	}, nil
}

func (s *service) CalculateCost(smartMeterId string) map[string]float64 {
	electricityReadings := s.pricePlans.GetReadings(smartMeterId)
	costs := map[string]float64{}
	for _, plan := range s.pricePlans.GetPricePlans() {
		costs[plan.PlanName] = calculateCost(electricityReadings, plan)
	}
	return costs
}

func calculateCost(electricityReadings []domain.ElectricityReading, pricePlan domain.PricePlan) float64 {
	average := calculateAverageReading(electricityReadings)
	timeElapsed := calculateTimeElapsed(electricityReadings)
	averagedCost := average / timeElapsed.Hours()
	return averagedCost * pricePlan.UnitRate
}

func calculateAverageReading(electricityReadings []domain.ElectricityReading) float64 {
	sum := 0.0
	for _, r := range electricityReadings {
		sum += r.Reading
	}
	return sum / float64(len(electricityReadings))
}

func calculateTimeElapsed(electricityReadings []domain.ElectricityReading) time.Duration {
	var first, last time.Time
	for _, r := range electricityReadings {
		if r.Time.Before(first) || (first == time.Time{}) {
			first = r.Time
		}
	}
	for _, r := range electricityReadings {
		if r.Time.After(last) || (last == time.Time{}) {
			last = r.Time
		}
	}
	return last.Sub(first)
}

func (s *service) RecommendPricePlans(smartMeterId string, limit uint64) (domain.PricePlanRecommendation, error) {
	consumptionsForPricePlans := s.CalculateCost(smartMeterId)
	if len(consumptionsForPricePlans) == 0 {
		return domain.PricePlanRecommendation{}, domain.ErrNotFound
	}
	var recommendations []domain.SingleRecommendation
	for k, v := range consumptionsForPricePlans {
		recommendations = append(recommendations, domain.SingleRecommendation{
			Key:   k,
			Value: v,
		})
	}
	sort.Slice(recommendations, func(i, j int) bool { return recommendations[i].Value < recommendations[j].Value })

	if limit > 0 {
		recommendations = recommendations[:limit]
	}

	return domain.PricePlanRecommendation{Recommendations: recommendations}, nil
}
