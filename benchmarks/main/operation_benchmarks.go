package main

import (
	"binarysearch/abstract/list"
	"binarysearch/benchmarks"
	"binarysearch/binarytree"
	"binarysearch/distribution"
	operations2 "binarysearch/operations"
	"binarysearch/utility"
	"flag"
)

func main() {
	operation := flag.String("operation", "", "")
	iterations := flag.Int("iterations", 10, "")
	samples := flag.Int("samples", 1000, "")
	scale := flag.Int("scale", 1_000_000, "")
	flag.Parse()

	benchmarks.TreeBenchmark{
		Scale:      *scale,
		Samples:    *samples,
		Iterations: *iterations,
		Operation: utility.Resolve(*operation, []list.Operation{
			&operations2.Insert{},
			//&operations2.InsertPersistent{},
			//&operations2.InsertDelete{},
			//&operations2.InsertDeletePersistent{},
			//&operations2.InsertDeleteCycles{},
			&operations2.InsertDeleteCyclesPersistent{},
			//&operations2.SplitJoin{},
		}),
		Distributions: []distribution.Distribution{
			&distribution.Uniform{},
			&distribution.Normal{},
			&distribution.Skewed{},
			&distribution.Zipf{},
			&distribution.Maximum{},
		},
		Strategies: []list.List{
			//&binarytree.AVLBottomUp{},
			//&binarytree.AVLJoinBased{},
			//&binarytree.AVLWeakTopDown{},
			//&binarytree.AVLWeakBottomUp{},
			//&binarytree.AVLWeakJoinBased{},
			//&binarytree.AVLRelaxedTopDown{},
			//&binarytree.AVLRelaxedBottomUp{},
			//&binarytree.RedBlackRelaxedBottomUp{},
			//&binarytree.RedBlackRelaxedTopDown{},
			&binarytree.LBSTBottomUp{},
			&binarytree.LBSTTopDown{},
			&binarytree.LBSTJoinBased{},
			&binarytree.LBSTRelaxed{},
			//&binarytree.TreapTopDown{},
			//&binarytree.TreapJoinBased{},
			//&binarytree.TreapFingerTree{},
			//&binarytree.Randomized{},
			//&binarytree.Zip{},
			//&binarytree.Splay{},
			//&binarytree.Conc{},
		},
	}.Run()
}
