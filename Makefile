GO = GOEXPERIMENT=arenas go

#
shell:
	docker compose run --interactive --tty --rm bash

# Confirms that the tests pass, then runs all benchmarks, then parses the log.
#benchmarks: run-benchmarks parse-benchmarks plot-benchmarks

# Disables assertions to avoid slowdown, then runs the animation builder.
animation: animate

# Run the test suite.
test: assertions-on measurements-off
	${GO} test ./trees

# Runs the test suite and collects code coverage.
coverage: assertions-off measurements-off
	${GO} test ./trees -coverprofile "coverage.out"
	go tool cover -html="coverage.out" -o "coverage.html"
	open coverage.html

profile: assertions-off measurements-off
	${GO} test ./trees -cpuprofile cpu.out
	go tool pprof cpu.out

animate: assertions-off measurements-off
	${GO} run main/animate/animate.go

assertions-on:
	${GO} run main/replace/replace.go -dir "./trees" -find "// assert(" -replace "assert("

assertions-off:
	${GO} run main/replace/replace.go -dir "./trees" -find "  assert(" -replace "  // assert("

measurements-on:
	${GO} run main/replace/replace.go -dir "./trees" -find "// measurement(" -replace "measurement("

measurements-off:
	${GO} run main/replace/replace.go -dir "./trees" -find "  measurement(" -replace "  // measurement("

sandbox:
	@${GO} run main/sandbox/sandbox.go
















benchmarks: operation-benchmarks balancer-benchmarks

measurements: operation-measurements balancer-measurements

BALANCERS := \
	Median \
	Height \
	Weight \
	Log \
	Cost \
	DSW \

balancer-measurements-%:
	${GO} run docs/benchmarks/main/balancer_measurements.go -strategy $*

balancer-measurements: assertions-off measurements-on
	$(MAKE) -j $(foreach balancer,$(BALANCERS),balancer-measurements-$(balancer))

balancer-benchmarks: assertions-off measurements-off
	${GO} run docs/benchmarks/main/balancer_benchmarks.go




OPERATIONS := \
	Insert \
	InsertPersistent \
	InsertDelete \
	InsertDeletePersistent \
	InsertDeleteCycles \
	InsertDeleteCyclesPersistent \
	InsertDeleteSearch \
	InsertDeleteSearchPersistent \


operation-benchmarks-%:
	${GO} run docs/benchmarks/main/operation_benchmarks.go -operation $*

operation-benchmarks: assertions-off measurements-off
	$(MAKE) $(foreach operation,$(OPERATIONS),operation-benchmarks-$(operation))

operation-measurements-%:
	${GO} run docs/benchmarks/main/operation_measurements.go -operation $*

operation-measurements: assertions-off measurements-on
	$(MAKE) -j $(foreach operation,$(OPERATIONS),operation-measurements-$(operation))






optimized-plots:
	npx --yes svgo --recursive --folder  "."

plot-index:
	${GO} run docs/benchmarks/main/index.go < docs/benchmarks/index.html

balancer-plots:
	gnuplot docs/benchmarks/plot/balancers.gnuplot

operation-plots:
	gnuplot docs/benchmarks/plot/operations.gnuplot

polytope:
	gnuplot docs/polytope/topdown.gnuplot

index:
	${GO} run docs/main/index/index.go < docs/index.html > index.html

article: optimized-plots plot-index index

publish: article assertions-on measurements-on