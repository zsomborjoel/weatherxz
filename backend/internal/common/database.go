package common

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := os.Getenv("DB_CONNECTION")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Stack().Err(err).Msg("db err: (Init)")
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
