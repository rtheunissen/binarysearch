package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/utility"
)

type WAVL struct {
   Tree
   RankBalanced
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

func (tree WAVL) buildL(l *Node, p *Node, r *Node, sl Size) (root *Node) {
   assert(tree.rank(l) >= tree.rank(r))
   if tree.rank(l) - tree.rank(r) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(l) + 1)
      return p
   }
   tree.persist(&l)
   l.r = tree.buildL(l.r, p, r, sl-l.s-1)
   return tree.balanceInsertR(l)
}

func (tree WAVL) buildR(l *Node, p *Node, r *Node, sl Size) (root *Node) {
   assert(tree.rank(r) >= tree.rank(l))
   if tree.rank(r) <= tree.rank(l) + 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(r) + 1)
      return p
   }
   tree.persist(&r)
   r.s = 1 + sl + r.s
   r.l = tree.buildR(l, p, r.l, sl)
   return tree.balanceInsertL(r)
}

func (tree *WAVL) build(l, p, r *Node, sl Size) *Node {
   if tree.rank(l) <= tree.rank(r) {
      return tree.buildR(l, p, r, sl)
   } else {
      return tree.buildL(l, p, r, sl)
   }
}

// Rebalancing is only necessary when p.r is 0-child, which is when the ranks
// are equal, given that both p and p.r are non-nil.
// TODO: is this AVL?
func (tree WAVL) balanceInsertL(p *Node) *Node {
   if tree.isZeroChild(p, p.l) {
      if tree.isOneChild(p, p.r) {
         tree.promote(p)
      } else if tree.isTwoChild(p.l, p.l.r) {
         tree.rotateR(&p)
         tree.demote(p.r)
      } else {
         tree.rotateLR(&p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.r)
      }
   }
   return p
}

// Rebalancing is only necessary when p.r is 0-child, which is when the ranks
// are equal, given that both p and p.r are non-nil.
//
// TODO: this is AVL
//
func (tree WAVL) balanceInsertR(p *Node) *Node {
   if tree.isZeroChild(p, p.r) {
      if tree.isOneChild(p, p.l) {
         tree.promote(p)
      } else if tree.isTwoChild(p.r, p.r.l) {
         tree.rotateL(&p)
         tree.demote(p.l)
      } else {
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.r)
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
   invariant(tree.rank(root) <= 2*int(utility.Log2(size))) // TODO: tighten?
   // TODO height bound
}

// TODO split into L and R?
func (tree WAVL) rebalanceOnDelete(p *Node) *Node {
   if p.isLeaf() && tree.isTwoTwo(p) {
      tree.demote(p)
      return p
   }
   if tree.isThreeChild(p, p.r) {
      if tree.isTwoChild(p, p.l) {
         tree.demote(p)

      } else if tree.isTwoTwo(p.l) {
         tree.demote(p.l)
         tree.demote(p)

      } else if tree.isOneChild(p.l, p.l.l) {
         tree.rotateR(&p)
         tree.promote(p)
         tree.demote(p.r)

         assert(tree.isTwoChild(p, p.l))
         assert(tree.isOneChild(p, p.r))

         if p.r.l == nil {
            assert(tree.isTwoTwo(p.r))
            tree.demote(p.r)
         }
      } else {
         tree.rotateLR(&p)
         tree.promote(p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.r)
         tree.demote(p.r)

         assert(tree.isTwoChild(p, p.l))
         assert(tree.isTwoChild(p, p.r))
      }
   } else if tree.isThreeChild(p, p.l) {
      if tree.isTwoChild(p, p.r) {
         tree.demote(p)

      } else if tree.isTwoTwo(p.r) {
         tree.demote(p.r)
         tree.demote(p)

      } else if tree.isOneChild(p.r, p.r.r) {
         tree.rotateL(&p)
         tree.promote(p)
         tree.demote(p.l)

         assert(tree.isOneChild(p, p.l))
         assert(tree.isTwoChild(p, p.r))

         if p.l.r == nil {
            assert(tree.isTwoTwo(p.l))
            tree.demote(p.l)
         }
      } else {
         tree.rotateRL(&p)
         tree.promote(p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.l)
         tree.demote(p.r)

         assert(tree.isTwoChild(p, p.l))
         assert(tree.isTwoChild(p, p.r))
      }
   }
   return p
}

// rank/2 <= h <= rank
func (tree WAVL) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   if p.isLeaf() {
      invariant(tree.rank(p) == 0)
   }
   invariant(tree.rank(p) > tree.rank(p.l))
   invariant(tree.rank(p) > tree.rank(p.r))

   invariant(tree.isOneChild(p, p.l) || tree.isTwoChild(p, p.l))
   invariant(tree.isOneChild(p, p.r) || tree.isTwoChild(p, p.r))

   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}
