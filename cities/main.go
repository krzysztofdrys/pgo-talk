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

	for c1 := 'A'; c1 <= 'Z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			for c3 := 'a'; c3 <= 'z'; c3++ {
				prefix := string(c1) + string(c2) + string(c3)
				f := city.AndFilter{
					P1: city.NamePrefixFilter{
						Prefix: prefix,
					},
					P2: city.PopulationFilter{
						Pop: 100000,
					},
				}
				r := city.Filter(dataset.Features, f)
				_, err := json.Marshal(r)
				if err != nil {
					panic(err)
				}
			}
		}
	}

}
