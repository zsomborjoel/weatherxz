package common

import (
	"io/ioutil"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/weatherxz/pkg/city"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&city.Country{})
	db.AutoMigrate(&city.City{})
	r := sqlLoad("/scripts/data/country_data_20221025_v1")
	log.Info().Msg(r[0])
	db.Exec(r[0])
}

func sqlLoad(path string) []string {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Error().Err(err)
	}

	return strings.Split(string(file), ";")
}
