package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/utility"
)

type BinaryTreeNode interface {
   Height() int
   TotalInternalPathLength(Size) Size
   TotalReferenceCount() Size
}

// Node
//
// A node is a simple structure that can be linked to other nodes. Each node has
// a left and right outgoing link to other nodes. TODO: parent, binary tree
// The node at the top of the tree with no parent is called the _root_ node.
//
//
//                                   (P)arent
//                                 ↙     ↘
//                               (L)eft  (R)ight
//
//
// A node is a container for a unit of information.
// When multiple nodes are linked together they form a binary tree.
// This structure allows the ability to organize information.
// One type of organization is a sequence - linear, list, etc.
//
// A sequence is implied by binary search tree, where the parent appears in the
// sequence after the left node and before the right node.
//
// This can be viewed as a binary search tree ordered by sequential position.
//
// Every node tracks its 0-based position relative to the start of its sequence,
// equal to the number of nodes in its left subtree. Given the total size of a
// node, we can calculate the sizes of both subtrees without referencing them.
//
//
//          Position:     0   1   2   3   4   5   6
//
//          Sequence:    [e,  x,  a,  m,  p,  l,  e]
//
//                                    3
//          ParseAnimation:                    (m)
//                            1   ↙       ↘   1
//                           (x)             (l)
//                        0 ↙   ↘ 0       0 ↙   ↘ 0
//                       (e)     (a)     (p)     (e)
//
//
// Notice the vertical projection of the sequence onto the tree, which follows
// the in-order traversal from the root, recursively left-self-right.
//

// Multiple trees may share the same node, allowing independent trees to be made
// up of common subtrees shared in memory. Making a change to one tree data
// in a new tree that shares most of the previous structure. The reference count
// of a node is the number of other trees that reference it, thus zero indicates
// that a node is only referenced by its own tree and is not shared by others.

type Node struct {
   ReferenceCounter

   l *Node // Pointers to the left and right subtrees.
   r *Node

   s uint64 // Size, usually of the left subtree and therefore also position.
   x uint64 // Data
   y uint64 // Rank
}

//func (p Node) Position() Size {
//   return p.s
//}
//
//func (p Node) SizeOfRightSubtree(size Size) Size {
//   return p.sizeR(size)
//}

//func (p Node) Count() Size {
//  return p.size()
//}
//
//func (p Node) Left() BinaryTreeNode {
//   return p.l
//}
//
//func (p Node) Right() BinaryTreeNode {
//   return p.r
//}
//
//func (p Node) Data() Data {
//   return p.x
//}

// https://aofa.cs.princeton.edu/online/slides/AA06-Trees.pdf
// https://dl.acm.org/doi/pdf/10.1145/126459.126463

// Data Structures and Algorithms in Python by Michael T. Goodrich, Roberto Tamassia, and Michael H. Goldwasser. This book provides an explanation and implementation of the algorithm in Python.
// Introduction to Algorithms by Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, and Clifford Stein. This book provides a detailed explanation of the algorithm and its analysis in Chapter 12.
// Algorithms by Sanjoy Dasgupta, Christos Papadimitriou, and Umesh Vazirani. This book provides an explanation and implementation of the algorithm in Python in Chapter 6.
//
// ??
func averagePathLength(p *Node, depth uint64, totalDepth *uint64, totalNodes *uint64) {
   if p == nil {
      return
   }
   *totalNodes = *totalNodes + 1
   *totalDepth = *totalDepth + depth
   averagePathLength(p.l, depth+1, totalDepth, totalNodes)
   averagePathLength(p.r, depth+1, totalDepth, totalNodes)
}

// long IPL(lbst t)
// { if (t == null) return 0;
//
//    return IPL(t->l) + IPL(t->r) + t->s;
//   }
func (p *Node) AveragePathLength() float64 {
   var totalDepth uint64
   var totalNodes uint64
   averagePathLength(p, 0, &totalDepth, &totalNodes)
   return float64(totalDepth) / float64(totalNodes)
}

//

