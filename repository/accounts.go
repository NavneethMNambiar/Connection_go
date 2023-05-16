package repository

type Accounts struct {
	smartMeterToPricePlanAccounts map[string]string
}

func NewAccounts(c *Accounts) {
	c = &Accounts{defaultSmartMeterToPricePlanAccounts()}
}

func defaultSmartMeterToPricePlanAccounts() map[string]string {
	return map[string]string{
		"smart-meter-0": "price-plan-0",
		"smart-meter-1": "price-plan-1",
		"smart-meter-2": "price-plan-0",
		"smart-meter-3": "price-plan-2",
		"smart-meter-4": "price-plan-1",
	}
}

func (a *Accounts) PricePlanIdForSmartMeterId(smartMeterId string) string {
	// TODO indicate missing value
	return a.smartMeterToPricePlanAccounts[smartMeterId]
}
