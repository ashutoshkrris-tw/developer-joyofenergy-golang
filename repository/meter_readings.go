package repository

import (
	"joi-energy-golang/domain"
	"time"
)

type MeterReadings struct {
	meterAssociatedReadings map[string][]domain.ElectricityReading
}

func NewMeterReadings(meterAssociatedReadings map[string][]domain.ElectricityReading) MeterReadings {
	return MeterReadings{meterAssociatedReadings: meterAssociatedReadings}
}

func (m *MeterReadings) GetReadings(smartMeterId string) []domain.ElectricityReading {
	v, ok := m.meterAssociatedReadings[smartMeterId]
	if !ok {
		return nil
	}
	return v
}

func (m *MeterReadings) GetReadingsWithinTime(smartMeterId string, from, to time.Time) []domain.ElectricityReading {
	filteredReadings := []domain.ElectricityReading{}
	v, ok := m.meterAssociatedReadings[smartMeterId]
	if !ok {
		return nil
	}
	for _, val := range v {
		if val.Time.After(from) && val.Time.Before(to) {
			filteredReadings = append(filteredReadings, val)
		}
	}
	return filteredReadings
}

func (m *MeterReadings) StoreReadings(smartMeterId string, electricityReadings []domain.ElectricityReading) {
	m.meterAssociatedReadings[smartMeterId] = append(m.meterAssociatedReadings[smartMeterId], electricityReadings...)
}
