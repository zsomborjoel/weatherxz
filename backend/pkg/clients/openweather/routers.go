package openweather

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WeatherLoadRegister(router *gin.RouterGroup) {
	router.GET("/load", WeatherLoad)
}

func WeatherLoad(c *gin.Context) {
	go FetchForAllCities()
	c.JSON(http.StatusOK, "")
}
