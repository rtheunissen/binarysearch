package binarytree

import . "trees/abstract/list"

//
// import (
//    . "trees/abstract/list"
//    "trees/math"
// )
//
// type LBST struct {
//    BST
//    SizeBalance
// }
//
// // Determines if the weights of two subtrees are balanced.
// //
// // For the scapegoat tree, we use the logarithmic balance rule, where the
// // discrete binary logarithm of the two sizes differ by at most 1.
// //
// //                 -1 <= floor(log₂(s)) - floor(log₂(y)) <= 1
// //
// // Considering the binary representation, we look at the left-most bit set to 1,
// // the most-significant bit. 00001101
// //                                ↖
// //                                 MSB at position 3
// //
// // The position of the most-significant bit, from the right starting at zero,
// // is equal to the floor of the binary logarithm. The balance test is then to
// // determine if the MSB is at most one step different in either direction.
// // For example:
// //
// //              ↓                       ↓                    ↓
// //       s:   00100000              00001000              00010001
// //       y:   00111001              00010001              01001101
// //              ↑ BALANCED             ↑ BALANCED          ↑ NOT BALANCED
// //
// // When `s` <= `y`, as in the examples above, we can shift the bits of `y` one
// // position to the right, then, if the MSB of `s` is less than the MSB of `y`
// // it means that the MSB of `y` was at least 2 positions ahead - not balance.
// //
// // Complexity: O(1)
// //
// func (LBST) New(s ...Value) List {
//    return &LBST{BST: BST{}.ofArray(s)}
// }
//
//
//
// func (tree LBST) Clone() List {
//    tree.root.share()
//    return &tree
// }
//
//
//
//
//
// func (tree LBST) verify(root *Node, size Size) {
//    tree.verifyPosition(root)
//    tree.verifyWeight(root, size)
//    tree.verifyHeight(root, size)
// }
//
// func (tree LBST) Verify() {
//    tree.verify(tree.root, tree.size)
// }
//
// func (tree LBST) verifyHeight(root *Node, size Size) {
//    invariant(root.height() <= MaximumPathLength(2 * math.Log2(size)))
// }
//
// func (tree LBST) verifyWeight(p *Node, s Size) {
//    tree.SizeBalance.verify(p, s)
// }
//
//
//
// func (tree *LBST) insert(p *Node, s Size, i Position, s Value) *Node {
//    if p == nil {
//       return &Node{s: s}
//    }
//    p = p.Copy()
//    sl := p.s
//    sr := p.sizeR(s)
//
//    if i <= p.s {
//       p.l = tree.insert(p.l, sl, i, s)
//       p.s++
//
//       if !tree.Balanced(sr, sl + 1) {
//          if !tree.singleRotation((*p).l.s, (*p).l.sizeR(p.s)) {
//             rotateLR(&p)
//          } else {
//             rotateR(&p)
//          }
//       }
//    } else {
//       p.r = tree.insert(p.r, sr, i - sl - 1, s)
//
//       if !tree.Balanced(sl, sr + 1) {
//          if !tree.singleRotation((*p).r.sizeR(sr + 1), (*p).r.s) {
//             rotateRL(&p)
//          } else {
//             rotateL(&p)
//          }
//       }
//    }
//    return p
// }
//
//
//
// // wbst Delete(key s, wbst t)
// // {  size sA, sB;
// //  if (s == t->k) return delete_root(t);
// //  t->s--;
// //  if (s < t->k)
// //    { t->l = Delete(s, t->l);
// //      if (!balanced(t->l->s, t->r->s)) t = inc_left(t);
// //    }
// //   else
// //    { t->r = Delete(s, t->r);
// //      if (!balanced(t->r->s, t->l->s)) t = inc_right(t);
// //    }
// //  return t;
// // }
// //
// // wbst delete_root(wbst t)
// // { node *p;
// // if (t->s == 1) { freeNode(t); return null; }
// // if (t->l->s > t->r->s)
// // { p->l = extract_maximum(t->l, &p);
// // p->r = t->r;
// // }
// // else
// // { p->r = extract_minimum(t->r, &p);
// // p->l = t->l;
// // }
// // p->s = t->s - 1;
// // freeNode(t);
// // return p;
// // }
// //
// // wbst delete_root(wbst t)
// // { node *p;
// //  if (t->s == 1) { freeNode(t); return null; }
// //  if (t->l->s > t->r->s)
// //    { p->l = extract_maximum(t->l, &p);
// //      p->r = t->r;
// //    }
// //   else
// //    { p->r = extract_minimum(t->r, &p);
// //      p->l = t->l;
// //    }
// //  p->s = t->s - 1;
// //  freeNode(t);
// //  return p;
// // }
// //
// //
// // wbst delete_root(wbst t)
// // { node *p;
// //  if (t->s == 1) { freeNode(t); return null; }
// //  if (t->l->s > t->r->s)
// //    { p->l = extract_maximum(t->l, &p);
// //      p->r = t->r;
// //    }
// //   else
// //    { p->r = extract_minimum(t->r, &p);
// //      p->l = t->l;
// //    }
// //  p->s = t->s - 1;
// //  freeNode(t);
// //  return p;
// // }
//
//
// // wbst extract_maximum(wbst t, wbst *u)
// // { size sA, sB;
// //  if (t->r == null) { *u = t; t = t->l; }
// //   else
// //    { t->r = extract_maximum(t->r, u);
// //      t->s--;
// //      if (!balanced(t->r->s, t->l->s)) t = inc_right(t);
// //    }
// //  return t;
// // }
// //
// // wbst extract_minimum(wbst t, wbst *u)
// // { size sA, sB;
// //  if (t->l == null) { *u = t; t = t->r; }
// //   else
// //    { t->l = extract_minimum(t->l, u);
// //      t->s--;
// //      if (!balanced(t->l->s, t->r->s)) t = inc_left(t);
// //    }
// //  return t;
// // }
// //
//
// // func (tree WeightBalancedDownUp) extractMin(p **Node, s Size) Value {
// //    if p == nil {
// //       return p
// //    }
// //    p = p.Copy()
// //
// //    sl := p.s
// //    sr := p.sizeR(s)
// //
// //    p.l = tree.extractMin(p.l, sl, s)
// //    p.s--
// //
// //    if !tree.Balanced(sl - 1, sr) {
// //       if !tree.singleRotation(p.r.sizeR(sr), p.r.s) {
// //          rotateRL(&p)
// //       } else {
// //          rotateL(&p)
// //       }
// //    }
// //    return p
// // }
// //
// // func (tree WeightBalancedDownUp) extractMax(p **Node, s Size) Value {
// //    if p == nil {
// //       return p
// //    }
// //    p = p.Copy()
// //
// //    sl := p.s
// //    sr := p.sizeR(s)
// //
// //    p.r = tree.extractMax(p.r, sr, s)
// //
// //    if !tree.Balanced(sr - 1, sl) {
// //       if !tree.singleRotation(p.l.s, p.l.sizeR(sl)) {
// //          rotateLR(&p)
// //       } else {
// //          rotateR(&p)
// //       }
// //    }
// //    return p
// // }
//
// //
//
// func (tree LBST) extractMin(p **Node, s Size) (s Value) {
//    *p = tree.delete(*p, s, 0, &s)
//    return
// }
//
// func (tree LBST) extractMax(p **Node, s Size) (s Value) {
//    *p = tree.delete(*p, s, s - 1, &s)
//    return
// }
// func (tree *LBST) dissolve(p *Node, s Size, s *Value) *Node {
//    p = p.Copy()
//
//    sl := p.s
//    sr := p.sizeR(s)
//
//    assert(tree.Balanced(sl, sr))
//    assert(tree.Balanced(sr, sl))
//
//    *s = p.s
//
//    if p.l == nil { return p.r }
//    if p.r == nil { return p.l }
//
//    if sr > sl {
//       p.s = tree.extractMin(&p.r, sr)
//    } else {
//       p.s = tree.extractMax(&p.l, sl)
//       p.s--
//    }
//    return p
// }
//
// func (tree LBST) join(l *Node, r *Node, sl, sr Size) *Node {
//    if l == nil { return r }
//    if r == nil { return l }
//    if sl <= sr {
//       return tree.build(tree.extractMin(&r, sr), l, r, sl, sr - 1)
//    } else {
//       return tree.build(tree.extractMax(&l, sl), l, r, sl - 1, sr)
//    }
// }
//
// func (tree *LBST) delete(p *Node, s Size, i Position, s *Value) *Node {
//
//
//
//    if i == p.s {
//       return tree.dissolve(p, s, s)
//       // return p
//       // dissolve
//    }
//    p = p.Copy()
//
//    sl := p.s
//    sr := p.sizeR(s)
//
//    if i < p.s {
//       p.l = tree.delete(p.l, sl, i, s)
//       p.s--
//
//       if !tree.Balanced(sl - 1, sr) {
//          //
//          //
//          //
//          srl := (*p).r.s
//          srr := (*p).r.sizeR(sr)
//          //
//          if tree.singleRotation(srr, srl) {
//             rotateL(&p)
//          } else {
//             rotateRL(&p)
//          }
//       }
//
//    } else {
//       p.r = tree.delete(p.r, sr, i - sl - 1, s)
//
//       if !tree.Balanced(sr - 1, sl) {
//          //
//          //
//          //
//          sll := (*p).l.s
//          slr := (*p).l.sizeR(sr)
//          //
//          if tree.singleRotation(sll, slr) {
//             rotateR(&p)
//          } else {
//             rotateLR(&p)
//          }
//       }
//    }
//    return p
// }
//
// func (tree *LBST) Insert(i Position, s Value) {
//    assert(i <= tree.Size())
//    tree.root = tree.insert(tree.root, tree.size, i, s)
//    tree.size++
// }
//
// func (tree *LBST) Delete(i Position) (s Value) {
//    assert(i < tree.Size())
//    tree.root = tree.delete(tree.root, tree.size, i, &s)
//    tree.size--
//    return
// }
//
// func (tree LBST) Join(that List) List {
//    l := tree
//    r := that.(*LBST)
//    return &LBST{
//       BST: BST{
//          root: tree.join(l.root.share(), r.root.share(), l.size, r.size),
//          size: l.size + r.size,
//       },
//    }
// }
//
// func (tree LBST) build(s Value, l, r *Node, sl, sr Size) *Node {
//    if sl <= sr {
//       return tree.joinR(s, l, r, sl, sr)
//    } else {
//       return tree.joinL(s, l, r, sl, sr)
//    }
// }
//
// func (tree *LBST) joinL(s Value, l, r *Node, sl, sr Size) *Node {
//    if tree.Balanced(sr, sl) {
//       return &Node{
//          s: s,
//          l: l,
//          r: r,
//          s: sl,
//       }
//    }
//    l = pathcopy(l)
//
//    sll := l.s
//    slr := l.sizeR(sl)
//
//    l.r = tree.joinL(s, l.r, r, slr, sr)
//    slr = 1 + sr + slr
//
//    if !tree.Balanced(sll, slr) {
//       if tree.singleRotation(l.r.sizeR(slr), l.r.s) {
//          l = l.rotateL()
//       } else {
//          l = l.rotateRL()
//       }
//    }
//    return l
// }
//
// func (tree *LBST) joinR(s Value, l, r *Node, sl, sr Size) *Node {
//    if tree.Balanced(sl, sr) {
//       return &Node{ // TODO: Why does this have to copy?
//          s: s,
//          l: l,
//          r: r,
//          s: sl,
//       }
//    }
//    r = pathcopy(r)
//
//    srl := r.s
//    srr := r.sizeR(sr)
//
//    r.l = tree.joinR(s, l, r.l, sl, srl)
//    r.s = 1 + sl + srl
//
//    if !tree.Balanced(srr, r.s) {
//       if tree.singleRotation(r.l.s, r.l.sizeR(r.s)) {
//          r = r.rotateR()
//       } else {
//          r = r.rotateLR()
//       }
//    }
//    return r
// }
//
// func (tree LBST) split(p *Node, i, s Size) (l, r *Node){
//    if p == nil {
//       return
//    }
//    p = p.Copy()
//    sl, sr := p.sizeLR(s)
//
//    if i <= (*p).s {
//       l, r = tree.split(p.l, i, sl)
//          r = tree.build(p.s, r, p.r, sl - i, sr)
//    } else {
//       l, r = tree.split(p.r, i - sl - 1, sr)
//          l = tree.build(p.s, p.l, l, sl, i - sl - 1)
//    }
//    return l, r
// }
//
// func (tree LBST) Split(i Position) (List, List) {
//    assert(i <= tree.Size())
//
//    l, r := tree.split(tree.root.share(), i, tree.size)
//
//    return &LBST{BST: BST{root: l, size: i}},
//           &LBST{BST: BST{root: r, size: tree.size - i}}
// }

