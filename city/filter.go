package city

import (
	"github.com/jftuga/geodist"
	"strings"
)

type Predicate interface {
	Check(City) bool
}

func Filter(cs []City, p Predicate) []City {
	var result []City
	for _, c := range cs {
		if p.Check(c) {
			result = append(result, c)
		}
	}
	return result
}

type NamePrefixFilter struct {
	Prefix string
}

func (f NamePrefixFilter) Check(c City) bool {
	return strings.HasPrefix(c.Properties.Name, f.Prefix)
}

type AndFilter struct {
	P1, P2 Predicate
}

func (af AndFilter) Check(c City) bool {
	b := af.P1.Check(c)
	if !b {
		return false
	}
	return af.P2.Check(c)
}

type PopulationFilter struct {
	Pop int
}

func (f PopulationFilter) Check(c City) bool {
	return c.Properties.Pop > f.Pop
}

type DistanceFilter struct {
	Lat, Lng float64
	Distance float64
}

func (df DistanceFilter) Check(c City) bool {
	p1 := geodist.Coord{
		Lat: c.Properties.Lat,
		Lon: c.Properties.Lon,
	}
	p2 := geodist.Coord{
		Lat: df.Lat,
		Lon: df.Lng,
	}
	_, d := geodist.HaversineDistance(p1, p2)
	return d < df.Distance
}
