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

balancer-measurements: assertions-off measurements-on
	${GO} run benchmarks/main/balancer_measurements.go

balancer-benchmarks: assertions-off measurements-off
	${GO} run benchmarks/main/balancer_benchmarks.go





define benchmark-operation
	${GO} run benchmarks/main/operation_benchmarks.go -operation $(1)
endef

operation-benchmarks: assertions-off measurements-off
	$(call benchmark-operation,Insert)
	$(call benchmark-operation,InsertPersistent)
	$(call benchmark-operation,InsertDelete)
	$(call benchmark-operation,InsertDeletePersistent)
	$(call benchmark-operation,InsertDeleteCycles)
	$(call benchmark-operation,InsertDeleteCyclesPersistent)
	#$(call benchmark-operation,SplitJoin)





operation-measurements: assertions-off measurements-on
	$(MAKE) -j \
		operation-measurement-insert \
		operation-measurement-insert-persistent \
		operation-measurement-insert-delete \
		operation-measurement-insert-delete-persistent \
		operation-measurement-insert-delete-cycles \
		operation-measurement-insert-delete-cycles-persistent \
		operation-measurement-split-join \

define measure-operation
	${GO} run benchmarks/main/operation_measurements.go -operation $(1)
endef

operation-measurement-insert:
	$(call measure-operation,Insert)

operation-measurement-insert-persistent:
	$(call measure-operation,InsertPersistent)

operation-measurement-insert-delete:
	$(call measure-operation,InsertDelete)

operation-measurement-insert-delete-persistent:
	$(call measure-operation,InsertDeletePersistent)

operation-measurement-insert-delete-cycles:
	$(call measure-operation,InsertDeleteCycles)

operation-measurement-insert-delete-cycles-persistent:
	$(call measure-operation,InsertDeleteCyclesPersistent)

operation-measurement-split-join:
	$(call measure-operation,SplitJoin)












plot-index:
	${GO} run benchmarks/main/index.go < benchmarks/svg/index.html

plot: plot-balancers plot-operations

plot-balancers: \
	plot-balancer-measurements \
	plot-balancer-benchmarks

plot-balancer-measurements:
	gnuplot benchmarks/plot/balancers/balancer_measurements.gnuplot

plot-balancer-benchmarks:
	gnuplot benchmarks/plot/balancers/balancer_benchmarks.gnuplot

plot-operations: \
	plot-operation-benchmarks \
	plot-operation-measurements \

plot-operation-benchmarks:
	gnuplot benchmarks/plot/operation_benchmarks.gnuplot

plot-operation-measurements:
	gnuplot benchmarks/plot/operation_measurements.gnuplot

overnight:
	$(MAKE) operation-measurements balancer-measurements -j
	$(MAKE) operation-benchmarks
	$(MAKE) balancer-benchmarks