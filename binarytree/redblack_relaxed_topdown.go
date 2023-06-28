package binarytree

import "binarysearch/abstract/list"

type RedBlackRelaxedTopDown struct {
   RedBlackTopDown
   RedBlackRelaxed
}

func (tree RedBlackRelaxedTopDown) Verify() {
   tree.Tree.verifySize(tree.root, tree.size)
   tree.RedBlackRelaxed.verifyRanks(tree.root)
   tree.RedBlackRelaxed.verifyHeight(tree.root)
}

func (RedBlackRelaxedTopDown) New() list.List {
   return &RedBlackRelaxedTopDown{}
}

func (tree *RedBlackRelaxedTopDown) Clone() list.List {
   return &RedBlackRelaxedTopDown{
      RedBlackTopDown: *tree.RedBlackTopDown.Clone().(*RedBlackTopDown),
   }
}

func (tree *RedBlackRelaxedTopDown) Insert(i list.Position, x list.Data) {
   tree.RedBlackTopDown.Insert(i, x)
}

func (tree *RedBlackRelaxedTopDown) Delete(i list.Position) (x list.Data) {
   return tree.Tree.Delete(i)
}

func (tree *RedBlackRelaxedTopDown) join(l, r *Node, sl list.Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) <= tree.rank(r) {
      return tree.RedBlackTopDown.build(l, tree.Tree.deleteMin(&r), r, sl)
   } else {
      return tree.RedBlackTopDown.build(l, tree.Tree.deleteMax(&l), r, sl-1)
   }
}

func (tree *RedBlackRelaxedTopDown) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*RedBlackRelaxedTopDown).root)
   return &RedBlackRelaxedTopDown{
      RedBlackTopDown: RedBlackTopDown{
         Tree: Tree{
            root:  tree.join(tree.root, other.(*RedBlackRelaxedTopDown).root, tree.size),
            size:  tree.size + other.(*RedBlackRelaxedTopDown).size,
            arena: tree.arena, // TODO maybe leave this nil? should probably have its own
         },
      },
   }
}

func (tree *RedBlackRelaxedTopDown) Split(i list.Position) (list.List, list.List) {
   l, r := tree.RedBlackTopDown.Split(i)
   return &RedBlackRelaxedTopDown{RedBlackTopDown: *l.(*RedBlackTopDown)},
          &RedBlackRelaxedTopDown{RedBlackTopDown: *r.(*RedBlackTopDown)}
}


//
//import (
//   . "binarysearch/abstract/list"
//)
//
//type RedBlackRelaxedTopDown struct {
//   RedBlackTopDown
//   RedBlackRelaxed
//}
//
//func (tree RedBlackRelaxedTopDown) Verify() {
//   tree.Tree.verifySize(tree.root, tree.size)
//   tree.RedBlackRelaxed.verifyRanks(tree.root)
//   tree.RedBlackRelaxed.verifyHeight(tree.root)
//}
//
//func (RedBlackRelaxedTopDown) New() List {
//   return &RedBlackRelaxedTopDown{}
//}
//
//// TODO: make all clone syntax the exact same to avoid inconsistency in results
//func (tree *RedBlackRelaxedTopDown) Clone() List {
//   return &RedBlackRelaxedTopDown{
//      RedBlackTopDown: *tree.RedBlackTopDown.Clone().(*RedBlackTopDown),
//   }
//}

