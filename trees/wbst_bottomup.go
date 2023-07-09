package trees

import (
   "bst/abstract/list"
)

func (tree WBSTBottomUp) verifyBalance(p *Node, s list.Size) {
   if p == nil {
      return
   }
   sl := p.sizeL()
   sr := p.sizeR(s)

   invariant(tree.isBalanced(sl, sr))
   invariant(tree.isBalanced(sr, sl))

   tree.verifyBalance(p.l, sl)
   tree.verifyBalance(p.r, sr)
}

func (tree WBSTBottomUp) verifyHeight(root *Node, size list.Size) {
   // Max height?
}

func (tree WBSTBottomUp) Verify() {
   tree.verifySizes()
   tree.verifyBalance(tree.root, tree.size)
   tree.verifyHeight(tree.root, tree.size)
}


func (WBSTBottomUp) isBalanced(x, y list.Size) bool {
   return 3 * (x + 1) >= (y + 1)
}

func (WBSTBottomUp) singleRotation(x, y list.Size) bool {
   return 2 * (x + 1) > (y + 1)
}

func (tree WBSTBottomUp) join(l *Node, r *Node, sl, sr list.Size) (k *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if sl <= sr {
      r = tree.extractMin(r, sr, &k)
      return tree.build(l, k, r, sl, sr-1)
   } else {
      l = tree.extractMax(l, sl, &k)
      return tree.build(l, k, r, sl-1, sr)
   }
}

func (tree WBSTBottomUp) extractMin(p *Node, s list.Size, x **Node) *Node {
   tree.persist(&p)
   if p.l == nil {
      *x = p
      p = p.r
      return p
   }
   sl := p.s
   sr := s - p.s - 1

   p.l = tree.extractMin(p.l, p.s, x)
   p.s--

   if !tree.isBalanced(sl-1, sr) {
      srl := (*p).r.s
      srr := sr - (*p).r.s - 1
      //
      if tree.singleRotation(srr, srl) {
         tree.rotateL(&p)
      } else {
         tree.rotateRL(&p)
      }
   }
   return p
}

func (tree WBSTBottomUp) extractMax(p *Node, s list.Size, x **Node) *Node {
   tree.persist(&p)
   if p.r == nil {
      *x = p
      p = p.l
      return p
   }
   sl := p.s
   sr := s - p.s - 1

   p.r = tree.extractMax(p.r, sr, x)
   if !tree.isBalanced(sr-1, sl) {
      if tree.singleRotation((*p).l.s, sl-(*p).l.s-1) {
         tree.rotateR(&p)
      } else {
         tree.rotateLR(&p)
      }
   }
   return p
}


func (tree WBSTBottomUp) Join(that list.List) list.List {
   l := tree
   r := that
   tree.share(l.root)
   tree.share(r.(*WBSTBottomUp).root)
   return &WBSTBottomUp{
      Tree{
         arena: tree.arena,
         root:  tree.join(l.root, r.(*WBSTBottomUp).root, l.size, r.(*WBSTBottomUp).size),
         size:  l.size + r.(*WBSTBottomUp).size,
      },
   }
}

func (tree WBSTBottomUp) build(l, p, r *Node, sl, sr list.Size) *Node {
   if sl <= sr { // TODO: consider == here?
      return tree.buildR(p, l, r, sl, sr)
   } else {
      return tree.buildL(p, l, r, sl, sr)
   }
}

func (tree *WBSTBottomUp) buildL(p *Node, l, r *Node, sl, sr list.Size) *Node {
   if tree.isBalanced(sr, sl) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.persist(&l)

   sll := l.s
   slr := sl - l.s - 1

   l.r = tree.buildL(p, l.r, r, slr, sr)
   slr = 1 + sr + slr

   if !tree.isBalanced(sll, slr) {

      srr := slr - l.r.s - 1
      srl := l.r.s

      if tree.singleRotation(srr, srl) {
         tree.rotateL(&l)
      } else {
         tree.rotateRL(&l)
      }
   }
   return l
}

func (tree *WBSTBottomUp) buildR(p *Node, l, r *Node, sl, sr list.Size) *Node {
   if tree.isBalanced(sl, sr) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.persist(&r)

   srl := r.s
   srr := sr - r.s - 1

   r.l = tree.buildR(p, l, r.l, sl, srl)
   r.s = 1 + sl + srl

   if !tree.isBalanced(srr, r.s) {
      if tree.singleRotation(r.l.s, r.s-r.l.s-1) {
         tree.rotateR(&r)
      } else {
         tree.rotateLR(&r)
      }
   }
   return r
}

