package api

import (
	"app/main/model"
	"encoding/json"
	"io"
	"net/http"
)

type Covid interface {
	OnUpdate([]model.Covid)
}

type CovidAPI struct {
	subscribers map[Covid]bool
}

func (covidAPI *CovidAPI) Subscribe(covid Covid) {
	covidAPI.subscribers[covid] = true
}

func (covidAPI *CovidAPI) Unsubscribe(covid Covid) {
	delete(covidAPI.subscribers, covid)
}

func (covidAPI *CovidAPI) Alert(data []model.Covid) {
	for covid := range covidAPI.subscribers {
		covid.OnUpdate(data)
	}
}

func (covidAPI CovidAPI) FetchCovidAPI() {
	res, _ := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")

	body, _ := io.ReadAll(res.Body)

	data := struct {
		Data []model.Covid
	}{}

	json.Unmarshal(body, &data)

	covidAPI.Alert(data.Data)
}

func NewCovidAPI() CovidAPI {
	return CovidAPI{map[Covid]bool{}}
}
