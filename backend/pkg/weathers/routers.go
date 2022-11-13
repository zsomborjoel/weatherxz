package weathers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WeathersRegister(r *gin.RouterGroup) {
	r.GET("/:cityId", WeathersRetrive)
	r.GET("/today/:cityId", WeatherRetrive)
}

func WeathersRetrive(c *gin.Context) {
	cityId := c.Param("cityId")
	cityIdUint, err := strconv.ParseUint(cityId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	weathers, err := GetAllWeather(&Weather{CityId: uint(cityIdUint)})
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	s := WeathersSerializer{c, weathers}
	c.JSON(http.StatusOK, s.Response())
}

func WeatherRetrive(c *gin.Context) {
	cityId := c.Param("cityId")
	cityIdUint, err := strconv.ParseUint(cityId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	weather, err := GetTodaysWeather(&Weather{CityId: uint(cityIdUint)})
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	s := WeatherSerializer{c, weather}
	c.JSON(http.StatusOK, s.Response())
}
