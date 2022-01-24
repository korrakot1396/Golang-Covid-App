package api_test

import (
	"app/main/api"
	"testing"
)

func FetchCovidAPI(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("API can not fetch")
			}
		}()

		covidAPI := api.NewCovidAPI()
		covidAPI.FetchCovidAPI()
	})

}
