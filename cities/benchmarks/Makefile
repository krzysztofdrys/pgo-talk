gcflags := -gcflags='github.com/...=-m=1'
count   := 100

all: nopgo.pprof pgo.pprof summary.txt

clean:
	rm -rf *.log *.times *.pprof summary.txt

nopgo.pprof:
	go test -timeout 2h -bench . -cpuprofile nopgo.pprof -pgo off -count $(count) $(gcflags) 2> nopgo.build.log

pgo.pprof: nopgo.pprof
	go test -timeout 2h -bench . -cpuprofile pgo.pprof -pgo nopgo.pprof  -count $(count) $(gcflags) 2> pgo.build.log

pgo_v2.pprof: pgo.pprof
	go test -timeout 2h -bench . -cpuprofile pgo_v2.pprof -pgo pgo.pprof  -count $(count) $(gcflags) 2> pgo_v2.build.log

nopgo.times:
	go test -timeout 2h -bench . -pgo off -count $(count) > nopgo.times

pgo.times: nopgo.pprof
	go test -timeout 2h -bench . -pgo nopgo.pprof  -count $(count) > pgo.times

pgo_v2.times: pgo.pprof
	go test -timeout 2h -bench . -pgo pgo.pprof  -count $(count) > pgo_v2.times

summary.txt: nopgo.times pgo.times pgo_v2.times
	benchstat nopgo.times pgo.times  pgo_v2.times > summary.txt
