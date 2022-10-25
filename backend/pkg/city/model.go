package city

import (
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name         string `gorm:"UNIQUE_INDEX:compositeindex;type:text;not null"`
	Alpha2Code   string
	Aplpha2Code  string
	NumbericCode int
}

type City struct {
	gorm.Model
	Name       string
	CoordLat   string
	CoordLon   string
	Population int
	TimeZone   int
	Sunrise    int
	Sunset     int
	CountryID  uint
}
