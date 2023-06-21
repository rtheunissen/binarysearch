package binarytree

import . "binarysearch/abstract/list"

type AVLRelaxedTopDown struct {
   AVLWeakTopDown
}

func (AVLRelaxedTopDown) New() List {
   return &AVLRelaxedTopDown{
      AVLWeakTopDown: *AVLWeakTopDown{}.New().(*AVLWeakTopDown),
   }
}

func (tree *AVLRelaxedTopDown) Clone() List {
   return &AVLRelaxedTopDown{
      AVLWeakTopDown: AVLWeakTopDown{
         WAVL: WAVL{
            Tree: tree.Tree.Clone(),
         },
      },
   }
}

func (tree AVLRelaxedTopDown) Verify() {
   tree.verifySizes()
   tree.verifyRanks(tree.root)
}

func (tree AVLRelaxedTopDown) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   invariant(tree.rank(p) >= p.height())
   invariant(tree.rank(p) > tree.rank(p.l))
   invariant(tree.rank(p) > tree.rank(p.r))

   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}

func (tree *AVLRelaxedTopDown) Delete(i Position) Data {
   // assert(i < tree.Size())
   x := tree.Tree.delete(&tree.root, tree.size, i)
   tree.size--
   return x
}

func (tree AVLRelaxedTopDown) buildL(l *Node, p *Node, r *Node, sl Size) (root *Node) {
   if tree.rank(l) - tree.rank(r) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(l) + 1)
      return p
   }
   tree.persist(&l)
   l.r = tree.build(l.r, p, r, sl-l.s-1)
   return tree.balanceInsertR(l)
}

func (tree AVLRelaxedTopDown) buildR(l *Node, p *Node, r *Node, sl Size) (root *Node) {
   if tree.rank(r) - tree.rank(l) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(r) + 1)
      return p
   }
   tree.persist(&r)
   r.s = 1 + sl + r.s
   r.l = tree.build(l, p, r.l, sl)
   return tree.balanceInsertL(r)
}

func (tree *AVLRelaxedTopDown) build(l, p, r *Node, sl Size) *Node {
   if tree.rank(l) <= tree.rank(r) {
      return tree.buildR(l, p, r, sl)
   } else {
      return tree.buildL(l, p, r, sl)
   }
}

func (tree AVLRelaxedTopDown) join(l, r *Node, sl Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) <= tree.rank(r) {
      return tree.build(l, tree.Tree.deleteMin(&r), r, sl)
   } else {
      return tree.build(l, tree.Tree.deleteMax(&l), r, sl-1)
   }
}

func (tree AVLRelaxedTopDown) Join(other List) List {
   l := tree
   r := other.(*AVLRelaxedTopDown)
   tree.share(l.root)
   tree.share(r.root)

   p := tree.join(l.root, r.root, l.size)

   return &AVLRelaxedTopDown{AVLWeakTopDown: AVLWeakTopDown{WAVL{Tree: Tree{arena: tree.arena, root: p, size: l.size + r.size}}}}
}


func (tree AVLRelaxedTopDown) split(p *Node, i, s Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.persist(&p)

   sl := p.s
   sr := s - p.s - 1

   if i <= (*p).s {
      l, r = tree.split(p.l, i, sl)
         r = tree.build(r, p, p.r, sl - i)
   } else {
      l, r = tree.split(p.r, i - sl - 1, sr)
         l = tree.build(p.l, p, l, sl)
   }
   return l, r
}

func (tree AVLRelaxedTopDown) Split(i Position) (List, List) {
   // assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)

   return &AVLRelaxedTopDown{AVLWeakTopDown{WAVL{Tree: Tree{arena: tree.arena, root: l, size: i}}}},
          &AVLRelaxedTopDown{AVLWeakTopDown{WAVL{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}}
}
