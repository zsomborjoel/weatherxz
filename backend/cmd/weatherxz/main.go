package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/zsomborjoel/weatherxz/internal/clients/openweather"
	"github.com/zsomborjoel/weatherxz/internal/common"
)

func init() {
	common.LoadEnvVariables()
	common.ConnectToDB()
	openweather.FetchWeatherForecast()
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	r := gin.Default()

	r.Run()

}
