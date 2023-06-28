package binarytree

import (
   "binarysearch/abstract/list"
   "binarysearch/distribution"
   "testing"
)

// TODO: can the tests use operations?
// TODO: do we need list tests really?
// TODO: can we not make them binarytree?
func TestBST(t *testing.T) {
   list.TestSuite{
      Scale: 100,
      Tests: []list.Test{
         list.TestNew,
         list.TestSelect,
         list.TestSelectAfterInsert,
         //list.TestSelectAfterInsertPersistent,
         list.TestUpdate,
         //list.TestUpdatePersistent,
         list.TestInsert,
         //list.TestInsertPersistent,
         list.TestDelete,
         //list.TestDeletePersistent,
         list.TestInsertDelete,
         //list.TestInsertDeletePersistent,
         list.TestSplit,
         list.TestJoin,
         list.TestJoinFromSplit,
         list.TestJoinAfterInsertDelete,
      },
      Distributions: []distribution.Distribution{
         &distribution.Uniform{},
         &distribution.Normal{},
         &distribution.Skewed{},
         &distribution.Zipf{},
         &distribution.Maximum{},
         //&distribution.Minimum{},
         //&distribution.BiModal{},
         //&distribution.Ascending{},
         //&distribution.Descending{},
         //&distribution.Parabolic{},
         //&distribution.Queue{},
         //&distribution.UShape{},
         //&distribution.Slope{},
      },
      Lists: []list.List{
         &AVLBottomUp{},
         &AVLWeakTopDown{},
         &AVLWeakBottomUp{},
         &AVLRelaxedTopDown{},
         &AVLRelaxedBottomUp{},
         &RedBlackBottomUp{},
         &RedBlackTopDown{},
         &RedBlackRelaxedBottomUp{},
         &RedBlackRelaxedTopDown{},
         //&LBSTBottomUp{},
         //&LBSTTopDown{},
         //&LBSTRelaxed{},
         //&TreapTopDown{},
         //&TreapFingerTree{},
         //&Randomized{},
         //&Zip{},
         //&Splay{},
         //&Conc{},
      },
   }.Run(t)
}