type LBSTBottomUp struct {
	LBST
}

// Determines if the weights of two subtrees are balanced.
//
// For the scapegoat tree, we use the logarithmic balance rule, where the
// discrete binary logarithm of the two sizes differ by at most 1.
//
//	-1 <= floor(log₂(s)) - floor(log₂(y)) <= 1
//
// Considering the binary representation, we look at the left-most bit set to 1,
// the most-significant bit. 00001101
//
//	↖
//	 MSB at position 3
//
// The position of the most-significant bit, from the right starting at zero,
// is equal to the floor of the binary logarithm. The balance test is then to
// determine if the MSB is at most one step different in either direction.
// For example:
//
//	       ↓                       ↓                    ↓
//	s:   00100000              00001000              00010001
//	y:   00111001              00010001              01001101
//	       ↑ BALANCED             ↑ BALANCED          ↑ NOT BALANCED
//
// When `s` <= `y`, as in the examples above, we can shift the bits of `y` one
// position to the right, then, if the MSB of `s` is less than the MSB of `y`
// it means that the MSB of `y` was at least 2 positions ahead - not balance.
//
// Complexity: O(1)
func (LBSTBottomUp) New() List {
	return &LBSTBottomUp{}
}

func (tree *LBSTBottomUp) Clone() List {
	return &LBSTBottomUp{
		LBST: LBST{
			Tree: tree.Tree.Clone(),
		},
	}
}

