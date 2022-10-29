package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/weatherxz/pkg/city"
)

type OpenWeatherForecastResponse struct {
	Cod     string   `json:"cod"`
	Message int      `json:"message"`
	Cnt     int      `json:"cnt"`
	List    []list   `json:"list"`
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

type weather struct {
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
	ThreeH float64 `json:"3h"`
}

type sys struct {
	Pod string `json:"pod"`
}

type snow struct {
	ThreeH float64 `json:"3h"`
}

type list struct {
	Dt         int       `json:"dt"`
	Main       main      `json:"main"`
	Weather    []weather `json:"weather"`
	Clouds     clouds    `json:"clouds"`
	Wind       wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`
	Rain       rain      `json:"rain,omitempty"`
	Sys        sys       `json:"sys"`
	DtTxt      string    `json:"dt_txt"`
	Snow       snow      `json:"snow,omitempty"`
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

func FetchForAllCities() {
	log.Info().Msg("WeatherForecast fetch for all cities started")

	cis, err := city.GetAllCity()
	if err != nil {
		log.Error().Stack().Err(err)
	}

	for _, ci := range cis {
		log.Info().Msg(string(ci.Name))

		var f OpenWeatherForecastResponse
		f.City.Coord.Lat = 6
		fmt.Println(f)
		//f := fetchWeatherForecast(ci.Name)
		ci = copyForecastToCity(f, ci)
		fmt.Println(ci.CoordLat)
	}
	log.Info().Msg("WeatherForecast fetch for all cities ended")
}

func copyForecastToCity(f OpenWeatherForecastResponse, ci city.City) (city.City) {
	ci.CoordLat = f.City.Coord.Lat
	ci.CoordLon = f.City.Coord.Lon
	ci.TimeZone = f.City.Timezone
	ci.Sunrise = f.City.Sunset
	return ci
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

	return
}
