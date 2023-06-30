package binarytree

type AVLRelaxed struct {
   RankBalanced
}

//func (tree RedBlackRelaxed) verifyHeight(p *Node) {
//   invariant(p.height() <= 2 * tree.rank(p) + 1)
//}

func (tree AVLRelaxed) verifyRanks(p *Node) {
  if p == nil {
     return
  }
  invariant(tree.rank(p) >= p.height())
  invariant(tree.rank(p) > tree.rank(p.l))
  invariant(tree.rank(p) > tree.rank(p.r))

  tree.verifyRanks(p.l)
  tree.verifyRanks(p.r)
}
