package binarytree

import (
   "binarysearch/abstract/list"
)

type RedBlackTopDown struct {
   Tree
   RedBlack
}

func (tree RedBlackTopDown) Verify() {
   tree.verifySize(tree.root, tree.size)
   tree.RedBlack.verifyRanks(tree.root)
   tree.RedBlack.verifyHeight(tree.root, tree.size)
}

func (RedBlackTopDown) New() list.List {
   return &RedBlackTopDown{}
}

func (tree *RedBlackTopDown) Clone() list.List {
   return &RedBlackTopDown{
      Tree: tree.Tree.Clone(),
   }
}

// This top-down insertion algorithm was translated and paraphrased from the
// _Deletion Without Rebalancing in Binary Search Trees_ paper referenced above.
func (tree *RedBlackTopDown) insert(p **Node, i list.Position, x list.Data) {
   //
   // "If the tree is empty, create a new node of rank zero containing the item
   //  to be inserted and make it the root, completing the insertion."
   //
   if *p == nil {
      tree.attach(p, x)
      return
   }
   tree.persist(p)
   //
   // "Otherwise, promote the root if it is 0,0."
   //
   if tree.isZeroZero(*p) {
      tree.promote(*p)
   }
   //
   // "This establishes the invariant for the main loop of the algorithm:
   //  *p is a non-nil node that is not a 0,0-node and not a 0-child."
   //
   for {
      assert(!tree.isZeroZero(*p))
      //
      // "From *p, take one step down the search path..."
      //
      if i <= (*p).sizeL() {
         //
         // LEFT
         //
         if (*p).l == nil {
            tree.attachL(*p, x)
            return
         }
         if !tree.isZeroChild(*p, (*p).l) && !tree.isZeroZero((*p).l) {
            tree.pathLeft(&p)
            continue
         }
         if tree.isZeroZero((*p).l) {
            tree.pathLeft(&p)
            tree.promote(*p)
            continue
         }
         // In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child
         assert(tree.isZeroChild(*p, (*p).l))
         assert(!tree.isZeroChild((*p).l, (*p).l.l))
         assert(!tree.isZeroChild((*p).l, (*p).l.l))

         if i <= (*p).l.sizeL() {
            if (*p).l.l == nil {
               tree.attachLL(*p, x)
               if tree.isZeroChild((*p).l, (*p).l.l) { // or is p.l.rank == 0 ?
                  tree.rotateR(p)
               }
               return
            }
            if !tree.isZeroZero((*p).l.l) {
               tree.pathLeft(&p)
               tree.pathLeft(&p)
               continue
            }
            if !tree.isOneChild((*p).l, (*p).l.l) {
              tree.pathLeft(&p)
              tree.pathLeft(&p)
              tree.promote(*p)
              continue
            }
            tree.rotateR(p)
            tree.pathLeft(&p)
            tree.promote(*p)
            continue

         } else {
            //
            // LEFT RIGHT
            //
            if (*p).l.r == nil {
               tree.attachLR(*p, x)
               if tree.isZeroChild((*p).l, (*p).l.r) { // or is p.l.rank == 0 ?
                  tree.rotateLR(p)
               }
               return
            }
            if !tree.isZeroZero((*p).l.r) {
               tree.pathLeft(&p)
               tree.pathRight(&p, &i)
               continue
            }
            if !tree.isOneChild((*p).l, (*p).l.r) {
              tree.pathLeft(&p)
              tree.pathRight(&p, &i)
              tree.promote(*p)
              continue
            }
            tree.rotateLR(p)
            tree.promote(*p)
            if i <= (*p).sizeL() {
               tree.pathLeft(&p) // LRL
            } else {
               tree.pathRight(&p, &i) // LRR
            }
         }
      } else {
         if (*p).r == nil {
            tree.attachR(*p, x)
            return
         }
         if !tree.isZeroZero((*p).r) && !tree.isZeroChild(*p, (*p).r) {
            tree.pathRight(&p, &i)
            continue
         }
         if tree.isZeroZero((*p).r) {
            tree.pathRight(&p, &i)
            tree.promote(*p)
            continue
         }
         if i > (*p).sizeL() + (*p).r.sizeL() + 1 {
            if (*p).r.r == nil {
               tree.attachRR(*p, x)
               if tree.isZeroChild((*p).r, (*p).r.r) { // or is p.r.rank == 0 ?
                  tree.rotateL(p)
               }
               return
            }
            if !tree.isZeroZero((*p).r.r) {
               tree.pathRight(&p, &i)
               tree.pathRight(&p, &i)
               continue
            }
            if !tree.isOneChild((*p).r, (*p).r.r) {
              tree.pathRight(&p, &i)
              tree.pathRight(&p, &i)
              tree.promote(*p)
              continue
            }
            tree.rotateL(p)
            tree.pathRight(&p, &i)
            tree.promote(*p)
            continue
         } else {
            //
            // RIGHT LEFT
            //
            if (*p).r.l == nil {
               tree.attachRL(*p, x)
               if tree.isZeroChild((*p).r, (*p).r.l) { // or is p.l.rank == 0 ?
                  tree.rotateRL(p)
               }
               return
            }
            if !tree.isZeroZero((*p).r.l) {
               tree.pathRight(&p, &i)
               tree.pathLeft(&p)
               continue
            }
            if !tree.isOneChild((*p).r, (*p).r.l) {
              tree.pathRight(&p, &i)
              tree.pathLeft(&p)
              tree.promote(*p)
              continue
            }
            tree.rotateRL(p)
            tree.promote(*p)
            if i > (*p).sizeL() {
               tree.pathRight(&p, &i) // RLR
            } else {
               tree.pathLeft(&p) // RLL
            }
         }
      }
   }
}

