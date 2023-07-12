package trees

import "bst/types/list"

type AVLRelaxedTopDown struct {
   AVLWeakTopDown
   AVLRelaxed
}

func (AVLRelaxedTopDown) New() list.List {
   return &AVLRelaxedTopDown{}
}

func (tree *AVLRelaxedTopDown) Clone() list.List {
   return &AVLRelaxedTopDown{
      AVLWeakTopDown: *tree.AVLWeakTopDown.Clone().(*AVLWeakTopDown),
   }
}

func (tree AVLRelaxedTopDown) Verify() {
   tree.verifySize(tree.root, tree.size)
   tree.verifyRanks(tree.root)
}

<<<<<<< HEAD


//
func (tree *AVLRelaxedTopDown) delete(p **Node, s list.Size, i list.Size) (x list.Data) {
   for {
      tree.persist(p)
      sl := (*p).sizeL()
      sr := (*p).sizeR(s)
      if i < sl {
         s = sl
         (*p).s = sl - 1
         p = &(*p).l
         continue
      }
      if i > sl {
         s = sr
         i = i - sl - 1
         p = &(*p).r
         continue
      }
      x := (*p).x
      *p = tree.join((*p).l, (*p).r, (*p).s)
      return x
   }
=======
func (tree *AVLRelaxedTopDown) Delete(i list.Position) list.Data {
   assert(i < tree.size)
   x := tree.Tree.delete(&tree.root, tree.size, i)
   tree.size--
   return x
>>>>>>> 35027895df6e025dbd2cb64c43b9cef058796b83
}

func (tree *AVLRelaxedTopDown) Delete(i list.Position) list.Data {
  assert(i < tree.size)
  x := tree.delete(&tree.root, tree.size, i)
  tree.size = tree.size - 1
  return x
}


//func (tree *AVLRelaxedTopDown) Delete(i list.Position) list.Data {
//   return tree.Tree.Delete(i)
//}

func (tree AVLRelaxedTopDown) buildL(l *Node, p *Node, r *Node, sl list.Size) (root *Node) {
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

func (tree AVLRelaxedTopDown) buildR(l *Node, p *Node, r *Node, sl list.Size) (root *Node) {
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

func (tree *AVLRelaxedTopDown) build(l, p, r *Node, sl list.Size) *Node {
   if tree.rank(l) <= tree.rank(r) {
      return tree.buildR(l, p, r, sl)
   } else {
      return tree.buildL(l, p, r, sl)
   }
}

func (tree AVLRelaxedTopDown) join(l, r *Node, sl list.Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) <= tree.rank(r) {
      return tree.build(l, tree.Tree.deleteMin(&r), r, sl)
   } else {
      return tree.build(l, tree.Tree.deleteMax(&l), r, sl-1)
   }
}

func (tree AVLRelaxedTopDown) Join(other list.List) list.List {
   l := tree
   r := other.(*AVLRelaxedTopDown)
   tree.share(l.root)
   tree.share(r.root)

   p := tree.join(l.root, r.root, l.size)

   return &AVLRelaxedTopDown{AVLWeakTopDown: AVLWeakTopDown{AVLWeakBottomUp{AVLBottomUp: AVLBottomUp{Tree: Tree{arena: tree.arena, root: p, size: l.size + r.size}}}}}
}


func (tree AVLRelaxedTopDown) split(p *Node, i, s list.Size) (l, r *Node) {
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

func (tree AVLRelaxedTopDown) Split(i list.Position) (list.List, list.List) {
   assert(i <= tree.size)
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)

   return &AVLRelaxedTopDown{AVLWeakTopDown: AVLWeakTopDown{AVLWeakBottomUp{AVLBottomUp: AVLBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}}}}},
          &AVLRelaxedTopDown{AVLWeakTopDown: AVLWeakTopDown{AVLWeakBottomUp{AVLBottomUp: AVLBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}}}
}
