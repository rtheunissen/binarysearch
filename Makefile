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
	${GO} run benchmarks/main/balancer_measurements.go > benchmarks/data/measurements/Balance

balancer-benchmarks: assertions-off measurements-off
	${GO} run benchmarks/main/balancer_benchmarks.go > benchmarks/data/Balance





define benchmark-operation
	${GO} run benchmarks/main/operation_benchmarks.go -operation $(1) > benchmarks/data/$(1)
endef

operation-benchmarks: assertions-off measurements-off
	$(call benchmark-operation,Insert)
	$(call benchmark-operation,InsertPersistent)
	$(call benchmark-operation,InsertDelete)
	$(call benchmark-operation,InsertDeletePersistent)
	$(call benchmark-operation,InsertDeleteCycles)
	$(call benchmark-operation,InsertDeleteCyclesPersistent)
	$(call benchmark-operation,SplitJoin)





operation-measurements: assertions-off measurements-on \
	operation-measurement-insert \
	operation-measurement-insert-persistent \
	operation-measurement-insert-delete \
	operation-measurement-insert-delete-persistent \
	operation-measurement-swell \
	operation-measurement-swell-persistent \
	operation-measurement-split-join \

define measure-operation
	${GO} run benchmarks/main/operation_measurements.go -operation $(1) > benchmarks/data/measurements/$(1)
endef

operation-measurement-insert:
	$(call measure-operation,Insert)

operation-measurement-insert-persistent:
	$(call measure-operation,InsertPersistent)

operation-measurement-insert-delete:
	$(call measure-operation,InsertDelete)

operation-measurement-insert-delete-persistent:
	$(call measure-operation,InsertDeletePersistent)

operation-measurement-swell:
	$(call measure-operation,InsertDeleteCycles)

operation-measurement-swell-persistent:
	$(call measure-operation,InsertDeleteCyclesPersistent)

operation-measurement-split-join:
	$(call measure-operation,SplitJoin)












plot-index:
	${GO} run benchmarks/main/index.go < benchmarks/index.html

plot: plot-balancers plot-operations

plot-balancers:
	gnuplot benchmarks/plot/balancers.gnuplot

plot-operations:
	gnuplot benchmarks/plot/operations.gnuplot


# make measurements -j && make benchmarks && make plot -j

overnight-operations:
	time $(MAKE) operation-measurements -j
	time $(MAKE) operation-benchmarks

overnight:
	time $(MAKE) operation-measurements -j
	time $(MAKE) balancer-measurements -j
	time $(MAKE) operation-benchmarks
	time $(MAKE) balancer-benchmarks