// TODO: maybe this can be moved to tree.pathcopy(**Node) when all trees
//       use the same node structure and most node functions are moved inline,
//       then we can consider `t` as a type label instead of `tree` and move the
//       scope to be as local as possibe within the file.
//
//       Using the reference means we just need to dereference to read the refs,
//      and if non-zero, we construct and assign. When it is zero, we only did
//      dereference and did not assign. Also, it means you only specify the Data
//      once, and also if the pattern is always A = pathcopy(A) then that A is
//      clearly redundant and can simplified to pathcopy(&A).
//
//

// TODO: Maybe put this on each tree where used
func rank(p *Node) uint64 {
   if p == nil {
      return 0
   } else {
      return p.y
   }
}

func (p Node) isLeaf() bool { // TODO conc can have its own that doesn't check p.r for nil
   return p.l == nil && p.r == nil
}

// Max: 2n - 1
// Min: 0
//
// Complexity: O(n)
//

func (p *Node) MaximumPathLength() int {
   return p.height()
}

func (p *Node) height() int {
   if p == nil {
      return -1
   }
   return 1 + utility.Max(p.l.height(), p.r.height())
}

// Counts the number of nodes reachable from p*, including itself.
func (p *Node) size() Size {
   if p == nil {
      return 0
   } else {
      return 1 + p.l.size() + p.r.size()
   }
}

// Returns the number of nodes in the left subtree of p*.
// TODO: This is not the case for all tree implementations - should this be up to the tree? Maybe mix it in?
func (p *Node) sizeL() Size {
   return p.s
}

// Returns the number of nodes in the right subtree of p*, given the s of p*.
func (p *Node) sizeR(s Size) Size {
   return s - p.s - 1
}

// Recursive in-order traversal of p*, left-self-right.
// This will call the given function for every Node in the sequence of p*.
// func (p *Node) inorder(visit func(Node)) {
//    if p == nil {
//       return
//    }
//    p.l.inorder(visit)
//    visit(*p)
//    p.r.inorder(visit)
// }

//func (p *Node) inorder2(visit func(Node)) {
//   if p == nil {
//      return
//   }
//   p.l.inorder2(visit)
//   visit(*p)
//   p.r.inorder2(visit)
//}

func (p *Node) inorder(visit func(Data)) {
   if p == nil {
      return
   }
   p.l.inorder(visit)
   visit(p.x)
   p.r.inorder(visit)
}

func (p *Node) rotateL() (r *Node) {
   // measurement(&rotations, 1)
   r = p.r
   p.r = r.l
   r.l = p
   r.s = r.s + p.s + 1
   return r
}

func (p *Node) rotateR() (l *Node) {
   // measurement(&rotations, 1)
   l = p.l
   p.l = l.r
   l.r = p
   p.s = p.s - l.s - 1
   return l
}

// Rotates the LEFT subtree LEFT, then rotates the root RIGHT.
func (p *Node) rotateLR() *Node {
   p.l = p.l.rotateL()
   return p.rotateR()
}

// Rotates the RIGHT subtree RIGHT, then rotates the root LEFT.
func (p *Node) rotateRL() *Node {
   p.r = p.r.rotateR()
   return p.rotateL()
}

// `i` is the number of nodes that will be attached to the left still
// so we are not tracking the total size at any point.
// we reduce the size of the left subtree of p by i - 1
//
// size left reduces by (i + 1)
//
// the resulting size of p is then (s - p.s - 1 + i)
// the size before
//
//
//    1. Link `p` to the left of `r`
//    2. Set r
//
//
//       (p)            (r)
//                     ↙
//
//
//

