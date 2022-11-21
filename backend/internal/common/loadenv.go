package common

import (
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Stack().Msg("Error loading .env file")
	}
}
