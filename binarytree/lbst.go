package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/utility"
)

type LBST struct {
   Tree
}

func (LBST) isBalanced(x, y Size) bool {
   return !utility.SmallerMSB(x, y >> 1)
}

func (LBST) singleRotation(x, y Size) bool {
   return !utility.SmallerMSB(x, y)
}

func (tree LBST) join2(l *Node, r *Node, sl, sr Size) (k *Node) {
   if l == nil {
      return r
   }
   if r == nil {
      return l
   }
   if sl <= sr {
      r = tree.extractMin(r, sr, &k)
      return tree.join3(l, k, r, sl, sr-1)
   } else {
      l = tree.extractMax(l, sl, &k)
      return tree.join3(l, k, r, sl-1, sr)
   }
}

func (tree LBST) extractMin(p *Node, s Size, x **Node) *Node {
   tree.copy(&p)
   if p.l == nil {
      *x = p
      p = p.r
      return p
   }
   sl := p.s
   sr := s - p.s - 1

   p.l = tree.extractMin(p.l, p.s, x)
   p.s--

   if !tree.isBalanced(sl-1, sr) {
      srl := (*p).r.s
      srr := sr - (*p).r.s - 1
      //
      if tree.singleRotation(srr, srl) {
         tree.rotateL(&p)
      } else {
         tree.rotateRL(&p)
      }
   }
   return p
}

func (tree LBST) extractMax(p *Node, s Size, x **Node) *Node {
   tree.copy(&p)
   if p.r == nil {
      *x = p
      p = p.l
      return p
   }
   sl := p.s
   sr := s - p.s - 1

   p.r = tree.extractMax(p.r, sr, x)
   if !tree.isBalanced(sr-1, sl) {
      if tree.singleRotation((*p).l.s, sl-(*p).l.s-1) {
         tree.rotateR(&p)
      } else {
         tree.rotateLR(&p)
      }
   }
   return p
}

func (tree LBST) Join(that LBST) LBST {
   l := tree
   r := that
   tree.share(l.root)
   tree.share(r.root)
   return LBST{
      Tree{
         arena: tree.arena,
         root:  tree.join2(l.root, r.root, l.size, r.size),
         size:  l.size + r.size,
      },
   }
}

func (tree LBST) join3(l, k, r *Node, sl, sr Size) *Node {
   //tree.pathcopy(&k) // optional?
   if sl <= sr {
      return tree.assembleRL(k, l, r, sl, sr)
   } else {
      return tree.assembleLR(k, l, r, sl, sr)
   }
}

func (tree *LBST) assembleLR(p *Node, l, r *Node, sl, sr Size) *Node {
   if tree.isBalanced(sr, sl) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.copy(&l)

   sll := l.s
   slr := sl - l.s - 1

   l.r = tree.assembleLR(p, l.r, r, slr, sr)
   slr = 1 + sr + slr

   if !tree.isBalanced(sll, slr) {

      srr := slr - l.r.s - 1
      srl := l.r.s

      if tree.singleRotation(srr, srl) {
         tree.rotateL(&l)
      } else {
         tree.rotateRL(&l)
      }
   }
   return l
}

func (tree *LBST) assembleRL(p *Node, l, r *Node, sl, sr Size) *Node {
   if tree.isBalanced(sl, sr) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.copy(&r)

   srl := r.s
   srr := sr - r.s - 1

   r.l = tree.assembleRL(p, l, r.l, sl, srl)
   r.s = 1 + sl + srl

   if !tree.isBalanced(srr, r.s) {
      if tree.singleRotation(r.l.s, r.s-r.l.s-1) {
         tree.rotateR(&r)
      } else {
         tree.rotateLR(&r)
      }
   }
   return r
}

func (tree LBST) Split(i Position) (LBST, LBST) {
   tree.share(tree.root)
   l, r := JoinBased{Tree: tree.Tree, Joiner: tree}.split(tree.root, i, tree.size)

   return LBST{Tree{arena: tree.arena, root: l, size: i}},
      LBST{Tree{arena: tree.arena, root: r, size: tree.size - i}}
}
