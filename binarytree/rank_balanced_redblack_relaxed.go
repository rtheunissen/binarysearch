package binarytree

import . "binarysearch/abstract/list"

type RedBlackRelaxed struct {
   Tree
   RankBalanced
}

func (tree *RedBlackRelaxed) Delete(i Position) (x Data) {
   // assert(i < tree.Size())
   x = tree.Tree.delete(&tree.root, tree.size, i)
   tree.size--
   return x
}

func (tree RedBlackRelaxed) split(p *Node, i, s Size) (l, r *Node) {
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

func (tree RedBlackRelaxed) Split(i Position) (RedBlackRelaxed, RedBlackRelaxed) {
   // assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)
   return RedBlackRelaxed{Tree: Tree{arena: tree.arena, root: l, size: i}},
      RedBlackRelaxed{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree RedBlackRelaxed) build(l, p, r *Node, sl Size) *Node {
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

func (tree RedBlackRelaxed) join(l, r *Node, sl Size) (p *Node) {
   if l == nil {
      return r
   }
   if r == nil {
      return l
   }
   if tree.rank(l) < tree.rank(r) {
      return tree.build(l, tree.deleteMin(&r), r, sl)
   } else {
      return tree.build(l, tree.deleteMax(&l), r, sl-1)
   }
}

func (tree RedBlackRelaxed) Join(other RedBlackRelaxed) RedBlackRelaxed {
   tree.share(tree.root)
   tree.share(other.root)
   return RedBlackRelaxed{
      Tree: Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, other.root, tree.size),
         size:  tree.size + other.size,
      },
   }
}

func (tree RedBlackRelaxed) fixL(p *Node) *Node {
   if isZeroChild(p, p.l) {
      if isZeroChild(p, p.r) {
         if isZeroChild(p.l, p.l.l) || isZeroChild(p.l, p.l.r) {
            promote(p)
            return p
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

func (tree RedBlackRelaxed) fixR(p *Node) *Node {
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

func (tree RedBlackRelaxed) verifyRanks(p *Node) {
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

func (tree RedBlackRelaxed) Verify() {
   tree.Tree.Verify()
   tree.verifyRanks(tree.root)
}
