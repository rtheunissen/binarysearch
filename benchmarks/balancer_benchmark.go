package benchmarks

import (
	"binarysearch/abstract/list"
	"binarysearch/binarytree"
	"binarysearch/distribution"
	"binarysearch/utility"
	"fmt"
	"time"
)

type BalancerBenchmark struct {
	Scale         int
	Samples       int
	Strategies    []binarytree.Balancer
	Distributions []distribution.Distribution
	Iterations    int
}

func (benchmark BalancerBenchmark) Run() {

	//
	for _, strategy := range benchmark.Strategies {

		//
		//
		header := []any{
			"Distribution",
			"Scale",
			"Size",
			"Duration",
		}
		stdout("#" + utility.NameOf(strategy))
		stdout(header...)

		step := benchmark.Scale / benchmark.Samples

		//
		for iteration := 1; iteration <= benchmark.Iterations; iteration++ {

			instance := binarytree.Splay{}.New().(*binarytree.Splay)

			stderr([]any{
				alignL(utility.NameOf(strategy)),
				alignR(fmt.Sprint(iteration, "/", benchmark.Iterations)),
			}...)

			for position := step; position <= benchmark.Scale; position += step {

				// Grow the tree.
				for instance.Size() < list.Size(position) {
					instance.Insert(0, 0)
				}

				for _, random := range benchmark.Distributions {

					// Randomize the tree.
					instance.Tree = instance.Tree.Randomize(random.New(uint64(iteration)))

					start := time.Now()

					instance.Tree = strategy.Restore(instance.Tree)

					duration := time.Since(start)

					row := []any{
						utility.NameOf(random),
						fmt.Sprint(benchmark.Scale),
						fmt.Sprint(instance.Size()),
						fmt.Sprint(duration.Nanoseconds()),
					}
					stdout(row...)
				}
			}
			instance.Free()
		}
		stdout()
		stdout()
	}
}
