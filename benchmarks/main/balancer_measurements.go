package main

import (
	"flag"
	"trees/benchmarks"
	"trees/binarytree"
	"trees/distribution"
)

func main() {
	samples := flag.Int("samples", 10_000, "")
	scale := flag.Int("scale", 1_000_000, "")
	flag.Parse()

	benchmarks.BalancerMeasurement{
		Samples: *samples,
		Scale:   *scale,
		Measurements: []binarytree.Measurement{
			&binarytree.PartitionCount{},
			&binarytree.PartitionDepth{},
			&binarytree.AveragePathLength{},
			&binarytree.MaximumPathLength{},
			&binarytree.Rotations{},
		},
		Distributions: []distribution.Distribution{
			&distribution.Uniform{},
			&distribution.Normal{},
			&distribution.Skewed{},
			&distribution.Zipf{},
			&distribution.Maximum{},
		},
		Strategies: []binarytree.Balancer{
			&binarytree.Median{},
			&binarytree.Height{},
			&binarytree.Weight{},
			&binarytree.DSW{},
		},
	}.Run()
}