func (tree *RedBlackTopDown) Insert(i list.Position, x list.Data) {
   assert(i <= tree.size)
   tree.insert(&tree.root, i, x)
   tree.size++
}

func (tree *RedBlackTopDown) Delete(i list.Position) (x list.Data) {
   assert(i < tree.size)
   tree.size = tree.size - 1
   tree.root = tree.delete(tree.root, i, &x)
   return x
}

func (tree *RedBlackTopDown) delete(p *Node, i list.Position, x *list.Data) *Node {
   tree.persist(&p)
   if i == p.s {
      *x = p.x
      defer tree.free(p)
      return tree.join(p.l, p.r, p.s)
   }
   if i < p.s {
      p.s = p.s - 1
      p.l = tree.delete(p.l, i, x)
      return tree.balanceDeleteL(p)
   } else {
      p.r = tree.delete(p.r, i-p.s-1, x)
      return tree.balanceDeleteR(p)
   }
}


func (tree RedBlackTopDown) balanceDeleteR(p *Node) *Node {
   //
   //
   //
   if tree.isZeroChild(p, p.r) {
      assert(tree.isOneChild(p.r, p.r.r))
      assert(tree.isOneChild(p.r, p.r.l))
      return p
   }
   //
   //
   //
   if tree.isOneChild(p, p.r) {
      return p
   }
   //
   //
   //
   if tree.isZeroChild(p, p.l) {
      assert(tree.isZeroChild(p, p.l))
      assert(tree.isTwoChild(p, p.r))
      assert(tree.isOneOne(p.l))
      //
      //
      //
      tree.rotateR(&p)
      assert(tree.isZeroChild(p, p.r))
      assert(tree.isOneChild(p, p.l))
      //
      //
      //
      if tree.isZeroChild(p.r.l, p.r.l.l) {
         assert(tree.isOneChild(p.r, p.r.l))
         assert(tree.isTwoChild(p.r, p.r.r))
         tree.rotateR(&p.r)
         tree.promote(p.r)
         tree.demote(p.r.r)
         return p
      }
      //
      //
      //
      if tree.isZeroChild(p.r.l, p.r.l.r) {
         assert(tree.isOneChild(p.r, p.r.l))
         assert(tree.isTwoChild(p.r, p.r.r))
         tree.rotateLR(&p.r)
         tree.promote(p.r)
         tree.demote(p.r.r)
         return p
      }
      //
      //
      //
      assert(tree.isOneChild(p.r, p.r.l))
      assert(tree.isTwoChild(p.r, p.r.r))
      assert(tree.isOneChild(p.r.l, p.r.l.l))
      assert(tree.isOneChild(p.r.l, p.r.l.r))
      tree.demote(p.r)
      return p

   } else {
      //
      //
      //
      assert(tree.isOneChild(p, p.l))
      assert(tree.isTwoChild(p, p.r))
      //
      //
      //
      if tree.isZeroChild(p.l, p.l.l) {
         tree.rotateR(&p)
         tree.promote(p)
         tree.demote(p.r)
         return p
      }
      //
      //
      //
      if tree.isZeroChild(p.l, p.l.r) {
         tree.rotateLR(&p)
         tree.promote(p)
         tree.demote(p.r)
         return p
      }
      //
      //
      //
      assert(tree.isOneOne(p.l))
      tree.demote(p)
      return p
   }
}

