package binarytree

import (
   "testing"
   "trees/abstract/list"
   "trees/distribution"
   "trees/tests"
)

// TODO: can the tests use operations?
// TODO: do we need list tests really?
// TODO: can we not make them binarytree?
func TestBST(t *testing.T) {
	tests.TestSuite{
		Scale: 100,
		Tests: []tests.Test{
			tests.TestNew,
			tests.TestSelect,
			tests.TestSelectAfterInsert,
			tests.TestSelectAfterInsertPersistent,
			tests.TestUpdate,
			tests.TestUpdatePersistent,
			tests.TestInsert,
			tests.TestInsertPersistent,
			tests.TestDelete,
			tests.TestDeletePersistent,
			tests.TestInsertDelete,
			tests.TestInsertDeletePersistent,
			tests.TestSplit,
			tests.TestJoin,
			tests.TestJoinFromSplit,
			tests.TestJoinAfterInsertDelete,
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
