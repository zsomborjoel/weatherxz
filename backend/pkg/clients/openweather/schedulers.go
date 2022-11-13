package openweather

import (
	"context"
	"fmt"
	"os"

	"github.com/procyon-projects/chrono"
	"github.com/rs/zerolog/log"
)

func ScheduleLoad() {
	cron := os.Getenv("OPEN_WEATHER_CRON")
	taskScheduler := chrono.NewDefaultTaskScheduler()
	log.Info().Msg(fmt.Sprintf("Scheduler OpenWeatherForecast initialized for cron [%s]", cron))

	_, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
		log.Info().Msg("Scheduled OpenWeatherForecast load")
		FetchForAllCities()
	}, cron)

	if err == nil {
		log.Error().Stack().Err(err)
	}
}
