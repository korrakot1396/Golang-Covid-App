package main

import (
	"app/main/api"
	service "app/main/services"
	"fmt"
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
