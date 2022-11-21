package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CitiesRegister(r *gin.RouterGroup) {
	r.GET("/:name", CityRetrive)
	r.GET("", CitiesRetrive)
}

func CityRetrive(c *gin.Context) {
	name := c.Param("name")
	city, err := FindOneCity(&City{Name: name})
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	s := CitySerializer{c, city}
	c.JSON(http.StatusOK, s.Response())
}

func CitiesRetrive(c *gin.Context) {
	cities, err := GetAllCity()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	s := CitiesSerializer{c, cities}
	c.JSON(http.StatusOK, s.Response())
}