//func (p *Node) deleteL(r **Node, s Size, i Position) (**Node, Size, Position) {
//   s = p.s
//   p.s--
//  *r = p
//   r = &p.l
//   return r, s, i
//}
//
//func (p *Node) deleteR(n **Node, s Size, i Position) (**Node, Size, Position) {
//   i = i - (*p).s - 1
//   s = s - (*p).s - 1
//  *n = p
//   n = &p.r
//   return n, s, i
//}
//
// // (*p).s = (*p).s + 1    // The size of the left subtree increases.
// //      p = &(*p).l       // The search path continues to the left.
// //      s = sl            // Track the subtree size going left.
// func (p *Node) pathL(r **Node, s Size, i Position) (**Node, Size, Position) {
//    s = p.s
//    p.s++
//   *r = p
//    return &p.l, s, i
// }
// //  i = i - sl - 1    // Skip all the nodes to the left, including p*
// //  p = &(*p).r       // The search path continues to the right.
// //  s = sr            // Track the subtree size going right.
// func (p *Node) pathR(n **Node, s Size, i Position) (**Node, Size, Position) {
//    i = i - (*p).s - 1
//    s = s - (*p).s - 1
//   *n = p
//    return &p.r, s, i
// }
// func (p *Node) insertLR(n **Node, s Size, i Position) (**Node, Size, Position) {
//    n, s, i = p.pathL(n, s, i)
//    return (*n).pathR(n, s, i)
// }
// func (p *Node) insertLL(n **Node, s Position, i Position) (**Node, Size, Position) {
//    n, s, i = p.pathL(n, s, i)
//    return (*n).pathL(n, s, i)
// }
// func (p *Node) insertRR(n **Node, s Size, i Position) (**Node, Size, Position) {
//    n, s, i = p.pathR(n, s, i)
//    return (*n).pathR(n, s, i)
// }
// func (p *Node) insertRL(n **Node, s Size, i Position) (**Node, Size, Position) {
//    n, s, i = p.pathR(n, s, i)
//    return (*n).pathL(n, s, i)
// }
//
//
// func (p *Node) isParentOf(child *Node) bool {
//    return child == p.l || child == p.r
// }

//
//func (p *Node) hasLeft() bool {
//   return p.l != nil
//}
//
//func (p *Node) hasRight() bool {
//   return p.r != nil
//}

func (tree Tree) verifySize(p *Node, s Size) Size {
   if p == nil {
      return 0
   }
   sl := tree.verifySize(p.l, p.sizeL())
   sr := tree.verifySize(p.r, p.sizeR(s))

   invariant(s == sl+sr+1)
   return s
}

// func partition2(root *Node, i Position) (*Node, *Node) {
//    p := root
//    n := Node{}
//    l := &n
//    r := &n
//    for {
//       if p == nil {
//          l.r = nil
//          r.l = nil
//          return n.r, n.l
//          // break
//       }
//       p = p.Copy()
//       if i == p.s {
//          l.r = p.l
//          r.l = p
//          p.l = nil
//          p.s = 0
//          return n.r, n.l
//       }
//       if i < p.s {
//          p.s = p.s - i
//          r.l = p
//          r = r.l
//          p = p.l
//       } else {
//          i = i - p.s - 1
//          l.r = p
//          l = l.r
//          p = p.r
//       }
//    }
// }

//func (tree *Tree) emplace(p **Node, i Position, n *Node) {
//   tree.split(*p, i, &n.l, &n.r)
//   n.s = i
//   *p = n
//}
//
// // func dissolve(tree joinable, p **Node, s Size, s *Data) {
// //    *s = (*p).s
// //    *p = tree.join((*p).l, (*p).r, (*p).s, (*p).sizeR(s))
// // }
//
// func join(l, r *Node, sl, sr Size) *Node {
//    // sl := (*p).s
//    // sr := (*p).sizeR(s)
//    if l == nil { return r }
//    if r == nil { return l }
//
//    if sl <= sr {
//       // return &No
//       r = r.Copy()
//       root := extractMin(&r)
//       root.s = sl
//       root.l = l
//       root.r = r
//       return root
//    } else {
//       l.s--
//       l = l.Copy()
//       root := extractMax(&l)
//       root.s = sl - 1
//       root.l = l
//       root.r = r
//
//       return root
//    }
//    //
//    //
//    // s := (*p).s
//    //
//    // if (*p).l == nil { *p = (*p).r; return s }
//    // if (*p).r == nil { *p = (*p).l; return s }
//    //
//    // if sl <= sr {
//    //    (*p).s = extractMin(&(*p).r)
//    // } else {
//    //    (*p).s = extractMax(&(*p).l).s
//    //    (*p).s--
//    // }
//    // return s
// }

//func lookup(p *Node, i Position) *Node {
//   for {
//      if i == p.i {
//         return p
//      }
//      if i < p.i {
//         p = p.l
//      } else {
//         i = i - p.i - 1
//         p = p.r
//      }
//   }
//}

func (tree *Tree) replacedByRightSubtree(p **Node) *Node {
   tree.copy(p)
   r := *p
   *p = (*p).r
   return r
}

