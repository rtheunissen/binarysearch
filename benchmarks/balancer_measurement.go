package benchmarks

import (
	"binarysearch/abstract/list"
	"binarysearch/binarytree"
	"binarysearch/distribution"
	"binarysearch/utility"
	"fmt"
)

type BalancerMeasurement struct {
	Scale         int
	Samples       int
	Strategies    []binarytree.Balancer
	Measurements  []binarytree.Measurement
	Distributions []distribution.Distribution
}

func (measurement BalancerMeasurement) Run() {

	//
	for _, strategy := range measurement.Strategies {

		//
		header := []any{
			"Distribution",
			"Scale",
			"Size",
		}

		//
		for _, measurement := range measurement.Measurements {
			header = append(header, utility.NameOf(measurement))
		}
		stdout("#" + utility.NameOf(strategy))
		stdout(header...)

		instance := binarytree.Splay{}.New().(*binarytree.Splay)

		step := measurement.Scale / measurement.Samples

		for position := step; position <= measurement.Scale; position += step {

			// Grow the tree if needed.
			for instance.Size() < list.Size(position) {
				instance.Insert(0, 0)
			}

			stderr([]any{
				alignL(utility.NameOf(strategy)),
				alignR(fmt.Sprintf("%d / %d", position, measurement.Scale)),
			}...)

			for _, random := range measurement.Distributions {

				// Randomize the tree.
				instance.Tree = instance.Tree.Randomize(random.New(uint64(position)))

				// Reset measurements.
				for _, measurement := range measurement.Measurements {
					measurement.Reset()
				}

				instance.Tree = strategy.Restore(instance.Tree)

				row := []any{
					utility.NameOf(random),
					fmt.Sprint(measurement.Scale),
					fmt.Sprint(position),
				}
				for _, measurement := range measurement.Measurements {
					row = append(row, fmt.Sprint(measurement.Measure(instance.Tree)))
				}
				stdout(row...)
			}
		}
		instance.Free()
		stdout()
		stdout()
	}
}
