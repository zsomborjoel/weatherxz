package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CitiesRegister(r *gin.RouterGroup) {
	r.GET("/:name")
}

func CityRetrive(c *gin.Context) {
	name := c.Param("name")
	city := FindOneCity(&City{Name: name})
	serializer := CitySerializer{c, city}
	c.JSON(http.StatusOK, serializer.Response())
}
