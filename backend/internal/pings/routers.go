package pings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/weatherxz/internal/common"
)

func PingRegister(router *gin.RouterGroup) {
	router.GET("", ping)
	router.GET("/db", pingDb)
}

func ping(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func pingDb(c *gin.Context) {
	db := common.GetDB()

	var result int
	db.Raw("SELECT 1").Scan(&result)

	if result == 1 {
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	c.Writer.WriteHeader(http.StatusBadRequest)
}
