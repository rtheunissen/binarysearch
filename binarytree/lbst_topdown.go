package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/utility"
)

type LBSTTopDown struct {
   LBST
}

func (LBSTTopDown) New() List {
   return &LBSTTopDown{}
}

func (tree *LBSTTopDown) Clone() List {
   return &LBSTTopDown{LBST{Tree: tree.Tree.Clone()}} // TODO: format
}

func (tree LBSTTopDown) verifyBalance(p *Node, s Size) {
   if p == nil {
      return
   }
   sl := p.sizeL()
   sr := p.sizeR(s)

   invariant(utility.Difference(utility.Log2(sl + 1), utility.Log2(sr + 1)) <= 1)

   tree.verifyBalance(p.l, sl)
   tree.verifyBalance(p.r, sr)
}

func (tree LBSTTopDown) verifyHeight(root *Node, size Size) {
   invariant(root.height() <= int(2*utility.Log2(size)))
}

func (tree LBSTTopDown) Verify() {
   tree.verifySizes()
   tree.verifyBalance(tree.root, tree.size)
   tree.verifyHeight(tree.root, tree.size)
}

func (tree *LBSTTopDown) isBalanced(x, y Size) bool {
   return tree.LBST.isBalanced(x + 1, y + 1)
}

func (tree *LBSTTopDown) singleRotation(x, y Size) bool {
   return tree.LBST.singleRotation(x + 1, y + 1)
}

func (tree *LBSTTopDown) insert(p **Node, s Size, i Position, x Data) {
   assert(i <= s)
   assert(s == (*p).size())
   for {
      if *p == nil {
         *p = tree.allocate(Node{x: x})
         return
      }
      tree.persist(p)

      sl := (*p).sizeL()
      sr := (*p).sizeR(s)

      assert(tree.isBalanced(sr, sl))
      assert(tree.isBalanced(sl, sr))

      if i <= (*p).s {
         if tree.isBalanced(sr, sl+1) {
            //
            // L BALANCED
            //
            tree.pathL(&p, &s)
         } else {
            if i <= (*p).l.s {
               if tree.singleRotation((*p).l.sizeL()+1, (*p).l.sizeR(sl)) {
                  //
                  // LL SINGLE
                  //
                  tree.rotateR(p)
                  tree.pathL(&p, &s)
               } else {
                  //
                  // LL DOUBLE
                  //
                  tree.rotateLR(p)
                  tree.pathL(&p, &s)
                  tree.pathL(&p, &s)
               }
            } else {
               if tree.singleRotation((*p).l.sizeL(), (*p).l.sizeR(sl)+1) {
                  //
                  // LR SINGLE
                  //
                  tree.rotateR(p)
                  tree.pathR(&p, &s, &i)
                  tree.pathL(&p, &s)
               } else {
                  if i <= (*p).l.s+(*p).l.r.s+1 {
                     //
                     // LRL DOUBLE
                     //
                     tree.rotateLR(p)
                     tree.pathL(&p, &s)
                     tree.pathR(&p, &s, &i)
                  } else {
                     //
                     // LRR DOUBLE
                     //
                     tree.rotateLR(p)
                     tree.pathR(&p, &s, &i)
                     tree.pathL(&p, &s)
                  }
               }
            }
         }
      } else {
         //
         // R BALANCED
         //
         if tree.isBalanced(sl, sr+1) {
            tree.pathR(&p, &s, &i)
            continue
         }
         if i > (*p).s+(*p).r.s+1 {
            if tree.singleRotation((*p).r.sizeR(sr)+1, (*p).r.sizeL()) {
               //
               // RR SINGLE
               //
               tree.rotateL(p)
               tree.pathR(&p, &s, &i)
            } else {
               //
               // RR DOUBLE
               //
               tree.rotateRL(p)
               tree.pathR(&p, &s, &i)
               tree.pathR(&p, &s, &i)
            }
         } else {
            if tree.singleRotation((*p).r.sizeR(sr), (*p).r.sizeL()+1) {
               //
               // RL SINGLE
               //
               tree.rotateL(p)
               tree.pathL(&p, &s)
               tree.pathR(&p, &s, &i)
            } else {
               if i > (*p).s+(*p).r.l.s+1 {
                  //
                  // RLR DOUBLE
                  //
                  tree.rotateRL(p)
                  tree.pathR(&p, &s, &i)
                  tree.pathL(&p, &s)
               } else {
                  //
                  // RLL DOUBLE
                  //
                  tree.rotateRL(p)
                  tree.pathL(&p, &s)
                  tree.pathR(&p, &s, &i)
               }
            }
         }
      }
   }
}