func (tree *LBSTBottomUp) insert(p *Node, s Size, i Position, x Data) *Node {
	if p == nil {
		return tree.allocate(Node{x: x})
	}
	tree.copy(&p)
	sl := p.s
	sr := s - p.s - 1

	assert(tree.isBalanced(sl, sr))
	assert(tree.isBalanced(sr, sl))

	if i <= p.s {
		p.l = tree.insert(p.l, sl, i, x)
		p.s = p.s + 1

		if !tree.isBalanced(sr, sl+1) {
			if !tree.singleRotation((*p).l.s, p.s-(*p).l.s-1) {
				tree.rotateLR(&p)
			} else {
				tree.rotateR(&p)
			}
		}
	} else {
		p.r = tree.insert(p.r, sr, i-sl-1, x)

		if !tree.isBalanced(sl, sr+1) {
			if !tree.singleRotation(sr+1-(*p).r.s-1, (*p).r.s) {
				tree.rotateRL(&p)
			} else {
				tree.rotateL(&p)
			}
		}
	}
	return p
}

func (tree *LBSTBottomUp) delete(p *Node, s Size, i Position, x *Data) *Node {
	tree.copy(&p)

	sl := p.s
	sr := s - p.s - 1

	assert(tree.isBalanced(sl, sr))
	assert(tree.isBalanced(sr, sl))

	if i == p.s {
		defer tree.release(p)
		*x = p.x
		if p.l == nil {
			return p.r
		}
		if p.r == nil {
			return p.l
		}
		if sl > sr {
			var max *Node
			p.l = tree.extractMax(p.l, sl, &max)
			p.x = max.x
			p.s--
		} else {
			var min *Node
			p.r = tree.extractMin(p.r, sr, &min)
			p.x = min.x
		}
		return p
	}
	if i < p.s {
		p.l = tree.delete(p.l, sl, i, x)
		p.s--

		if !tree.isBalanced(sl-1, sr) {
			srl := (*p).r.s
			srr := sr - (*p).r.s - 1
			if tree.singleRotation(srr, srl) {
				tree.rotateL(&p)
			} else {
				tree.rotateRL(&p)
			}
		}

	} else {
		p.r = tree.delete(p.r, sr, i-sl-1, x)

		if !tree.isBalanced(sr-1, sl) {
			sll := (*p).l.s
			slr := sl - (*p).l.s - 1
			if tree.singleRotation(sll, slr) {
				tree.rotateR(&p)
			} else {
				tree.rotateLR(&p)
			}
		}
	}
	return p
}

