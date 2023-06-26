package binarytree

import . "binarysearch/abstract/list"

type AVLWeakTopDown struct {
   WAVL
}

func (AVLWeakTopDown) New() List {
   return &AVLWeakTopDown{}
}

func (tree *AVLWeakTopDown) Clone() List {
   return &AVLWeakTopDown{
      WAVL{
         Tree: tree.Tree.Clone(),
      },
   }
}

// This top-down insertion algorithm was translated and paraphrased from the
// _Deletion Without Rebalancing in Binary Search Trees_ paper referenced above.
func (tree *AVLWeakTopDown) insert(p **Node, i Position, x Data) {
   //
   // "If the tree is empty, create a new node containing the item to be inserted
   //  and make it the root, completing the insertion."
   //
   if *p == nil {
      tree.attach(p, x)
      return
   }
   tree.persist(p)
   //
   // "Otherwise, tree.promote the root if it is not 1,1."
   //
   // TODO: ^ this contradicts?
   //
   if tree.isOneOne(*p) {
      tree.promote(*p)
   }
   // "This establishes the invariant for the main loop for the algorithm:
   //  *p is a non-nil node that is not a 1,1-node."
   //
   for {
      // assert(!tree.isOneOne(*p))

      // "From *p, take one step down the search path..."
      //
      if i <= (*p).s {
         //
         // LEFT
         //
         // "If the next node on the search path is nil, replace it by a new
         //  node containing the item to be inserted, completing the insertion.
         //
         //  The new node cannot be a 0-child since the parent is not a 1,1-node
         //  and hence has positive rank."
         //
         if (*p).l == nil {
            tree.attachL(*p, x)
            return
         }
         //
         // "If the next node on the search path is not a 1,1-node, continue."
         //
         if !tree.isOneOne((*p).l) {
            tree.pathLeft(&p)
            continue
         }
         // "If the next node on the search path is not a 1-child, tree.promote it,
         //  then continue to the next step."
         //
         if !tree.isOneChild(*p, (*p).l) {
            tree.pathLeft(&p)
            tree.promote(*p)
            continue
         }
         // "In the remaining cases, the next node is a 1,1-node and a 1-child."
         //
         // assert(tree.isOneOne((*p).l) && tree.isOneChild(*p, (*p).l))
         //
         // "From this node, take one further step down the search path..."
         //
         if i <= (*p).l.s {
            //
            // LEFT LEFT
            //
            // "If this node is nil, replace it by a new node containing the
            //  item to be inserted. If the new node and its parent are both
            //  left children, or, symmetrically, both right children, do a
            //  rotate step, completing the insertion."
            //
            if (*p).l.l == nil {
               tree.attachLL(*p, x)
               tree.rotateR(p)
               tree.promote(*p)
               tree.demote((*p).r)
               return
            }
            //
            // "If the new node is a right child and its parent a left child, or
            //  symmetrically if the new node is a left child and its parent a
            //  right child, do a double rotate step, completing the insertion."
            //
            //  ^That is not the case here because we know this is a left-left.
            //
            // "If this node is not a 1,1-node, continue with both search steps."
            //
            if !tree.isOneOne((*p).l.l) {
               tree.pathLeft(&p)
               tree.pathLeft(&p)
               continue
            }
            // "Otherwise tree.promote the new node and its parent, making its parent
            //  a 0-child, then do a rotate or double rotate step to make all
            //  rank differences positive."
            tree.rotateR(p)
            tree.promote(*p)
            tree.demote((*p).r)
            tree.pathLeft(&p)
            tree.promote(*p)
            continue

         } else {
            //
            // LEFT RIGHT
            //
            // "If the new node is a right child and its parent a left child, or
            //  symmetrically if the new node is a left child and its parent a
            //  right child, do a double rotate step, completing the insertion."
            //
            // ^That is the case here because we know this is a left-right step,
            //  which requires a double rotation, follows by the right and left
            //  steps down the search path after the rotation.
            //
            if (*p).l.r == nil {
               tree.attachLR(*p, x)
               tree.rotateLR(p)
               tree.promote(*p)
               tree.demote((*p).r)
               return
            }
            //
            // "If this node is not a 1,1-node, continue with both search steps."
            //
            if !tree.isOneOne((*p).l.r) {
               tree.pathLeft(&p)
               tree.pathRight(&p, &i)
               continue
            }
            // "Otherwise tree.promote the new node and its parent, making its parent
            //  a 0-child, then do a rotate or double rotate step to make all
            //  rank differences positive."
            //
            tree.rotateLR(p)
            tree.promote(*p)
            tree.promote(*p)
            tree.demote((*p).r)
            //
            // "If a double rotation is done, take one further step down the
            //  search path after the rotation. Ths completes the step."
            //
            if i <= (*p).s {
               tree.pathLeft(&p) // LRL
            } else {
               tree.pathRight(&p, &i) // LRR
            }
         }
      } else {
         //
         // RIGHT
         //
         // Comments follow symmetrically from above.
         //
         if (*p).r == nil {
            tree.attachR(*p, x)
            return
         }
         if !tree.isOneOne((*p).r) {
            tree.pathRight(&p, &i)
            continue
         }
         if !tree.isOneChild(*p, (*p).r) {
            tree.pathRight(&p, &i)
            tree.promote(*p)
            continue
         }

         if i > (*p).s+(*p).r.s+1 {
            //
            // RIGHT RIGHT
            //
            if (*p).r.r == nil {
               tree.attachRR(*p, x)
               tree.rotateL(p)
               tree.promote(*p)
               tree.demote((*p).l)
               return
            }
            if !tree.isOneOne((*p).r.r) {
               tree.pathRight(&p, &i)
               tree.pathRight(&p, &i)
               continue
            }
            tree.rotateL(p)
            tree.promote(*p)
            tree.demote((*p).l)
            tree.pathRight(&p, &i)
            tree.promote(*p)
            continue

         } else {
            //
            // RIGHT LEFT
            //
            if (*p).r.l == nil {
               tree.attachRL(*p, x)
               tree.rotateRL(p)
               tree.demote((*p).l)
               tree.promote(*p)
               return
            }
            if !tree.isOneOne((*p).r.l) {
               tree.pathRight(&p, &i)
               tree.pathLeft(&p)
               continue
            }
            tree.rotateRL(p)
            tree.promote(*p)
            tree.promote(*p)
            tree.demote((*p).l)

            if i > (*p).s {
               tree.pathRight(&p, &i) // RLR
            } else {
               tree.pathLeft(&p) // RLL
            }
         }
      }
   }
}

