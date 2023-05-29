package binarytree

import (
	"trees/random"
)

type TreapFingerTree struct {
	Tree
	random.Source
}

func (TreapFingerTree) New() List {
	return &TreapFingerTree{
		Source: random.New(random.Uint64()),
	}
}

func (tree *TreapFingerTree) Clone() List {
	return &TreapFingerTree{
		Tree:   tree.Tree.Clone(),
		Source: tree.Source,
	}
}

//func (p *Node) flip(s Size) {
//   p.s = s - p.s - 1
//}

//func (p *Node) flipL(s Size) {
//   assert(s == p.count())
//   for p != nil {
//      p = p.Copy()
//      sl := p.s
//      p.s = s - p.s - 1
//      p = p.l
//      s = sl
//   }
//}

func (tree *TreapFingerTree) reverseL(p *Node, g *Node, s Size) *Node {
	assert(s == p.size())
	for {
		if p == nil {
			return g
		}
		tree.copy(&p)
		sl := p.s
		p.s = s - p.s - 1
		l := p.l
		p.l = g
		g = p
		p = l
		s = sl
	}
}

func (tree *TreapFingerTree) reverseR(p *Node, g *Node) *Node {
	for {
		if p == nil {
			return g
		}
		tree.copy(&p)
		r := p.r
		p.r = g
		g = p
		p = r
	}
}

//
//func (tree TreapTopDown) toTreapFingerTree() *TreapFingerTree {
//   ftree := TreapFingerTree{
//      TreapTopDown: tree,
//   }
//   ftree.Source = tree.Source
//   ftree.arena = tree.arena
//   ftree.root = tree.root
//   ftree.size = tree.size
//
//   if tree.root == nil {
//      return &ftree
//   }
//   tree.pathcopy(&tree.root)
//   ftree.root.l = ftree.reverseL(shared(tree.root.l), nil, tree.root.i)
//   ftree.root.r = ftree.reverseR(shared(tree.root.r), nil)
//   return &ftree
//}

func (tree *TreapFingerTree) randomRank() uint64 {
	return tree.Uint64()
}

func (tree *TreapFingerTree) rotateParentLeftOnRightSpine(p *Node) {
	tree.copy(&p.r)
	r := p.r // parent on the spine
	p.r = r.r
	r.r = p.l
	p.l = r
	p.s = p.s + r.s + 1
	// measurement(&rotations, 1)
}

func (tree *TreapFingerTree) rotateParentRightOnLeftSpine(p *Node) {
	tree.copy(&p.l)
	l := p.l
	p.l = l.l
	l.l = p.r
	p.r = l
	p.s = p.s + l.s + 1
	l.s = p.s - l.s - 1
	// measurement(&rotations, 1)
}

func (tree *TreapFingerTree) rotateRightIntoRoot(l *Node) {
	assert(l.l == nil)

	p := tree.root

	tree.appendR(&p.r, p)

	l.l = p.l
	p.l = l.r
	l.r = p.r
	p.r = nil
	l.s = p.s - l.s - 1
	p.s = p.s - l.s - 1

	tree.root = l // TODO: consider returning this, accepting p?, not tree
	// measurement(&rotations, 1)
}

func (tree *TreapFingerTree) rotateLeftIntoRoot(r *Node) {
	assert(r.r == nil)

	p := tree.root

	tree.appendL(&p.l, p)

	r.r = p.r
	p.r = r.l
	r.l = p.l
	p.l = nil
	r.s = r.s + p.s + 1
	p.s = r.s - p.s - 1

	tree.root = r
	// measurement(&rotations, 1)
}

func (tree *TreapFingerTree) rotateUpR(p *Node) *Node {
	for {
		if p.r == nil {
			if p.y > tree.root.y {
				tree.rotateLeftIntoRoot(p)
				return nil
			}
		} else {
			if p.y > p.r.y {
				tree.rotateParentLeftOnRightSpine(p)
				continue
			}
		}
		return p
	}
}

func (tree *TreapFingerTree) rotateUpL(p *Node) *Node {
	for {
		if p.l == nil {
			if p.y > tree.root.y {
				tree.rotateRightIntoRoot(p)
				return nil
			}
		} else {
			if p.y > p.l.y {
				tree.rotateParentRightOnLeftSpine(p)
				continue
			}
		}
		return p
	}
}

func (tree *TreapFingerTree) rotateDownL(p *Node) {
	for p.r != nil && p.r.y > p.y {
		tree.copy(&p.r)
		r := p.r
		p.r = r.l
		r.l = p.l
		p.l = r
		r.s = p.s - r.s - 1
		p.s = p.s - r.s - 1
	}
	// measurement(&rotations, 1)
}

func (tree *TreapFingerTree) rotateDownR(p *Node) {
	for p.l != nil && p.l.y > p.y {
		tree.copy(&p.l)
		l := p.l
		p.l = l.r
		l.r = p.r
		p.r = l
		p.s = p.s - l.s - 1
	}
	// measurement(&rotations, 1)
}

//	func (tree *TreapFingerTree) setRoot(p *Node) {
//	  tree.root = p
//	}
func (tree *TreapFingerTree) setHead(p *Node) {
	tree.root.l = p
}

//	func (tree *TreapFingerTree) setSize(s Size) {
//	  tree.size = s
//	}
//
//	func (tree *TreapFingerTree) setHeadSize(s Size) {
//	  tree.root.i = s
//	}
func (tree *TreapFingerTree) setTail(p *Node) {
	tree.root.r = p
}
func (tree *TreapFingerTree) getHead() (p *Node) {
	return tree.root.l
}
func (tree *TreapFingerTree) getTail() (p *Node) {
	return tree.root.r
}

func (tree *TreapFingerTree) insertAsLast(x Data) {
	tree.copy(&tree.root)
	p := tree.allocate(Node{x: x, y: tree.randomRank()})
	p.r = tree.getTail()
	p = tree.rotateUpR(p)
	tree.setTail(p)
	tree.size++
}

func (tree *TreapFingerTree) insertAsFirst(x Data) {
	tree.copy(&tree.root)
	p := tree.allocate(Node{x: x, y: tree.randomRank()})
	p.l = tree.getHead()
	p = tree.rotateUpL(p)
	tree.setHead(p)
	tree.root.s++
	tree.size++
}

func (tree TreapFingerTree) Select(i Position) Data {
	assert(i < tree.Size())
	switch {
	case i < tree.root.s:
		return tree.accessFromHead(i)
	case i > tree.root.s:
		return tree.accessFromTail(tree.size - i - 1)
	default:
		return tree.root.x
	}
}

func (tree *TreapFingerTree) Update(i Position, x Data) {
	assert(i < tree.Size())
	switch {
	case i < tree.root.s:
		tree.updateFromHead(x, i)
	case i > tree.root.s:
		tree.updateFromTail(x, tree.size-i-1)
	default:
		tree.copy(&tree.root)
		tree.root.x = x
	}
}

func (tree TreapFingerTree) accessFromHead(i Position) Data {
	p := tree.getHead()
	for {
		if i == 0 {
			return p.x
		}
		if i > p.s {
			i = i - p.s - 1
			p = p.l
		} else {
			return tree.lookup(p.r, i-1)
		}
	}
}

