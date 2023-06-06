package binarytree

import (
   "arena"
   "binarysearch/abstract/list"
)

//type ExteriorHeightsAlongTheSpines = [2][]int

//type InteriorHeightsAlongTheSpines = [2][]int

//type SymmetricWeightPerLevel = [2][]Size

// TODO I don't think this interface has any Data
type BinaryTree interface {
   Root() *Node
   Size() list.Size

   InteriorHeightsAlongTheSpines() [2][]int
   ExteriorHeightsAlongTheSpines() [2][]int
   SymmetricWeightPerLevel() [2][]list.Size
}

// The internal path length of a binary tree that has n internal nodes is
// sum (i=1 to n) leveli where leveli is the length of the path
// from the root to internal node i .
//
// In a complete binary tree
// one leveli is 0 ,
// two leveli s are 1 ,
// four leveli s are 2 ,
// eight leveli s are 4 , and so on.

// That is, the external path length of a full binary tree equals the
// internal path length plus twice the number of internal nodes

// // Returns the total number of links between p* and all nodes reachable BST p*.
//func InternalPathLength(p BinaryTreeNode, s Size) Size {
//   if reflect.DataOf(p).IsNil() {
//      return 0
//   }
//   // assert(p.Count() == s)
//   iL := InternalPathLength(p.Left(), p.SizeOfLeftSubtree(s))
//   iR := InternalPathLength(p.Right(), p.SizeOfRightSubtree(s))
//   return s - 1 + iL + iR // TODO reference this
//}
//
//func Count(p BinaryTreeNode) Size {
//   if p == nil {
//      return 0
//   } else {
//      return 1 + Count(p.Left()) + Count(p.Right())
//   }
//}

//
// // Returns the average path length of all nodes reachable BST p*.
//func AveragePathLength(p BinaryTreeNode, s Size) float64 {
//   if s == 0 {
//      return 0
//   }
//   return float64(p.InternalPathLength(s)) / float64(s)
//}

type Tree struct {
   arena *arena.Arena
   root *Node
   size list.Size
}

// Clone creates a shallow copy of the tree and shares its root with the copy.
func (tree *Tree) Clone() Tree {
   if tree.arena == nil {
      tree.arena = arena.NewArena()
   }
   tree.share(tree.root)
   return *tree
}

func (tree Tree) Free() {
   if tree.arena != nil {
      tree.arena.Free()
   }
}

//
//func (tree Tree) Values() []Data {
//   values := make([]Data, 0, tree.size)
//   tree.Each(func(x Data) {
//      values = append(values, x)
//   })
//   return values
//}
//
//func (tree *Tree) build(Datas []Data) *Node {
//   if len(Datas) == 0 {
//      return nil
//   }
//   m := len(Datas) - (len(Datas) >> 1)
//
//   return tree.allocate(Node{
//      x: Datas[m],
//      s: Size(m),
//      l: tree.build(Datas[:m]),
//      r: tree.build(Datas[m+1:]),
//   })
//}

func (tree *Tree) insert(p **Node, s list.Size, i list.Position, x list.Data) {
   for {
      if *p == nil {
         *p = tree.allocate(Node{x: x})
         return
      }
      tree.copy(p)
      sl := (*p).s
      sr := s - (*p).s - 1

      if i > sl {
         p = insertR(*p, &i)
         s = sr
      } else {
         p = insertL(*p)
         s = sl
      }
   }
}

func (tree *Tree) Insert(i list.Size, x list.Data) {
   tree.insert(&tree.root, tree.size, i, x)
   tree.size++
}

//func (tree *Tree) Split(i Size) (Tree, Tree) {
//   // assert(i <= tree.size)
//
//   var l, r *Node
//   tree.partition(tree.root, i, &l, &r)
//
//   return Tree{arena: tree.arena, root: l, size: i},
//          Tree{arena: tree.arena, root: r, size: tree.size - i}
//}

func (tree Tree) Verify() {
   tree.verifySize(tree.root, tree.size)
}

//
//func (tree Tree) Root() *Node {
//  return tree.root
//}

func (tree Tree) Size() list.Size {
   return tree.size
}

func (tree Tree) Root() *Node {
   return tree.root
}

func (tree Tree) lookup(p *Node, i list.Size) list.Data {
   for {
      if i == p.s {
         return p.x
      }
      if i < p.s {
         p = p.l
      } else {
         i = i - p.s - 1
         p = p.r
      }
   }
}

//func (tree *Tree) Select(i Size) Data {
//  // assert(i < tree.Size())
//  return tree.lookup(tree.root, i)
//}

//func (tree *Tree) Update(i Size, x Data) {
//   // assert(i < tree.Size())
//   tree.pathcopy(&tree.root)
//   tree.update(tree.root, i, x)
//}

func (tree Tree) Each(visit func(list.Data)) {
   tree.root.inorder(visit)
}

