package distance

import (
	"testing"

	"github.com/jftuga/geodist"
)

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		distance()
	}
}

func distance() {
	getDistance()
}

func getDistance() float64 {
	var elPaso = geodist.Coord{Lat: 31.7619, Lon: 106.4850}
	var stLouis = geodist.Coord{Lat: 38.6270, Lon: 90.1994}
	_, km := geodist.HaversineDistance(elPaso, stLouis)
	return km
}