func (tree *TreapFingerTree) updateFromTail(x Data, i Position) {
	tree.copy(&tree.root)
	tree.copy(&tree.root.r)

	p := tree.root.r
	for {
		if i == 0 {
			p.x = x
			return
		}
		if i > p.s {
			tree.copy(&p.r)
			i = i - p.s - 1
			p = p.r
		} else {
			tree.copy(&p.l)
			tree.update(p.l, p.s-i, x)
			return
		}
	}
}

func (tree *TreapFingerTree) updateFromHead(x Data, i Position) {
	tree.copy(&tree.root)
	tree.copy(&tree.root.l)

	p := tree.root.l
	for {
		if i == 0 {
			p.x = x
			return
		}
		if i > p.s {
			tree.copy(&p.l)
			i = i - p.s - 1
			p = p.l
		} else {
			tree.copy(&p.r)
			tree.update(p.r, i-1, x)
			return
		}
	}
}

func (tree TreapFingerTree) accessFromTail(i Position) Data {
	p := tree.getTail()
	for {
		if i == 0 {
			return p.x
		}
		if i > p.s {
			i = i - p.s - 1
			p = p.r
		} else {
			return tree.lookup(p.l, p.s-i)
		}
	}
}

func (tree *TreapFingerTree) insert(p **Node, i Position, n *Node) {
	for {
		if *p == nil {
			*p = n
			return
		}
		if (*p).y <= n.y {
			n.l, n.r = tree.Tree.split(*p, i)
			n.s = i
			*p = n
			return
		}
		tree.copy(p)
		if i <= (*p).s {
			p = insertL(*p)
		} else {
			p = insertR(*p, &i)
		}
	}
}

func (tree *TreapFingerTree) insertFromHead(x Data, i Position) {
	tree.copy(&tree.root)
	tree.copy(&tree.root.l)
	tree.root.s++
	tree.size++

	n := tree.allocate(Node{x: x, y: tree.randomRank()})
	p := tree.root.l
	for {
		//
		if i > p.s {
			tree.copy(&p.l)
			i = i - p.s - 1
			p = p.l
			continue
		}
		//
		if rank(n) > rank(p) {
			p.r, n.r = tree.Tree.split(p.r, i)
			n.s = p.s - i
			p.s = i
			n.l = p.l
			p.l = tree.rotateUpL(n)
			return
		}
		//
		p.s++
		tree.insert(&p.r, i, n)
		return
	}
}

func (tree TreapFingerTree) split(i Position) (Tree, Tree) {
	assert(i <= tree.size)
	tree.share(tree.root)
	if i == 0 {
		return Tree{arena: tree.arena},
			Tree{arena: tree.arena, root: tree.root, size: tree.size}
	}
	if i == tree.size {
		return Tree{arena: tree.arena, root: tree.root, size: tree.size},
			Tree{arena: tree.arena}
	}
	if i <= tree.root.s {
		return tree.splitFromHead(i)
	} else {
		return tree.splitFromTail(i)
	}
}

func (tree TreapFingerTree) Split(i Position) (List, List) {
	assert(i <= tree.size)
	l, r := tree.split(i)
	return &TreapFingerTree{Tree: l, Source: tree.Source},
		&TreapFingerTree{Tree: r, Source: tree.Source}
}

func (tree TreapFingerTree) splitFromHead(i Position) (Tree, Tree) {
	assert(i <= tree.size)
	assert(i <= tree.root.s)
	/////

	// TODO: I think there is a bug here when i == n

	//////
	tree.copy(&tree.root)

	p := tree.root
	d := i
	for d > (p.l.s + 1) {
		d = d - (p.l.s + 1)
		tree.copy(&p.l)
		p = p.l
	}
	tree.copy(&p.l)
	g := p.l
	p.l = nil
	p = g
	g = p.l

	sl := d - 1
	sr := p.s - sl

	l, r := tree.Tree.split(p.r, sl)

	L := Tree{arena: tree.arena, root: p, size: i}
	R := Tree{arena: tree.arena, root: tree.root, size: tree.size - i}

	L.root.s = i - d
	L.root.l = tree.root.l
	L.root.r = tree.reverseR(l, nil)

	R.root.s = tree.root.s - i
	R.root.r = tree.root.r
	R.root.l = tree.reverseL(r, g, sr)

	return L, R
}

func (tree TreapFingerTree) splitFromTail(i Position) (Tree, Tree) {
	assert(i < tree.size)
	assert(i > tree.root.s)

	tree.copy(&tree.root)

	p := tree.root
	d := tree.size - i

	for d > (p.r.s + 1) {
		d = d - (p.r.s + 1)
		tree.copy(&p.r)
		p = p.r
	}
	tree.copy(&p.r)
	g := p.r
	p.r = nil
	p = g
	g = p.r

	sr := d - 1
	sl := p.s - sr

	l, r := tree.Tree.split(p.l, sl)

	L := Tree{arena: tree.arena}
	R := Tree{arena: tree.arena}

	R.root = p
	R.size = tree.size - i
	R.root.s = sr
	R.root.l = tree.reverseL(r, nil, sr)
	R.root.r = tree.root.r

	L.root = tree.root
	L.size = i
	L.root.l = tree.getHead()
	L.root.r = tree.reverseR(l, g)

	return L, R
}

//
//
//

func (tree *TreapFingerTree) insertFromTail(x Data, i Position) {
	tree.copy(&tree.root)
	tree.copy(&tree.root.r)

	tree.size++

	n := tree.allocate(Node{x: x, y: tree.randomRank()})
	p := tree.root.r
	for {
		//
		if i > p.s {
			tree.copy(&p.r)
			i = i - p.s - 1
			p = p.r
			continue
		}
		//
		if rank(n) > rank(p) {

			n.l, p.l = tree.Tree.split(p.l, p.s-i)
			n.s = p.s - i
			p.s = i
			n.r = p.r
			p.r = tree.rotateUpR(n)
			return
		}
		//
		p.s++
		tree.insert(&p.l, p.s-i-1, n)
		return
	}
}

// create new node with random rank
// if this rank is less than the root's rank, we need a new root node and the
//
//	old root node will
func (tree *TreapFingerTree) Insert(i Position, x Data) {
	assert(i <= tree.Size())
	if tree.root == nil {
		tree.root = tree.allocate(Node{x: x, y: tree.randomRank()})
		tree.size = 1
		return
	}
	if i <= tree.root.s {
		if i == 0 {
			tree.insertAsFirst(x)
		} else {
			tree.insertFromHead(x, i-1)
		}
	} else {
		if i == tree.size {
			tree.insertAsLast(x)
		} else {
			tree.insertFromTail(x, tree.size-i-1)
		}
	}
}

