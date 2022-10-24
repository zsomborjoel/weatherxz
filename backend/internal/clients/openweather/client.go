package openweather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/rs/zerolog/log"
)

type OpenWeatherForecastResponse struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []List `json:"list"`
	City    City   `json:"city"`
}

type Main struct {
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

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Rain struct {
	ThreeH float64 `json:"3h"`
}

type Sys struct {
	Pod string `json:"pod"`
}

type Snow struct {
	ThreeH float64 `json:"3h"`
}

type List struct {
	Dt         int       `json:"dt"`
	Main       Main      `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`
	Rain       Rain      `json:"rain,omitempty"`
	Sys        Sys       `json:"sys"`
	DtTxt      string    `json:"dt_txt"`
	Snow       Snow      `json:"snow,omitempty"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type City struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

func FetchWeatherForecast() {
	log.Info().Msg("Info message")
	ofu := os.Getenv("OPEN_WEATHER_URL")
	appid := os.Getenv("OPEN_WEATHER_APP_ID")
	city := "hanoi"

	base, err := url.Parse(ofu)
	if err != nil {
		log.Error().Err(err)
	}

	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", appid)
	base.RawQuery = params.Encode()

	response, err := http.Get(base.String())
	if err != nil {
		log.Error().Err(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err)
	}

	log.Info().Msg(string(data))

	var owfr OpenWeatherForecastResponse
	json.Unmarshal(data, &owfr)

	log.Info().Msg(owfr.City.Name)
}