func (tree *AVLWeakTopDown) dissolve(p **Node, x *Data) {
   tree.persist(p)
   defer tree.free(*p)
   *x = (*p).x
   *p = tree.join((*p).l, (*p).r, (*p).s)
}

func (tree *AVLWeakTopDown) Update(i Size, x Data) {
   // assert(i < tree.size)
   tree.persist(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *AVLWeakTopDown) Select(i Size) Data {
   // assert(i < tree.size)
   return tree.lookup(tree.root, i)
}

func (tree *AVLWeakTopDown) Insert(i Position, x Data) {
   // assert(i <= tree.size)
   tree.size++
   tree.insert(&tree.root, i, x)
}

func (tree *AVLWeakTopDown) Delete(i Position) (x Data) {
   // assert(i < tree.size)
   x = tree.delete(&tree.root, i)
   tree.size--
   return
}

func (tree AVLWeakTopDown) join(l, r *Node, sl Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) <= tree.rank(r) {
      return tree.build(l, tree.deleteMin(&r), r, sl)
   } else {
      return tree.build(l, tree.deleteMax(&l), r, sl-1)
   }
}

func (tree AVLWeakTopDown) Join(other List) List {
   l := tree
   r := other.(*AVLWeakTopDown)
   tree.share(l.root)
   tree.share(r.root)
   return &AVLWeakTopDown{
      WAVL{
         Tree: Tree{
            arena: tree.arena,
            root:  tree.join(l.root, r.root, l.size),
            size:  l.size + r.size,
         },
      },
   }
}