func (tree *Tree) replacedByLeftSubtree(p **Node) *Node {
   tree.copy(p)
   l := *p
   *p = (*p).l
   return l
}

//
//func (tree *Tree) deleteMin2(p **Node) *Node {
//   for {
//      if (*p).l == nil {
//         tree.pathcopy(p)
//         r := *p
//         *p = (*p).r
//         return r
//         // return replacedByRightSubtree(p)
//      }
//      tree.pathcopy(p)
//      (*p).i--
//      p = &(*p).l
//   }
//}
//

func (tree *Tree) deleteMin2(p *Node) (root *Node, deleted *Node) {
   // assert(p != nil)
   n := Node{}
   l := &n
   for {
      tree.copy(&p)
      if p.l == nil {
         l.l = p.r
         break
      }
      p.s = p.s - 1
      l.l = p
      l = l.l
      p = p.l
   }
   return n.l, p
}

func (tree *Tree) deleteMin(p **Node) *Node {
   for {
      tree.copy(p)
      if (*p).l == nil {
         r := *p
         *p = (*p).r
         return r
      }
      (*p).s--
      p = &(*p).l
   }
}

func (tree *Tree) deleteMax(p **Node) *Node {
   for {
      if (*p).r == nil {
         tree.copy(p)
         l := *p
         *p = (*p).l
         return l
      }
      tree.copy(p)
      p = &(*p).r
   }
}

//
// func (Node) joinLr(l, r *Node, sl, sr Size) *Node {
//    if l.sizeR(sl) < sr {
//       return &Node{
//          s: extractMax(&l),
//          s: sl - 1,
//          l: l,
//          r: r,
//       }
//    }
//    l = l.Copy()
//    l.r = Node{}.joinLr(l.r, r, l.sizeR(sl), sr)
//    return l
// }
//
// func (Node) joinRl(l, r *Node, sl, sr Size) *Node {
//    if r.s < sl {
//       return &Node{
//          s: extractMin(&r),
//          s: sl,
//          l: l,
//          r: r,
//       }
//    }
//    r = r.Copy()
//    r.l = Node{}.joinRl(l, r.l, sl, r.s)
//    r.s = sl + r.s
//    return r
// }

func (tree *Tree) update(p *Node, i Position, x Data) {
   for {
      if i == p.s {
         p.x = x
         return
      }
      if i < p.s {
         tree.copy(&p.l)
         p = p.l
      } else {
         tree.copy(&p.r)
         i = i - p.s - 1
         p = p.r
      }
   }
}

// func (Node) update(root **Node, i Position, s Data) {
//    p := root
//    for {
//       *p = (*p).copy()
//       if i == (*p).s {
//          (*p).s = s
//          return
//       }
//       if i < (*p).s {
//          p = &(*p).l
//       } else {
//          i = i - (*p).s - 1
//          p = &(*p).r
//       }
//    }
// }

//
// func (p *Node) unzip(i Position) (*Node, *Node) {
//    n := Node{}
//    l := &n
//    r := &n
//    for {
//       if p == nil {
//          return n.r, n.l
//       }
//       mutable(&p)
//       if i == p.s {
//          l.r = p.l
//          r.l = p
//          p.l = nil
//          p.s = 0
//          return n.r, n.l
//       }
//       if i < p.s {
//          p.s = p.s - i // There will be `i` nodes attached to l**.
//          r.l = p
//          r = r.l
//          p = p.l
//
//       } else {
//          l.r = p
//          i = i - p.s - 1 // Skip all the nodes to the left, incl. p*.
//          l = l.r
//          p = p.r
//       }
//    }
// }
//
// func (p *Node) unzip2(i Position) (*Node, *Node) {
//    n := Node{}
//    l := &n
//    r := &n
//    for p != nil {
//       mutable(&p)
//       if i <= p.s {
//          p.s = p.s - i           // There will be `i` nodes attached to l**.
//          r.l = p
//          r = r.l
//          p = p.l
//
//       } else {
//          l.r = p
//          i = i - p.s - 1         // Skip all the nodes to the left, incl. p*.
//          l = l.r
//          p = p.r
//       }
//    }
//    l.r = nil
//    r.l = nil
//    return n.r, n.l
// }

// func (Node) split(p *Node, i Position, l, r **Node) {
//    *l, *r = p.unzip(i)
// }

