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
      RedBlackBottomUp: *tree.RedBlackBottomUp.Clone().(*RedBlackBottomUp),
   }
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
   assert(i <= tree.size)
   l, r := tree.RedBlackBottomUp.Split(i)
   return &RedBlackRelaxedBottomUp{RedBlackBottomUp: *l.(*RedBlackBottomUp)},
          &RedBlackRelaxedBottomUp{RedBlackBottomUp: *r.(*RedBlackBottomUp)}
}
