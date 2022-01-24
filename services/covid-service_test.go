package summary_test

import (
	"app/main/model"
	service "app/main/services"
	"reflect"
	"testing"
)

func TestCovidService(t *testing.T) {
	float32Ptr := func(f float32) *float32 {
		return &f
	}

	t.Run("Success", func(t *testing.T) {
		s := service.NewCovidAPISummary(service.API)
		s.OnUpdate([]model.Covid{
			{Province: "A", Age: float32Ptr(0)},
			{Province: "B", Age: float32Ptr(28)},
			{Province: "A", Age: float32Ptr(75.9)},
			{Province: "A", Age: float32Ptr(91)},
			{Province: "B", Age: nil},
			{Province: "A", Age: float32Ptr(1)},
		})

		got, ok := s.Calculate().(service.Response)
		expect := service.Response{
			Province: map[string]int{
				"A": 4,
				"B": 2,
			},
			AgeGroup: map[string]int{
				"0-30":  2,
				"31-60": 2,
				"61+":   1,
				"N/A":   1,
			},
		}

		if !ok {
			t.Errorf("CovidAPI not working")
		} else if !reflect.DeepEqual(got, expect) {
			t.Errorf("CovidAPI got wrong answer")
		}
	})

}