//
// type BST struct {
//    List
//    Root *Node
//    Size Size
//    Strategy
// }
//
// type BST = BST
//
// type Strategy interface {
//
//    //
//    new(...Data) *BST
//
//    //
//    access(p **Node, i Position) Data
//
//    //
//    update(p **Node, i Position, s Data)
//
//    //
//    delete(p **Node, i Position, s Size) *Node
//
//    //
//    insert(p **Node, i Position, s Size, s Data)
//
//    //
//    join(l *Node, r *Node, sl, sr Size) *Node
//
//    //
//    split(p *Node, i Position, s Size, l, r **Node)
//
//    //
//    each(p *Node, visit func(Node))
//
//    //
//    filter(p *Node, predicate func(*Node) bool) (*Node, Size)
//
//    //
//    apply(p *Node, mutator func(Data) Data) *Node
//
//    //
//    verify(p *Node, s Size)
// }
// //
// // // TODO delete this
// func (BST) FromStrategy(strategy Strategy) *BST {
//    return strategy.new()
// }
// //
// // func (t BST) New(s ...Data) abstract.List {
// //    return t.Strategy.new(s...)
// // }
// //
// // // Verifies that the tree structure is valid according to the strategy in use.
// // func (tree BST) Verify() {
// //    tree.Strategy.verify(tree.Root, tree.Size)
// // }
// //
// func (tree BST) Len() Size {
//    return tree.Size
// }
// //
// // // Clones a BST
// // func (tree *BST) Clone() abstract.List {
// //    return &BST{
// //       Root:     tree.Root.reference(),
// //       Size:     tree.Size,
// //       Strategy: tree.Strategy,
// //    }
// // }
// //
// // func (tree BST) MaximumPathLength() Size {
// //    return tree.Root.MaximumPathLength()
// // }
// //
// // // Creates a tree from an existing sequence.
// // // TODO BST from array is not a great pattern, because some strategies have their own "new"
// func (tree BST) ofArray(array []Data) *BST {
//    debug(tree.Strategy != nil)
//    tree.Root = Node{}.ofArray(array)
//    tree.Size = Size(len(array))
//    return &tree
// }
// //
// // func (tree BST) Apply(mutator func (Data) Data) abstract.List {
// //    tree.Root = tree.Strategy.apply(tree.Root, func(s Data) Data {
// //       return mutator(s)
// //    })
// //    return &tree
// // }
// //
// // func (tree BST) Each(visit func(Data)) {
// //    tree.Strategy.each(tree.Root, func(p Node) {
// //       visit(p.s)
// //    })
// // }
// //
// // func (tree BST) Array() []Data {
// //    array := make([]Data, tree.Size)
// //    index := 0
// //    tree.Strategy.each(tree.Root, func(p Node) {
// //       array[index] = p.s; index++
// //    })
// //    return array
// // }
// //
// // func (tree *BST) Access(i Position) (s Data) {
// //    debug(i < tree.Size)
// //    return tree.Strategy.access(&tree.Root, i)
// // }
// //
// // func (tree *BST) Update(i Position, s Data) {
// //    debug(i < tree.Size)
// //    tree.Strategy.update(&tree.Root, i, s)
// // }
// // func (tree *BST) Set2(i Position, s Data) {
// //    debug(i < tree.Size)
// //    tree.Strategy.update(&tree.Root, i, s)
// // }
// //
// // func (tree *BST) Insert(i Position, s Data) {
// //    debug(i <= tree.Size)
// //    tree.Strategy.insert(&tree.Root, i, tree.Size, s)
// //    tree.Size++
// // }
// //
// // func (tree *BST) Delete(i Position) (s Data) {
// //    debug(i < tree.Size)
// //    s = tree.Strategy.delete(&tree.Root, i, tree.Size).s
// //    tree.Size--
// //    return
// // }
// //
// // func (tree *BST) Join(rest abstract.List) abstract.List {
// //    if tree.Size() == 0 { return rest }
// //    if rest.Size() == 0 { return tree }
// //
// //    // Clone the two joining trees to preserve them.
// //    l := tree.Clone().(*BST)
// //    r := rest.Clone().(*BST)
// //
// //    // Join the two roots into the root of the New tree.
// //    return &BST{
// //       Strategy: tree.Strategy,
// //       Root:     tree.Strategy.join(l.Root, r.Root, l.Size, r.Size),
// //       Size:     l.Size + r.Size,
// //    }
// // }
// //
// // func (tree *BST) Split(i Position) (abstract.List, abstract.List) {
// //    debug(i <= tree.Size)
// //
// //    // Create New containers for the left and right partitions.
// //    l := tree.Clone().(*BST)
// //    r := tree.Clone().(*BST)
// //
// //    // Split the tree at `i` into the left and right partitions.
// //    tree.Strategy.split(tree.Root, i, tree.Size, &l.Root, &r.Root)
// //
// //    l.Size = i
// //    r.Size = tree.Size - i
// //
// //    return l, r
// // }

