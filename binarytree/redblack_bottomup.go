//
//In a ranked binary tree obeying the red-black rule,
//the 0-children are the red nodes, the 1-children are the black nodes. All missing nodes have rank
//difference 1 and are black. The rank of a node is the number of black nodes on a path from the node
//to a leaf, not counting the node itself: this number is independent of the path. Some authors require
//that the root of a red-black tree be black, others allow it to be either red or black. In our formulation,
//the root has no rank difference, and hence no color. Since all rank differences are 0 or 1, we can
//store the balance information in one bit per node, indicating whether its rank difference is zero (it is
//red) or one (it is black).
//
//[Guibas and Sedgewick 1978],
//
//Red-Black Rule: All rank differences are 0 or 1, and no parent of a 0-child is a 0-child

package binarytree

import (
   "binarysearch/abstract/list"
   "math"
)

type RedBlackBottomUp struct {
   Tree
   RankBalanced
}

func (RedBlackBottomUp) New() list.List {
   return &RedBlackBottomUp{}
}

func (tree *RedBlackBottomUp) Clone() list.List {
   return &RedBlackBottomUp{
      Tree: tree.Tree.Clone(),
   }
}

func (tree RedBlackBottomUp) Verify() {
   tree.verifySize(tree.root, tree.size)
   tree.verifyRanks(tree.root)
   tree.verifyHeight(tree.root, tree.size)
}

func (tree RedBlackBottomUp) verifyHeight(p *Node, s list.Size) {
   invariant(p.height() <= 2 * tree.rank(tree.root) + 1)
   invariant(p.height() <= 2 * int(math.Floor(math.Log2(float64(s)))))
}

func (tree RedBlackBottomUp) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   if p.isLeaf() {
      invariant(tree.rank(p) == 0)
   }
   invariant(tree.rank(p) >= tree.rank(p.l))
   invariant(tree.rank(p) >= tree.rank(p.r))

   // All rank differences are 0 or 1.
   invariant(tree.isZeroChild(p, p.l) || tree.isOneChild(p, p.l))
   invariant(tree.isZeroChild(p, p.r) || tree.isOneChild(p, p.r))

   // No parent of a 0-child is a 0-child.
   invariant(tree.isOneChild(p, p.l) || tree.isOneChild(p.l, p.l.l))
   invariant(tree.isOneChild(p, p.l) || tree.isOneChild(p.l, p.l.r))
   invariant(tree.isOneChild(p, p.r) || tree.isOneChild(p.r, p.r.l))
   invariant(tree.isOneChild(p, p.r) || tree.isOneChild(p.r, p.r.r))

   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}

func (tree *RedBlackBottomUp) Delete(i list.Position) (x list.Data) {
   assert(i < tree.size)
   tree.size = tree.size - 1
   tree.root = tree.delete(tree.root, i, &x)
   return x
}

func (tree *RedBlackBottomUp) delete(p *Node, i list.Position, x *list.Data) *Node {
   tree.persist(&p)
   if i == p.s {
      *x = p.x
      defer tree.free(p)
      return tree.join(p.l, p.r, p.s)
   }
   if i < p.s {
      p.s = p.s - 1
      p.l = tree.delete(p.l, i, x)
      return tree.balanceDeleteL(p)
   } else {
      p.r = tree.delete(p.r, i-p.s-1, x)
      return tree.balanceDeleteR(p)
   }
}


