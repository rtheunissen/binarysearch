package binarytree

import . "binarysearch/abstract/list"

type AVLWeakJoinBased struct {
   WAVL
}

func (AVLWeakJoinBased) New() List {
   return &AVLWeakJoinBased{}
}

func (tree *AVLWeakJoinBased) Select(i Size) Data {
   assert(i < tree.Size())
   return tree.lookup(tree.root, i)
}

func (tree *AVLWeakJoinBased) Update(i Size, x Data) {
   assert(i < tree.Size())
   tree.copy(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *AVLWeakJoinBased) Insert(i Position, x Data) {
   assert(i <= tree.Size())
   tree.root = JoinBased{Tree: tree.Tree, Joiner: tree}.insert(tree.root, i, tree.size, tree.allocate(Node{x: x}))
   tree.size++
}

func (tree *AVLWeakJoinBased) Delete(i Position) (x Data) {
   assert(i < tree.Size())
   tree.root = JoinBased{Tree: tree.Tree, Joiner: tree}.delete(tree.root, i, tree.size, &x)
   tree.size--
   return
}

func (tree *AVLWeakJoinBased) Clone() List {
   return &AVLWeakJoinBased{
      WAVL{
         Tree: tree.Tree.Clone(),
      },
   }
}

func (tree AVLWeakJoinBased) Split(i Position) (List, List) {
   assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := JoinBased{Tree: tree.Tree, Joiner: tree}.splitToBST(tree.root, i, tree.size)
   return &AVLWeakJoinBased{WAVL{Tree: l}},
      &AVLWeakJoinBased{WAVL{Tree: r}}
}

func (tree AVLWeakJoinBased) extractMin(p *Node, min **Node) *Node {
   if p.l == nil {
      *min = tree.replacedByRightSubtree(&p)
      return p
   }
   tree.copy(&p)
   p.s--
   p.l = tree.extractMin(p.l, min)
   return tree.rebalanceOnDelete(p)
}

func (tree AVLWeakJoinBased) extractMax(p *Node, max **Node) *Node {
   if p.r == nil {
      *max = tree.replacedByLeftSubtree(&p)
      return p
   }
   tree.copy(&p)
   p.r = tree.extractMax(p.r, max)
   return tree.rebalanceOnDelete(p)
}

func (tree AVLWeakJoinBased) join2(l, r *Node, sl, sr Size) (p *Node) {
   return tree.join(l, r, sl)
}

func (tree AVLWeakJoinBased) join3(l, p, r *Node, sl, sr Size) *Node {
   return tree.build(l, p, r, sl)
}

func (tree AVLWeakJoinBased) join(l, r *Node, sl Size) (p *Node) {
   if l == nil {
      return r
   }
   if r == nil {
      return l
   }

   if tree.rank(l) <= tree.rank(r) {
      return tree.build(l, p, tree.extractMin(r, &p), sl)
   } else {
      return tree.build(tree.extractMax(l, &p), p, r, sl-1)
   }
}

func (tree AVLWeakJoinBased) Join(other List) List {
   l := tree
   r := other.(*AVLWeakJoinBased)

   tree.share(l.root)
   tree.share(r.root)
   return &AVLWeakJoinBased{
      WAVL{
         Tree: Tree{
            arena: tree.arena,
            root:  tree.join(l.root, r.root, l.size),
            size:  l.size + r.size,
         },
      },
   }
}