func (tree *WBSTBottomUp) split(p *Node, i, s list.Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.persist(&p)

   sl := p.s
   sr := s - p.s - 1

   if i <= (*p).s {
      l, r = tree.split(p.l, i, sl)
      r = tree.build(r, p, p.r, sl-i, sr)
   } else {
      l, r = tree.split(p.r, i-sl-1, sr)
      l = tree.build(p.l, p, l, sl, i-sl-1)
   }
   return l, r
}

func (tree *WBSTBottomUp) Split(i list.Position) (list.List, list.List) {
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)

   return &WBSTBottomUp{Tree{arena: tree.arena, root: l, size: i}},
          &WBSTBottomUp{Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

//
// import (
//    . "trees/abstract/list"
//    "trees/math"
// )
//
// type WBST struct {
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
// func (WBST) New(s ...Value) List {
//    return &WBST{BST: BST{}.ofArray(s)}
// }
//
//
//
// func (tree WBST) Clone() List {
//    tree.root.share()
//    return &tree
// }
//
//
//
//
//
// func (tree WBST) verify(root *Node, size Size) {
//    tree.verifyPosition(root)
//    tree.verifyWeight(root, size)
//    tree.verifyHeight(root, size)
// }
//
// func (tree WBST) Verify() {
//    tree.verify(tree.root, tree.size)
// }
//
// func (tree WBST) verifyHeight(root *Node, size Size) {
//    invariant(root.height() <= MaximumPathLength(2 * math.Log2(size)))
// }
//
// func (tree WBST) verifyWeight(p *Node, s Size) {
//    tree.SizeBalance.verify(p, s)
// }
//
//
//
// func (tree *WBST) insert(p *Node, s Size, i Position, s Value) *Node {
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
// func (tree WBST) extractMin(p **Node, s Size) (s Value) {
//    *p = tree.delete(*p, s, 0, &s)
//    return
// }
//
// func (tree WBST) extractMax(p **Node, s Size) (s Value) {
//    *p = tree.delete(*p, s, s - 1, &s)
//    return
// }
// func (tree *WBST) dissolve(p *Node, s Size, s *Value) *Node {
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
// func (tree WBST) join(l *Node, r *Node, sl, sr Size) *Node {
//    if l == nil { return r }
//    if r == nil { return l }
//    if sl <= sr {
//       return tree.build(tree.extractMin(&r, sr), l, r, sl, sr - 1)
//    } else {
//       return tree.build(tree.extractMax(&l, sl), l, r, sl - 1, sr)
//    }
// }
//
// func (tree *WBST) delete(p *Node, s Size, i Position, s *Value) *Node {
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
// func (tree *WBST) Insert(i Position, s Value) {
//    assert(i <= tree.size)
//    tree.root = tree.insert(tree.root, tree.size, i, s)
//    tree.size++
// }
//
// func (tree *WBST) Delete(i Position) (s Value) {
//    assert(i < tree.size)
//    tree.root = tree.delete(tree.root, tree.size, i, &s)
//    tree.size--
//    return
// }
//
// func (tree WBST) Join(that List) List {
//    l := tree
//    r := that.(*WBST)
//    return &WBST{
//       BST: BST{
//          root: tree.join(l.root.share(), r.root.share(), l.size, r.size),
//          size: l.size + r.size,
//       },
//    }
// }
//
// func (tree WBST) build(s Value, l, r *Node, sl, sr Size) *Node {
//    if sl <= sr {
//       return tree.joinR(s, l, r, sl, sr)
//    } else {
//       return tree.joinL(s, l, r, sl, sr)
//    }
// }
//
// func (tree *WBST) joinL(s Value, l, r *Node, sl, sr Size) *Node {
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
// func (tree *WBST) joinR(s Value, l, r *Node, sl, sr Size) *Node {
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
// func (tree WBST) split(p *Node, i, s Size) (l, r *Node){
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
// func (tree WBST) Split(i Position) (List, List) {
//    assert(i <= tree.size)
//
//    l, r := tree.split(tree.root.share(), i, tree.size)
//
//    return &WBST{BST: BST{root: l, size: i}},
//           &WBST{BST: BST{root: r, size: tree.size - i}}
// }

type WBSTBottomUp struct {
   Tree
}

// Determines if the weights of two subtrees are balanced.
//
// For the scapegoat tree, we use the logarithmic balance rule, where the
// discrete binary logarithm of the two sizes differ by at most 1.
//
//   -1 <= floor(log₂(s)) - floor(log₂(y)) <= 1
//
// Considering the binary representation, we look at the left-most bit set to 1,
// the most-significant bit. 00001101
//
//   ↖
//    MSB at position 3
//
// The position of the most-significant bit, from the right starting at zero,
// is equal to the floor of the binary logarithm. The balance test is then to
// determine if the MSB is at most one step different in either direction.
// For example:
//
//          ↓                       ↓                    ↓
//   s:   00100000              00001000              00010001
//   y:   00111001              00010001              01001101
//          ↑ BALANCED             ↑ BALANCED          ↑ NOT BALANCED
//
// When `s` <= `y`, as in the examples above, we can shift the bits of `y` one
// position to the right, then, if the MSB of `s` is less than the MSB of `y`
// it means that the MSB of `y` was at least 2 positions ahead - not balance.
//
// Complexity: O(1)
func (WBSTBottomUp) New() list.List {
   return &WBSTBottomUp{}
}

func (tree *WBSTBottomUp) Clone() list.List {
   return &WBSTBottomUp{
      Tree: tree.Tree.Clone(),
   }
}

func (tree *WBSTBottomUp) insert(p *Node, s list.Size, i list.Position, x list.Data) *Node {
   if p == nil {
      return tree.allocate(Node{x: x})
   }
   tree.persist(&p)
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

func (tree *WBSTBottomUp) delete(p *Node, s list.Size, i list.Position, x *list.Data) *Node {
   tree.persist(&p)

   sl := p.s
   sr := s - p.s - 1

   assert(tree.isBalanced(sl, sr))
   assert(tree.isBalanced(sr, sl))

   if i == p.s {
      defer tree.free(p)
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

func (tree *WBSTBottomUp) Select(i list.Size) list.Data {
   assert(i < tree.size)
   return tree.lookup(tree.root, i)
}

func (tree *WBSTBottomUp) Update(i list.Size, x list.Data) {
   assert(i < tree.size)
   tree.persist(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *WBSTBottomUp) Insert(i list.Position, x list.Data) {
   assert(i <= tree.size)
   tree.root = tree.insert(tree.root, tree.size, i, x)
   tree.size++
}

func (tree *WBSTBottomUp) Delete(i list.Position) (x list.Data) {
   assert(i < tree.size)
   tree.root = tree.delete(tree.root, tree.size, i, &x)
   tree.size--
   return
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
// func (tree *WBSTDownUp) dissolve(p *Node, s Size) *Node {
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
// func (tree WBSTDownUp) extractMin(p **Node, s Size) (s Value) {
//       *p = tree.delete(*p, s, 0, &s)
//    return
// }
//
// func (tree WBSTDownUp) extractMax(p **Node, s Size) (s Value) {
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
// func (tree WBSTDownUp) Join(that List) List {
//    l := tree
//    r := that.(*WBSTDownUp)
//    return &WBSTDownUp{
//       BST: BST{
//          root: tree.join(l.root.share(), r.root.share(), l.size, r.size),
//          size: l.size + r.size,
//       },
//    }
// }
//
// func (tree WBSTDownUp) build(s Value, l, r *Node, sl, sr Size) *Node {
//    if sl <= sr {
//       return tree.joinR(s, l, r, sl, sr)
//    } else {
//       return tree.joinL(s, l, r, sl, sr)
//    }
// }
//
// func (tree *WBSTDownUp) joinL(s Value, l, r *Node, sl, sr Size) *Node {
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
// func (tree *WBSTDownUp) joinR(s Value, l, r *Node, sl, sr Size) *Node {
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
// func (tree WBSTDownUp) split(p *Node, i, s Size) (l, r *Node) {
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
// func (tree WBSTDownUp) Split(i Position) (List, List) {
//    assert(i <= tree.size)
//
//    l, r := tree.split(tree.root.share(), i, tree.size)
//
//    return &WBSTDownUp{BST: BST{root: l, size: i}},
//       &WBSTDownUp{BST: BST{root: r, size: tree.size - i}}
// }

