package main

import (
   "bst/abstract/list"
   "bst/abstract/list/operations"
   "bst/trees"
   "bst/trees/animations"
   "bst/utility/console"
   "bst/utility/random"
   "bst/utility/random/distribution"
   "flag"
   "os"
)

func main() {
   flag.Parse()
   console.Animate(prompt())
}

func prompt() console.Animation {
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
   return console.Choose[console.Animation]("Animation",
      &animations.ExteriorHeights{BinaryTreeAnimation: animation},
      &animations.InteriorHeights{BinaryTreeAnimation: animation},
      &animations.WeightsPerLevel{BinaryTreeAnimation: animation},
   )
}

func chooseOperation() list.Operation {
   return console.Choose[list.Operation]("Operation",
      &operations.Insert{},
      &operations.InsertPersistent{},
      &operations.InsertDelete{},
      &operations.InsertDeletePersistent{},
      &operations.InsertDeleteCycles{},
      &operations.InsertDeleteCyclesPersistent{},
      &operations.InsertDeleteSearch{},
      &operations.InsertDeleteSearchPersistent{},
      &operations.InsertDeleteSplitJoin{},
      &operations.InsertDeleteSplitJoinPersistent{},
   )
}

func chooseDistribution() random.Distribution {
   return console.Choose[random.Distribution]("Distribution",
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
   return console.Choose[list.List]("Strategy",
      &trees.WBSTTopDown{},
      &trees.WBSTBottomUp{},
      &trees.AVLTopDown{},
      &trees.AVLBottomUp{},
      &trees.AVLWeakTopDown{},
      &trees.AVLWeakBottomUp{},
      &trees.AVLRelaxedTopDown{},
      &trees.AVLRelaxedBottomUp{},
      &trees.RedBlackRelaxedBottomUp{},
      &trees.RedBlackRelaxedTopDown{},
      &trees.LBSTBottomUp{},
      &trees.LBSTTopDown{},
      &trees.LBSTRelaxed{},
      &trees.TreapTopDown{},
      &trees.TreapFingerTree{},
      &trees.Randomized{},
      &trees.Zip{},
      &trees.Splay{},
      &trees.Conc{},
   )
}
