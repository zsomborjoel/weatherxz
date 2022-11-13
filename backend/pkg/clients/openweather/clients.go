package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/weatherxz/pkg/cities"
	"github.com/zsomborjoel/weatherxz/pkg/weathers"
)

type OpenWeatherForecastResponse struct {
	Cod     string   `json:"cod"`
	Message int      `json:"message"`
	Cnt     int      `json:"cnt"`
	List    []List   `json:"list"`
	City    cityinfo `json:"city"`
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type weatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type clouds struct {
	All int `json:"all"`
}

type wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type rain struct {
	OneH float64 `json:"1h"`
}

type sys struct {
	Pod string `json:"pod"`
}

type snow struct {
	OneH float64 `json:"1h"`
}

type List struct {
	Dt         int           `json:"dt"`
	Main       main          `json:"main"`
	Weather    []weatherInfo `json:"weather"`
	Clouds     clouds        `json:"clouds"`
	Wind       wind          `json:"wind"`
	Visibility int           `json:"visibility"`
	Pop        float64       `json:"pop"`
	Rain       rain          `json:"rain,omitempty"`
	Sys        sys           `json:"sys"`
	DtTxt      string        `json:"dt_txt"`
	Snow       snow          `json:"snow,omitempty"`
}

type coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type cityinfo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Coord      coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

const ITERATION_LEVEL = 50

func FetchForAllCities() {
	log.Info().Msg("WeatherForecast fetch for all cities started")

	currCi, err := cities.GetAllCity()
	if err != nil {
		log.Error().Stack().Err(err)
	}

	var savableWe = []weathers.Weather{}
	for i, ci := range currCi {
		if i%ITERATION_LEVEL == 0 && len(savableWe) > 0 {
			weathers.SaveAll(savableWe)
			savableWe = []weathers.Weather{}

			log.Info().Msg("Forecast Fetch - sleep")
			time.Sleep(1 * time.Minute)
		}

		log.Info().Msg(fmt.Sprintf("Currently processed city: [%s]", ci.Name))

		f := fetchWeatherForecast(ci.Name)

		copyToCity(&f, &ci)
		err := cities.UpdateCity(&ci)
		if err != nil {
			log.Error().Stack().Err(err)
		}

		for _, l := range f.List {
			var w weathers.Weather
			copyToWeather(&l, &w)
			w.CityId = ci.ID
			savableWe = append(savableWe, w)
		}
	}

	if len(savableWe) > 0 {
		err = weathers.SaveAll(savableWe)
		if err != nil {
			log.Error().Stack().Err(err)
		}
	}

	log.Info().Msg("WeatherForecast fetch for all cities ended")
}

func copyToCity(f *OpenWeatherForecastResponse, ci *cities.City) {
	ci.CoordLat = f.City.Coord.Lat
	ci.CoordLon = f.City.Coord.Lon
	ci.Population = f.City.Population
	ci.TimeZone = f.City.Timezone
	ci.Sunrise = f.City.Sunrise
	ci.Sunset = f.City.Sunset
}

func copyToWeather(l *List, w *weathers.Weather) {
	lastWCond := l.Weather[len(l.Weather)-1]

	w.DateTime = l.Dt
	w.Temp = l.Main.Temp
	w.FeelsLike = l.Main.FeelsLike
	w.TempMin = l.Main.TempMin
	w.TempMax = l.Main.TempMax
	w.Pressure = l.Main.Pressure
	w.PreassureSeaLevel = l.Main.SeaLevel
	w.PreassureGroundLevel = l.Main.GrndLevel
	w.Humidity = l.Main.Humidity
	w.WeatherConditionId = lastWCond.ID
	w.WeatherConditionMain = lastWCond.Main
	w.WeatherConditionDesc = lastWCond.Description
	w.WeatherConditionIcon = lastWCond.Icon
	w.Clouds = l.Clouds.All
	w.WindSpeed = l.Wind.Speed
	w.WindDeg = l.Wind.Deg
	w.WindGust = l.Wind.Gust
	w.Visibility = l.Visibility
	w.ProbabilityOfPrecipitation = l.Pop
	w.RainVolume = l.Rain.OneH
	w.SnowVolume = l.Snow.OneH
	w.PartOfDay = l.Sys.Pod
}

func fetchWeatherForecast(city string) (f OpenWeatherForecastResponse) {
	ofu := os.Getenv("OPEN_WEATHER_URL")
	aid := os.Getenv("OPEN_WEATHER_APP_ID")

	base, err := url.Parse(ofu)
	if err != nil {
		log.Error().Stack().Err(err)
	}

	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", aid)
	base.RawQuery = params.Encode()

	res, err := http.Get(base.String())
	if err != nil {
		log.Error().Stack().Err(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Stack().Err(err)
	}

	var owfr OpenWeatherForecastResponse
	json.Unmarshal(data, &owfr)

	return owfr
}
