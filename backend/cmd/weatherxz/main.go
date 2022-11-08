package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/zsomborjoel/weatherxz/pkg/cities"
	"github.com/zsomborjoel/weatherxz/pkg/clients/openweather"
	"github.com/zsomborjoel/weatherxz/pkg/common"
	"github.com/zsomborjoel/weatherxz/pkg/pings"
	"github.com/zsomborjoel/weatherxz/pkg/weathers"
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

	r := gin.Default()

	v1 := r.Group("/api")
	pings.PingRegister(v1.Group("/ping"))
	openweather.WeatherLoadRegister(v1.Group("/weathers"))

	r.Run()
}
