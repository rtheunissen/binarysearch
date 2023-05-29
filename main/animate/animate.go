package main

import (
	"flag"
	"os"
	"trees/abstract/list"
	"trees/binarytree/animations"
	console2 "trees/console"
	"trees/distribution"
	operations2 "trees/operations"
)

func main() {
	flag.Parse()
	console2.Animate(prompt())
}

func prompt() console2.Animation {
	//var binaryTree BST = &Splay{} // TODO what about an unbalanced tree? Can BST be List?
	//var distribution number.Distribution = &Uniform{}

	// Choose a strategy and set up a new tree for the animation.

	//switch operation.(type) {
	//   //
	//   // // No need to prompt for these if the operation is generic.
	//   //
	//   //case *operations.BalanceWeight:
	//   case *operations.BalanceHeight:
	//   case *operations.Randomize:
	//default:
	animation := animations.BinaryTreeAnimation{
		Writer:       os.Stdout,
		Operation:    chooseOperation(),
		List:         chooseStrategy(),
		Distribution: chooseDistribution().New(123),
		Size:         1_000_000,
		Height:       40,
	}
	return console2.Choose[console2.Animation]("Animation",
		&animations.ExteriorHeights{BinaryTreeAnimation: animation},
		&animations.InteriorHeights{BinaryTreeAnimation: animation},
		&animations.WeightsPerLevel{BinaryTreeAnimation: animation},
	)
}

func chooseOperation() list.Operation {
	return console2.Choose[list.Operation]("Operation",
		&operations2.Insert{},
		&operations2.Delete{},
		&operations2.InsertDeleteCycles{},
		&operations2.SplitJoin{},
	)
}

func chooseDistribution() distribution.Distribution {
	return console2.Choose[distribution.Distribution]("Distribution",
		&distribution.Uniform{},
		&distribution.Normal{},
		&distribution.Skewed{},
		&distribution.Minimum{},
		&distribution.Maximum{},
		&distribution.Queue{},
		&distribution.Parabolic{},
		&distribution.Slope{},
		&distribution.UShape{},
		&distribution.Median{},
		&distribution.Ascending{},
		&distribution.Descending{},
		&distribution.BiModal{},
		&distribution.Zipf{},
	)
}

func chooseStrategy() list.List {
	return console2.Choose[list.List]("Strategy",
		&AVLBottomUp{},
		&AVLJoinBased{},
		&AVLWeakTopDown{},
		&AVLWeakBottomUp{},
		&AVLWeakJoinBased{},
		&AVLRelaxedTopDown{},
		&AVLRelaxedBottomUp{},
		&RedBlackRelaxedBottomUp{},
		&RedBlackRelaxedTopDown{},
		&LBSTBottomUp{},
		&LBSTTopDown{},
		&LBSTJoinBased{},
		&LBSTRelaxed{},
		&TreapTopDown{},
		&TreapJoinBased{},
		&TreapFingerTree{},
		&Randomized{},
		&Zip{},
		&Splay{},
		&Conc{},
	)
}
