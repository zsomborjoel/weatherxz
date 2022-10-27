package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/zsomborjoel/weatherxz/pkg/city"
	"github.com/zsomborjoel/weatherxz/pkg/common"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	city.AutoMigrate()
	common.LoadSQLFile("/home/zsomborjoel/Git/weatherxz/backend/scripts/data/country_data_20221025_v1.sql")
	common.LoadSQLFile("/home/zsomborjoel/Git/weatherxz/backend/scripts/data/city_data_20221027_v1.sql")
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	common.LoadEnvVariables()

	db := common.Init()
	Migrate(db)

	r := gin.Default()

	r.Run()
}
