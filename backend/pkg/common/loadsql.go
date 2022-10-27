package common

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

const Delimeter = ";"

func LoadSQLFile(sqlFile string) {
	db := GetDB()

	file, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Error().Stack().Err(err)
	}

	var executed bool
	for _, q := range strings.Split(string(file), Delimeter) {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		db.Exec(q)
		executed = true
	}

	if executed {
		log.Info().Msg("Migration File executed: " + sqlFile)
	}
}
