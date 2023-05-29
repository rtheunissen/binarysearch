package main

import (
	"flag"
	"trees/benchmarks"
	"trees/binarytree"
	"trees/distribution"
)

func main() {
	iterations := flag.Int("iterations", 100, "")
	samples := flag.Int("samples", 100, "")
	scale := flag.Int("scale", 1_000_000, "")
	flag.Parse()

	benchmarks.BalancerBenchmark{
		Samples:    *samples,
		Scale:      *scale,
		Iterations: *iterations,
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
