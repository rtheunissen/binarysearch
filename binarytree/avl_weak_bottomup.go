package binarytree

import . "binarysearch/abstract/list"

type AVLWeakBottomUp struct {
   WAVL // TODO: just use tree
}

func (AVLWeakBottomUp) New() List {
   return &AVLWeakBottomUp{}
}

func (tree *AVLWeakBottomUp) Clone() List {
   return &AVLWeakBottomUp{
      WAVL{
         Tree: tree.Tree.Clone(),
      },
   }
}

func (tree *AVLWeakBottomUp) insert(p *Node, i Position, x Data) *Node {
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

//
// "Deletion of a leaf may convert its parent, previously a 1,2 node
//  into a 2,2 leaf, violating the rank rule. In this case we begin
//  by demoting the parent, which may make it a 3-child."
//
// func (tree AVLWeakBottomUp) rebalanceBottomUpDeletingLeft(p *Node) *Node {
//    return tree.rebalanceDeletingLeft(p)
// }

//
// "Deletion of a leaf may convert its parent, previously a 1,2 node
//  into a 2,2 leaf, violating the rank rule. In this case we begin
//  by demoting the parent, which may make it a 3-child."
//
// func (tree AVLWeakBottomUp) rebalanceBottomUpDeletingRight(p *Node) *Node {
//    return tree.rebalanceDeletingRight(p)
// }

func (tree AVLWeakBottomUp) delete(p *Node, i Position, x *Data) *Node {
   tree.persist(&p)
   if i == p.s {
      *x = p.x
      defer tree.free(p)
      return tree.join(p.l, p.r, p.s)
   }
   if i < p.s {
      p.s = p.s - 1
      p.l = tree.delete(p.l, i, x)
   } else {
      p.r = tree.delete(p.r, i-p.s-1, x)
   }
   return tree.rebalanceOnDelete(p)
}

func (tree *AVLWeakBottomUp) Delete(i Position) (x Data) {
   assert(i < tree.size)
   tree.root = tree.delete(tree.root, i, &x)
   tree.size = tree.size - 1
   return
}

func (tree *AVLWeakBottomUp) Insert(i Position, x Data) {
   assert(i <= tree.size)
   tree.size = tree.size + 1
   tree.root = tree.insert(tree.root, i, x)
}

func (tree AVLWeakBottomUp) extractMin(p *Node, min **Node) *Node {
   if p.l == nil {
      *min = tree.replacedByRightSubtree(&p)
      return p
   }
   tree.persist(&p)
   p.s--
   p.l = tree.extractMin(p.l, min)
   return tree.rebalanceOnDelete(p)
}

// TODO: rename to deleteMin and sort out conflicts with relaxed which embeds it
func (tree AVLWeakBottomUp) extractMax(p *Node, max **Node) *Node {
   if p.r == nil {
      *max = tree.replacedByLeftSubtree(&p)
      return p
   }
   tree.persist(&p)
   p.r = tree.extractMax(p.r, max)
   return tree.rebalanceOnDelete(p)
}

func (tree AVLWeakBottomUp) join(l, r *Node, sl Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) <= tree.rank(r) {
      return tree.build(l, p, tree.extractMin(r, &p), sl)
   } else {
      return tree.build(tree.extractMax(l, &p), p, r, sl-1)
   }
}

func (tree AVLWeakBottomUp) Join(that List) List {
   tree.share(tree.root)
   tree.share(that.(*AVLWeakBottomUp).root)
   return &AVLWeakBottomUp{
      WAVL{
         Tree: Tree{
            arena: tree.arena,
            root:  tree.join(tree.root, that.(*AVLWeakBottomUp).root, tree.size),
            size:  tree.size + that.(*AVLWeakBottomUp).size,
         },
      },
   }
}


func (tree AVLWeakBottomUp) split(p *Node, i, s Size) (l, r *Node) {
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

func (tree AVLWeakBottomUp) Split(i Position) (List, List) {
   assert(i <= tree.size)
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)

   return &AVLWeakBottomUp{WAVL{Tree: Tree{arena: tree.arena, root: l, size: i}}},
          &AVLWeakBottomUp{WAVL{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}
}