func (tree *LBSTTopDown) delete(p **Node, s Size, i Position) (deleted *Node) {
   assert(i < s)
   assert(s == (*p).size())
   for {
      tree.persist(p)

      sl := (*p).s
      sr := s - (*p).s - 1

      assert(tree.isBalanced(sl, sr))
      assert(tree.isBalanced(sr, sl))

      if i == (*p).s {
         defer tree.free(*p)
         x := *p
         *p = tree.join((*p).l, (*p).r, sl, sr)
         return x
      }
      if i <= (*p).s {
         if tree.isBalanced(sl-1, sr) {
            //
            // L BALANCED
            //
            tree.deleteL(&p, &s)
         } else {
            if tree.singleRotation(sr-(*p).r.s-1, (*p).r.s) {
               //
               // L SINGLE
               //
               tree.rotateL(p)
               tree.deleteL(&p, &s)
               tree.deleteL(&p, &s)
            } else {
               //
               // L DOUBLE
               //
               tree.rotateRL(p)
               tree.deleteL(&p, &s)
               tree.deleteL(&p, &s)
            }
         }
      } else {
         if tree.isBalanced(sr-1, sl) {
            //
            // R BALANCED
            //
            tree.deleteR(&p, &s, &i)
         } else {
            if tree.singleRotation((*p).l.s, sl-(*p).l.s-1) {
               //
               // R SINGLE
               //
               tree.rotateR(p)
               tree.deleteR(&p, &s, &i)
               tree.deleteR(&p, &s, &i)
            } else {
               //
               // R DOUBLE
               //
               tree.rotateLR(p)
               tree.deleteR(&p, &s, &i)
               tree.deleteR(&p, &s, &i)
            }
         }
      }
   }
}

func (tree *LBSTTopDown) rebalanceR(p **Node, sr Size) {
   if tree.singleRotation(sr-(*p).r.s-1, (*p).r.s) {
      tree.rotateL(p)
   } else {
      tree.rotateRL(p)
   }
}

func (tree *LBSTTopDown) rebalanceL(p **Node, sl Size) {
   if tree.singleRotation((*p).l.s, sl-(*p).l.s-1) { // R SINGLE
      tree.rotateR(p)
   } else { // R DOUBLE
      tree.rotateLR(p)
   }
}

func (tree *LBSTTopDown) pathL(p ***Node, s *Size) {
   *s = (**p).s
   (**p).s++
   *p = &(**p).l
}

func (tree *LBSTTopDown) deleteL(p ***Node, s *Size) {
   *s = (**p).s
   (**p).s--
   *p = &(**p).l
}

func (tree *LBSTTopDown) pathR(p ***Node, s *Size, i *Position) {
   *s = *s - (**p).s - 1
   *i = *i - (**p).s - 1
   *p = &(**p).r
}

func (tree *LBSTTopDown) deleteR(p ***Node, s *Size, i *Position) {
   *s = *s - (**p).s - 1
   *i = *i - (**p).s - 1
   *p = &(**p).r
}

func (tree *LBSTTopDown) Insert(i Position, x Data) {
   assert(i <= tree.size)
   tree.insert(&tree.root, tree.size, i, x)
   tree.size++
}

func (tree *LBSTTopDown) Delete(i Position) (x Data) {
   assert(i < tree.size)
   x = tree.delete(&tree.root, tree.size, i).x
   tree.size--
   return
}

func (tree *LBSTTopDown) Join(that List) List {
   l := tree
   r := that.(*LBSTTopDown)
   tree.share(l.root)
   tree.share(r.root)
   return &LBSTTopDown{
      LBST{
         Tree: Tree{
            arena: tree.arena,
            root:  tree.join(l.root, r.root, l.size, r.size),
            size:  l.size + r.size,
         },
      },
   }
}
func (tree *LBSTTopDown) deleteMin(p **Node, s Size) *Node {
   return tree.delete(p, s, 0)
}

func (tree *LBSTTopDown) deleteMax(p **Node, s Size) *Node {
   return tree.delete(p, s, s-1)
}

func (tree *LBSTTopDown) join(l *Node, r *Node, sl, sr Size) *Node {
   if l == nil { return r }
   if r == nil { return l }
   if sl <= sr {
      return tree.build(l, tree.deleteMin(&r, sr), r, sl, sr-1)
   } else {
      return tree.build(l, tree.deleteMax(&l, sl), r, sl-1, sr)
   }
}

func (tree *LBSTTopDown) build(l, p, r *Node, sl, sr Size) *Node {
   if sl <= sr {
      return tree.buildR(p, l, r, sl, sr)
   } else {
      return tree.buildL(p, l, r, sl, sr)
   }
}

func (tree *LBSTTopDown) buildL(p *Node, l, r *Node, sl, sr Size) *Node {
   if tree.isBalanced(sr, sl) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.persist(&l)
   l.r = tree.buildL(p, l.r, r, sl-l.s-1, sr)
   if !tree.isBalanced(l.s, sl+sr-l.s) {
      tree.rebalanceR(&l, sr+sl-l.s)
   }
   return l
}

func (tree *LBSTTopDown) buildR(p *Node, l, r *Node, sl, sr Size) *Node {
   if tree.isBalanced(sl, sr) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.persist(&r)
   r.l = tree.buildR(p, l, r.l, sl, r.s)
   r.s = 1 + sl + r.s
   if !tree.isBalanced(sl+sr-r.s, r.s) {
      tree.rebalanceL(&r, r.s)
   }
   return r
}

func (tree LBSTTopDown) split(p *Node, i, s Size) (l, r *Node) {
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

func (tree LBSTTopDown) Split(i Position) (List, List) {
   assert(i <= tree.size)
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)

   return &LBSTTopDown{LBST{Tree: Tree{arena: tree.arena, root: l, size: i}}},
          &LBSTTopDown{LBST{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}
}