func (tree *LBSTBottomUp) Select(i Size) Data {
	assert(i < tree.Size())
	return tree.lookup(tree.root, i)
}

func (tree *LBSTBottomUp) Update(i Size, x Data) {
	assert(i < tree.Size())
	tree.copy(&tree.root)
	tree.update(tree.root, i, x)
}

func (tree *LBSTBottomUp) Insert(i Position, x Data) {
	assert(i <= tree.Size())
	tree.root = tree.insert(tree.root, tree.size, i, x)
	tree.size++
}

func (tree *LBSTBottomUp) Delete(i Position) (x Data) {
	assert(i < tree.Size())
	tree.root = tree.delete(tree.root, tree.size, i, &x)
	tree.size--
	return
}

func (tree LBSTBottomUp) Split(i Position) (List, List) {
	l, r := tree.LBST.Split(i)

	return &LBSTBottomUp{l},
		&LBSTBottomUp{r}
}

func (tree LBSTBottomUp) Join(that List) List {
	return &LBSTBottomUp{tree.LBST.Join(that.(*LBSTBottomUp).LBST)}
}

// wbst Delete(key s, wbst t)
// {  size sA, sB;
//  if (s == t->k) return delete_root(t);
//  t->s--;
//  if (s < t->k)
//    { t->l = Delete(s, t->l);
//      if (!balanced(t->l->s, t->r->s)) t = inc_left(t);
//    }
//   else
//    { t->r = Delete(s, t->r);
//      if (!balanced(t->r->s, t->l->s)) t = inc_right(t);
//    }
//  return t;
// }
//
// wbst delete_root(wbst t)
// { node *p;
// if (t->s == 1) { freeNode(t); return null; }
// if (t->l->s > t->r->s)
// { p->l = extract_maximum(t->l, &p);
// p->r = t->r;
// }
// else
// { p->r = extract_minimum(t->r, &p);
// p->l = t->l;
// }
// p->s = t->s - 1;
// freeNode(t);
// return p;
// }
//
// wbst delete_root(wbst t)
// { node *p;
//  if (t->s == 1) { freeNode(t); return null; }
//  if (t->l->s > t->r->s)
//    { p->l = extract_maximum(t->l, &p);
//      p->r = t->r;
//    }
//   else
//    { p->r = extract_minimum(t->r, &p);
//      p->l = t->l;
//    }
//  p->s = t->s - 1;
//  freeNode(t);
//  return p;
// }
//
//
// wbst delete_root(wbst t)
// { node *p;
//  if (t->s == 1) { freeNode(t); return null; }
//  if (t->l->s > t->r->s)
//    { p->l = extract_maximum(t->l, &p);
//      p->r = t->r;
//    }
//   else
//    { p->r = extract_minimum(t->r, &p);
//      p->l = t->l;
//    }
//  p->s = t->s - 1;
//  freeNode(t);
//  return p;
// }