func (tree RedBlackBottomUp) balanceDeleteR(p *Node) *Node {
   //
   //
   //
   if tree.isZeroChild(p, p.r) {
      assert(tree.isOneChild(p.r, p.r.r))
      assert(tree.isOneChild(p.r, p.r.l))
      return p
   }
   //
   //
   //
   if tree.isOneChild(p, p.r) {
      return p
   }
   //
   //
   //
   if tree.isZeroChild(p, p.l) {
      assert(tree.isZeroChild(p, p.l))
      assert(tree.isTwoChild(p, p.r))
      assert(tree.isOneOne(p.l))
      //
      //
      //
      tree.rotateR(&p)
      assert(tree.isZeroChild(p, p.r))
      assert(tree.isOneChild(p, p.l))
      //
      //
      //
      if tree.isZeroChild(p.r.l, p.r.l.l) {
         assert(tree.isOneChild(p.r, p.r.l))
         assert(tree.isTwoChild(p.r, p.r.r))
         tree.rotateR(&p.r)
         tree.promote(p.r)
         tree.demote(p.r.r)
         return p
      }
      //
      //
      //
      if tree.isZeroChild(p.r.l, p.r.l.r) {
         assert(tree.isOneChild(p.r, p.r.l))
         assert(tree.isTwoChild(p.r, p.r.r))
         tree.rotateLR(&p.r)
         tree.promote(p.r)
         tree.demote(p.r.r)
         return p
      }
      //
      //
      //
      assert(tree.isOneChild(p.r, p.r.l))
      assert(tree.isTwoChild(p.r, p.r.r))
      assert(tree.isOneChild(p.r.l, p.r.l.l))
      assert(tree.isOneChild(p.r.l, p.r.l.r))
      tree.demote(p.r)
      return p

   } else {
      //
      //
      //
      assert(tree.isOneChild(p, p.l))
      assert(tree.isTwoChild(p, p.r))
      //
      //
      //
      if tree.isZeroChild(p.l, p.l.l) {
         tree.rotateR(&p)
         tree.promote(p)
         tree.demote(p.r)
         return p
      }
      //
      //
      //
      if tree.isZeroChild(p.l, p.l.r) {
         tree.rotateLR(&p)
         tree.promote(p)
         tree.demote(p.r)
         return p
      }
      //
      //
      //
      assert(tree.isOneOne(p.l))
      tree.demote(p)
      return p
   }
}

func (tree RedBlackBottomUp) balanceDeleteL(p *Node) *Node {
   if tree.isZeroChild(p, p.l) {
      if tree.isOneOne(p.l.r) && tree.isOneOne(p.l.l) {
         return p
      }
   }
   if tree.isZeroChild(p, p.l) || tree.isOneChild(p, p.l) {
      return p
   }
   if tree.isOneChild(p, p.r) {
      if tree.isOneChild(p.r, p.r.l) && tree.isOneChild(p.r, p.r.r) {
         tree.demote(p)
         return p
      }
      if tree.isZeroChild(p.r, p.r.r) {
         tree.rotateL(&p)
         tree.promote(p)
         tree.demote(p.l)
         return p
      } else {
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.l)
         return p
      }
   } else {
      tree.rotateL(&p)
      if tree.isOneChild(p.l.r, p.l.r.r) && tree.isOneChild(p.l.r, p.l.r.l) {
         tree.demote(p.l)
         return p
      }
      if tree.isZeroChild(p.l.r, p.l.r.r) {
         tree.rotateL(&p.l)
         tree.promote(p.l)
         tree.demote(p.l.l)
         return p
      }
      tree.rotateRL(&p.l)
      tree.promote(p.l)
      tree.demote(p.l.l)
      return p
   }
}

func (tree *RedBlackBottomUp) Insert(i list.Position, x list.Data) {
   assert(i <= tree.size)
   tree.size = tree.size + 1
   tree.root = tree.insert(tree.root, i, x)
   return
}

func (tree *RedBlackBottomUp) insert(p *Node, i list.Position, x list.Data) *Node {
   if p == nil {
      return tree.allocate(Node{x: x})
   }
   tree.persist(&p)
   if i <= p.s {
      p.s = p.s + 1
      p.l = tree.insert(p.l, i, x)
      return tree.balanceInsertL(p)
   } else {
      p.r = tree.insert(p.r, i-p.s-1, x)
      return tree.balanceInsertR(p)
   }
}