func (tree RedBlackTopDown) balanceDeleteL(p *Node) *Node {
   if tree.isZeroChild(p, p.l) {
      if tree.isOneOne(p.l.r) && tree.isOneOne(p.l.l) {
         return p
      }
   }
   if tree.isZeroChild(p, p.l) || tree.isOneChild(p, p.l) {
      return p
   }
   if tree.isOneChild(p, p.r) {
      if tree.isOneChild(p.r, p.r.l) && tree.isOneChild(p.r, p.r.r) {
         tree.demote(p)
         return p
      }
      if tree.isZeroChild(p.r, p.r.r) {
         tree.rotateL(&p)
         tree.promote(p)
         tree.demote(p.l)
         return p
      } else {
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.l)
         return p
      }
   } else {
      tree.rotateL(&p)
      if tree.isOneChild(p.l.r, p.l.r.r) && tree.isOneChild(p.l.r, p.l.r.l) {
         tree.demote(p.l)
         return p
      }
      if tree.isZeroChild(p.l.r, p.l.r.r) {
         tree.rotateL(&p.l)
         tree.promote(p.l)
         tree.demote(p.l.l)
         return p
      }
      tree.rotateRL(&p.l)
      tree.promote(p.l)
      tree.demote(p.l.l)
      return p
   }
}


func (tree RedBlackTopDown) split(p *Node, i, s list.Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.persist(&p)

   sl := p.s
   sr := s - p.s - 1

   if i <= (*p).s {
      l, r = tree.split(p.l, i, sl)
         r = tree.build(r, p, p.r, sl-i)
   } else {
      l, r = tree.split(p.r, i-sl-1, sr)
         l = tree.build(p.l, p, l, sl)
   }
   return l, r
}

func (tree *RedBlackTopDown) Split(i list.Position) (list.List, list.List) {
   assert(i <= tree.size)
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)
   return &RedBlackTopDown{Tree: Tree{arena: tree.arena, root: l, size: i}},
          &RedBlackTopDown{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree *RedBlackTopDown) deleteMin(p *Node, min **Node) *Node {
   tree.persist(&p)
   if p.l == nil {
      *min = p
      return p.r
   }
   p.s = p.s - 1
   p.l = tree.deleteMin(p.l, min)
   return tree.balanceDeleteL(p)
}


func (tree *RedBlackTopDown) deleteMax(p *Node, max **Node) *Node {
   tree.persist(&p)
   if p.r == nil {
      *max = p
      return p.l
   }
   p.r = tree.deleteMax(p.r, max)
   return tree.balanceDeleteR(p)
}

// TODO: Refactor, simplify.
func (tree *RedBlackTopDown) build(l, p, r *Node, sl list.Size) *Node {
   if tree.rank(l) == tree.rank(r) {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(l))
      if l == nil || tree.isZeroChild(l, l.l) || tree.isZeroChild(l, l.r) ||
         r == nil || tree.isZeroChild(r, r.l) || tree.isZeroChild(r, r.r) {
         tree.promote(p)
      }
      return p
   }
   if tree.rank(l) < tree.rank(r) {
      tree.persist(&r)
      r.s = 1 + sl + r.s
      r.l = tree.build(l, p, r.l, sl)
      if tree.isZeroChild(r, r.l) {
         if tree.isZeroChild(r, r.r) {
            if tree.isZeroChild(r.l, r.l.l) {
               tree.promote(r)
               return r
            }
         } else {
            if tree.isZeroChild(r.l, r.l.l) {
               tree.rotateR(&r)
               return r
            }
         }
      }
      return r
   } else {
      tree.persist(&l)
      l.r = tree.build(l.r, p, r, sl-l.s-1)
      if tree.isZeroChild(l, l.r) {
         if tree.isZeroChild(l, l.l) {
            if tree.isZeroChild(l.r, l.r.r) {
               tree.promote(l)
               return l
            }
         } else {
            if tree.isZeroChild(l.r, l.r.r) {
               tree.rotateL(&l)
               return l
            }
         }
      }
      return l
   }
}

func (tree *RedBlackTopDown) join(l, r *Node, sl list.Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) < tree.rank(r) {
      return tree.build(l, p, tree.deleteMin(r, &p), sl)
   } else {
      return tree.build(tree.deleteMax(l, &p), p, r, sl-1)
   }
}

func (tree *RedBlackTopDown) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*RedBlackTopDown).root)
   return &RedBlackTopDown{
      Tree: Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, other.(*RedBlackTopDown).root, tree.size),
         size:  tree.size + other.(*RedBlackTopDown).size,
      },
   }
}
