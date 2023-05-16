package repository

import (
	"joi-energy-golang/domain"
)

type PriceRepository interface {
	GetPricePlans() []domain.PricePlan
	GetReadings(smartMeterId string) []domain.ElectricityReading
}

type pricePlans struct {
	pricePlans    []domain.PricePlan
	meterReadings *MeterReadings
}

func NewPricePlans(newPrices []domain.PricePlan, meterReadings *MeterReadings) PriceRepository {
	return &pricePlans{
		pricePlans:    newPrices,
		meterReadings: meterReadings,
	}
}

func (p *pricePlans) GetPricePlans() []domain.PricePlan {
	return p.pricePlans
}

func (p *pricePlans) GetReadings(smartMeterId string) []domain.ElectricityReading {
	return p.meterReadings.GetReadings(smartMeterId)
}
