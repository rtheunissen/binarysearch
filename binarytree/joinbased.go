package binarytree

import . "binarysearch/abstract/list"

// TODO just inline all of this stuff

type Joiner interface {
   join3(l, p, r *Node, sl, sr Size) *Node
   join2(l, r *Node, sl, sr Size) *Node
}

type JoinBased struct {
   Tree
   Joiner
}

func (tree JoinBased) insert(p *Node, i Position, s Size, x *Node) *Node {
   if p == nil {
      return x
   }
   tree.copy(&p)

   sl := p.s
   sr := s - p.s - 1

   if i <= p.s {
      p.s++
      return tree.join3(tree.insert(p.l, i, sl, x), p, p.r, sl+1, sr)
   } else {
      return tree.join3(p.l, p, tree.insert(p.r, i-sl-1, sr, x), sl, sr+1)
   }
}

func (tree JoinBased) delete(p *Node, i Position, s Size, x *Data) *Node {
   tree.copy(&p)

   sl := p.s
   sr := s - p.s - 1

   if i == p.s {
      *x = p.x
      defer tree.release(p)
      return tree.join2(p.l, p.r, sl, sr)
   }
   if i < p.s {
      p.s--
      return tree.join3(tree.delete(p.l, i, sl, x), p, p.r, sl-1, sr)
   } else {
      return tree.join3(p.l, p, tree.delete(p.r, i-sl-1, sr, x), sl, sr-1)
   }
}

func (tree JoinBased) split(p *Node, i, s Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.copy(&p)

   sl := p.s
   sr := s - p.s - 1

   if i <= (*p).s {
      l, r = tree.split(p.l, i, sl)
      r = tree.join3(r, p, p.r, sl-i, sr)
   } else {
      l, r = tree.split(p.r, i-sl-1, sr)
      l = tree.join3(p.l, p, l, sl, i-sl-1)
   }
   return l, r
}

func (tree JoinBased) splitToBST(p *Node, i, s Size) (Tree, Tree) {
   l, r := tree.split(p, i, s)
   return Tree{arena: tree.arena, root: l, size: i},
      Tree{arena: tree.arena, root: r, size: s - i}
}

//func (tree JoinBased) join(l, r *Node, sl, sr Size) *Node {
//   return tree.join2(l, r, sl, sr)
//}