//
//func (tree *Tree) dissolve(p **Node, s Size) Data {
//   defer tree.release(*p)
//   x := (*p).x
//
//   if (*p).l == nil { *p = (*p).r; return x }
//   if (*p).r == nil { *p = (*p).l; return x }
//
//   if (*p).sizeL() <= (*p).sizeR(s) {
//      m := tree.deleteMin(&(*p).r)
//      m.l = (*p).l
//      m.r = (*p).r
//      m.s = (*p).s
//      m.y = (*p).y
//      *p = m
//   } else {
//      m := tree.deleteMax(&(*p).l)
//      m.l = (*p).l
//      m.r = (*p).r
//      m.s = (*p).s - 1
//      m.y = (*p).y
//      *p = m
//   }
//   return x
//}

//func (tree *Tree) dissolve(p *Node, s Size) *Node {
//   if p.l == nil { tree.release(p); return p.r }
//   if p.r == nil { tree.release(p); return p.l }
//   if p.sizeL() < p.sizeR(s) {
//      defer tree.release(p)
//      var x *Node
//      p.r, x = tree.deleteMin2(p.r)
//      p.x = x.x
//   } else {
//      defer tree.release(p)
//      p.x = tree.deleteMax(&p.l).x
//      p.s = (*p).s - 1
//   }
//   return p
//}

func (tree *Tree) dissolve(p *Node, s list.Size) *Node {
   defer tree.release(p)
   if p.l == nil {
      return p.r
   }
   if p.r == nil {
      return p.l
   }
   if p.sizeL() <= p.sizeR(s) {
      r := tree.deleteMin(&p.r)
      r.l = p.l
      r.r = p.r
      r.s = p.s
      r.y = p.y
      return r
   } else {
      l := tree.deleteMax(&p.l)
      l.r = p.r
      l.l = p.l
      l.s = p.s - 1
      l.y = p.y
      return l
   }
}

func (tree *Tree) delete(p **Node, s list.Size, i list.Size) (x list.Data) {
   for {
      tree.copy(p)
      if i == (*p).s {
         x := (*p).x
         *p = tree.dissolve(*p, s)
         return x
      }
      if i < (*p).s {
         s = (*p).s
         (*p).s = (*p).s - 1
         p = &(*p).l
      } else {
         i = i - (*p).s - 1
         s = s - (*p).s - 1
         p = &(*p).r
      }
   }
}

//func (p *Node) deleteL(r **Node, s Size, i Position) (**Node, Size, Position) {
//   s = p.i
//   p.i--
//   *r = p
//   r = &p.l
//   return r, s, i
//}
//
//func (p *Node) deleteR(n **Node, s Size, i Position) (**Node, Size, Position) {
//   i = i - (*p).i - 1
//   s = s - (*p).i - 1
//   *n = p
//   n = &p.r
//   return n, s, i
//}
// Deletes the node at position `i` from the tree.
// Returns the data that was in the deleted Data.
//func (tree *Tree) Delete(i Size) Data {
//   // assert(i < tree.Size())
//   x := tree.delete(&tree.root, tree.size, i)
//   tree.size--
//   return x
//}

func (tree Tree) verifySizes() {
   tree.verifySize(tree.root, tree.size)
}

func (tree Tree) InteriorHeightsAlongTheSpines() (h [2][]int) {
   if tree.root == nil {
      return
   }
   //
   for l := tree.root.l; l != nil; l = l.l {
      h[0] = append(h[0], l.r.height()+1)
   }
   for r := tree.root.r; r != nil; r = r.r {
      h[1] = append(h[1], r.l.height()+1)
   }

   // Reverse the left spine.
   i := 0
   j := len(h[0]) - 1
   for i < j {
      h[0][i], h[0][j] = h[0][j], h[0][i]
      i++
      j--
   }
   return
}

func (tree Tree) ExteriorHeightsAlongTheSpines() (h [2][]int) {
   if tree.root == nil {
      return
   }
   //
   for l := tree.root.l; l != nil; l = l.l {
      h[0] = append(h[0], l.height())
   }
   for r := tree.root.r; r != nil; r = r.r {
      h[1] = append(h[1], r.height())
   }

   // Reverse the left spine.
   i := 0
   j := len(h[0]) - 1
   for i < j {
      h[0][i], h[0][j] = h[0][j], h[0][i]
      i++
      j--
   }
   return
}

func (tree Tree) countNodesPerLevel(p *Node, counter *[]list.Size, level int) {
   if p == nil {
      return
   }
   // Add more levels to the counter as needed on the way down.
   if len(*counter) <= level {
      *counter = append(*counter, 0)
   }
   (*counter)[level]++
   tree.countNodesPerLevel(p.l, counter, level+1)
   tree.countNodesPerLevel(p.r, counter, level+1)
}

func (tree Tree) SymmetricWeightPerLevel() (weights [2][]list.Size) {
   if tree.root == nil {
      return
   }
   tree.countNodesPerLevel(tree.root.l, &weights[0], 0)
   tree.countNodesPerLevel(tree.root.r, &weights[1], 0)
   return
}
