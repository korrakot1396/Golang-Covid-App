package main

import (
	"app/main/api"
	service "app/main/services"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		apiCovid = service.NewCovidAPISummary(service.API)
	)

	covidAPI := api.NewCovidAPI()

	covidAPI.Subscribe(&apiCovid)
	covidAPI.FetchCovidAPI()

	app := gin.Default()

	app.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      ".",
		Extension: ".html",
	})

	app.GET("/covid/summary", func(c *gin.Context) {
		c.JSON(http.StatusOK, apiCovid.Calculate())
	})

	app.Run()
}
