package benchmarks

import (
	"github.com/krzysztofdrys/pgo-talk/cities/city"
	"testing"
)

func BenchmarkDistance(b *testing.B) {
	b.StopTimer()
	dataset, err := city.Read()
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		city.Filter(dataset.Features, city.DistanceFilter{
			Lat:      51.12199803,
			Lng:      17.03799962,
			Distance: 100,
		})
	}
}

func BenchmarkComplexFilterDistance(b *testing.B) {
	b.StopTimer()
	dataset, err := city.Read()
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		city.Filter(dataset.Features, city.AndFilter{
			P1: city.DistanceFilter{
				Lat:      51.12199803,
				Lng:      17.03799962,
				Distance: 100,
			},
			P2: city.PopulationFilter{Pop: 100000},
		})
	}
}

func BenchmarkClosest(b *testing.B) {
	b.StopTimer()
	dataset, err := city.Read()
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		city.ClosestCities(51.12199803, 17.03799962, 1000, city.HaversineDistance{}, dataset.Features)
	}
}
