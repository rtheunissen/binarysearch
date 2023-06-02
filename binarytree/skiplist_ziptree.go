package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/random"
   "math/bits"
)

type Zip struct {
   Tree
   random.Source
}

func (Zip) New() List {
   return &Zip{
      Source: random.New(random.Uint64()),
   }
}

func (tree *Zip) Clone() List {
   return &Zip{
   	Tree:   tree.Tree.Clone(),
   	Source: tree.Source, // TODO: copy?
   }
}

func (tree *Zip) randomRank() uint64 {
   return uint64(bits.LeadingZeros64(tree.Uint64()))
}

func (tree *Zip) unzip(p *Node, i Position, l, r **Node) {
   //p.partition(i, &l, &r)
   for p != nil {
   	tree.copy(&p)
   	if i <= p.s {
   		*r = p
   		p.s = p.s - i
   		r = &p.l
   		p = *r
   	} else {
   		*l = p
   		i = i - p.s - 1
   		l = &p.r
   		p = *l
   	}
   }
   *l = nil
   *r = nil
}

// When the new node's rank is greater than the rank of the current node,
// we know for sure that we can insert the new node at the current level.
//
// Otherwise, the new rank is less than or equal to the current rank.
//
//	When branching LEFT: if the ranks are equal, a split at the current
//	node would make it the right child of the new node, where an equal
//	rank would be valid. Keep searching if the new rank is less than.
//
//	When branching RIGHT: if the ranks are equal, a split at the current
//	node would make it the left child of the new node, where an equal
//	rank would NOT be valid.
func (tree *Zip) Insert(i Position, x Data) {
   tree.size++

   p := &tree.root                                      // parent, pointer
   n := tree.allocate(Node{x: x, y: tree.randomRank()}) // new node

   for *p != nil {
   	if n.y > (*p).y { // New rank is greater, insert here.
   		break
   	}
   	if i <= (*p).s {
   		if n.y == (*p).y { // Branching left and ranks are equal.
   			break
   		}
   		tree.copy(p)
   		(*p).s = (*p).s + 1 // Increase the size of the left subtree.
   		p = &(*p).l         // Path left.
   	} else {
   		tree.copy(p)
   		i = i - ((*p).s + 1) // Skip the current node and left subtree.
   		p = &(*p).r          // Path right.
   	}
   }
   // assert(rank(n) >= rank(*p))
   tree.unzip(*p, i, &n.l, &n.r) // Unzip the path into the new node.
   n.s = i                       // Set the size of the left subtree.
   *p = n                        // Write the new node to the path.
}

func (tree *Zip) zip(l, r *Node, sl Size) (root *Node) {
   // assert(sl == l.size())
   p := &root
   for {
   	if l == nil {
   		*p = r
   		return
   	}
   	if r == nil {
   		*p = l
   		return
   	}

   	if rank(l) >= rank(r) {
   		tree.copy(&l)
   		sl = sl - l.s - 1 //l.sizeR(sl)
   		*p = l
   		p = &l.r
   		l = *p
   	} else {
   		tree.copy(&r)
   		r.s = r.s + sl
   		*p = r
   		p = &r.l
   		r = *p
   	}
   }
}

//
//func (tree *Zip) zipRecursive(x, y *Node, sl Size) *Node {
//   // assert(sl == x.count())
//   if x == nil { return y }
//   if y == nil { return x }
//   if rank(x) < rank(y) {
//      tree.pathcopy(&y)
//      y.i = y.i + sl
//      y.l = tree.zipRecursive(x, y.l, sl)
//      return y
//   } else {
//      tree.pathcopy(&x)
//      sl = sl - x.i - 1
//      x.r = tree.zipRecursive(x.r, y, sl)
//      return x
//   }
//}
//
//func (tree *Zip) deleteRecursive(p *Node, i Position, x *Value) *Node {
//   if p == nil {
//      return nil
//   }
//   tree.pathcopy(&p)
//   if i == p.i {
//      // assert(p.i == p.l.count())
//      *x = p.x
//      return tree.zipRecursive(p.l, p.r, p.i)
//   }
//   if i < p.i {
//      p.i = p.i - 1
//      if i == p.l.i {
//         *x = p.l.x
//         tree.pathcopy(&p.l)
//         tree.pathcopy(&p.l.r)
//         tree.pathcopy(&p.l.l)
//         p.l = tree.zipRecursive(p.l.l, p.l.r, p.l.i)
//         return p
//      } else {
//         p.l = tree.deleteRecursive(p.l, i, x)
//         return p
//      }
//   } else {
//      if (i - (p.i + 1)) == p.r.i {
//         *x = p.r.x
//         tree.pathcopy(&p.r)
//         tree.pathcopy(&p.r.r)
//         tree.pathcopy(&p.r.l)
//         p.r = tree.zipRecursive(p.r.l, p.r.r, p.r.i)
//         return p
//      } else {
//         p.r = tree.deleteRecursive(p.r, i - (p.i + 1), x)
//         return p
//      }
//   }
//}

