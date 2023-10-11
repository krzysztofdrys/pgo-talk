gcflags := -gcflags='github.com/...=-m=1'
count   := 100

all: nopgo.tests.pprof pgo.tests.pprof pgo_v2.tests.pprof summary.tests.txt

clean:
	rm -rf *.log *.times *.pprof summary.txt

nopgo.tests.pprof:
	go test -timeout 2h -bench . -cpuprofile nopgo.tests.pprof -pgo off -count $(count) $(gcflags) 2> nopgo.tests.build.log

pgo.tests.pprof: nopgo.tests.pprof
	go test -timeout 2h -bench . -cpuprofile pgo.tests.pprof -pgo nopgo.tests.pprof  -count $(count) $(gcflags) 2> pgo.tests.build.log

pgo_v2.tests.pprof: pgo.tests.pprof
	go test -timeout 2h -bench . -cpuprofile pgo_v2.tests.pprof -pgo pgo.tests.pprof  -count $(count) $(gcflags) 2> pgo_v2.tests.build.log

nopgo.tests.times:
	go test -timeout 2h -bench . -pgo off -count $(count) > nopgo.tests.times

pgo.tests.times: nopgo.tests.pprof
	go test -timeout 2h -bench . -pgo nopgo.tests.pprof  -count $(count) > pgo.tests.times

pgo_v2.tests.times: pgo.tests.pprof
	go test -timeout 2h -bench . -pgo pgo.tests.pprof  -count $(count) > pgo_v2.tests.times

summary.tests.txt: nopgo.tests.times pgo.tests.times pgo_v2.tests.times
	benchstat nopgo.tests.times pgo.tests.times  pgo_v2.tests.times > summary.tests.txt
