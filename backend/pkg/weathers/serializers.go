package weathers

import "github.com/gin-gonic/gin"

type WeatherResponse struct {
	ID                         uint    `json:"id"`
	DateTime                   int     `json:"DateTime"`
	Temp                       float64 `json:"temp"`
	FeelsLike                  float64 `json:"feelsLike"`
	TempMin                    float64 `json:"tempMin"`
	TempMax                    float64 `json:"tempMax"`
	Pressure                   int     `json:"pressure"`
	PreassureSeaLevel          int     `json:"preassureSeaLevel"`
	PreassureGroundLevel       int     `json:"preassureGroundLevel"`
	Humidity                   int     `json:"humidity"`
	WeatherConditionMain       string  `json:"weatherConditionMain"`
	WeatherConditionDesc       string  `json:"weatherConditionDesc"`
	Clouds                     int     `json:"clouds"`
	WindSpeed                  float64 `json:"windSpeed"`
	WindDeg                    int     `json:"windDeg"`
	WindGust                   float64 `json:"windGust"`
	Visibility                 int     `json:"visibility"`
	ProbabilityOfPrecipitation float64 `json:"probabilityOfPrecipitation"`
	RainVolume                 float64 `json:"rainVolume"`
	SnowVolume                 float64 `json:"snowVolume"`
	PartOfDay                  string  `json:"partOfDay"`
}

type WeatherSerializer struct {
	C *gin.Context
	Weather
}

func (s *WeatherSerializer) Response() WeatherResponse {
	r := WeatherResponse{
		ID:                         s.Model.ID,
		DateTime:                   s.DateTime,
		Temp:                       s.Temp,
		FeelsLike:                  s.FeelsLike,
		TempMin:                    s.TempMin,
		TempMax:                    s.TempMax,
		Pressure:                   s.Pressure,
		PreassureSeaLevel:          s.PreassureSeaLevel,
		PreassureGroundLevel:       s.PreassureGroundLevel,
		Humidity:                   s.Humidity,
		WeatherConditionMain:       s.WeatherConditionMain,
		WeatherConditionDesc:       s.WeatherConditionDesc,
		Clouds:                     s.Clouds,
		WindSpeed:                  s.WindSpeed,
		WindDeg:                    s.WindDeg,
		WindGust:                   s.WindGust,
		Visibility:                 s.Visibility,
		ProbabilityOfPrecipitation: s.ProbabilityOfPrecipitation,
		RainVolume:                 s.RainVolume,
		SnowVolume:                 s.SnowVolume,
		PartOfDay:                  s.PartOfDay,
	}
	return r
}

type WeathersSerializer struct {
	C        *gin.Context
	Weathers []Weather
}

func (s *WeathersSerializer) Response() []WeatherResponse {
	r := []WeatherResponse{}
	for _, weather := range s.Weathers {
		s := WeatherSerializer{s.C, weather}
		r = append(r, s.Response())
	}
	return r
}
