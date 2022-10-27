package city

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
	CoordLat   string `gorm:"type:varchar(50)"`
	CoordLon   string `gorm:"type:varchar(50)"`
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
