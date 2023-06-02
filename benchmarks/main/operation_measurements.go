package main

import (
   "binarysearch/abstract/list"
   "binarysearch/benchmarks"
   "binarysearch/binarytree"
   "binarysearch/distribution"
   operations "binarysearch/operations"
   "binarysearch/utility"
   "flag"
)

func main() {
   operation := flag.String("operation", "", "")
   samples := flag.Int("samples", 1_000, "")
   scale := flag.Int("scale", 1_000_000, "")
   flag.Parse()

   benchmarks.OperationMeasurement{
      Scale:   *scale,
      Samples: *samples,
   	Operation: utility.Resolve(*operation, []list.Operation{
   		&operations.Insert{},
   		&operations.InsertPersistent{},
   		&operations.InsertDelete{},
   		&operations.InsertDeletePersistent{},
   		&operations.InsertDeleteCycles{},
   		&operations.InsertDeleteCyclesPersistent{},
   		&operations.SplitJoin{},
   	}),
   	Measurements: []binarytree.Measurement{
   		&binarytree.PartitionCount{},
   		&binarytree.PartitionDepth{},
   		&binarytree.AveragePathLength{},
   		&binarytree.MaximumPathLength{},
   		&binarytree.Rotations{},
   		&binarytree.Allocations{},
   	},
   	Distributions: []distribution.Distribution{
   		&distribution.Uniform{},
   		&distribution.Normal{},
   		&distribution.Skewed{},
   		&distribution.Zipf{},
   		&distribution.Maximum{},
   	},
   	Strategies: []list.List{
   		&binarytree.AVLBottomUp{},
   		&binarytree.AVLJoinBased{},
   		&binarytree.AVLWeakTopDown{},
   		&binarytree.AVLWeakBottomUp{},
   		&binarytree.AVLWeakJoinBased{},
   		&binarytree.AVLRelaxedTopDown{},
   		&binarytree.AVLRelaxedBottomUp{},
   		&binarytree.RedBlackRelaxedTopDown{},
   		&binarytree.RedBlackRelaxedBottomUp{},
   		&binarytree.LBSTBottomUp{},
   		&binarytree.LBSTTopDown{},
   		&binarytree.LBSTJoinBased{},
   		&binarytree.LBSTRelaxed{},
   		&binarytree.TreapTopDown{},
   		&binarytree.TreapJoinBased{},
   		&binarytree.TreapFingerTree{},
   		&binarytree.Randomized{},
   		&binarytree.Zip{},
   		&binarytree.Splay{},
   		&binarytree.Conc{},
   	},
   }.Run()
}