func (tree *Zip) delete(p **Node, i Position, x *Data) {
   for {
   	if i == (*p).s {
   		//
   		// Found the node to delete.
   		//
   		defer tree.release(*p)
   		*x = (*p).x
   		if (*p).l == nil && (*p).r == nil {
   			*p = nil
   		} else {
   			tree.copy(p)
   			*p = tree.zip((*p).l, (*p).r, (*p).s)
   		}
   		return
   	}
   	tree.copy(p)
   	if i < (*p).s {
   		(*p).s = (*p).s - 1 // Decrease the size of the left subtree.
   		p = &(*p).l         // Path left.
   	} else {
   		i = i - ((*p).s + 1) // Skip the current node and left subtree.
   		p = &(*p).r          // Path right.
   	}
   }
}

func (tree *Zip) Delete(i Position) (x Data) {
   // assert(i < tree.Size())
   //tree.pathcopy(&tree.root)
   //tree.root = tree.deleteRecursive(tree.root, i, &x)
   tree.delete(&tree.root, i, &x)
   tree.size--
   return
}

func (tree Zip) split(i Size) (Tree, Tree) {
   // assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := tree.Tree.split(tree.root, i)

   return Tree{arena: tree.arena, root: l, size: i},
   	Tree{arena: tree.arena, root: r, size: tree.size - i}
}

func (tree *Zip) Split(i Position) (List, List) {
   l, r := tree.split(i)

   return &Zip{l, tree.Source},
   	&Zip{r, tree.Source}
}

func (tree *Zip) Select(i Size) Data {
   // assert(i < tree.Size())
   return tree.lookup(tree.root, i)
}

func (tree *Zip) Update(i Size, x Data) {
   // assert(i < tree.Size())
   tree.copy(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *Zip) Join(that List) List {
   tree.share(tree.root)
   tree.share(that.(*Zip).root)

   root := tree.zip(tree.root, that.(*Zip).root, tree.size)
   size := tree.size + that.(*Zip).size

   return &Zip{Tree{arena: tree.arena, root: root, size: size}, tree.Source}
}

func (tree *Zip) verifyRanks(p *Node) {
   if p == nil {
   	return
   }
   invariant(p.l == nil || p.y > p.l.y)
   invariant(p.r == nil || p.y > p.r.y || p.y == p.r.y)

   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}

//func (tree *Zip) maxHeapify(p *Node) {
//   if p == nil {
//      return
//   }
//   tree.maxHeapify(p.l)
//   tree.maxHeapify(p.r)
//
//   m := p
//   if p.l != nil && rank(p.l) > rank(m) { m = p.l }
//   if p.r != nil && rank(p.r) > rank(m) { m = p.r }
//   if m != p {
//      y := p.y
//      p.y = m.y
//      m.y = y
//      tree.maxHeapify(m)
//   } else {
//      // p is either greater than or equal to both left and right children
//      // so we need to increase p
//      p.y++
//   }
//}

func (tree *Zip) Verify() {
   tree.verifySizes()
   tree.verifyRanks(tree.root)
}

//func (tree *Zip) assignRanks(p *Node) {
//   if p == nil {
//      return
//   }
//   p.y = tree.randomRank()
//   tree.assignRanks(p.l)
//   tree.assignRanks(p.r)
//   return
//  //if p == nil {
//  //   return
//  //}
//  //tree.assignRanks(p.l)
//  //tree.assignRanks(p.r)
//  //
//  //if p.l == nil && p.r == nil {
//  //   p.y = 0
//  //} else {
//  //   if rank(p.l) == rank(p.r) {
//  //      p.y = rank(p.l) + 1
//  //   } else if rank(p.l) < rank(p.r) {
//  //      p.y = rank(p.r) + (tree.Random.Uint64() & 1)
//  //   } else {
//  //      p.y = rank(p.l) + 1
//  //   }
//  //}
//}