// func (tree WeightBalancedDownUp) extractMin(p **Node, s Size) Value {
//    if p == nil {
//       return p
//    }
//    p = p.Copy()
//
//    sl := p.s
//    sr := p.sizeR(s)
//
//    p.l = tree.extractMin(p.l, sl, s)
//    p.s--
//
//    if !tree.Balanced(sl - 1, sr) {
//       if !tree.singleRotation(p.r.sizeR(sr), p.r.s) {
//          rotateRL(&p)
//       } else {
//          rotateL(&p)
//       }
//    }
//    return p
// }
//
// func (tree WeightBalancedDownUp) extractMax(p **Node, s Size) Value {
//    if p == nil {
//       return p
//    }
//    p = p.Copy()
//
//    sl := p.s
//    sr := p.sizeR(s)
//
//    p.r = tree.extractMax(p.r, sr, s)
//
//    if !tree.Balanced(sr - 1, sl) {
//       if !tree.singleRotation(p.l.s, p.l.sizeR(sl)) {
//          rotateLR(&p)
//       } else {
//          rotateR(&p)
//       }
//    }
//    return p
// }

//

//
// This method removes the node p from its tree.
// It is replaced by either its inorder successor or predecessor.
//
// The dissolved node is written, and the root returned.
//
//
//
//
// func (tree *LBSTDownUp) dissolve(p *Node, s Size) *Node {
//    // *dissolved = p
//    return tree.join(p.l, p.r, p.s, p.sizeR(s))
//    //
//    // sl := p.s
//    // sr := p.sizeR(s)
//    //
//      assert(tree.Balanced(sl, sr))
//      assert(tree.Balanced(sr, sl))
//    //
//    // *dissolved = p // yeah man, this is the node is that going bye bye!
//    //
//    // if p.l == nil { return p.r }
//    // if p.r == nil { return p.l }
//    //
//    // if sr > sl {
//    //    max := tree.deleteMax(&p.l, sl)
//    //    max.l = p.l
//    //    max.r = p.r
//    //    max.s = sl
//    //    return max
//    //
//    //    // p.s =
//    //    // root, dissolved = tree.deleteMin(p.r, sr, &p.s)
//    // } else {
//    //    // root, dissolved = tree.deleteMax(p.l, sl, &p.s)
//    //    // p.s--
//    //
//    // }
// }
//
// func (tree LBSTDownUp) extractMin(p **Node, s Size) (s Value) {
//       *p = tree.delete(*p, s, 0, &s)
//    return
// }
//
// func (tree LBSTDownUp) extractMax(p **Node, s Size) (s Value) {
//    *p = tree.delete(*p, s, s-1, &s)
//    return
// }

