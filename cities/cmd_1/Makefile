gcflags := -gcflags='github.com/...=-m=2' -gcflags='json=-m=1'

all: pgo.pprof results.md std.pprof

clean:
	rm cities results.md cities.pgo *.pprof *.log

cities: std.build.log
	go build -pgo off -v -o cities .

cities.pgo: std.pprof pgo.build.log
	go build -pgo std.pprof -v -o cities.pgo .

std.build.log:
	go build -pgo off $(gcflags) -a . 2> std.build.log

pgo.build.log: std.pprof
	go build -pgo std.pprof $(gcflags) -a . 2> pgo.build.log

std.pprof: cities
	./cities -cpuprofile std.pprof > /dev/null

pgo.pprof: cities.pgo
	./cities.pgo -cpuprofile pgo.pprof > /dev/null

results.md: cities.pgo cities
	hyperfine \
		--export-markdown results.md \
		--warmup 1 \
		'./cities' \
		'./cities.pgo'
