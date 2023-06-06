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

import . "binarysearch/abstract/list"

type RedBlackBottomUp struct {
   Tree
}

func (RedBlackBottomUp) New() List {
   return &RedBlackBottomUp{}
}

func (tree *RedBlackBottomUp) Clone() List {
   return &RedBlackBottomUp{
      Tree: tree.Tree.Clone(),
   }
}

func (tree *RedBlackBottomUp) Select(i Size) Data {
   assert(i < tree.Size())
   return tree.lookup(tree.root, i)
}

func (tree *RedBlackBottomUp) Update(i Size, x Data) {
   assert(i < tree.Size())
   tree.copy(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *RedBlackBottomUp) rank(p *Node) int {
   return p.rank()
}

func (tree *RedBlackBottomUp) Delete(i Position) (x Data) {
   assert(i < tree.size)
   tree.size = tree.size - 1
   tree.root = tree.delete(tree.root, i, &x)
   return x
}

func (tree *RedBlackBottomUp) delete(p *Node, i Position, x *Data) *Node {
   tree.copy(&p)
   if i == p.s {
      *x = p.x
      defer tree.release(p)
      return tree.join(p.l, p.r, p.s)
   }
   if i < p.s {
      p.s = p.s - 1
      p.l = tree.delete(p.l, i, x)
      return tree.fixL(p)
   } else {
      p.r = tree.delete(p.r, i-p.s-1, x)
      return tree.fixR(p)
   }
}

func (tree *RedBlackBottomUp) insert(p *Node, i Position, x Data) *Node {
   if p == nil {
      return tree.allocate(Node{x: x})
   }
   tree.copy(&p)
   if i <= p.s {
      p.s = p.s + 1
      p.l = tree.insert(p.l, i, x)
      return tree.fixL(p)
   } else {
      p.r = tree.insert(p.r, i-p.s-1, x)
      return tree.fixR(p)
   }
}

func (tree *RedBlackBottomUp) Insert(i Position, x Data) {
   assert(i <= tree.size)
   tree.size = tree.size + 1
   tree.root = tree.insert(tree.root, i, x)
   return
}

func (tree RedBlackBottomUp) split(p *Node, i, s Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.copy(&p)

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

func (tree RedBlackBottomUp) Split(i Position) (List, List) {
   assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)
   return &RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}},
          &RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree RedBlackBottomUp) build(l, p, r *Node, sl Size) *Node {
   if tree.rank(l) == tree.rank(r) {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(p.l.rank() + 1)
      return p
   }
   if tree.rank(l) < tree.rank(r) {
      tree.copy(&r)
      r.s = 1 + sl + r.s
      r.l = tree.build(l, p, r.l, sl)
      return tree.fixL(r)
   } else {
      tree.copy(&l)
      l.r = tree.build(l.r, p, r, sl-l.s-1)
      return tree.fixR(l)
   }
}

func (tree RedBlackBottomUp) join(l, r *Node, sl Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) < tree.rank(r) {
      return tree.build(l, tree.deleteMin(&r), r, sl)
   } else {
      return tree.build(l, tree.deleteMax(&l), r, sl-1)
   }
}

func (tree RedBlackBottomUp) Join(other List) List {
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

func (tree RedBlackBottomUp) fixL(p *Node) *Node {
   if isZeroChild(p, p.l) {
      if isZeroChild(p, p.r) {
         if isZeroChild(p.l, p.l.l) || isZeroChild(p.l, p.l.r) {
            promote(p)
         }
      } else {
         if isZeroChild(p.l, p.l.l) {
            return p.rotateR()
         }
         if isZeroChild(p.l, p.l.r) {
            return p.rotateLR()
         }
      }
   }
   return p
}

func (tree RedBlackBottomUp) fixR(p *Node) *Node {
   if isZeroChild(p, p.r) {
      if isZeroChild(p, p.l) {
         if isZeroChild(p.r, p.r.r) || isZeroChild(p.r, p.r.l) {
            promote(p)
         }
      } else {
         if isZeroChild(p.r, p.r.r) {
            return p.rotateL()
         }
         if isZeroChild(p.r, p.r.l) {
            return p.rotateRL()
         }
      }
   }
   return p
}

func (tree RedBlackBottomUp) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   invariant(tree.rank(p) >= tree.rank(p.l))
   invariant(tree.rank(p) >= tree.rank(p.r))

   // No zero-child has a zero-child as a parent
   if isZeroChild(p, p.l) {
      invariant(!isZeroChild(p.l, p.l.l))
      invariant(!isZeroChild(p.l, p.l.r))
   }
   if isZeroChild(p, p.r) {
      invariant(!isZeroChild(p.r, p.r.l))
      invariant(!isZeroChild(p.r, p.r.r))
   }
   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}

func (tree RedBlackBottomUp) Verify() {
   tree.Tree.Verify()
   tree.verifyRanks(tree.root)
}
