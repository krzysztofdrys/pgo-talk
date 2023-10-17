---
marp: true
paginate: false
html: true
---
# Profile-guided optimization
Krzysztof Dryś
2023-10-12
<!-- theme: leibniz -->
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
5. Deploy new binary to prod,
6. Observe performance gains,
7. GOTO (3)

---
# Let's see if this actually works

---
# Haversine formula

The haversine formula determines the great-circle distance between two points on a sphere given their longitudes and latitudes. Important in navigation, it is a special case of a more general formula in spherical trigonometry, the law of haversines, that relates the sides and angles of spherical triangles. 

(from [wikipedia](https://en.wikipedia.org/wiki/Haversine_formula))

Implemented by [github.com/jftuga/geodist](https://github.com/jftuga/geodist):

```go
var elPaso = geodist.Coord{Lat: 31.7619, Lon: 106.4850}
var stLouis = geodist.Coord{Lat: 38.6270, Lon: 90.1994}
miles, km = geodist.HaversineDistance(elPaso, stLouis)
fmt.Printf("[Haversine] El Paso to St. Louis:  %.3f m, %.3f km\n", miles, km)
```
---
# Demo

This time for real.

---
# Results

All tests were run on `n2-highcpu-4` machine with Intel Cascade Lake CPU.

# Results

```
goos: linux
goarch: amd64
pkg: github.com/krzysztofdrys/pgo-talk/benchmarks/distance
cpu: Intel(R) Xeon(R) CPU @ 2.80GHz
           │ nopgo.tests.times │       pgo.tests.times       │     pgo_v2.tests.times      │
           │      sec/op       │   sec/op     vs base        │   sec/op     vs base        │
Distance-2         125.0n ± 0%   118.4n ± 0%  -5.28% (n=100)   118.3n ± 0%  -5.36% (n=100)
```

---
# Distance (parallel)

```go
func BenchmarkDistance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			distance()
		}
	})
}

func distance() {
	var elPaso = geodist.Coord{Lat: 31.7619, Lon: 106.4850}
	var stLouis = geodist.Coord{Lat: 38.6270, Lon: 90.1994}
	geodist.HaversineDistance(elPaso, stLouis)
}
```
---
# Summary

```go
goos: linux
goarch: amd64
pkg: github.com/krzysztofdrys/pgo-talk/benchmarks/distance_paralel
cpu: Intel(R) Xeon(R) CPU @ 2.80GHz
           │ nopgo.tests.times │       pgo.tests.times       │     pgo_v2.tests.times      │
           │      sec/op       │   sec/op     vs base        │   sec/op     vs base        │
Distance-2         62.65n ± 0%   59.31n ± 0%  -5.33% (n=100)   59.29n ± 0%  -5.36% (n=100)
```

---

# JSON marshalling

```go
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
Json-2         5.263m ± 0%   4.574m ± 1%  -13.11% (n=100)   4.714m ± 0%  -10.44% (n=100)

```

---
# JSON marshalling (reprise)


```go
import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkJson(b *testing.B) {
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
pkg: github.com/krzysztofdrys/pgo-talk/benchmarks/json-iterator
cpu: Intel(R) Xeon(R) CPU @ 2.80GHz
       │ nopgo.tests.times │       pgo.tests.times       │     pgo_v2.tests.times      │
       │      sec/op       │   sec/op     vs base        │   sec/op     vs base        │
Json-2         5.151m ± 0%   4.739m ± 0%  -8.00% (n=100)   4.695m ± 0%  -8.84% (n=100)

```

---
# Rendering HTML from Markdown

```go
func BenchmarkMarkdownRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md := markdown.New(
			markdown.XHTMLOutput(true),
			markdown.Typographer(true),
			markdown.Linkify(true),
			markdown.Tables(true),
		)

		var buf bytes.Buffer
		if err := md.Render(&buf, src); err != nil {
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
pkg: github.com/krzysztofdrys/pgo-talk/benchmarks/markdown
cpu: Intel(R) Xeon(R) CPU @ 2.80GHz
                 │ nopgo.tests.times │       pgo.tests.times       │         pgo_v2.tests.times          │
                 │      sec/op       │   sec/op     vs base        │   sec/op     vs base                │
MarkdownRender-2         76.27µ ± 0%   74.46µ ± 0%  -2.37% (n=100)   76.05µ ± 0%  -0.29% (p=0.044 n=100)
```

Note: [article on go.dev blog](https://go.dev/blog/pgo) reports ~3.8% improvement for web server running this converter 🤷‍. ️ 

---
# Q&A

Part where I ask questions and I answer them

---
# Will profiles from version `1.0.1` work for `1.0.2`? 

Yes, mostly.

---
# Source stability

Specifically, Go uses line offsets within functions (e.g., call on 5th line of function foo).

Many common changes will not break matching, including:

- Changes in a file outside of a hot function (adding/changing code above or below the function).
- Moving a function to another file in the same package (the compiler ignores source filenames altogether). 

Some changes that may break matching:

- Changes within a hot function (may affect line offsets).
- Renaming the function (and/or type for methods) (changes symbol name).
- Moving the function to another package (changes symbol name).

(from [go documentation](https://go.dev/doc/pgo#source-stability))

---
# AutoFDO: Automatic Feedback-Directed Optimization for Warehouse-Scale Applications

It's a [paper](https://storage.googleapis.com/pub-tools-public-publication-data/pdf/45290.pdf) from Google.  It's only 12 pages!

Key points:
- advocates for AutoFDO flow,
- stale source usually means about 90% of performance gains from PGO,
- getting the right profiles is key.

---
# Getting the right profiles is key

Using PGO compiler will only apply optimisations to paths, which are "hot" in the profile. Things which might not be affected:
- startup operations,
- batch jobs, which didn't happen when you were taking taking the profile,
- traffic spikes using different endpoints,
- new code.

---
# What optimisations are enabled by PGO?

Right now PGO supports two optimisations:
- function inlining,
- devirtualisation.

More will (most likely come) resulting more in better results.

---
# Where do I keep prod pprof data?

Preferably in git, right where you do `go build`. If you name file `default.pgo`, then go will pick it up automatically.

- This gives you repeatable builds,
- You can update `default.pgo` every day/week/month with fresh data,
- Whenever you update `default.pgo`, you need to rebuild every package.

---
# Thank you!
