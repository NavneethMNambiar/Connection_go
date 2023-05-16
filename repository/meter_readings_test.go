package repository

import (
	"joi-energy-golang/domain"
	// "reflect"
	"sort"
	// "testing"
	"time"
)

// func TestMeterReadings_GetWeeklyReadings(t *testing.T) {
// 	tests := []struct {
// 		testName string
// 		//fields initReadings
// 		smartMeterName string
// 		want           float64
// 	}{
// 		{
// 			testName:       "firstTest",
// 			smartMeterName: "fistName",
// 			want:           10,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.testName, func(t *testing.T) {
// 			m := &MeterReadings{
// 				meterAssociatedReadings: initReadings(tt.smartMeterName),
// 			}
// 			if got := m.GetWeeklyReadings(tt.smartMeterName); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetWeeklyReadings() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func initReadings(meter string) map[string][]domain.ElectricityReading {
	readings := make([]domain.ElectricityReading, 7)
	meterReading := make(map[string][]domain.ElectricityReading)
	now := time.Now()
	for i := range readings {
		electricityReading := domain.ElectricityReading{
			Time:    now.Add(time.Duration(i*-10) * time.Second),
			Reading: 10,
		}
		readings[i] = electricityReading
	}
	sort.Slice(readings, func(i, j int) bool { return readings[i].Time.Before(readings[j].Time) })
	meterReading[meter] = readings
	return meterReading
}
