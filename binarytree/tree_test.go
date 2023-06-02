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
   		list.TestSelectAfterInsertPersistent,
   		list.TestUpdate,
   		list.TestUpdatePersistent,
   		list.TestInsert,
   		list.TestInsertPersistent,
   		list.TestDelete,
   		list.TestDeletePersistent,
   		list.TestInsertDelete,
   		list.TestInsertDeletePersistent,
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
   	},
   	Lists: []list.List{
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
   	},
   }.Run(t)
}
