package cities

import "github.com/gin-gonic/gin"

type CityResponse struct {
	Name       string  `json:"name"`
	CoordLat   float64 `json:"coordLat"`
	CoordLon   float64 `json:"coordLon"`
	Population int     `json:"population"`
	TimeZone   int     `json:"timeZone"`
	Sunrise    int     `json:"sunrise"`
	Sunset     int     `json:"sunset"`
}

type CitySerializer struct {
	C *gin.Context
	City
}

func (s *CitySerializer) Response() CityResponse {
	response := CityResponse{
		Name:       s.Name,
		CoordLat:   s.CoordLat,
		CoordLon:   s.CoordLon,
		Population: s.Population,
		TimeZone:   s.TimeZone,
		Sunrise:    s.Sunrise,
		Sunset:     s.Sunset,
	}
	return response
}

type CitiesSerializer struct {
	C      *gin.Context
	Cities []City
}

func (s *CitiesSerializer) Response() []CityResponse {
	return nil
}
