package main

import (
	"encoding/json"
	"flag"
	"github.com/krzysztofdrys/pgo-talk/cities/city"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "")
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	dataset, err := city.Read()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000000; i++ {
		cs := city.ClosestCities(51.12199803, 17.03799962, 1000, city.HaversineDistance{}, dataset.Features)
		_, err := json.Marshal(cs)
		if err != nil {
			panic(err)
		}
	}

}
