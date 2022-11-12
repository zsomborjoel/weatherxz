package cities

import (
	"github.com/zsomborjoel/weatherxz/pkg/common"
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);unique;not null"`
	Alpha2Code  string `gorm:"type:varchar(2);not null"`
	Alpha3Code  string `gorm:"type:varchar(3);not null"`
	NumericCode int    `gorm:"not null"`
	Cities      []City `gorm:"foreignKey:CountryID"`
}

type City struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255);unique;not null"`
	CoordLat   float64
	CoordLon   float64
	Population int
	TimeZone   int
	Sunrise    int
	Sunset     int
	CountryID  uint
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Country{})
	db.AutoMigrate(&City{})
}

func FindOneCity(condition interface{}) (City, error) {
	db := common.GetDB()
	var c City
	err := db.Where(condition).First(&c).Error
	return c, err
}

func FindOneCountryById(countryId uint) (Country, error) {
	db := common.GetDB()
	var c Country
	err := db.Find(&c, countryId).Error
	return c, err
}

func GetAllCountry() ([]Country, error) {
	db := common.GetDB()
	var countries []Country
	err := db.Find(&countries).Error
	return countries, err
}

func GetAllCity() ([]City, error) {
	db := common.GetDB()
	var cities []City
	err := db.Find(&cities).Error
	return cities, err
}

func UpdateCity(city *City) error {
	db := common.GetDB()
	err := db.Model(&city).Updates(city).Error
	return err
}
