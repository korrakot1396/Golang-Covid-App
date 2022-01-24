package summary

import (
	"app/main/model"
	"sync"
)

type covidServiceSolution interface {
	calculate([]model.Covid) interface{}
}

type CovidAPI struct {
	data   []model.Covid
	result interface{}
	mutex  sync.Mutex
	covidServiceSolution
}

func (s *CovidAPI) Calculate() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.result == nil {
		s.result = s.covidServiceSolution.calculate(s.data)
	}
	return s.result
}

func (s *CovidAPI) OnUpdate(data []model.Covid) {
	s.mutex.Lock()
	s.data = data
	s.result = nil
	s.mutex.Unlock()
}

type CovidAPISummary int

const (
	API CovidAPISummary = iota
)

func NewCovidAPISummary(covidAPISummary CovidAPISummary) CovidAPI {
	return CovidAPI{covidServiceSolution: covidSummary{}}
}