////
//func (t *TreapFingerTreeOld) removeFromHead(i int) (s Data) {
//  invariant(i >= 0)
//
//  //
//  //
//  if i == 0 {
//     return t.Shift()
//  }
//
//  //
//  t.size = t.size - 1
//  t.root = t.root.copy().withRelativePosition(t.root.w - 1)
//  t.head = t.head.copy()
//
//  //
//  p := t.head
//  d := i
//
//  //
//  //
//  for {
//     //visualize("↗")
//     invariant(d > 0)
//     invariant(p.isL())
//
//     //
//     //
//     if d + p.w < 0 {
//        //visualize("|↘")
//        p.decreaseInternalWeightOfL()
//        return (&Treap{}).remove(&p.r, d, p.sizeLR()+1).s
//     }
//
//     //
//     //
//     if d + p.w == 0 {
//        dissolveOnLeftSpine(p, &s)
//        return
//     }
//
//     //
//     //
//     d = d + p.w
//     p = shadow(&p.l)
//  }
//}
//
////
////
//func (t *TreapFingerTreeOld) removeFromTail(i int) (s Data) {
//  invariant(i >= 0)
//
//  //
//  if i == t.size - 1 {
//     return t.Pop()
//  }
//
//  //
//  t.size = t.size - 1
//  t.tail = t.tail.copy()
//
//  //
//  p := t.tail
//  d := i - t.size
//
//  //
//  //
//  for {
//     //visualize("↖")
//     invariant(d < 0)
//     invariant(p.isR())
//
//     //
//     //
//     if d + p.w > 0 {
//        //visualize("|↙")
//        p.decreaseInternalWeightOfR()
//        return (&Treap{}).remove(&p.l, d, p.sizeRL()+1).s
//     }
//
//     //
//     //
//     if d + p.w == 0 {
//        dissolveOnRightSpine(p, &s)
//        return
//     }
//
//     //
//     //
//     d = d + p.w
//     p = shadow(&p.r)
//  }
//}
//
////
////
//func dissolveOnLeftSpine(l *Node, s *Data) {
//  //visualize("!")
//
//  p  := copyOf(l.l)
//  g  := p.l
//  p.l = l
//  l.l = nil
//  *s  = p.s
//  *l  = *reverseL((&Treap{}).dissolveWithoutSize(p), g)
//}
//
////
////
//func dissolveOnRightSpine(r *Node, s *Data) {
//  //visualize("!")
//  invariant(r.isR())
//
//  p  := r.r.copy()
//  g  := p.r
//  p.r = r
//  r.r = nil
//  *s  = p.s
//  *r  = *reverseR((&Treap{}).dissolveWithoutSize(p), g)
//}
//

func (tree *TreapFingerTree) deleteFirst(x *Data) {
	//println("deleteFirst")
	defer tree.release(tree.root.l)
	*x = tree.root.l.x
	tree.copy(&tree.root)
	tree.copy(&tree.root.l)
	tree.root.l = tree.reverseL(tree.root.l.r, tree.root.l.l, tree.root.l.s)
	tree.root.s--
}

func (tree *TreapFingerTree) deleteLast(x *Data) {
	//println("deleteLast")
	defer tree.release(tree.root.r)
	*x = tree.root.r.x
	tree.copy(&tree.root)
	tree.copy(&tree.root.r)
	tree.root.r = tree.reverseR(tree.root.r.l, tree.root.r.r)
}
func (tree TreapFingerTree) delete(p **Node, i Position, x *Data) {
	for {
		if i == (*p).s {
			*x = (*p).x
			if (*p).l == nil && (*p).r == nil {
				*p = nil
			} else {
				tree.copy(p) // TODO: should we instead share the left and right here?
				//defer tree.release(*p)
				*p = tree.join((*p).l, (*p).r, (*p).s)
			}
			return
		}
		tree.copy(p)
		if i < (*p).s {
			p = deleteL(*p)
		} else {
			p = deleteR(*p, &i)
		}
	}
}
func (tree *TreapFingerTree) deleteFromHead(i Position, x *Data) {
	//println("deleteFromHead")
	if i == 0 {
		tree.deleteFirst(x)
		return
	}
	tree.copy(&tree.root)
	tree.copy(&tree.root.l)
	tree.root.s--
	p := tree.root.l
	for i > p.s+1 {
		tree.copy(&p.l)
		i = i - p.s - 1
		p = p.l
	}
	if i < p.s+1 {
		tree.delete(&p.r, i-1, x)
		p.s--
		return
	}
	tree.copy(&p.l)
	g := p.l
	defer tree.release(g)
	*x = g.x
	p.r = tree.join(p.r, g.r, p.s)
	p.l = g.l
	p.s = p.s + g.s
	tree.rotateDownL(p)
}

