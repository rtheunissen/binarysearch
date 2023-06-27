package binarytree

import . "binarysearch/abstract/list"

type RedBlackRelaxedBottomUp struct {
   RedBlackBottomUp
   RedBlackRelaxed
}

func (tree RedBlackRelaxedBottomUp) Verify() {
   tree.Tree.verifySize(tree.root, tree.size)
   tree.RedBlackRelaxed.verifyRanks(tree.root)
   tree.RedBlackRelaxed.verifyHeight(tree.root)
}

func (RedBlackRelaxedBottomUp) New() List {
   return &RedBlackRelaxedBottomUp{}
}

func (tree *RedBlackRelaxedBottomUp) Clone() List {
   return &RedBlackRelaxedBottomUp{
      RedBlackBottomUp: RedBlackBottomUp{
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
      return tree.balanceInsertL(p)
   } else {
      p.r = tree.insert(p.r, i-p.s-1, x)
      return tree.balanceInsertR(p)
   }
}

func (tree *RedBlackRelaxedBottomUp) Insert(i Position, x Data) {
   // assert(i <= tree.Size())
   tree.size = tree.size + 1
   tree.root = tree.insert(tree.root, i, x)
   return
}

func (tree RedBlackRelaxedBottomUp) Join(other List) List {
   tree.share(tree.root)
   tree.share(other.(*RedBlackRelaxedBottomUp).root)
   return &RedBlackRelaxedBottomUp{
      RedBlackBottomUp: RedBlackBottomUp{
         Tree: Tree{
            arena: tree.arena,
            root:  tree.join(tree.root, other.(*RedBlackRelaxedBottomUp).root, tree.size),
            size:  tree.size + other.(*RedBlackRelaxedBottomUp).size,
         },
      },
   }
}

func (tree RedBlackRelaxedBottomUp) Split(i Position) (List, List) {
   // assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)
   return &RedBlackRelaxedBottomUp{RedBlackBottomUp: RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}}},
          &RedBlackRelaxedBottomUp{RedBlackBottomUp: RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}
}
