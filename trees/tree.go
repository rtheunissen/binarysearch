package trees

import (
   "arena"
   "bst/abstract/list"
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
//   assert(p.Count() == s)
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
//
//func (tree *LBSTRelaxed) partition(p *Node, i uint64) *Node {
//   assert(i < p.size())
//   measurement(&partitionCount, 1)
//   n := Node{s: i}
//   l := &n
//   r := &n
//   for i != p.s {
//      measurement(&partitionDepth, 1)
//      tree.persist(&p)
//      if i < p.s {
//         p.s = p.s - i - 1
//         r.l = p
//         r = r.l
//         p = p.l
//      } else {
//         i = i - p.s - 1
//         l.r = p
//         l = l.r
//         p = p.r
//      }
//   }
//   tree.persist(&p)
//   l.r = p.l
//   r.l = p.r
//   p.l = n.r
//   p.r = n.l
//   p.s = n.s
//   return p
//}
// Moves the node at position `i` in the tree of `p` to its root and returns the
// resulting root. This algorithm is identical to splay with no rotation steps.
//

// Clone creates a shallow copy of the tree and shares its root with the copy.
func (tree *Tree) Clone() Tree {
   if tree.arena == nil {
      tree.arena = arena.NewArena()
   }
   tree.share(tree.root)
   return *tree
}

// Deletes the node at position `i` from the tree.
// Returns the data that was in the deleted value.
func (tree *Tree) Delete(i list.Position) list.Data {
   assert(i < tree.size)
   x := tree.delete(&tree.root, tree.size, i)
   tree.size--
   return x
}


func (tree *Tree) Select(i list.Size) list.Data {
   assert(i < tree.size)
   return tree.lookup(tree.root, i)
}

func (tree *Tree) Update(i list.Size, x list.Data) {
   assert(i < tree.size)
   tree.persist(&tree.root)
   tree.update(tree.root, i, x)
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
      tree.persist(p)
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
//   assert(i <= tree.size)
//
//   var l, r *Node
//   tree.partition(tree.root, i, &l, &r)
//
//   return Tree{arena: tree.arena, root: l, size: i},
//          Tree{arena: tree.arena, root: r, size: tree.size - i}
//}



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
//  assert(i < tree.size)
//  return tree.lookup(tree.root, i)
//}

//func (tree *Tree) Update(i Size, x Data) {
//   assert(i < tree.size)
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
//
//func (tree *Tree) joinLr(l, r *Node, sl, sr list.Size) *Node {
//   n := Node{}
//   p := &n
//   for {
//      if l.sizeR(sl) <= sr  {
//         z := l.y
//         m := tree.deleteMax(&l)
//         m.l = l
//         m.r = r
//         m.s = sl - 1
//         m.y = z
//         p.r = m
//         break
//      }
//      tree.persist(&l)
//      p.r = l
//      p = p.r
//      sl = sl - p.sizeL() - 1
//      l = l.r
//   }
//   return n.r
//   //if l == nil {
//   //   return r
//   //}
//   //if sl <= sr {
//   //   p := tree.deleteMax(&l)
//   //   p.l = l
//   //   p.r = r
//   //   p.s = sl - 1
//   //   return p
//   //}
//   //tree.persist(&l)
//   //l.r = tree.joinLr(l.r, r, sl - l.s - 1, sr)
//   //return l
//}
//
//func (tree *Tree) joinlR(l, r *Node, sl, sr list.Size) *Node {
//   n := Node{}
//   p := &n
//   for {
//      // func (tree *Tree) deleteMin2(p *Node) (root, min *Node) {
//      //   n := Node{}
//      //   l := &n
//      //   for {
//      //      tree.persist(&p)
//      //      if p.l == nil {
//      //         l.l = p.r
//      //         return n.l, p
//      //      }
//      //      p.s = p.s - 1
//      //      l.l = p
//      //      l = l.l
//      //      p = p.l
//      //   }
//      //}
//      if r.sizeL() <= sl  {
//         var m *Node
//         z := r.y
//         r, m = tree.deleteMin2(r)
//         m.l = l
//         m.r = r
//         m.s = sl
//         m.y = z
//         p.l = m
//         break
//      }
//      tree.persist(&r)
//      p.l = r
//      p = p.l
//      r.s = r.s + sl
//      r = r.l
//   }
//   return n.l
//   //
//   //
//   //if r == nil {
//   //   return l
//   //}
//   //if sr <= sl {
//   //   p := tree.deleteMin(&r)
//   //   p.l = l
//   //   p.r = r
//   //   p.s = sl
//   //   return p
//   //}
//   //tree.persist(&r)
//   //r.l = tree.joinlR(l, r.l, sl, r.s)
//   //r.s = sl + r.s
//   //return r
//}
//
//func (tree *Tree) join(l, r *Node, sl, sr list.Size) *Node {
//   if l == nil {
//      return r
//   }
//   if r == nil {
//      return l
//   }
//   if sl <= sr {
//      tree.persist(&l)
//      tree.persist(&r)
//      p := tree.deleteMin(&r)
//      p.l = l
//      p.r = r
//      p.s = sl
//      return p
//   } else {
//      tree.persist(&l)
//      tree.persist(&r)
//      p := tree.deleteMax(&l)
//      p.r = r
//      p.l = l
//      p.s = sl - 1
//      return p
//   }
//   //if l == nil {
//   //   return r
//   //}
//   //if sl < sr {
//   //   return tree.joinlR(l, r, sl, sr)
//   //} else {
//   //   return tree.joinLr(l, r, sl, sr)
//   //}
//}

func (tree *Tree) dissolve(p *Node, s list.Size) *Node {
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
      tree.persist(p)
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


func (tree *Tree) split(p *Node, i uint64) (*Node, *Node) {
   n := Node{}
   l := &n
   r := &n
   for p != nil {
      tree.persist(&p)
      if i <= p.s {
        p.s = p.s - i
        r.l = p
          r = r.l
          p = p.l
      } else {
          i = i - p.s - 1
        l.r = p
          l = l.r
          p = p.r
      }
   }
   l.r = nil
   r.l = nil
   return n.r, n.l
}

func (tree Tree) Split(i list.Size, split func(p *Node, i, s list.Size) (l, r *Node)) (Tree, Tree) {
   tree.share(tree.root)
   l, r := split(tree.root, i, tree.size)
   return Tree{arena: tree.arena, root: l, size: i}, // TODO: new arenas?
          Tree{arena: tree.arena, root: r, size: tree.size - i}
}

func (tree Tree) Join(other Tree, join func(l, r *Node, sl list.Size) *Node) Tree {
   tree.share(tree.root)
   tree.share(other.root)
   return Tree{
      arena: tree.arena, // TODO: remove?
      size:  tree.size + other.size,
      root:  join(tree.root, other.root, tree.size),
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
//   assert(i < tree.size)
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