func (tree *TreapFingerTree) join(l, r *Node, sl Size) (root *Node) {
	assert(sl == l.size())
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
		if l.y >= r.y {
			tree.copy(&l)
			sl = sl - l.s - 1
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

func (tree *TreapFingerTree) deleteFromTail(i Position, x *Data) {
	//println("deleteFromTail")
	if i == tree.size-1 {
		tree.deleteLast(x)
		return
	}
	i = tree.size - i - 1
	tree.copy(&tree.root)
	tree.copy(&tree.root.r)
	p := tree.root.r
	for i > p.s+1 {
		tree.copy(&p.r) // TODO I think this could be nulll?
		i = i - p.s - 1
		p = p.r
	}
	//println("found inflection", p.s)
	if i < p.s+1 {
		//println("descend")
		tree.delete(&p.l, p.s-i, x)
		p.s--
		return
	}
	//println("delete on the right spine")
	tree.copy(&p.r)
	g := p.r
	defer tree.release(g)
	*x = g.x
	p.l = tree.join(g.l, p.l, g.s)
	p.r = g.r
	p.s = p.s + g.s
	tree.rotateDownR(p)
}

//func (tree TreapFingerTree) toTreap() (treap TreapTopDown) {
//   treap = tree.TreapTopDown
//   if treap.root == nil {
//      return
//   }
//   tree.root.l = tree.reverseL(shared(tree.root.l), nil, tree.root.i)
//   tree.root.r = tree.reverseR(shared(tree.root.r), nil)
//   return
//}

func (tree *TreapFingerTree) reverseL2(p *Node, g *Node) *Node {
	s := Size(0)
	for {
		if p == nil {
			return g
		}
		tree.copy(&p)
		s = s + p.s + 1
		p.s = s - p.s - 1
		l := p.l
		p.l = g
		g = p
		p = l
	}
}
func (tree *TreapFingerTree) deleteRoot(v *Data) {
	//println("deleteRoot")
	tree.copy(&tree.root)
	*v = tree.root.x

	// To treap
	tree.root.l = tree.reverseL2(tree.root.l, nil)
	tree.root.r = tree.reverseR(tree.root.r, nil)

	// Dissolve root
	tree.root = tree.join(tree.root.l, tree.root.r, tree.root.s)

	if tree.root == nil {
		return
	}
	// To finger tree
	//tree.pathcopy(&tree.root)
	tree.root.l = tree.reverseL(tree.root.l, nil, tree.root.s)
	tree.root.r = tree.reverseR(tree.root.r, nil)
}

func (tree *TreapFingerTree) Delete(i Position) (v Data) {
	assert(i < tree.size)
	switch {
	case i < tree.root.s:
		tree.deleteFromHead(i, &v)
		tree.size--
	case i > tree.root.s:
		tree.deleteFromTail(i, &v)
		tree.size--
	default:
		tree.deleteRoot(&v)
		tree.size--
	}
	return
}

//func (t *TreapFingerTreeOld) Remove(i int) (s Data){
//  switch {
//  case i < t.root.sizeRL(): return t.removeFromHead(i)
//  case i > t.root.sizeRL(): return t.removeFromTail(i)
//  default:
//     return t.removeRoot()
//  }
//}
//
//
//func (tree *FingerTreeDisjointUniform) Delete(i Position) (s Data) {
//   if i < tree.head.size {
//      if i == 0 {
//         tree.shift(&s)
//         return s
//      }
//      tree.deleteFromHead(i, &s)
//      tree.distribute()
//      return s
//   } else {
//      if i == tree.head.size + tree.tail.size - 1 {
//         tree.pop(&s)
//         return s
//      }
//      tree.deleteFromTail(tree.head.size + tree.tail.size - i - 1, &s)
//      tree.distribute()
//      return s
//   }
//}
//
//func (tree *FingerTreeDisjointUniform) deleteFromHead(i Position, s *Data) {
//
//   tree.head.root = pathcopy(tree.head.root)
//   tree.head.size--
//   p := tree.head.root
//   for i > p.s + 1 {
//      p.l = p.l.Copy()
//      i = i - p.s - 1
//      p = p.l
//   }
//   if i < p.s + 1 {
//      Treap{}.delete(&p.r, i - 1, s)
//      p.s--
//      return
//   }
//   p.l = p.l.Copy()
//   g := p.l
//   *s = g.s
//   p.r = Treap{}.join(p.r, g.r, p.s)
//   p.l = g.l
//   p.s = p.s + g.s
//   tree.rotateDownL(p)
//}
//func (tree *FingerTreeDisjointUniform) deleteFromTail(i Position, s *Data) {
//
//   tree.tail.root = pathcopy(tree.tail.root)
//   tree.tail.size--
//   p := tree.tail.root
//   for {
//      if i == p.s + 1 {
//         p.r = p.r.Copy()
//         g := p.r
//         *s = g.s
//         p.l = Treap{}.join(g.l, p.l, g.s)
//         p.r = g.r
//         p.s = p.s + g.s
//         tree.rotateDownR(p)
//         return
//      }
//      if i < p.s + 1 {
//         Treap{}.delete(&p.l, p.s - i, s)
//         p.s--
//         return
//      }
//      i = i - p.s - 1
//      p.r = p.r.Copy()
//      p = p.r
//   }
//}
//func (tree *FingerTreeDisjointUniform) pop(s *Data) {
//
//   *s = tree.tail.root.s
//
//   tree.tail.root = tree.tail.root.copy()
//   tree.tail.root = tree.reverseR(tree.tail.root.l, tree.tail.root.r)
//   tree.tail.size = tree.tail.size - 1
//   tree.distribute()
//}
//
//
//func (tree *FingerTreeDisjointUniform) shift(s *Data) {
//
//   *s = tree.head.root.s
//
//   tree.head.root = tree.head.root.copy()
//   tree.head.root = tree.reverseL(tree.head.root.r, tree.head.root.l, tree.head.root.s)
//   tree.head.size = tree.head.size - 1
//   tree.distribute()
//}

func (tree *TreapFingerTree) joinUp(o *TreapFingerTree) *Node {
	tree.copy(&tree.root)
	tree.copy(&o.root)

	l := tree.root.r
	r := o.root.l
	s := Size(0) // size of p

	var p *Node
	for {
		if r == nil {
			tree.root.r = o.root.r
			o.root.l = p
			o.root.r = l
			o.root.s = s

			tree.appendR(&tree.root.r, tree.rotateUpR(o.root))
			return tree.root
		}
		if l == nil {
			o.root.l = tree.root.l
			o.root.s = tree.size + o.root.s
			tree.root.l = r
			tree.root.r = p
			tree.root.s = s

			tree.appendL(&o.root.l, tree.rotateUpL(tree.root))
			return o.root
		}
		if l.y < r.y { // TODO: how does <= affect things? Should we prefer larger size?
			tree.copy(&l)
			s = s + l.s + 1
			g := l.r
			l.r = p
			p = l
			l = g
		} else {
			tree.copy(&r)
			s = s + r.s + 1
			r.s = s - r.s - 1
			g := r.l
			r.l = p
			p = r
			r = g
		}
	}
}

func (tree *TreapFingerTree) Join(that List) List {
	if tree.Size() == 0 {
		return that.Clone()
	} // TODO: can we avoid this?
	if that.Size() == 0 {
		return tree.Clone()
	} // TODO: can we avoid this?

	l := tree.Clone().(*TreapFingerTree) // TODO: can we avoid this?
	r := that.Clone().(*TreapFingerTree) // TODO: can we avoid this?

	return &TreapFingerTree{
		Tree: Tree{
			arena: tree.arena,
			root:  l.joinUp(r),
			size:  l.size + r.size,
		},
		Source: tree.Source,
	}
}

func (tree TreapFingerTree) eachFromHead(p *Node, visit func(Data)) {
	if p == nil {
		return
	}
	visit(p.x)
	p.r.inorder(visit)
	tree.eachFromHead(p.l, visit)
}

func (tree TreapFingerTree) eachFromTail(p *Node, visit func(Data)) {
	if p == nil {
		return
	}
	tree.eachFromTail(p.r, visit)
	p.l.inorder(visit)
	visit(p.x)
}

// func (tree FingerTreap) head() *Node {
//    return tree.root.l
// }
//
// func (tree FingerTreap) tail() *Node {
//    return tree.root.r
// }

func (tree TreapFingerTree) inorder(p *Node, visit func(Data)) {
	if p == nil {
		return
	}
	tree.eachFromHead(p.l, visit)
	visit(p.x)
	tree.eachFromTail(p.r, visit)
}

func (tree TreapFingerTree) Each(visit func(Data)) {
	tree.inorder(tree.root, visit)
}

// func (p *Node) preorder(visit func(Data)) {
//    if p == nil {
//       return
//    }
//    visit(p.s)
//    p.l.preorder(visit)
//    p.r.preorder(visit)
// }
//
// func (p *Node) postorder(visit func(Data)) {
//    if p == nil {
//       return
//    }
//    p.l.postorder(visit)
//    p.r.postorder(visit)
//    visit(p.s)
// }

func (tree TreapFingerTree) verifyRanks() {
	if tree.root == nil {
		return
	}
	l := tree.root.l
	r := tree.root.r
	for ; l != nil; l = l.l {
		TreapTopDown{}.verifyMaxRankHeap(l.r)
		invariant(rank(l) >= rank(l.r))
		invariant(rank(l) <= rank(l.l) || l.l == nil)
	}
	for ; r != nil; r = r.r {
		TreapTopDown{}.verifyMaxRankHeap(r.l)
		invariant(rank(r) >= rank(r.l))
		invariant(rank(r) <= rank(r.r) || r.r == nil)
	}
	invariant(rank(tree.root) >= rank(l))
	invariant(rank(tree.root) >= rank(r))
}

func (tree TreapFingerTree) verifyPositions() {
	if tree.root == nil {
		return
	}
	// The root's size must be equal to the size of the left subtree.
	invariant(tree.root.s == tree.getHead().size())

	// Verify internal positions along the spines.
	for l := tree.getHead(); l != nil; l = l.l {
		tree.verifySize(l.r, l.s)
	}
	for r := tree.getTail(); r != nil; r = r.r {
		tree.verifySize(r.l, r.s)
	}
}

func (tree TreapFingerTree) Verify() {
	invariant(tree.size == tree.root.size())
	tree.verifyPositions()
	tree.verifyRanks()
}

//package structures
//
//import (
//  . "trees/pkg/abstract/list"
//  "trees/pkg/utility"
//  "golang.org/x/exp/rand"
//)
//
//
//
//// An implementation of a persistent list using a finger search tree (TreapFingerTreeOld).
////
//// The first node of the sequence is the head, the last node is the tail. The
//// head is therefore the leftmost node in the tree and the tail the rightmost.
////
//// Starting at the head, nodes along the left spine of the tree have their left
//// pointers reversed to point to their parent instead. Symmetrically along the
//// right spine starting from the tail, nodes have their right pointers reversed
//// to point to their parent. This provides access to internal nodes of the tree
//// by ascending along the spine from the head or tail and then descend inwards.
////
//// The root has no children. Datas access to nodes other than the root must start
//// from the head or the tail. There is no need to ascend both spines because the
//// position to search for can be compared to the relative position of the root
//// which is an absolute position in the sequence.
////
////
////                               ROOT
////                                   ↘ +4
////                                    (m)
////                            -2               +2
////                            (s)             (l)
////                        -1 ↗   ↘ +1     -1 ↙   ↖ +1
////                        (e)     (a)     (p)     (e)
////                       ↗                           ↖
////                   HEAD                             TAIL
////
////
//// Another way to visualize this structure is to use a horizontal spine layout.
//// Notice that the relative position of every node along the spine indicates how
//// many nodes will be skipped if a search continues along in the same direction:
////
////
////   HEAD                              ROOT                              TAIL
////
////   -1  -2    -4            -8        +16        +8            +4    +2  +1
////    ○ → ○  →  ○ --------- → ○         ◎         ○ ← --------- ○  ←  ○ ← ○
////        ↓     ↓             ↓                   ↓             ↓     ↓
////        ○     ○             ○                   ○             ○     ○
////            ↙   ↘        ↙     ↘             ↙     ↘        ↙   ↘
////           ○     ○     ○         ○         ○         ○     ○     ○
////                     ↙   ↘     ↙   ↘     ↙   ↘     ↙   ↘
////                    ○     ○   ○     ○   ○     ○   ○     ○
////
////
////
//// Operations:
////
////    From        O(n)        Creates a new instance from existing Datas.
////    Get         O(lg n)     Returns a Data at position.
////    Set         O(lg n)     Updates a Data at position.
////    Insert      O(lg n)     Inserts a Data at position, increasing length.
////    Remove      O(lg n)     Removes a Data at position, reducing length.
////    Split       O(lg n)     Partition by position, in-place on the left.
////    Join        O(lg n)     Concatenate by appending another TreapFingerTreeOld at the end.
////    Push        O(1)        Adds a Data after the last Data.
////    Unshift     O(1)        Adds a Data before the first Data.
////    Pop         O(1)        Removes the last Data.
////    Shift       O(1)        Removes the first Data.
////
////
////
//// Reading list:
////
////  - Implement an immutable deque as a balanced binary tree? (2010)
////    https://stackoverflow.com/q/3271256
////
////  - Functional Set Operations with Treaps (2001)
////    Dan Blandford, Guy Blelloch
////    https://www.cs.cmu.edu/afs/cs/project/pscico/pscico/papers/fingertrees/
////    https://www.cs.cmu.edu/afs/cs/project/pscico/pscico/src/fingertrees/
////
////  - An O(n log log n)-Time Algorithm for Triangulating a Simple Polygon (1986)
////    Appendix: Finger search trees
////    R. E. Tarjan, C. J. Van Wyk
////    https://api.semanticscholar.org/CorpusID:4981331
////  - Finger trees: a simple general-purpose data structure (2006)
////    R. Hinze, R. Paterson
////    https://api.semanticscholar.org/CorpusID:6881581
////
////  - Finger Search Trees (2005)
////    G. Brodal
////    https://api.semanticscholar.org/CorpusID:5694716
////
////
////
//// Symbols:
////
////       t   tree
////       o   other tree
////       i   index, offset, 0-based
////       d   relative position, distance
////       s   data
////
////
//type TreapFingerTreeOld struct {
//  Tree
//  root *Node // s ?
//  head *Node // l
//  tail *Node // r
//  size int   // w
//  // z another int available here
//}
//
//func (t *TreapFingerTreeOld) New() List {
//  return &TreapFingerTreeOld{}
//}
//
//// Creates a new instance of TreapFingerTreeOld by building a Treap then converting it to an TreapFingerTreeOld.
////
//// TODO: Build recursively from the head and tail towards the root.
//// TODO: Explore building by push without copy, aka transience.
////
//func (t *TreapFingerTreeOld) From(arr []Data) List {
//  return Treap{}.From(arr).(*Treap).ToFST()
//}
//
//// Returns the size of the tree, the length of its sequence.
//func (t *TreapFingerTreeOld) Len() int {
//  return t.size
//}
//
//// Creates a shallow copy of an TreapFingerTreeOld.
//// TODO just use structs where it makes sense to.
//func (t TreapFingerTreeOld) Clone() List {
//  /// TODO extract this to a node method
//  if t.root != nil {
//     t.root.rc++
//  }
//  ///////
//  if t.head != nil {
//     t.head.rc++
//  }
//  if t.tail != nil {
//     t.tail.rc++
//  }
//  return &t
//}
//
//// Returns the Data at i.
//func (t *TreapFingerTreeOld) Get(i int) Data {
//  assert(i < t.size)
//  return t.seekTo(i).s
//}
//
//// Replaces the Data at i, returns t.
//func (t *TreapFingerTreeOld) Set(i int, s Data) {
//  assert(i < t.size) // TODO move these assertions to the list impl?
//  t.shadowTo(i).withData(s)
//}
//
////
//func (t *TreapFingerTreeOld) seekTo(i int) *Node {
//  switch {
//     case i < t.root.sizeRL(): return t.seekFromHead(i)
//     case i > t.root.sizeRL(): return t.seekFromTail(i)
//  default:
//     return t.root
//  }
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) shadowTo(i int) *Node {
//  switch {
//     case i < t.root.sizeRL(): return t.shadowFromHead(i)
//     case i > t.root.sizeRL(): return t.shadowFromTail(i)
//  default:
//     return shadow(&t.root)
//  }
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) seekFromHead(i int) *Node {
//  return t.seekL(t.head, i)
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) seekFromTail(i int) *Node {
//  return t.seekR(t.tail, i - t.size + 1)
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) shadowFromHead(i int) *Node {
//  return t.seekShadowL(shadow(&t.head), i)
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) shadowFromTail(i int) *Node {
//  return t.seekShadowR(shadow(&t.tail), i - t.size + 1)
//}
//
////
////
////
////
////
////
////
////
////
// func (t *TreapFingerTreeOld) seekL(p *Node, i int) *Node {
//  for {
//     invariant(i >= 0)
//     invariant(p.isL())
//     //
//     if i == 0 {
//        return p
//     }
//     //
//     if i + p.w < 0 { // d - (sizeLR(p) + 1) < 0
//        return seekTo(i, p.r)
//     }
//     //
//     i = i + p.w
//     p = p.l
//  }
// }
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) seekR(p *Node, d int) *Node {
//  for {
//     invariant(d <= 0)
//     invariant(p.isR())
//
//     //
//     if d == 0 {
//        return p
//     }
//     //
//     if d + p.w > 0 { // d - (sizeRL(p) + 1) < 0
//        return seekTo(d, p.l)
//     }
//     //
//     d = d + p.w
//     p = p.r
//  }
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) seekShadowL(p *Node, d int) *Node {
//  for {
//     invariant(d >= 0)
//     invariant(p.isL())
//     //
//     if d == 0 {
//        return p
//     }
//     invariant(d > 0)
//     //
//     if d + p.w < 0 { // d - (sizeLR(p) + 1) < 0
//        return shadowTo(d, &p.r)
//     }
//     //
//     d = d + p.w
//     p = shadow(&p.l)
//  }
//}
//
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) seekShadowR(p *Node, d int) *Node {
//  for {
//     invariant(d <= 0)
//     invariant(p.isR())
//
//     //
//     if d == 0 {
//        return p
//     }
//     //
//     if d + p.w > 0 { // d - (sizeRL(p) + 1) < 0
//        return shadowTo(d, &p.l)
//     }
//     //
//     d = d + p.w
//     p = shadow(&p.r)
//  }
//}
//
//func generateRandomRank() Rank {
//  return random.Uint64()
//  //var level Rank
//  //for random.Uint64() & 1 == 0 {
//  //   level++
//  //}
//  ////fmt.Println(level)
//  //return level
//}
//
//// Creates a new node containing a given data Data with a fixed random rank.
//func (TreapFingerTreeOld) createNodeWithData(s Data) *Node {
//  return &Node{
//     s: s,
//     z: generateRandomRank(),
//  }
//}
//func (Treap) createNodeWithData(s Data) *Node {
//  return &Node{
//     s: s,
//     z: generateRandomRank(),
//  }
//}
//
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) Push(s Data) {
//  if t.isEmpty() {
//     t.root = t.createNodeWithData(s).withRelativePosition(+1)
//  } else {
//     t.tail = t.rotateUpR(t.createNodeWithData(s).withRelativePosition(+1).withR(t.tail))
//  }
//  t.size++
//}
//
//
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) Unshift(s Data) {
//  if t.isEmpty() {
//     t.root = t.createNodeWithData(s).withRelativePosition(+1)
//  } else {
//     //
//     t.root = copyOf(t.root); t.root.increaseInternalWeightOfR()
//     t.head = t.rotateUpL(t.createNodeWithData(s).withRelativePosition(-1).withL(t.head))
//  }
//  t.size++
//}
//
////
////
////
////
////
////
////
////
////

//
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) Shift() (s Data) {
//  assert(!t.isEmpty())
//  if t.head != nil {
//     s = t.head.s
//     t.dissolveHead()
//  } else {
//     s = t.root.s
//     t.removeRoot()
//  }
//  return s
//}
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) rotateUpR(p *Node) *Node {
//  invariant(p.isR())
//  //
//  //
//  for p.hasR() {
//     if rankOf(p) <= rankOf(p.r) {
//        return p
//     }
//     t.rotateLeftOnRightSpine(p)
//  }
//  //
//  //
//  if rankOf(p) <= rankOf(t.root) {
//     return p
//  }
//  t.rotateLeftIntoRoot(p)
//  return nil
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) rotateUpL(p *Node) *Node {
//  invariant(p.isL())
//  //
//  //
//  for p.hasL() {
//     if rankOf(p) <= rankOf(p.l) {
//        return p
//     }
//     t.rotateRightOnLeftSpine(p)
//  }
//  //
//  //
//  if rankOf(p) <= rankOf(t.root) {
//     return p
//  }
//  t.rotateRightIntoRoot(p)
//  return nil
//}
//
//
//
//
//
////
////
////
////
////
////
////
////
////
//
//func (t *TreapFingerTreeOld) rotateRightIntoRoot(p *Node) {
//  //visualize("↻")
//  invariant(p.isL())
//  invariant(!p.hasL())
//  invariant(!t.isEmpty())
//  //
//  //
//  g := copyOf(t.root)
//  r := p.r
//  //
//  //
//  appendR(&t.tail, g)
//
//  g.l = r; t.root = p; p.r = nil
//  rotateRelativePositions(g, p, r)
//}
//
////
////
////
////
////
////
////
////
////
//
//func (t *TreapFingerTreeOld) rotateLeftIntoRoot(p *Node) {
//  //visualize("↺")
//  invariant(p.isR())
//  invariant(!p.hasR())
//  invariant(!t.isEmpty())
//  // invariant rank of p is greater than root
//  g := copyOf(t.root)
//  l := p.l
//
//  //
//  //
//  appendL(&t.head, g)
//
//  g.r = l
//  t.root = p
//  p.l = nil
//  rotateRelativePositions(g, p, l)
//}
//
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) rotateLeftOnRightSpine(p *Node)  {
//  invariant(p.isR())
//  //visualize("↺")
//  g := p.r.copy()
//  //
//  // as
//  p.r, p.l, g.r = g.r, g, p.l
//  rotateRelativePositions(g, p, g.r)
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) rotateRightOnLeftSpine(p *Node) {
//  invariant(p.isL())
//  //visualize("↻")
//  g := p.l.copy()
//  //
//  // as
//  p.l, p.r, g.l = g.l, g, p.r
//  rotateRelativePositions(g, p, g.l)
//}
//
//
//
////
////
//func (t *TreapFingerTreeOld) Insert(i int, s Data) {
//  assert(i <= t.size)
//
//  // takes care of the null root case also.
//  if i == 0 {
//     t.Unshift(s)
//     return
//  }
//  if i == t.size {
//     t.Push(s)
//     return
//  }
//  if i < t.root.w {
//     t.insertFromHead(s, i)
//  } else {
//     t.insertFromTail(s, i)
//  }
//}
//
////
//func (t *TreapFingerTreeOld) insertOnRightSpine(p *Node, n *Node, d int) {
//  invariant(!n.hasL())
//  invariant(!n.hasR())
//  invariant(d < 0)
//
//  n.l, p.l = unzip(p.l, d, p.sizeRL())
//
//  n.w = d + p.w + 1 // TODO revise
//  p.w = -d
//
//  //
//  if n.hasL() {
//     n.l.toL(n.sizeRL())
//  }
//  n.r = p.r
//  p.r = t.rotateUpR(n)
//}
//
////
//func (t *TreapFingerTreeOld) insertOnLeftSpine(p *Node, n *Node, d int) {
//  invariant(!n.hasL())
//  invariant(!n.hasR())
//  invariant(d >= 0)
//  //
//  //
//  p.r, n.r = unzip(p.r, d, p.sizeLR())
//
//  //
//  n.w = p.w + d // TODO revise
//  p.w = -d - 1
//
//  //
//  if n.hasR() {
//     n.r.toR(n.sizeLR())
//  }
//  n.l = p.l
//  p.l = t.rotateUpL(n)
//}
//
//
////
////
//func (t *TreapFingerTreeOld) insertFromTail(s Data, i int) {
//  //visualize("↖")
//
//  t.tail = copyOf(t.tail)
//
//  n := t.createNodeWithData(s)
//  p := t.tail
//  d := i - t.size
//
//  //
//  t.size = t.size + 1
//
//  for {
//     //visualize("↖")
//     invariant(d <= 0)
//     invariant(p.isR())
//
//     if d + p.w < 0 {
//        d = d + p.w
//        p = shadow(&p.r)
//        continue
//     }
//     if rankOf(p) < rankOf(n) {
//        t.insertOnRightSpine(p, n, d)
//        return
//     }
//     //visualize("|↙")
//     (&Treap{}).insert(&p.l, n, d, p.sizeRL()); p.increaseInternalWeightOfR()
//     return
//  }
//}
//
//func (t *TreapFingerTreeOld) insertFromHead(s Data, i int) {
//  //visualize("↗")
//  invariant(i >= 0)
//
//  //
//  t.head = copyOf(t.head)
//  t.root = copyOf(t.root)
//  t.size = t.size + 1
//
//  t.root.increaseInternalWeightOfR()
//  //
//  n := t.createNodeWithData(s)
//  p := t.head
//  d := i - 1
//
//  for {
//     //visualize("↗")
//     invariant(d >= 0)
//     invariant(p.isL())
//
//     //
//     //
//     if d + p.w >= 0  {
//        d = d + p.w
//        p = shadow(&p.l)
//        continue
//     }
//     //
//     //
//     if rankOf(p) < rankOf(n) {
//        t.insertOnLeftSpine(p, n, d)
//        return
//     }
//     //
//     //
//     //visualize("|↘")
//     (&Treap{}).insert(&p.r, n, d, p.sizeLR()); p.increaseInternalWeightOfL()
//     return
//  }
//}
//
//
////
//
//
////
////
//func (t *TreapFingerTreeOld) Remove(i int) (s Data){
//  switch {
//  case i < t.root.sizeRL(): return t.removeFromHead(i)
//  case i > t.root.sizeRL(): return t.removeFromTail(i)
//  default:
//     return t.removeRoot()
//  }
//}
//
//
//func (t *TreapFingerTreeOld) toBST() *Treap {
//  if t.isEmpty() {
//     return &Treap{}
//  }
//  return &Treap{
//     size: t.size,
//     root: copyOf(t.root).
//        withL(reverseL(t.head.copyOrNil(), nil)).
//        withR(reverseR(t.tail.copyOrNil(), nil)),
//  }
//}
//
//
//
////
////
//func (t *TreapFingerTreeOld) isEmpty() bool {
//  return t.root == nil
//}
//
////
////
//func inorderL(p *Node, fn func(*Node)) {
//  if p == nil {
//     return
//  }
//  fn(p)
//  inorder(p.r, fn)
//  inorderL(p.l, fn)
//}
//
////
//func inorderR(p *Node, fn func(*Node)) {
//  if p == nil {
//     return
//  }
//  inorderR(p.r, fn)
//  inorder(p.l, fn)
//  fn(p)
//}
////
//func (t *TreapFingerTreeOld) inorder(fn func(*Node)) {
//  if t.isEmpty() {
//     return
//  }
//  inorderL(t.head, fn)
//  fn(t.root)
//  inorderR(t.tail, fn)
//}
//
////
//func (t *TreapFingerTreeOld) Each(fn func(i int, s Data)) {
//  i := 0
//  t.inorder(func(n *Node) {
//     fn(i, n.s)
//     i++
//  })
//}
//
////
//func (t *TreapFingerTreeOld) Array() []Data {
//  a := make([]Data, t.size)
//  i := 0
//  t.inorder(func(n *Node) {
//     a[i] = n.s
//     i++
//  })
//  return a
//}
//
////
////
//func (t *TreapFingerTreeOld) Validate() {
//  if t.isEmpty() {
//     assert(t.head == nil)
//     assert(t.tail == nil)
//     assert(t.size == 0)
//     return
//  }
//  t.verifyRelationPositionInvariant()
//  t.verifyRankHeapInvariant()
//  t.verifyRoot()
//}
//
//
//
////
////
//func (t *TreapFingerTreeOld) verifyRanksAlongLeftSpine(p *Node) {
//  for ; p != nil; p = p.l {
//     if p.hasR() {
//        assert(rankOf(p) >= rankOf(p.r))
//        p.r.verifyRankHeapInvariant()
//     }
//     if p.hasL() {
//        assert(rankOf(p) <= rankOf(p.l))
//     }
//  }
//}
//
////
////
//func (t *TreapFingerTreeOld) verifyRanksAlongRightSpine(p *Node) {
//  for ; p != nil; p = p.r {
//     if p.hasL() {
//        assert(rankOf(p) >= rankOf(p.l))
//        p.l.verifyRankHeapInvariant()
//     }
//     if p.hasR() {
//        assert(rankOf(p) <= rankOf(p.r))
//     }
//  }
//}
//
////
//func (t *TreapFingerTreeOld) verifyRankHeapInvariant() {
//  t.verifyRanksAlongLeftSpine(t.head)
//  t.verifyRanksAlongRightSpine(t.tail)
//
//  if t.root != nil {
//     //
//     //
//     if t.head != nil {
//        for p := t.head; p.hasL(); p = p.l {
//           assert(rankOf(p) <= rankOf(t.root))
//        }
//     }
//     //
//     //
//     if t.tail != nil {
//        for p := t.tail; p.hasR(); p = p.r {
//           assert(rankOf(p) <= rankOf(t.root))
//        }
//     }
//  }
//}
//
////
//func (t *TreapFingerTreeOld) verifyRoot() {
//  assert(t.root.l == nil)
//  assert(t.root.r == nil)
//  assert(t.root.sizeRL() == t.head.size())
//}
//
////
//func (t *TreapFingerTreeOld) verifyRelationPositionInvariant() {
//  if t.root != nil {
//     assert(t.root.isR())
//     assert(t.head.size() == t.root.sizeRL())
//  }
//  for p := t.head; p != nil; p = p.l {
//     assert(p.isL())
//     p.verifyRelativePositionInvariant()
//  }
//  for p := t.tail; p != nil; p = p.r {
//     assert(p.isR())
//     p.verifyRelativePositionInvariant()
//  }
//}
//
//
//func (t *TreapFingerTreeOld) hasHead() bool {
//  return t.head != nil
//}
//
//func (t *TreapFingerTreeOld) hasTail() bool {
//  return t.tail != nil
//}
//
//func (t *TreapFingerTreeOld) hasRoot() bool {
//  return t.root != nil
//}
//
////
////
////
////
////
////
////
////
////
//func (t *TreapFingerTreeOld) dissolveHead() {
//  invariant(t.hasHead())
//
//  //
//  t.root = copyOf(t.root).withRelativePosition(t.root.sizeRL())
//  //
//  p := t.head
//  g := t.head.l
//  r := t.head.r
//  //
//  //
//  if t.head.hasR() {
//     t.head = reverseL(copyOf(r).toL(p.sizeLR()), g)
//  } else {
//     t.head = g
//  }
//  t.size--
//}
//
//
////
////
////
////
////
////
////
////
//
//
////
////
//func (t *TreapFingerTreeOld) Split(i Index) (List, List) {
//  assert(i <= t.size)
//  //
//  //
//  if i == t.size { return t, &TreapFingerTreeOld{} }
//  if i == 0      { return &TreapFingerTreeOld{}, t }
//
//  tmp := t.Clone().(*TreapFingerTreeOld) // get this of this all
//
//  if i <= t.root.sizeRL() {
//     return tmp, tmp.splitFromHead(i)
//  } else {
//     return tmp, tmp.splitFromTail(i)
//  }
//}
//
////
////

//
//// Can we use _dissolve here?
////
//func (t *TreapFingerTreeOld) removeRoot() (s Data) {
//  invariant(t.root != nil)
//  s = t.root.s
//  //
//  if t.size == 1 {
//     t.size = 0
//     t.root = nil
//     return
//  }
//  //
//  //
//  if t.head == nil {
//     // TODO use tree construction here?
//     pr := followR(&t.tail)
//     t.head = reverseL((*pr).l.copyOrNil(), nil)
//     t.root = copyOf(*pr).withL(nil).withR(nil)
//     t.size = t.size - 1
//     *pr = nil
//     return
//  }
//  //
//  //
//  if t.tail == nil {
//     pl := followL(&t.head)
//     t.tail = reverseR((*pl).r.copyOrNil(), nil)
//     t.root = copyOf(*pl).withL(nil).withR(nil).toR(t.root.sizeRL())
//     t.size = t.size - 1
//     *pl = nil
//     return
//  }
//  //
//  //
//  pl := followL(&t.head)
//  pr := followR(&t.tail)
//  //
//  //
//  sl := t.root.sizeRL()
//  sr := t.root.sizeRR(t.size)
//  //
//
//  //
//  l := (*pl).copy().toR(sl)
//  r := (*pr).copy().toL(sr)
//
//  var p *Node
//
//  //
//  // TODO is it not possible to already have a p on hand here?
//  (&Treap{}).joinRL(&p, l, r, sl + sr)
//  //
//  //
//  if rankOf(l) >= rankOf(r) {
//     *pr = reverseR(p.r, nil)
//     *pl = nil
//  } else {
//     *pl = reverseL(p.l, nil)
//     *pr = nil
//  }
//  //
//  //
//  t.root = p.withL(nil).withR(nil)
//  t.size = t.size - 1
//  return
//}
//
//
//// TODO: if instead we use the same, does there exist a general algorithm?
////       if so or otherwise, does it simplify anything?
////
////
////
//func (t *TreapFingerTreeOld) splitFromHead(i int) (o *TreapFingerTreeOld) {
//  //
//  //
//  //
//  g := &t.head
//  d := i
//  for {
//     invariant(d > 0)
//     //
//     //
//     if d - ((*g).sizeLR() + 1) > 0 {
//        d -= (*g).sizeLR() + 1
//        g = &(shadow(g).l)
//        continue
//     }
//     //
//     //
//     //
//     //
//     p := shadow(g)
//     l, r := unzip(p.r, d - 1, p.sizeLR())
//     *g = nil
//     //
//     //
//     o = &TreapFingerTreeOld{
//        root: copyOf(t.root).withRelativePosition(t.root.w - i),
//        head: reverseL(r, p.l),
//        tail: t.tail,
//        size: t.size - i,
//     }
//     //
//     //
//     *t = TreapFingerTreeOld{
//        root: p.withRelativePosition(i - (d - 1)).withL(nil).withR(nil),
//        tail: reverseR(l, nil),
//        head: t.head,
//        size: i,
//     }
//     return
//  }
//}
//
////
////
//func (t *TreapFingerTreeOld) splitFromTail(i int) (o *TreapFingerTreeOld) {
//  //
//  //
//  g := &t.tail
//  d := i - t.size
//  for {
//     invariant(d < 0)
//     //
//     //
//     if d + 1 + (*g).sizeRL() < 0 {
//        d += (*g).sizeRL() + 1
//        g = &shadow(g).r
//        continue
//     }
//     //
//     //
//     p := shadow(g)
//     l, r := unzip(p.l, d, p.sizeRL())
//     *g = nil
//
//     o = &TreapFingerTreeOld{
//        root: copyOf(p).withW(-d).withL(nil).withR(nil),
//        head: reverseL(r, nil),
//        tail: t.tail,
//        size: t.size - i,
//     }
//     //
//     //
//     *t = TreapFingerTreeOld{
//        root: t.root,
//        tail: reverseR(l, p.r),
//        head: t.head,
//        size: i,
//     }
//     return
//  }
//}
//
//
//func (t Treap) DepthAlongTheSpines() (depths [2][]int) {
//  if t.root == nil {
//     return
//  }
//  traverseL(t.root.l, func(p *Node) {
//     depths[0] = append(depths[0], 1 + p.r.depth())
//  })
//  traverseR(t.root.r, func(p *Node) {
//     depths[1] = append(depths[1], 1 + p.l.depth())
//  })
//  return
//}
//
//
////
//func (t TreapFingerTreeOld) WeightAlongTheSpinesLog2() (weights [2][]int) {
//  var l []*Node
//  var r []*Node
//
//  traverseL(t.head, func(p *Node) {
//     l = append(l, p)
//  })
//  traverseR(t.tail, func(p *Node) {
//     r = append(r, p)
//  })
//  for i := len(l) - 1; i >= 0; i-- {
//     weights[0] = append(weights[0], 1 + utility.Log2(l[i].sizeLR() + 1))
//  }
//  for i := len(r) - 1; i >= 0; i-- {
//     weights[1] = append(weights[1], 1 + utility.Log2(r[i].sizeRL() + 1))
//  }
//  return
//}
//
//func (t TreapFingerTreeOld) DepthAlongTheSpines() (depths [2][]int) {
//  var l []*Node
//  var r []*Node
//  traverseL(t.head, func(p *Node) {
//     l = append(l, p)
//  })
//  traverseR(t.tail, func(p *Node) {
//     r = append(r, p)
//  })
//  for i := len(l) - 1; i >= 0; i-- {
//     depths[0] = append(depths[0], 1 + l[i].r.depth())
//  }
//  for i := len(r) - 1; i >= 0; i-- {
//     depths[1] = append(depths[1], 1 + r[i].l.depth())
//  }
//  return
//}
