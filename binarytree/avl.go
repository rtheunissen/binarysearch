package binarytree

import (
   "binarysearch/abstract/list"
   "math"
)

type AVLBottomUp struct {
   Tree
   RankBalanced
}

func (AVLBottomUp) New() list.List {
   return &AVLBottomUp{}
}

func (tree *AVLBottomUp) Clone() list.List {
   return &AVLBottomUp{
      Tree: tree.Tree.Clone(),
   }
}

func (tree *AVLBottomUp) Verify() {
   tree.verifySizes(tree.root, tree.size)
   tree.verifyHeight(tree.root, tree.size)
   tree.verifyRanks(tree.root)
}

func (tree *AVLBottomUp) verifySizes(p *Node, s list.Size) list.Size {
   if p == nil {
      return 0
   }
   sl := tree.verifySizes(p.l, p.sizeL())
   sr := tree.verifySizes(p.r, p.sizeR(s))

   invariant(s == sl + sr + 1)
   return s
}

func (tree *AVLBottomUp) verifyHeight(p *Node, s list.Size) {
   if p == nil {
      return
   }
   invariant(tree.rank(p) == tree.root.height())
   invariant(tree.rank(p) <= int(1.44043 * math.Log2(float64(s))))
}

func (tree *AVLBottomUp) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   //
   // AVL rule: Every node is 1,1 or 1,2.
   //
   invariant(tree.isOneChild(p, p.l) || tree.isTwoChild(p, p.l))
   invariant(tree.isOneChild(p, p.r) || tree.isTwoChild(p, p.r))
   invariant(tree.isOneChild(p, p.l) || tree.isOneChild(p, p.r))
   invariant(tree.isOneChild(p, p.r) || tree.isOneChild(p, p.l))

   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}

func (tree *AVLBottomUp) Insert(i list.Position, x list.Data) {
   tree.root = tree.insert(tree.root, i, x)
   tree.size = tree.size + 1
}

func (tree *AVLBottomUp) insert(p *Node, i list.Position, x list.Data) *Node {
   if p == nil {
      return tree.allocate(Node{x: x})
   }
   tree.copy(&p)
   if i <= p.s {
      p.s = p.s + 1
      p.l = tree.insert(p.l, i, x)
      return tree.balanceInsertL(p)
   } else {
      p.r = tree.insert(p.r, i - p.s - 1, x)
      return tree.balanceInsertR(p)
   }
}

func (tree *AVLBottomUp) balanceInsertL(p *Node) *Node {
   //
   //
   //
   if tree.isZeroChild(p, p.l) {
      //
      //
      //
      if tree.isOneChild(p, p.r) {
         tree.promote(p)
         return p
      }
      //
      //
      //
      if tree.isTwoChild(p.l, p.l.r) {
         tree.rotateR(&p)
         tree.demote(p.r)
         return p
      }
      //
      //
      //
      tree.rotateLR(&p)
      tree.promote(p)
      tree.demote(p.l)
      tree.demote(p.r)
      return p
   }
   return p
}

func (tree *AVLBottomUp) balanceInsertR(p *Node) *Node {
   //
   //
   //
   if tree.isZeroChild(p, p.r) {
      //
      //
      //
      if tree.isOneChild(p, p.l) {
         tree.promote(p)
         return p
      }
      //
      //
      //
      if tree.isTwoChild(p.r, p.r.l) {
         tree.rotateL(&p)
         tree.demote(p.l)
         return p
      }
      //
      //
      //
      tree.rotateRL(&p)
      tree.promote(p)
      tree.demote(p.l)
      tree.demote(p.r)
      return p
   }
   return p
}

