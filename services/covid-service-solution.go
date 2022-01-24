package summary

import (
	"app/main/model"
)

type (
	covidSummary struct{}

	Response struct {
		Province map[string]int
		AgeGroup map[string]int
	}
)

func (c covidSummary) calculate(data []model.Covid) interface{} {
	res := Response{
		map[string]int{}, map[string]int{},
	}
	for _, c := range data {
		if c.Province != "" {
			res.Province[c.Province]++
		}
		if c.Age == nil {
			res.AgeGroup["N/A"]++
		} else if *c.Age >= 0 && *c.Age <= 30 {
			res.AgeGroup["0-30"]++
		} else if *c.Age >= 31 && *c.Age <= 60 {
			res.AgeGroup["31-60"]++
		} else {
			res.AgeGroup["61+"]++
		}
	}
	return res
}
