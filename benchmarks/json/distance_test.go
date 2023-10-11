package distance

import (
	"encoding/json"
	"testing"

	"github.com/krzysztofdrys/pgo-talk/city"
)

func BenchmarkJson(b *testing.B) {
	b.StopTimer()
	// Reads 1.1M of json data
	cs, err := city.Read()
	if err != nil {
		panic(err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(cs)
		if err != nil {
			panic(err)
		}
	}
}
