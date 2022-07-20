package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// configuration for static files and templates
	engine.LoadHTMLFiles("templates/index.html")
	engine.StaticFile("/favicon.ico", "favicon.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Weather Service",
		})
	})

	engine.GET("/api/weather", queryWeather) // get weather for city

	ConnectDatabase()
	engine.Run(port())
}

func queryWeather(c *gin.Context) {
	weather := GetOrRetrieveWeather(c.Query("city"))
	c.JSON(http.StatusOK, weather)
}