// This top-down insertion algorithm was translated and paraphrased from the
// _Deletion Without Rebalancing in Binary Search Trees_ paper referenced above.
//func (tree *RedBlackRelaxedTopDown) insert(p **Node, i Position, x Data) {
   ////
   //// "If the tree is empty, create a new node of rank zero containing the item
   ////  to be inserted and make it the root, completing the insertion."
   ////
   //if *p == nil {
   //   tree.attach(p, x)
   //   return
   //}
   //tree.persist(p)
   ////
   //// "Otherwise, promote the root if 0,0."
   ////
   //if tree.isZeroZero(*p) {
   //   tree.promote(*p)
   //}
   ////
   //// "This establishes the invariant for the main loop of the algorithm:
   ////  *p is a non-nil node that is not a 0,0-node and not a 0-child.
   //for {
   //   assert(!tree.isZeroZero(*p))
   //   //
   //   // "From *p, take one step down the search path..."
   //   //
   //   if i <= (*p).s {
   //      //
   //      // LEFT
   //      //
   //      // "If the next node on the search path is nil, replace it by a new
   //      //  node of rank 0 containing the item to be inserted. This completes
   //      //  the insertion: the new node may be a 0-child, but *p is not."
   //      //
   //      if (*p).l == nil {
   //         tree.attachL(*p, x)
   //         return
   //      }
   //      if !tree.isZeroZero((*p).l) && !tree.isZeroChild(*p, (*p).l) {
   //         tree.pathLeft(&p)
   //         continue
   //      }
   //      if tree.isZeroZero((*p).l) {
   //         tree.pathLeft(&p)
   //         tree.promote(*p)
   //         continue
   //      }
   //      // In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child
   //      assert(tree.isZeroChild(*p, (*p).l))
   //      assert(!tree.isZeroChild((*p).l, (*p).l.l))
   //      assert(!tree.isZeroChild((*p).l, (*p).l.l))
   //
   //      // From y, take one step down the
   //      // search path to z. If z is null, replace z by a new node of rank 0 containing the item to
   //      // be inserted; if the new node is a 0-child, do a rotate or double rotate step to restore
   //      // the rank rule (Figure 6(d), (e), and (f)). This completes the insertion.
   //      if i <= (*p).l.s {
   //         //
   //         // LEFT LEFT
   //         //
   //         // "If this node is nil, replace it by a new node containing the
   //         //  item to be inserted. "
   //         //
   //         if (*p).l.l == nil {
   //            tree.attachLL(*p, x)
   //            if tree.isZeroChild((*p).l, (*p).l.l) { // or is p.l.rank == 0 ?
   //               tree.rotateR(p)
   //            }
   //            return
   //         }
   //         //If z is not a 0,0-node, replace w by y and x by z, completing the
   //         // step (Figure 6(g)).
   //         if !tree.isZeroZero((*p).l.l) {
   //            tree.pathLeft(&p)
   //            tree.pathLeft(&p)
   //            continue
   //         }
   //         // If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
   //         // x by z, completing the step (Figure 6(g)).
   //         if !tree.isOneChild((*p).l, (*p).l.l) {
   //            tree.pathLeft(&p)
   //            tree.pathLeft(&p)
   //            tree.promote(*p)
   //            continue
   //         }
   //         // Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).
   //         //  (i) Node z is a 0,0-node and a 1-child, and y and z are both left or
   //         //both right children: promote z, do a rotate step, and replace w by z and x by the child of z along the search
   //         //path.
   //         tree.rotateR(p)
   //         tree.pathLeft(&p)
   //         tree.promote(*p)
   //         continue
   //
   //      } else {
   //         //
   //         // LEFT RIGHT
   //         //
   //         if (*p).l.r == nil {
   //            tree.attachLR(*p, x)
   //            if tree.isZeroChild((*p).l, (*p).l.r) { // or is p.l.rank == 0 ?
   //               tree.rotateLR(p)
   //            }
   //            return
   //         }
   //         // In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child.
   //
   //         //If z is not a 0,0-node, replace w by y and x by z, completing the
   //         // step (Figure 6(g)).
   //         if !tree.isZeroZero((*p).l.r) {
   //            tree.pathLeft(&p)
   //            tree.pathRight(&p, &i)
   //            continue
   //         }
   //         // If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
   //         // x by z, completing the step (Figure 6(g)).
   //         if !tree.isOneChild((*p).l, (*p).l.r) {
   //            tree.pathLeft(&p)
   //            tree.pathRight(&p, &i)
   //            tree.promote(*p)
   //            continue
   //         }
   //         // Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).
   //
   //         // (j) Node z is a 0,0-node and a 1-child, and exactly one of y and z is a left child: promote z, do a double
   //         //            //rotate step, replace w by whichever of x and y is on the search path from z after the rotations, and replace x
   //         //            //by the child of the new w on the search path.
   //         tree.rotateLR(p)
   //         tree.promote(*p)
   //         //
   //         // "If a double rotation is done, take one further step down the
   //         //  search path after the rotation. Ths completes the step."
   //         //
   //         if i <= (*p).s {
   //            tree.pathLeft(&p) // LRL
   //         } else {
   //            tree.pathRight(&p, &i) // LRR
   //         }
   //      }
   //   } else {
   //      //
   //      // RIGHT
   //      //
   //      // "If the next node on the search path is nil, replace it by a new
   //      //  node of rank 0 containing the item to be inserted. This completes
   //      //  the insertion: the new node may be a 0-child, but *p is not."
   //      //
   //      if (*p).r == nil {
   //         tree.attachR(*p, x)
   //         return
   //      }
   //      if !tree.isZeroZero((*p).r) && !tree.isZeroChild(*p, (*p).r) {
   //         tree.pathRight(&p, &i)
   //         continue
   //      }
   //      if tree.isZeroZero((*p).r) {
   //         tree.pathRight(&p, &i)
   //         tree.promote(*p)
   //         continue
   //      }
   //      // In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child
   //
   //      // From y, take one step down the
   //      // search path to z. If z is null, replace z by a new node of rank 0 containing the item to
   //      // be inserted; if the new node is a 0-child, do a rotate or double rotate step to restore
   //      // the rank rule (Figure 6(d), (e), and (f)). This completes the insertion.
   //      if i > (*p).s + (*p).r.s + 1 {
   //         //
   //         // RIGHT RIGHT
   //         //
   //         // "If this node is nil, replace it by a new node containing the
   //         //  item to be inserted. "
   //         //
   //         if (*p).r.r == nil {
   //            tree.attachRR(*p, x)
   //            if tree.isZeroChild((*p).r, (*p).r.r) { // or is p.r.rank == 0 ?
   //               tree.rotateL(p)
   //            }
   //            return
   //         }
   //         // In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child.
   //
   //         //If z is not a 0,0-node, replace w by y and x by z, completing the
   //         // step (Figure 6(g)).
   //         if !tree.isZeroZero((*p).r.r) {
   //            tree.pathRight(&p, &i)
   //            tree.pathRight(&p, &i)
   //            continue
   //         }
   //         // If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
   //         // x by z, completing the step (Figure 6(g)).
   //         if !tree.isOneChild((*p).r, (*p).r.r) {
   //            tree.pathRight(&p, &i)
   //            tree.pathRight(&p, &i)
   //            tree.promote(*p)
   //            continue
   //         }
   //         // Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).
   //         //  (i) Node z is a 0,0-node and a 1-child, and y and z are both left or
   //         //both right children: promote z, do a rotate step, and replace w by z and x by the child of z along the search
   //         //path.
   //         tree.rotateL(p)
   //         tree.pathRight(&p, &i)
   //         tree.promote(*p)
   //         continue
   //
   //      } else {
   //         //
   //         // RIGHT LEFT
   //         //
   //         if (*p).r.l == nil {
   //            tree.attachRL(*p, x)
   //            if tree.isZeroChild((*p).r, (*p).r.l) { // or is p.l.rank == 0 ?
   //               tree.rotateRL(p)
   //            }
   //            return
   //         }
   //         // In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child.
   //
   //         //If z is not a 0,0-node, replace w by y and x by z, completing the
   //         // step (Figure 6(g)).
   //         if !tree.isZeroZero((*p).r.l) {
   //            tree.pathRight(&p, &i)
   //            tree.pathLeft(&p)
   //            continue
   //         }
   //         // If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
   //         // x by z, completing the step (Figure 6(g)).
   //         if !tree.isOneChild((*p).r, (*p).r.l) {
   //            tree.pathRight(&p, &i)
   //            tree.pathLeft(&p)
   //            tree.promote(*p)
   //            continue
   //         }
   //         // Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).
   //
   //         // (j) Node z is a 0,0-node and a 1-child, and exactly one of y and z is a left child: promote z, do a double
   //         //            //rotate step, replace w by whichever of x and y is on the search path from z after the rotations, and replace x
   //         //            //by the child of the new w on the search path.
   //         tree.rotateRL(p)
   //         tree.promote(*p)
   //         //
   //         // "If a double rotation is done, take one further step down the
   //         //  search path after the rotation. Ths completes the step."
   //         //
   //         if i > (*p).s {
   //            tree.pathRight(&p, &i) // RLR
   //         } else {
   //            tree.pathLeft(&p) // RLL
   //         }
   //      }
   //   }
   //}
