package distance

import (
	"encoding/json"
	"testing"
)

type Address struct {
	Name, City, PostalCode, Country string
	AddressLines                    []string
	Labels                          map[string]string
}

func BenchmarkJson(b *testing.B) {
	v := Address{
		Name:         "Wędrówki",
		City:         "Wrocław",
		AddressLines: []string{"Podwale", "37/38,"},
		PostalCode:   "50-040",
		Country:      "Poland",
		Labels: map[string]string{
			"type":      "pub",
			"free_beer": "sometimes",
		},
	}

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
}
