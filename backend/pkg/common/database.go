package common

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/rs/zerolog/log"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := os.Getenv("DB_CONNECTION")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("db err: (Init) ")
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}