func (tree RedBlackBottomUp) balanceInsertL(p *Node) *Node {
   if tree.isZeroChild(p, p.l) {
      if tree.isZeroChild(p, p.r) {
         if tree.isZeroChild(p.l, p.l.l) || tree.isZeroChild(p.l, p.l.r) {
            tree.promote(p)
            return p
         }
      } else {
         if tree.isZeroChild(p.l, p.l.l) {
            tree.rotateR(&p)
            return p
         }
         if tree.isZeroChild(p.l, p.l.r) {
            tree.rotateLR(&p)
            return p
         }
      }
   }
   return p
}

func (tree RedBlackBottomUp) balanceInsertR(p *Node) *Node {
   if tree.isZeroChild(p, p.r) {
      if tree.isZeroChild(p, p.l) {
         if tree.isZeroChild(p.r, p.r.r) || tree.isZeroChild(p.r, p.r.l) {
            tree.promote(p)
            return p
         }
      } else {
         if tree.isZeroChild(p.r, p.r.r) {
            tree.rotateL(&p)
            return p
         }
         if tree.isZeroChild(p.r, p.r.l) {
            tree.rotateRL(&p)
            return p
         }
      }
   }
   return p
}


func (tree RedBlackBottomUp) split(p *Node, i, s list.Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.persist(&p)

   sl := p.s
   sr := s - p.s - 1

   if i <= (*p).s {
      l, r = tree.split(p.l, i, sl)
      r = tree.build(r, p, p.r, sl-i)
   } else {
      l, r = tree.split(p.r, i-sl-1, sr)
      l = tree.build(p.l, p, l, sl)
   }
   return l, r
}

func (tree RedBlackBottomUp) Split(i list.Position) (list.List, list.List) {
   assert(i <= tree.size)
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)
   return &RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}},
      &RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree *RedBlackBottomUp) deleteMin(p *Node, min **Node) *Node {
   tree.persist(&p)
   if p.l == nil {
      *min = p
      return p.r
   }
   p.s = p.s - 1
   p.l = tree.deleteMin(p.l, min)
   return tree.balanceDeleteL(p)
}


func (tree *RedBlackBottomUp) deleteMax(p *Node, max **Node) *Node {
   tree.persist(&p)
   if p.r == nil {
      *max = p
      return p.l
   }
   p.r = tree.deleteMax(p.r, max)
   return tree.balanceDeleteR(p)
}

// TODO: we can do if rank is within 1?
func (tree RedBlackBottomUp) build(l, p, r *Node, sl list.Size) *Node {
   if tree.rank(l) == tree.rank(r) {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(p.l))
      //
      //
      //
      if (l == nil || tree.isZeroChild(l, l.l) || tree.isZeroChild(l, l.r))||
         (r == nil || tree.isZeroChild(r, r.l) || tree.isZeroChild(r, r.r)) {
         tree.promote(p)
      }
      return p
   }
   if tree.rank(l) < tree.rank(r) {
      tree.persist(&r)
      r.s = 1 + sl + r.s
      r.l = tree.build(l, p, r.l, sl)
      return tree.balanceInsertL(r)
   } else {
      tree.persist(&l)
      l.r = tree.build(l.r, p, r, sl-l.s-1)
      return tree.balanceInsertR(l)
   }
}

func (tree RedBlackBottomUp) join(l, r *Node, sl list.Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) < tree.rank(r) {
      return tree.build(l, p, tree.deleteMin(r, &p), sl)
   } else {
      return tree.build(tree.deleteMax(l, &p), p, r, sl-1)
   }
}

func (tree RedBlackBottomUp) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*RedBlackBottomUp).root)
   return &RedBlackBottomUp{
      Tree: Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, other.(*RedBlackBottomUp).root, tree.size),
         size:  tree.size + other.(*RedBlackBottomUp).size,
      },
   }
}
