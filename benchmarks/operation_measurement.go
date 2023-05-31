package benchmarks

import (
	"binarysearch/abstract/list"
	"binarysearch/binarytree"
	"binarysearch/distribution"
	"binarysearch/utility"
	"fmt"
)

type OperationMeasurement struct {
	Scale         int
	Samples       int
	Operation     list.Operation
	Distributions []distribution.Distribution
	Strategies    []list.List
	Measurements  []binarytree.Measurement
}

func (measurement OperationMeasurement) Run() {
	if measurement.Operation == nil {
		return
	}

	//
	for _, strategy := range measurement.Strategies {

		//
		header := []any{
			"Distribution",
			"Scale",
			"Size",
			"Step",
			"Position",
		}
		//
		for _, measurement := range measurement.Measurements {
			header = append(header, utility.NameOf(measurement))
		}
		//
		stdout("#" + utility.NameOf(strategy))
		stdout(header...)

		//
		for _, random := range measurement.Distributions {
			//
			//
			//
			stderr(
				alignL(utility.NameOf(measurement.Operation)),
				alignL(utility.NameOf(strategy)),
				alignL(utility.NameOf(random)))

			//
			instance := measurement.Operation.Setup(strategy, list.Size(measurement.Scale))

			//
			access := random.New(1)

			//
			step := measurement.Scale / measurement.Samples

			//
			for position := step; position <= measurement.Scale; position = position + step {

				//
				for _, measurement := range measurement.Measurements {
					measurement.Reset()
				}
				//
				for i := 0; i < step; i++ {
					instance, _ = measurement.Operation.Update(instance, access)
				}
				//
				row := []any{
					utility.NameOf(access),
					fmt.Sprint(measurement.Scale),
					fmt.Sprint(instance.Size()),
					fmt.Sprint(step),
					fmt.Sprint(position),
				}
				//
				for _, measurement := range measurement.Measurements {
					row = append(row, fmt.Sprint(measurement.Measure(instance.(binarytree.BinaryTree))))
				}
				stdout(row...)
			}
			instance.Free()
		}
		stdout()
		stdout()
	}
}
