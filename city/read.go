package city

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed cities.geojson
var data []byte

func Read() (Dataset, error) {
	var dataset Dataset
	if err := json.Unmarshal(data, &dataset); err != nil {
		return dataset, fmt.Errorf("failed to read dataset: %w", err)
	}
	return dataset, nil
}
