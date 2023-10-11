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



---
# Will profiles from version `1.0.1` work for `1.0.2`? 

Yes, mostly.

---
# What optimisations are 
