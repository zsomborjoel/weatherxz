package weathers

import (
	"github.com/zsomborjoel/weatherxz/pkg/common"
	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model
	DateTime                   int
	Temp                       float64
	FeelsLike                  float64
	TempMin                    float64
	TempMax                    float64
	Pressure                   int
	PreassureSeaLevel          int
	PreassureGroundLevel       int
	Humidity                   int
	WeatherConditionId         int
	WeatherConditionMain       string
	WeatherConditionDesc       string
	WeatherConditionIcon       string
	Clouds                     int
	WindSpeed                  float64
	WindDeg                    int
	WindGust                   float64
	Visibility                 int
	ProbabilityOfPrecipitation float64
	RainVolume                 float64
	SnowVolume                 float64
	PartOfDay                  string
	CityId                     uint
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Weather{})
}

func SaveAll(data interface{}) error {
	db := common.GetDB()
	err := db.Create(data).Error
	return err
}

func GetAllWeather(condition interface{}) ([]Weather, error) {
	db := common.GetDB()
	var countries []Weather
	err := db.Where(condition).Find(&countries).Error
	return countries, err
}
