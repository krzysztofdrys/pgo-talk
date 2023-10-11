---
marp: true
paginate: false
html: true
---
# Profile-guided optimization
Krzysztof Dryś
2023-10-12
<!-- theme: gaia -->
<style>
img[alt~="center"] {
  display: block;
  margin: 0 auto;
}
</style>
---
# Improve performance of you app by up to 7% using this one simple trick

---
# PGO lifecycle

1. `go build .`
2. Run your application in production,
3. `curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"` to profile your application,
4. `go build -pgo cpu.pprof .`
5. Deploy new profile to prod,
6. Observe performance gains.

---
# Demo

---
# Haversine formula

The haversine formula determines the great-circle distance between two points on a sphere given their longitudes and latitudes. Important in navigation, it is a special case of a more general formula in spherical trigonometry, the law of haversines, that relates the sides and angles of spherical triangles. [from wikipedia](https://en.wikipedia.org/wiki/Haversine_formula)

Implemented by [github.com/jftuga/geodist](https://github.com/jftuga/geodist).

```go
var elPaso = geodist.Coord{Lat: 31.7619, Lon: 106.4850}
var stLouis = geodist.Coord{Lat: 38.6270, Lon: 90.1994}
miles, km = geodist.HaversineDistance(elPaso, stLouis)
fmt.Printf("[Haversine] El Paso to St. Louis:  %.3f m, %.3f km\n", miles, km)
```

---
# Results

```
goos: linux
goarch: amd64
pkg: github.com/krzysztofdrys/pgo-talk/benchmarks/distance_paralel
cpu: Intel(R) Xeon(R) CPU @ 2.80GHz
           │ nopgo.tests.times │       pgo.tests.times       │     pgo_v2.tests.times      │
           │      sec/op       │   sec/op     vs base        │   sec/op     vs base        │
Distance-2         92.10n ± 0%   89.97n ± 0%  -2.31% (n=100)   89.94n ± 0%  -2.35% (n=100)
```

---
# JSON marshalling

```go
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
```
---
# Results

```
goos: linux
goarch: amd64
pkg: github.com/krzysztofdrys/pgo-talk/benchmarks/json
cpu: Intel(R) Xeon(R) CPU @ 2.80GHz
       │ nopgo.tests.times │       pgo.tests.times        │      pgo_v2.tests.times      │
       │      sec/op       │   sec/op     vs base         │   sec/op     vs base         │
Json-2         4.302m ± 0%   3.748m ± 0%  -12.90% (n=100)   3.816m ± 0%  -11.32% (n=100)

```

---
# Will profiles from version `1.0.1` work for `1.0.2`? 

Yes, mostly.

---
# What optimisations are 