func (tree *AVLBottomUp) delete(p *Node, i list.Position, x *list.Data) *Node {
   tree.copy(&p)
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

func (tree *AVLBottomUp) Delete(i list.Position) (x list.Data) {
   // assert(i < tree.size)
   tree.root = tree.delete(tree.root, i, &x)
   tree.size--
   return
}

func (tree *AVLBottomUp) balanceDeleteR(p *Node) *Node {
   //
   //
   //
   if tree.isTwoTwo(p) {
      tree.demote(p)
      return p
   }
   if tree.isThreeChild(p, p.r) {
      //
      //
      //
      if tree.isTwoChild(p.l, p.l.l) {
         tree.rotateLR(&p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.r)
         tree.demote(p.r)
      } else {
         //
         //
         //
         if tree.isTwoChild(p.l, p.l.r) {
            tree.rotateR(&p)
            tree.demote(p.r)
            tree.demote(p.r)
         } else {
            tree.rotateR(&p)
            tree.promote(p)
            tree.demote(p.r)
         }
      }
   }
   return p
}

func (tree *AVLBottomUp) balanceDeleteL(p *Node) *Node {
   //
   //
   //
   if tree.isTwoTwo(p) {
      tree.demote(p)
      return p
   }
   //
   //
   //
   if tree.isThreeChild(p, p.l) {
      //
      //
      //
      if tree.isTwoChild(p.r, p.r.r) {
         //
         //
         //
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.r)
         tree.demote(p.l)
         tree.demote(p.l)
      } else {
         //
         //
         //
         if tree.isTwoChild(p.r, p.r.l) {
            //
            //
            //
            tree.rotateL(&p)
            tree.demote(p.l)
            tree.demote(p.l)
         } else {
            //
            //
            //
            tree.rotateL(&p)
            tree.promote(p)
            tree.demote(p.l)
         }
      }
   }
   return p
}


func (tree *AVLBottomUp) deleteMin(p *Node, min **Node) *Node {
   tree.copy(&p)
   if p.l == nil {
      *min = p
      return p.r
   }
   p.s = p.s - 1
   p.l = tree.deleteMin(p.l, min)
   return tree.balanceDeleteL(p)
}

func (tree *AVLBottomUp) deleteMax(p *Node, r **Node) *Node {
   tree.copy(&p)
   if p.r == nil {
      *r = p
      return p.l
   }
   p.r = tree.deleteMax(p.r, r)
   return tree.balanceDeleteR(p)
}

func (tree *AVLBottomUp) build(l, p, r *Node, sl list.Size) *Node {
   if tree.rank(l) < tree.rank(r) {
      return tree.buildL(l, p, r, sl)
   } else {
      return tree.buildR(l, p, r, sl)
   }
}

func (tree *AVLBottomUp) buildR(l, p, r *Node, sl list.Size) *Node {
   if tree.rank(l) - tree.rank(r) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(l) + 1)
      return p
   }
   tree.copy(&l)
   l.r = tree.buildR(l.r, p, r, sl-l.s-1)
   return tree.balanceInsertR(l)
}

func (tree *AVLBottomUp) buildL(l, p, r *Node, sl list.Size) *Node {
   if tree.rank(r) - tree.rank(l) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(r) + 1)
      return p
   }
   tree.copy(&r)
   r.s = 1 + sl + r.s
   r.l = tree.buildL(l, p, r.l, sl)
   return tree.balanceInsertL(r)
}

func (tree *AVLBottomUp) joinR(l, r *Node, sl list.Size) (p *Node) {
   if tree.rank(l) <= tree.rank(r) + 1 {
      return tree.build(tree.deleteMax(l, &p), p, r, sl-1)
   }
   tree.copy(&l)
   l.r = tree.joinR(l.r, r, sl-l.s-1)
   return tree.balanceInsertR(l)
}

func (tree *AVLBottomUp) joinL(l, r *Node, sl list.Size) (p *Node) {
   if tree.rank(r) <= tree.rank(l) + 1 {
      return tree.build(l, p, tree.deleteMin(r, &p), sl)
   }
   tree.copy(&r)
   r.s = sl + r.s
   r.l = tree.joinL(l, r.l, sl)
   return tree.balanceInsertL(r)
}

func (tree *AVLBottomUp) join(l, r *Node, sl list.Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) < tree.rank(r) {
      return tree.joinL(l, r, sl)
   } else {
      return tree.joinR(l, r, sl)
   }
}

func (tree *AVLBottomUp) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*AVLBottomUp).root)
   return &AVLBottomUp{
      Tree: Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, other.(*AVLBottomUp).root, tree.size),
         size:  tree.size + other.(*AVLBottomUp).size,
      },
   }
}

func (tree *AVLBottomUp) split(p *Node, i, s list.Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.copy(&p)
   if i <= (*p).s {
      l, r = tree.split(p.l, i, p.s)
         r = tree.build(r, p, p.r, p.s - i)
   } else {
      l, r = tree.split(p.r, i - p.s - 1, s - p.s - 1)
         l = tree.build(p.l, p, l, p.s)
   }
   return l, r
}

func (tree *AVLBottomUp) Split(i list.Position) (list.List, list.List) {
   // assert(i <= tree.size)
   tree.share(tree.root)

   l, r := tree.split(tree.root, i, tree.size)

   return &AVLBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}},
          &AVLBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}
