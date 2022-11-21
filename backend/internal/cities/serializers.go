package cities

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CityResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	CoordLat    float64 `json:"coordLat"`
	CoordLon    float64 `json:"coordLon"`
	Population  int     `json:"population"`
	TimeZone    int     `json:"timeZone"`
	Sunrise     int     `json:"sunrise"`
	Sunset      int     `json:"sunset"`
	CountryName string  `json:"countryName"`
}

type CitySerializer struct {
	C *gin.Context
	City
}

func (s *CitySerializer) Response() CityResponse {
	c, err := FindOneCountryById(s.CountryID)

	// We don't care if country exists for particular city
	if err != nil {
		log.Info().Stack().Err(err)
	}

	r := CityResponse{
		ID:          s.Model.ID,
		Name:        s.Name,
		CoordLat:    s.CoordLat,
		CoordLon:    s.CoordLon,
		Population:  s.Population,
		TimeZone:    s.TimeZone,
		Sunrise:     s.Sunrise,
		Sunset:      s.Sunset,
		CountryName: c.Name,
	}
	return r
}

type CitiesSerializer struct {
	C      *gin.Context
	Cities []City
}

func (s *CitiesSerializer) Response() []CityResponse {
	r := []CityResponse{}
	for _, city := range s.Cities {
		s := CitySerializer{s.C, city}
		r = append(r, s.Response())
	}
	return r
}
