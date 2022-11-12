package weathers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WeathersRegister(r *gin.RouterGroup) {
	r.GET("", WeathersRetrive)
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
