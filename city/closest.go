package city

import (
	"github.com/jftuga/geodist"
	"math"
)

type CityAndDistance struct {
	City     City    `json:"name"`
	Distance float64 `json:"distance"`
}

type closestCities []CityAndDistance

func (cs closestCities) push(cd CityAndDistance) {
	if cd.Distance > cs[len(cs)-1].Distance {
		return
	}
	cs[len(cs)-1] = cd
	for i := len(cs) - 1; i > 0; i-- {
		if cs[i].Distance < cs[i-1].Distance {
			cs[i], cs[i-1] = cs[i-1], cs[i]
		}
	}

}

type Distance interface {
	Distance(c City, lat, lng float64) float64
}

type HaversineDistance struct{}

func (hv HaversineDistance) Distance(c City, lat, lng float64) float64 {
	p1 := geodist.Coord{
		Lat: c.Properties.Lat,
		Lon: c.Properties.Lon,
	}
	p2 := geodist.Coord{
		Lat: lat,
		Lon: lng,
	}
	_, d := geodist.HaversineDistance(p1, p2)
	return d
}

func ClosestCities(lat, lng float64, n int, d Distance, cs []City) []CityAndDistance {
	result := make(closestCities, n)
	for i := range result {
		result[i].Distance = math.Inf(1)
	}
	for _, c := range cs {
		d := d.Distance(c, lat, lng)
		result.push(CityAndDistance{Distance: d, City: c})
	}
	return result
}