//}

//func (tree *RedBlackRelaxedTopDown) Insert(i Position, x Data) {
//   assert(i <= tree.size)
//   tree.size++
//   tree.insert(&tree.root, i, x)
//}

//
//func (tree RedBlackRelaxedTopDown) Join(other List) List {
//   tree.share(tree.root)
//   tree.share(other.(*RedBlackRelaxedTopDown).root)
//   return &RedBlackRelaxedTopDown{
//      RedBlackTopDown: RedBlackTopDown{
//         Tree: Tree{
//            arena: tree.arena,
//            root:  tree.join(tree.root, other.(*RedBlackRelaxedTopDown).root, tree.size),
//            size:  tree.size + other.(*RedBlackRelaxedTopDown).size,
//         },
//      },
//   }
//}
//
//func (tree RedBlackRelaxedTopDown) Split(i Position) (List, List) {
//   assert(i <= tree.size)
//   tree.share(tree.root)
//   l, r := tree.split(tree.root, i, tree.size)
//   return &RedBlackRelaxedTopDown{RedBlackTopDown: RedBlackTopDown{Tree: Tree{arena: tree.arena, root: l, size: i}}},
//          &RedBlackRelaxedTopDown{RedBlackTopDown: RedBlackTopDown{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}
//}
//
//func (tree RedBlackRelaxedTopDown) split(p *Node, i, s Size) (l, r *Node) {
//   if p == nil {
//      return
//   }
//   tree.persist(&p)
//
//   sl := p.s
//   sr := s - p.s - 1
//
//   if i <= (*p).s {
//      l, r = tree.split(p.l, i, sl)
//      r = tree.build(r, p, p.r, sl-i)
//   } else {
//      l, r = tree.split(p.r, i-sl-1, sr)
//      l = tree.build(p.l, p, l, sl)
//   }
//   return l, r
//}
//
//func (tree RedBlackRelaxedTopDown) build(l, p, r *Node, sl Size) *Node {
//   if tree.rank(l) == tree.rank(r) {
//      p.l = l
//      p.r = r
//      p.s = sl
//      p.y = uint64(tree.rank(p.l) + 1)
//      return p
//   }
//   if tree.rank(l) < tree.rank(r) {
//      tree.persist(&r)
//      r.s = 1 + sl + r.s
//      r.l = tree.build(l, p, r.l, sl)
//      return tree.balanceInsertL(r)
//   } else {
//      tree.persist(&l)
//      l.r = tree.build(l.r, p, r, sl-l.s-1)
//      return tree.balanceInsertR(l)
//   }
//}
//
//func (tree RedBlackRelaxedTopDown) join(l, r *Node, sl Size) (p *Node) {
//   if l == nil {
//      return r
//   }
//   if r == nil {
//      return l
//   }
//   if tree.rank(l) < tree.rank(r) {
//      return tree.build(l, tree.RedBlackTopDown.Tree.deleteMin(&r), r, sl)
//   } else {
//      return tree.build(l, tree.RedBlackTopDown.Tree.deleteMax(&l), r, sl-1)
//   }
//}