//func searchL(p *Node, i Position) bool {
//   return i <= p.s
//}
//
//func searchLL(p *Node, i Position) bool {
//   return i <= p.l.s
//}

//func searchR(p *Node, i Position) bool {
//   return i > p.i
//}
//
//func searchRR(p *Node, i Position) bool {
//   return i > p.i+ p.r.i+ 1
//}

func insertL(p *Node) **Node {
   p.s++
   return &p.l
}

func insertR(p *Node, i *Position) **Node {
   *i = *i - p.s - 1
   return &p.r
}

func deleteL(p *Node) **Node {
   //println("deleteL")
   p.s--
   return &p.l
}

func deleteR(p *Node, i *Position) **Node {
   //println("deleteR")
   *i = *i - p.s - 1
   return &p.r
}

/*
func (tree *Tree) pathL(p *Node) **Node {
   // assert(p.l != nil)
   tree.copy(&p.l)
   return pathL(p)
}

func (tree *Tree) pathR(p *Node, i *Position) **Node {
   // assert(p.r != nil)
   tree.copy(&p.r)
   return pathR(p, i)
}
*/
// TODO: these are nuts
func (tree *Tree) pathLeft(p ***Node) {
   // assert((**p).l != nil)
   tree.copy(&(**p).l)
   *p = insertL(**p)
}
func (tree *Tree) pathRight(p ***Node, i *Position) {
   // assert((**p).r != nil)
   tree.copy(&(**p).r)
   *p = insertR(**p, i)
}
func (tree *Tree) attach(p **Node, x Data) {
   *p = tree.allocate(Node{x: x})
}
func (tree *Tree) attachL(p *Node, x Data) {
   p.s++
   p.l = tree.allocate(Node{x: x})
}

func (tree *Tree) attachLL(p *Node, x Data) {
   tree.copy(&p.l)
   p.s++
   p.l.s++
   p.l.l = tree.allocate(Node{x: x})
}
func (tree *Tree) attachRR(p *Node, x Data) {
   tree.copy(&p.r)
   p.r.r = tree.allocate(Node{x: x})
}
func (tree *Tree) attachLR(p *Node, x Data) {
   tree.copy(&p.l)
   p.s++
   p.l.r = tree.allocate(Node{x: x})
}

func (tree *Tree) attachRL(p *Node, x Data) {
   tree.copy(&p.r)
   p.r.s++
   p.r.l = tree.allocate(Node{x: x})
}

func (tree *Tree) attachR(p *Node, x Data) {
   p.r = tree.allocate(Node{x: x})
}

func pathDeletingRightIgnoringIndex(p *Node) **Node {
   return &p.r
}

func (tree Tree) rotateL(p **Node) {
   tree.copy(&(*p).r)
   *p = (*p).rotateL()
}

func (tree Tree) rotateR(p **Node) {
   tree.copy(&(*p).l)
   *p = (*p).rotateR()
}

func (tree Tree) rotateRL(p **Node) {
   tree.copy(&(*p).r)
   tree.copy(&(*p).r.l)
   *p = (*p).rotateRL()
}

func (tree Tree) rotateLR(p **Node) {
   tree.copy(&(*p).l)
   tree.copy(&(*p).l.r)
   *p = (*p).rotateLR()
}

