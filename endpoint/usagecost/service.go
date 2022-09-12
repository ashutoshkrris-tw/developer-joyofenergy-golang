package usagecost

import (
	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

type Service interface {
	CalculateLastWeekUsageCost(smartMeterId string) (domain.UsageCost, error)
}

type service struct {
	logger        *logrus.Entry
	meterReadings *repository.MeterReadings
	pricePlans    *repository.PricePlans
	accounts      *repository.Accounts
}

func NewService(
	logger *logrus.Entry,
	meterReadings *repository.MeterReadings,
	pricePlans *repository.PricePlans,
	accounts *repository.Accounts,
) Service {
	return &service{
		logger:        logger,
		meterReadings: meterReadings,
		pricePlans:    pricePlans,
		accounts:      accounts,
	}
}

func (s *service) CalculateLastWeekUsageCost(smartMeterId string) (domain.UsageCost, error) {
	pricePlanId := s.accounts.PricePlanIdForSmartMeterId(smartMeterId)
	usageCost := s.pricePlans.UsageCostOfLastWeekElectricityReadings(smartMeterId, pricePlanId)
	return domain.UsageCost{
		Cost: usageCost,
	}, nil
}

// smartMeterId
// price plan -> tariff
// readings -> avg, hours
//usagecost = tariff * avg * hours
