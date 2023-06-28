package binarytree

import "binarysearch/abstract/list"

type RedBlackRelaxedBottomUp struct {
   RedBlackBottomUp
   RedBlackRelaxed
}

func (tree RedBlackRelaxedBottomUp) Verify() {
   tree.Tree.verifySize(tree.root, tree.size)
   tree.RedBlackRelaxed.verifyRanks(tree.root)
   tree.RedBlackRelaxed.verifyHeight(tree.root)
}

func (RedBlackRelaxedBottomUp) New() list.List {
   return &RedBlackRelaxedBottomUp{}
}

func (tree *RedBlackRelaxedBottomUp) Clone() list.List {
   return &RedBlackRelaxedBottomUp{
      RedBlackBottomUp: *tree.RedBlackBottomUp.Clone().(*RedBlackBottomUp),
   }
}

func (tree *RedBlackRelaxedBottomUp) Insert(i list.Position, x list.Data) {
   tree.RedBlackBottomUp.Insert(i, x)
}

func (tree *RedBlackRelaxedBottomUp) Delete(i list.Position) (x list.Data) {
   return tree.Tree.Delete(i)
}

func (tree *RedBlackRelaxedBottomUp) join(l, r *Node, sl list.Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) <= tree.rank(r) {
      return tree.RedBlackBottomUp.build(l, tree.Tree.deleteMin(&r), r, sl)
   } else {
      return tree.RedBlackBottomUp.build(l, tree.Tree.deleteMax(&l), r, sl-1)
   }
}

func (tree *RedBlackRelaxedBottomUp) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*RedBlackRelaxedBottomUp).root)
   return &RedBlackRelaxedBottomUp{
      RedBlackBottomUp: RedBlackBottomUp{
         Tree: Tree{
            arena: tree.arena, // TODO maybe leave this nil? should probably have its own
            root:  tree.join(tree.root, other.(*RedBlackRelaxedBottomUp).root, tree.size),
            size:  tree.size + other.(*RedBlackRelaxedBottomUp).size,
         },
      },
   }
}

func (tree RedBlackRelaxedBottomUp) Split(i list.Position) (list.List, list.List) {
   l, r := tree.RedBlackBottomUp.Split(i)
   return &RedBlackRelaxedBottomUp{RedBlackBottomUp: *l.(*RedBlackBottomUp)},
          &RedBlackRelaxedBottomUp{RedBlackBottomUp: *r.(*RedBlackBottomUp)}
}
