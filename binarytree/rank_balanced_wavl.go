package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/utility"
)

type WAVL struct {
   Tree
   RankBalanced
}

func (tree WAVL) calculateRank(p *Node) {
   p.y = uint64(utility.Max(p.l.rank(), p.r.rank())) + 1
}

//func (tree WAVL) assignRanksFromScratch(p *Node) {
//   if p == nil {
//      return
//   }
//   tree.assignRanksFromScratch(p.l)
//   tree.assignRanksFromScratch(p.r)
//   tree.calculateRank(p)
//}

func (tree WAVL) Verify() {
   tree.verifySizes()
   tree.verifyRanks(tree.root)
   tree.verifyHeight(tree.root, tree.size)
}

func (tree WAVL) joinL(l *Node, p *Node, r *Node, sl Size) (root *Node) {
   if tree.rank(l)-tree.rank(r) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      tree.calculateRank(p)
      return p
   }
   tree.copy(&l)
   l.r = tree.joinL(l.r, p, r, sl-l.s-1)
   return tree.rebalanceBottomUpAfterInsertingRight(l)
}

func (tree WAVL) joinR(l *Node, p *Node, r *Node, sl Size) (root *Node) {
   if tree.rank(r)-tree.rank(l) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      tree.calculateRank(p)
      return p
   }
   tree.copy(&r)
   r.s = 1 + sl + r.s
   r.l = tree.joinR(l, p, r.l, sl)
   return tree.rebalanceBottomUpAfterInsertingLeft(r)
}

func (tree WAVL) build(l, p, r *Node, sl Size) *Node {
   if tree.rank(l) <= tree.rank(r) {
      return tree.joinR(l, p, r, sl)
   } else {
      return tree.joinL(l, p, r, sl)
   }
}

// Rebalancing is only necessary when p.r is 0-child, which is when the ranks
// are equal, given that both p and p.r are non-nil.
func (tree WAVL) rebalanceBottomUpAfterInsertingLeft(p *Node) *Node {
   if isZeroChild(p, p.l) {
      if isOneChild(p, p.r) {
         promote(p)
      } else if isTwoChild(p.l, p.l.r) { // SINGLE ROTATION
         tree.rotateR(&p)
         demote(p.r)
      } else { // DOUBLE ROTATION
         tree.rotateLR(&p)
         promote(p)
         demote(p.l)
         demote(p.r)
      }
   }
   return p
}

// Rebalancing is only necessary when p.r is 0-child, which is when the ranks
// are equal, given that both p and p.r are non-nil.
func (tree WAVL) rebalanceBottomUpAfterInsertingRight(p *Node) *Node {
   if isZeroChild(p, p.r) {
      if isOneChild(p, p.l) {
         promote(p)
      } else if isTwoChild(p.r, p.r.l) { // SINGLE ROTATION
         tree.rotateL(&p)
         demote(p.l)
      } else { // DOUBLE ROTATION
         tree.rotateRL(&p)
         promote(p)
         demote(p.l)
         demote(p.r)
      }
   }
   return p
}

//func (tree WAVL) verify(root *Node, size Size) {
//   tree.verifyHeight(root, size)
//   tree.verifyRanks(root)
//}

func (tree WAVL) verifyHeight(root *Node, size Size) {
   if root == nil {
      return
   }
   height := root.height()

   invariant(tree.rank(root) >= height || height == 0)
   invariant(tree.rank(root) <= 2*height)
   invariant(tree.rank(root) <= 2*int(utility.Log2(size)))
}

func (tree WAVL) rebalanceOnDelete(p *Node) *Node {
   if p.isLeaf() && isTwoTwo(p) {
      demote(p)
      return p
   }
   if isThreeChild(p, p.r) {
      if isTwoChild(p, p.l) {
         demote(p)

      } else if isTwoTwo(p.l) {
         demote(p.l)
         demote(p)

      } else if isOneChild(p.l, p.l.l) {
         tree.rotateR(&p)
         promote(p)
         demote(p.r)

         // assert(isTwoChild(p, p.l))
         // assert(isOneChild(p, p.r))

         if p.r.l == nil {
            // assert(isTwoTwo(p.r))
            demote(p.r)
         }
      } else {
         tree.rotateLR(&p)
         promote(p)
         promote(p)
         demote(p.l)
         demote(p.r)
         demote(p.r)

         // assert(isTwoChild(p, p.l))
         // assert(isTwoChild(p, p.r))
      }
   } else if isThreeChild(p, p.l) {
      if isTwoChild(p, p.r) {
         demote(p)

      } else if isTwoTwo(p.r) {
         demote(p.r)
         demote(p)

      } else if isOneChild(p.r, p.r.r) {
         tree.rotateL(&p)
         promote(p)
         demote(p.l)

         // assert(isOneChild(p, p.l))
         // assert(isTwoChild(p, p.r))

         if p.l.r == nil {
            // assert(isTwoTwo(p.l))
            demote(p.l)
         }
      } else {
         tree.rotateRL(&p)
         promote(p)
         promote(p)
         demote(p.l)
         demote(p.l)
         demote(p.r)

         // assert(isTwoChild(p, p.l))
         // assert(isTwoChild(p, p.r))
      }
   }
   return p
}

func (tree WAVL) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   if p.isLeaf() {
      invariant(tree.rank(p) == 0)
   }
   invariant(tree.rank(p) > tree.rank(p.l))
   invariant(tree.rank(p) > tree.rank(p.r))

   invariant(isOneChild(p, p.l) || isTwoChild(p, p.l))
   invariant(isOneChild(p, p.r) || isTwoChild(p, p.r))

   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}
