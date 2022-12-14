package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/zsomborjoel/weatherxz/internal/cities"
	"github.com/zsomborjoel/weatherxz/internal/clients/openweather"
	"github.com/zsomborjoel/weatherxz/internal/common"
	"github.com/zsomborjoel/weatherxz/internal/pings"
	"github.com/zsomborjoel/weatherxz/internal/weathers"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	cities.AutoMigrate()
	weathers.AutoMigrate()
	common.LoadSQLFile("/home/zsomborjoel/Git/weatherxz/backend/scripts/data/country_data_20221025_v1.sql")
	common.LoadSQLFile("/home/zsomborjoel/Git/weatherxz/backend/scripts/data/city_data_20221027_v1.sql")
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	common.LoadEnvVariables()

	db := common.Init()
	Migrate(db)

	openweather.ScheduleLoad()
	r := gin.Default()
	r.Use(common.CORSMiddleware())

	v1 := r.Group("/api")
	pings.PingRegister(v1.Group("/ping"))
	openweather.WeatherLoadRegister(v1.Group("/weatherload"))
	cities.CitiesRegister(v1.Group("/cities"))
	weathers.WeathersRegister(v1.Group("/weathers"))

	r.Run()
}