// // Creates a new path starting at p* by descending to the left until there is
// // no further left child, returning a pointer to the last node on that path.
//
//   func followL(p **Node) **Node {
//      if *p != nil {
//         for (*p).hasL() {
//            p = &(shadow(p)).l
//         }
//      }
//      return p
//   }
//
// // Creates a new path starting at p* by descending to the right until there is
// // no further right child, returning a pointer to the last node on that path.
//
//   func followR(p **Node) **Node {
//      if *p != nil {
//         for (*p).hasR() {
//            p = &(shadow(p)).r
//         }
//      }
//      return p
//   }
//
// // Creates a new path starting at p* by descending to the right until there is
// // no right child, then appends n* to the right of the last node on the path.
// //
// //    1. p* → ◯ → ◯ → ◯
// //    2. p* → ◯ → ◯ → ◯ → n*
// //
func (tree *Tree) appendR(p **Node, n *Node) {
   for *p != nil {
      tree.copy(p)
      p = &(*p).r
   }
   *p = n
}

// // Creates a new path starting at p* by descending to the left until there is
// // no left child, then appends n* to the left of the last node on the path.
func (tree *Tree) appendL(p **Node, n *Node) {
   for *p != nil {
      tree.copy(p)
      p = &(*p).l
   }
   *p = n
}

//// TODO put this under linked list ?
//func (tree *Tree) truncateL(p **Node) *Node {
//   // assert(*p != nil)
//   tree.copy(p)
//   for (*p).l != nil {
//      p = &(*p).l
//      tree.copy(p)
//   }
//   q := *p
//   *p = nil
//   return q
//}
//
//func (tree *Tree) truncateR(p **Node) *Node {
//   // assert(*p != nil)
//   tree.copy(p)
//   for (*p).r != nil {
//      p = &(*p).r
//      tree.copy(p)
//   }
//   q := *p
//   *p = nil
//   return q
//}

//
//func flipSize(p *Node, s Size) {
//   p.i = s - p.i - 1
//}

//
//// Prepends p* to a right path at g**, returning the previous right child of p*.
////
////    1. p* → r*                        p* has a right child r*.
////
////       g** → ◯ → ◯ → ◯                g** is the head of a path going right.
////
////    2. g** → p* → ◯ → ◯ → ◯           p* is prepended to the path at g**, and
////                                      g** is adjusted to have p* as the head.
////    3. Return r*
////
//func prependR(p *Node, g **Node) (r *Node) {
//   r, p.r, *g = p.r, *g, p
//   return
//}
//
//// Prepends p* to a left path at g**, returning the previous left child of p*.
//func prependL(p *Node, g **Node) (l *Node) {
//   l, p.l, *g = p.l, *g, p
//   return
//}

// Reverses the order of the nodes along the right spine of p*.
// Example:
//
//       p*
//       ↓
//      (A) → (B) → (C) → (D) → nil
//
// A given node r* replaces the nil right child of the last node on the path as
// it becomes the right child of p*, the initial head of the path.
//
// Returns the last node of the initial path, ie. the head of the reversed path.
//

//
//                        (A) → r*     Prepend (A) to  r*, p* becomes (B)
//                  (B) → (A) → r*     Prepend (B) to (A), p* becomes (C)
//            (C) → (B) → (A) → r*     Prepend (C) to (B), p* becomes (D)
//      (D) → (C) → (B) → (A) → r*     Prepend (D) to (C), p* becomes nil
//       ↑
//       return
//
//func reverseR(p *Node, r *Node) *Node {
//   for p != nil {
//      p = p.Copy()
//       p = prependR(p, &r).copyOrNil()
//   }
//   return r
//}
//
//// Reverses the order of the nodes along the left spine of p*.
//func reverseL(p *Node, l *Node) *Node {
//   for p != nil {
//       p = p.Copy()
//       p = prependL(p, &l).copyOrNil()
//   }
//   return l
//}
//
// func (tree FingerTree) toTreap() *Treap {
//    if tree.root == nil {
//       return &tree.Treap
//    }
//    tree.root = tree.root.Copy()
//    tree.root.l = reverseL(tree.root.l.copyOrNil(), nil)
//    tree.root.r = reverseR(tree.root.r.copyOrNil(), nil)
//
//    p := tree.root.l // this is now the shoulder as connected, but they are all sizeL
//    s := tree.root.s
//    for p != nil {
//       p.s = s - p.s - 1
//       s = p.s
//       p = p.l
//    }
//    return &tree.Treap
// }
//
