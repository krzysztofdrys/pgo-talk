gcflags := -gcflags='github.com/...=-m=1'
count   := 100

all: nopgo.pprof pgo.pprof summary.md

clean:
	rm -rf *.log *.times *.pprof summary.txt

nopgo.pprof:
	go test -bench . -cpuprofile nopgo.pprof -pgo off -count $(count) $(gcflags) 2> nopgo.build.log > nopgo.times

pgo.pprof: nopgo.pprof
	go test -bench . -cpuprofile pgo.pprof -pgo nopgo.pprof  -count $(count) $(gcflags) 2> pgo.build.log > pgo.times

pgo_v2.pprof: pgo.pprof
	go test -bench . -cpuprofile pgo_v2.pprof -pgo pgo.pprof  -count $(count) $(gcflags) 2> pgo_v2.build.log > pgo_v2.times

summary.md: nopgo.pprof pgo.pprof pgo_v2.pprof
	benchstat nopgo.times pgo.times  pgo_v2.times > summary.md
