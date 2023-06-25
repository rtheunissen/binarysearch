package binarytree

import . "binarysearch/abstract/list"

type RedBlackRelaxedBottomUp struct {
   RedBlackRelaxed
}

func (RedBlackRelaxedBottomUp) New() List {
   return &RedBlackRelaxedBottomUp{}
}

func (tree *RedBlackRelaxedBottomUp) Clone() List {
   return &RedBlackRelaxedBottomUp{
      RedBlackRelaxed: RedBlackRelaxed{
         Tree: tree.Tree.Clone(),
      },
   }
}

func (tree *RedBlackRelaxedBottomUp) insert(p *Node, i Position, x Data) *Node {
   if p == nil {
      return tree.allocate(Node{x: x})
   }
   tree.persist(&p)
   if i <= p.s {
      p.s = p.s + 1
      p.l = tree.insert(p.l, i, x)
      return tree.fixL(p)
   } else {
      p.r = tree.insert(p.r, i-p.s-1, x)
      return tree.fixR(p)
   }
}

func (tree *RedBlackRelaxedBottomUp) Insert(i Position, x Data) {
   assert(i <= tree.Size())
   tree.size = tree.size + 1
   tree.root = tree.insert(tree.root, i, x)
   return
}

func (tree *RedBlackRelaxedBottomUp) Select(i Size) Data {
   assert(i < tree.Size())
   return tree.lookup(tree.root, i)
}

func (tree *RedBlackRelaxedBottomUp) Update(i Size, x Data) {
   assert(i < tree.Size())
   tree.persist(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *RedBlackRelaxedBottomUp) Join(other List) List {
   return &RedBlackRelaxedBottomUp{
      tree.RedBlackRelaxed.Join(other.(*RedBlackRelaxedBottomUp).RedBlackRelaxed),
   }
}

func (tree *RedBlackRelaxedBottomUp) Split(i Position) (List, List) {
   l, r := tree.RedBlackRelaxed.Split(i)
   return &RedBlackRelaxedBottomUp{l},
          &RedBlackRelaxedBottomUp{r}
}