func (tree AVLWeakTopDown) split(p *Node, i, s Size) (l, r *Node) {
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


func (tree AVLWeakTopDown) Split(i Position) (List, List) {
   // assert(i <= tree.size)
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)

   return &AVLWeakTopDown{WAVL{Tree: Tree{arena: tree.arena, root: l, size: i}}},
          &AVLWeakTopDown{WAVL{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}
}

// "In a deletion, if the current node is 2,2 or it is 1,2 and its 1-child
//
//   is 2,2, we can force a reset on the next search step by demoting the
//   current node in the former case, or the current node and its 1-child
//   in the latter, and rebalancing top-down from the safe node."
func (tree AVLWeakTopDown) resetSafeNode(p *Node) bool {
   if tree.isTwoTwo(p) {
      tree.demote(p)
      return true
   }
   if tree.isTwoChild(p, p.l) && tree.isTwoTwo(p.r) {
      // assert(tree.isOneChild(p, p.r))
      tree.persist(&p.r)
      tree.demote(p)
      tree.demote(p.r)
      return true
   }
   if tree.isTwoChild(p, p.r) && tree.isTwoTwo(p.l) {
      // assert(tree.isOneChild(p, p.l))
      tree.persist(&p.l)
      tree.demote(p)
      tree.demote(p.l)
      return true
   }
   return false // Could not reset the safe node.
}

func (tree AVLWeakTopDown) rebalanceTopDownOnDelete(p **Node) {
   *p = tree.rebalanceOnDelete(*p)
}

// "Deletion of a unary node converts the child that replaces it
//
//   into a 2- or 3-child; the latter violates the rank rule."
func (tree AVLWeakTopDown) rebalanceAfterDissolve(g **Node, p **Node) {
   //
   // "Deletion of a leaf may convert its parent, previously a 1,2 node
   //  into a 2,2 leaf, violating the rank rule. In this case we begin
   //  by demoting the parent, which may make it a 3-child."
   //
   if (*p).isLeaf() && tree.isTwoTwo(*p) {
      tree.demote(*p)
      tree.rebalanceTopDownOnDelete(g)
   } else {
      tree.rebalanceTopDownOnDelete(p)
   }
}

func (tree AVLWeakTopDown) delete(p **Node, i Position) (x Data) {
   //
   // Deleting the root does not require any rebalancing steps because a join
   // will always produce a valid tree.
   //
   if (*p).s == i {
      tree.dissolve(p, &x)
      return
   }
   // This is the parent of the current node; the grandparent and "safe node".
   g := p
   for {
      tree.persist(p)
      if tree.resetSafeNode(*p) {
         tree.rebalanceTopDownOnDelete(g)
      }
      if i < (*p).s {
         //
         // LEFT
         //
         l := deleteL(*p)
         if (*l).s == i {
            //
            // The node delete is on the left; path to it then replace it by a
            // join of its subtrees, followed by a rebalancing step.
            //
            tree.dissolve(l, &x)
            tree.rebalanceAfterDissolve(g, p)
            return
         }
         g = p
         p = l

      } else {
         //
         // RIGHT
         //
         r := deleteR(*p, &i)
         if (*r).s == i {
            //
            // The node delete is on the right; path to it then replace it by a
            // join of its subtrees, followed by a rebalancing step.
            //
            tree.dissolve(r, &x)
            tree.rebalanceAfterDissolve(g, p)
            return
         }
         g = p
         p = r
      }
   }
}
func (tree AVLWeakTopDown) deleteMax(p **Node) (max *Node) {
   g := p
   if (*p).r == nil {
      return tree.replacedByLeftSubtree(p)
   }
   for {
      tree.persist(p)
      right := pathDeletingRightIgnoringIndex(*p)
      if tree.resetSafeNode(*p) {
         tree.rebalanceTopDownOnDelete(g)
      }
      if (*p).r.r == nil {
         max = tree.replacedByLeftSubtree(right)
         tree.rebalanceAfterDissolve(g, p)
         return
      }
      g = p
      p = right
   }
}

func (tree AVLWeakTopDown) deleteMin(p **Node) (min *Node) {
   g := p
   if (*p).l == nil {
      return tree.replacedByRightSubtree(p)
   }
   for {
      tree.persist(p)
      left := deleteL(*p)
      if tree.resetSafeNode(*p) {
         tree.rebalanceTopDownOnDelete(g)
      }
      if (*p).l.l == nil {
         min = tree.replacedByRightSubtree(left)
         tree.rebalanceAfterDissolve(g, p)
         return
      }
      g = p
      p = left
   }
}
