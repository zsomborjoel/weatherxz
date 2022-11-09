package openweather

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WeatherLoadRegister(r *gin.RouterGroup) {
	r.GET("/load", weatherLoad)
}

func weatherLoad(c *gin.Context) {
	go FetchForAllCities()
	c.Writer.WriteHeader(http.StatusOK)
}