// wbst extract_maximum(wbst t, wbst *u)
// { size sA, sB;
//  if (t->r == null) { *u = t; t = t->l; }
//   else
//    { t->r = extract_maximum(t->r, u);
//      t->s--;
//      if (!balanced(t->r->s, t->l->s)) t = inc_right(t);
//    }
//  return t;
// }
//
// wbst extract_minimum(wbst t, wbst *u)
// { size sA, sB;
//  if (t->l == null) { *u = t; t = t->r; }
//   else
//    { t->l = extract_minimum(t->l, u);
//      t->s--;
//      if (!balanced(t->l->s, t->r->s)) t = inc_left(t);
//    }
//  return t;
// }
//

//
// func (tree LBSTDownUp) Join(that List) List {
//    l := tree
//    r := that.(*LBSTDownUp)
//    return &LBSTDownUp{
//       BST: BST{
//          root: tree.join(l.root.share(), r.root.share(), l.size, r.size),
//          size: l.size + r.size,
//       },
//    }
// }
//
// func (tree LBSTDownUp) build(s Value, l, r *Node, sl, sr Size) *Node {
//    if sl <= sr {
//       return tree.joinR(s, l, r, sl, sr)
//    } else {
//       return tree.joinL(s, l, r, sl, sr)
//    }
// }
//
// func (tree *LBSTDownUp) joinL(s Value, l, r *Node, sl, sr Size) *Node {
//    if tree.Balanced(sr, sl) {
//       return &Node{
//          s: s,
//          l: l,
//          r: r,
//          s: sl,
//       }
//    }
//    l = pathcopy(l)
//
//    sll := l.s
//    slr := l.sizeR(sl)
//
//    l.r = tree.joinL(s, l.r, r, slr, sr)
//    slr = 1 + sr + slr
//
//    if !tree.Balanced(sll, slr) {
//       if tree.singleRotation(l.r.sizeR(slr), l.r.s) {
//          l = l.rotateL()
//       } else {
//          l = l.rotateRL()
//       }
//    }
//    return l
// }
//
// func (tree *LBSTDownUp) joinR(s Value, l, r *Node, sl, sr Size) *Node {
//    if tree.Balanced(sl, sr) {
//       return &Node{ // TODO: Why does this have to copy?
//          s: s,
//          l: l,
//          r: r,
//          s: sl,
//       }
//    }
//    r = pathcopy(r)
//
//    srl := r.s
//    srr := r.sizeR(sr)
//
//    r.l = tree.joinR(s, l, r.l, sl, srl)
//    r.s = 1 + sl + srl
//
//    if !tree.Balanced(srr, r.s) {
//       if tree.singleRotation(r.l.s, r.l.sizeR(r.s)) {
//          r = r.rotateR()
//       } else {
//          r = r.rotateLR()
//       }
//    }
//    return r
// }
//
// func (tree LBSTDownUp) split(p *Node, i, s Size) (l, r *Node) {
//    if p == nil {
//       return
//    }
//    p = p.Copy()
//    sl, sr := p.sizeLR(s)
//
//    if i <= (*p).s {
//       l, r = tree.split(p.l, i, sl)
//       r = tree.build(p.s, r, p.r, sl-i, sr)
//    } else {
//       l, r = tree.split(p.r, i-sl-1, sr)
//       l = tree.build(p.s, p.l, l, sl, i-sl-1)
//    }
//    return l, r
// }
//
// func (tree LBSTDownUp) Split(i Position) (List, List) {
//    assert(i <= tree.Size())
//
//    l, r := tree.split(tree.root.share(), i, tree.size)
//
//    return &LBSTDownUp{BST: BST{root: l, size: i}},
//       &LBSTDownUp{BST: BST{root: r, size: tree.size - i}}
// }
