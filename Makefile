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
	${GO} test -count 1 ./binarytree

# Runs the test suite and collects code coverage.
coverage: assertions-off measurements-off
	${GO} test ./binarytree -coverprofile "coverage.out"
	go tool cover -html="coverage.out" -o "coverage.html"
	open coverage.html

profile: assertions-off measurements-off
	${GO} test ./binarytree -cpuprofile cpu.out
	go tool pprof cpu.out

animate: assertions-off measurements-off
	${GO} run main/animate/animate.go

assertions-on:
	${GO} run main/replace/replace.go -dir "./binarytree" -find "// assert(" -replace "assert("

assertions-off:
	${GO} run main/replace/replace.go -dir "./binarytree" -find "  assert(" -replace "  // assert("

measurements-on:
	${GO} run main/replace/replace.go -dir "./binarytree" -find "// measurement(" -replace "measurement("

measurements-off:
	${GO} run main/replace/replace.go -dir "./binarytree" -find "  measurement(" -replace "  // measurement("

sandbox:
	${GO} run main/sandbox/sandbox.go
















benchmarks: operation-benchmarks balancer-benchmarks

measurements: operation-measurements balancer-measurements


BALANCERS := \
	Median \
	Height \
	HalfSize \
	LogSize \
	HalfWeight \
	LogWeight \
	Cost \
	DSW

balancer-measurements-%:
	${GO} run benchmarks/main/balancer_measurements.go -strategy $*

balancer-measurements: assertions-off measurements-on
	$(MAKE) -j $(foreach balancer,$(BALANCERS),balancer-measurements-$(balancer))

balancer-benchmarks: assertions-off measurements-off
	${GO} run benchmarks/main/balancer_benchmarks.go




OPERATIONS := \
	Insert \
	InsertPersistent \
	InsertDelete \
	InsertDeletePersistent \
	InsertDeleteCycles \
	InsertDeleteCyclesPersistent \
	SplitJoin \


operation-benchmarks-%:
	${GO} run benchmarks/main/operation_benchmarks.go -operation $*

operation-benchmarks: assertions-off measurements-off
	$(MAKE) -j $(foreach operation,$(OPERATIONS),operation-benchmarks-$(operation))

operation-measurements-%:
	${GO} run benchmarks/main/operation_measurements.go -operation $*

operation-measurements: assertions-off measurements-on
	$(MAKE) -j $(foreach operation,$(OPERATIONS),operation-measurements-$(operation))






optimized-plots:
	npx --yes svgo --recursive --folder "benchmarks/svg"

plot-index:
	${GO} run benchmarks/main/index.go < benchmarks/index.html

balancer-plots:
	gnuplot benchmarks/plot/balancers.gnuplot

operation-plots:
	gnuplot benchmarks/plot/operations.gnuplot
