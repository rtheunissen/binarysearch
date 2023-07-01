package binarytree

import (
   "binarysearch/abstract/list"
)

// AVLBottomUp
//
// This a standard, recursive, bottom-up implementation of an AVL tree
// using the rank-balanced framework of Haeupler, Sen, and Tarjan [].
//
// Using ranks makes it easy to reason about the height of each subtree and
// provides an intuitive way to adjust ranks after rotations. Balancing is
// annotated in one direction only since the algorithms are symmetrical.
//
// A choice was made to not unify the symmetric cases using the direction-based
// technique of Ben Pfaff and others because it makes the logic more difficult
// to follow even though there would be less code overall.
//
// Storing ranks rather than rank differences takes up an entire integer field,
// but it makes `join` easier to implement and is consistent with the other
// rank-balanced implementations. It is possible to store only rank differences
// instead of ranks to use only 1 bit for the balancing information, if needed.
//
type AVLBottomUp struct {
   AVL
   Tree
}

func (tree *AVLBottomUp) New() list.List {
   return &AVLBottomUp{}
}

func (tree *AVLBottomUp) Clone() list.List {
   return &AVLBottomUp{Tree: tree.Tree.Clone()}
}

func (tree *AVLBottomUp) Verify() {
   tree.verifySize(tree.root, tree.size)
   tree.AVL.verifyRanks(tree.root, tree.size)
}

func (tree *AVLBottomUp) Insert(i list.Position, x list.Data) {
   tree.root = tree.insert(tree.root, i, x)
   tree.size = tree.size + 1
}

func (tree *AVLBottomUp) insert(p *Node, i list.Position, x list.Data) *Node {
   if p == nil {
      return tree.allocate(Node{x: x})
   }
   tree.persist(&p)
   if i <= p.sizeL() {
      p.s = p.sizeL() + 1
      p.l = tree.insert(p.l, i, x)
      return tree.balanceInsertL(p)
   } else {
      p.r = tree.insert(p.r, i - (p.sizeL() + 1), x)
      return tree.balanceInsertR(p)
   }
}

func (tree *AVLBottomUp) balanceInsertL(p *Node) *Node {
   //
   // After inserting to the left, there is no need to balance if the height of
   // the left subtree is not equal to the height of its parent, i.e. the height
   // of the left subtree did not increase.
   //
   if !tree.isZeroChild(p, p.l) {
      return p
   }
   //
   // Otherwise, the height of the left subtree increased.
   //
   // The AVL rule is that every node is 1,1 or 1,2 and the left subtree is now
   // a 0-child, so we would like to make it a 1-child. It would not make sense
   // to demote the left subtree because we know that it is balanced and valid.
   //
   // Our only options are to either promote the parent or to rotate to somehow
   // resolve the invariant. Promoting the parent would change the left subtree
   // from a 0-child to a 1-child, and the right subtree from either a 1-child
   // to a 2-child or from a 2-child to a 3-child.
   //
   // Therefore, we can promote the parent to make the left subtree a 1-child
   // only if the right subtree is currently a 1-child becoming a 2-child,
   // which results in the parent becoming a 1,2-node, restoring the invariant.
   //
   if tree.isOneChild(p, p.r) {
      tree.promote(p)
      return p
   }
   //
   // The parent is a 0,2-node because we could not promote it without creating
   // a 3-child in the right subtree. The only way to resolve this is to rotate,
   // and we know that we need to rotate to the right because the left subtree
   // must have increased in height because we inserted somewhere to the left.
   //
   if tree.isTwoChild(p.l, p.l.r) {
      //
      //                                  2
      //                          ╭───────┴───────╮
      //                          2               0   ← 2-child
      //                      ╭───┴───╮
      //                      1       0
      //                    ╭─╯
      //                    0
      //
      // Consider what a right rotation would do here: the parent with rank 2
      // is pushed down to the right, pulling its left subtree with rank 2 up
      // into its place, and the right subtree with rank 0 at p.l.r will move
      // sideways to the right to become the left subtree of the current parent.
      //
      //                        AFTER A RIGHT ROTATION
      //
      //                                  2
      //                              ╭───┴───╮
      //                              1       2   ← Should have rank 1
      //                            ╭─╯     ╭─┴─╮
      //                            0       0   0
      //
      // This creates a valid AVL-rule structure, but the height of the right
      // subtree is actually 1 when its rank is 2, so we need to demote it.
      //
      //                  AFTER A RIGHT ROTATION AND DEMOTION
      //
      //                                  2
      //                              ╭───┴───╮
      //                              1       1
      //                            ╭─╯     ╭─┴─╮
      //                            0       0   0
      //
      tree.rotateR(&p)
      tree.demote(p.r)
      return p
   }
   // assert(tree.isOneChild(p.l, p.l.r))
   //
   // The right subtree of the left subtree is a 1-child, which prevents us from
   // making a simple right rotation followed by a demotion. If we did that, the
   // left subtree would have a height of 0 and the right subtree a height of 2,
   // which by AVL rules is a height difference > 1. Considering the ranks, we
   // would need to promote the parent to have its rank match its height of 3,
   // but that would create a 3-child on the left.
   //
   // No promotion or demotion would resolve this because the structure itself
   // is not valid - the rotation does not help us to resolve the invariant.
   //
   //            CURRENT TREE               ROTATE RIGHT AND DEMOTE?
   //
   //                  2                               2
   //          ╭───────┴───────╮               ╭───────┴───────╮
   //          2               0               0               1
   //      ╭───┴───╮                                       ╭───┴───╮
   //      0       1                                       1       0
   //            ╭─╯                                     ╭─╯
   //            0                                       0
   //
   //                      1
   // The problem is the ╭─╯ subtree that is creating a tree of height 3.
   //                    0
   //
   // We can pull that up before we rotate to the right by first rotating the
   // left subtree left. Let's take a look at what that looks like for now:
   //
   //            CURRENT TREE             ROTATE THE LEFT SUBTREE LEFT
   //
   //                  2                               2
   //          ╭───────┴───────╮               ╭───────┴───────╮
   //          2               0               1               0
   //      ╭───┴───╮                       ╭───╯
   //      0       1                       2
   //            ╭─╯                     ╭─┴─╮
   //            0                       0   0
   //
   // We now have a parent with height 3 as before, but the left subtree has a
   // height of 2 and the right subtree has a height of 0. Rotating to the right
   // decreases the height of the left subtree by 1 to make it 1, and also
   // increases the height of the right subtree by 1 to make it 1.
   //
   // Let's take a look at just doing that, without making rank changes for now:
   //
   //                 AFTER DOING A LEFT-RIGHT ROTATION
   //
   //                               1
   //                           ╭───┴───╮
   //                           2       2
   //                         ╭─┴─╮     ╰─╮
   //                         0   0       0
   //
   // The structure looks good, but the ranks are not right because they should
   // be equal the height at each node. Notice that the new parent should have
   // a rank of 2 and the left and right subtrees should both have a rank of 1.
   //
   // Promote the parent and demote both subtrees to resolve the rank invariant.
   //
   //               AFTER ONE PROMOTION AND TWO DEMOTIONS
   //
   //                               2
   //                           ╭───┴───╮
   //                           1       1
   //                         ╭─┴─╮     ╰─╮
   //                         0   0       0
   //
   tree.rotateLR(&p)
   tree.promote(p)
   tree.demote(p.l)
   tree.demote(p.r)
   return p
}

// Symmetric
func (tree *AVLBottomUp) balanceInsertR(p *Node) *Node {
   if tree.isZeroChild(p, p.r) {
      if tree.isOneChild(p, p.l) {
         tree.promote(p)
      } else if tree.isTwoChild(p.r, p.r.l) {
         tree.rotateL(&p)
         tree.demote(p.l)
      } else {
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.r)
      }
   }
   return p
}

func (tree *AVLBottomUp) Delete(i list.Position) (x list.Data) {
   // assert(i < tree.size)
   tree.root = tree.delete(tree.root, i, &x)
   tree.size = tree.size - 1
   return x
}

func (tree *AVLBottomUp) delete(p *Node, i list.Position, x *list.Data) *Node {
   tree.persist(&p)
   if i == p.sizeL() {
      *x = p.x
      defer tree.free(p)
      return tree.join(p.l, p.r, p.s)
   }
   if i < p.sizeL() {
      p.s = p.sizeL() - 1
      p.l = tree.delete(p.l, i, x)
      return tree.balanceDeleteL(p)
   } else {
      p.r = tree.delete(p.r, i - (p.sizeL() + 1), x)
      return tree.balanceDeleteR(p)
   }
}

func (tree *AVLBottomUp) balanceDeleteL(p *Node) *Node {
   //
   // Deleting a node should not increase the height of the tree, and therefore
   // should not increase the rank of a node, except as needed after a rotation.
   // The general intuition is to decrease ranks because the height of the tree
   // is decreasing as nodes are deleted.
   //
   // Recall that the AVL rule is that every node must be 1,1 or 1,2.
   //
   // The first case is simple: we have a 2,2 node that was previously 1,2 but
   // the left subtree has decreased in height after a deletion. The height of
   // the left subtree is now equal to the height of the right subtree, so the
   // node is height-balanced, but the 2,2 parent is not allowed.
   //
   // To resolve the invariant, demote the parent to create a valid 1,1 node.
   //
   //    DELETE DECREASES HEIGHT     PARENT BECOMES 2,2      DEMOTE TO 1,1
   //
   //                3                        3                     2
   //            ╭───┴───╮                ╭───┴───╮             ╭───┴───╮
   //            2       1                1       1             1       1
   //          ╭─┴─╮   ╭─╯              ╭─┴─╮   ╭─╯           ╭─┴─╮   ╭─╯
   //          1   0   0                0   0   0             0   0   0
   //        ╭─╯
   //     →  0
   //
   if tree.isTwoTwo(p) {
      tree.demote(p)
      return p
   }
   //
   // Otherwise, it is possible that the left subtree was a 2-child before the
   // deletion, which would make it a 3-child if the deletion decreased height.
   //
   if tree.isThreeChild(p, p.l) {
      //
      //                                3
      //                        ╭───────┴───────╮
      //            3-child →   0               2
      //                                    ╭───┴───╮
      //                                    1       0
      //                                  ╭─╯
      //                                  0
      //
      // In this case, the right subtree must be a 1-child because the node was
      // previously a 2,1-node and the 2-node is now a 3-node.
      //
      // assert(tree.isOneChild(p, p.r))
      //
      // Demoting the parent is not possible because that would make the right
      // subtree a 0-child. Consider that the 3,1-node situation means that the
      // left subtree has become too much shorter than the right subtree.
      //
      // This requires a rotation to the left to increase the height of the left
      // subtree and decrease the height of the right subtree.
      //
      // Let's try a left rotation and see what happens.
      //
      //            ORIGINAL                      ROTATE PARENT LEFT
      //
      //               3                                     2
      //       ╭───────┴───────╮                     ╭───────┴───────╮
      //       0               2                     3               0
      //                   ╭───┴───╮             ╭───┴───╮
      //                   1       0             0       1
      //                 ╭─╯                           ╭─╯
      //                 0                             0
      //
      // This rotation did not help, because the height of the left subtree is
      // now 3 and the height of the right subtree is 1. The structure is not a
      // valid AVL tree so no amount of rank adjustments can fix it.
      //
      // Take a look at the right subtree of the right subtree: if that node is
      // a 2-child, it suggests that its sibling must be a 1-child. The 1-child
      // sibling is then the subtree with the greater height because its rank is
      // closer to its parent.
      //
      if tree.isTwoChild(p.r, p.r.r) {
         //
         //                             3
         //                     ╭───────┴───────╮
         //                     0               2
         //                                 ╭───┴───╮
         //                     1-child →   1       0   ← 2-child
         //                               ╭─╯
         //                               0
         //
         // assert(tree.isOneChild(p, p.r))
         // assert(tree.isOneChild(p.r, p.r.l))
         //
         // We get a valid AVL structure by first rotating the right subtree to
         // the right, and then rotating the parent left.
         //
         //
         //       ROTATE RIGHT SUBTREE RIGHT        ROTATE PARENT LEFT
         //
         //               3                                 1
         //       ╭───────┴───────╮                     ╭───┴───╮
         //       0               1                     3       2
         //                   ╭───┴───╮               ╭─┴─╮     ╰─╮
         //                   0       2               0   0       0
         //                           ╰─╮
         //                             0
         //
         // The only thing to do is fix the ranks after the rotations:
         //
         //    Consider how the height of each node has changed: the rank of the
         //    left subtree must change from 3 to 1 because its height is 1, the
         //    parent must change from 1 to 2 and the right subtree from 2 to 1.
         //
         //    Promote the parent once, demote the right subtree once, and
         //    demote the left subtree twice. This restores the rank invariant.
         //
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.r)
         tree.demote(p.l)
         tree.demote(p.l)
         return p
      }
      // Otherwise, the right subtree of the right subtree must be a 1-child,
      // given that it is not a 2-child, which allows the left subtree of the
      // right subtree to be either a 1-child or a 2-child.
      //
      // assert(tree.isThreeChild(p, p.l))
      // assert(tree.isOneChild(p.r, p.r.r))
      //
      //                              3
      //                      ╭───────┴───────╮
      //          3-child →   0               2
      //                                  ╭───┴───╮
      //           1-child or 2-child →   ?       1   ← 1-child
      //                                          ╰─╮
      //                                            0
      //
      // There was no need to first rotate the right subtree to the right,
      // because the right subtree of the right child is a 1-child, which
      // indicates that it either has the same or greater height than its
      // sibling, the left subtree of the right subtree of the parent.
      //
      if tree.isTwoChild(p.r, p.r.l) {
         //
         //                              3
         //                      ╭───────┴───────╮
         //          3-child →   0               2
         //                                  ╭───┴───╮
         //                      2-child →   0       1   ← 1-child
         //                                          ╰─╮
         //                                            0
         //
         // Rotating the parent to the left results in a valid structure, but
         // the ranks are not correct. The height of the subtree is 2, so the
         // rank of the parent is correct. The height of the new left subtree
         // is 1 but its rank is 3, so we demote the left subtree twice.
         //
         // This restores the rank invariant.
         //
         //                      ROTATE PARENT LEFT
         //
         //                              2
         //                          ╭───┴───╮
         //            2x Demote →   3       1
         //                        ╭─┴─╮     ╰─╮
         //                        0   0       0
         //
         //
         //                   DEMOTE LEFT SUBTREE TWICE
         //
         //                              2
         //                          ╭───┴───╮
         //            2x Demote →   3       1
         //                        ╭─┴─╮     ╰─╮
         //                        0   0       0
         //
         //                              2
         //                          ╭───┴───╮
         //                          1       1
         //                        ╭─┴─╮     ╰─╮
         //                        0   0       0
         tree.rotateL(&p)
         tree.demote(p.l)
         tree.demote(p.l)
         return p

      } else {
         // assert(tree.isOneChild(p.r, p.r.l))
         // assert(tree.isOneChild(p.r, p.r.r))
         //
         //                              3
         //                      ╭───────┴───────╮
         //          3-child →   0               2
         //                                  ╭───┴───╮
         //                      1-child →   1       1   ← 1-child
         //                                  ╰─╮     ╰─╮
         //                                    0       0
         //
         // Rotating the parent to the left results in a valid structure, but
         // the ranks are not correct. The rank of the parent is 2, but its
         // height is actually 3. The rank of the left subtree is 3 but its
         // height is actually 2. To fix this, promote the parent and demote the
         // left subtree, which restores the rank invariant.
         //
         //                      ROTATE PARENT LEFT
         //
         //                              2
         //                      ╭───────┴───────╮
         //                      3               1
         //                  ╭───┴───╮           ╰───╮
         //                  0       1               0
         //                          ╰─╮
         //                            0
         //
         //
         //               PROMOTE PARENT, DEMOTE LEFT SUBTREE
         //
         //                              2   ← Promote
         //                      ╭───────┴───────╮
         //           Demote →   3               1
         //                  ╭───┴───╮           ╰───╮
         //                  0       1               0
         //                          ╰─╮
         //                            0
         //
         //                              3
         //                      ╭───────┴───────╮
         //                      2               1
         //                  ╭───┴───╮           ╰───╮
         //                  0       1               0
         //                          ╰─╮
         //                            0
         tree.rotateL(&p)
         tree.promote(p)
         tree.demote(p.l)
         return p
      }
   }
   return p
}

// Symmetric
func (tree *AVLBottomUp) balanceDeleteR(p *Node) *Node {
   if tree.isTwoTwo(p) {
      tree.demote(p)
      return p
   }
   if tree.isThreeChild(p, p.r) {
      if tree.isTwoChild(p.l, p.l.l) {
         tree.rotateLR(&p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.r)
         tree.demote(p.r)
      } else {
         if tree.isTwoChild(p.l, p.l.r) {
            tree.rotateR(&p)
            tree.demote(p.r)
            tree.demote(p.r)
         } else {
            tree.rotateR(&p)
            tree.promote(p)
            tree.demote(p.r)
         }
      }
   }
   return p
}

// Returns the result of deleting the left-most node of p.
func (tree *AVLBottomUp) deleteMin(p *Node, min **Node) *Node {
   tree.persist(&p)
   if p.l == nil {
      *min = p
      return p.r
   }
   p.s = p.s - 1
   p.l = tree.deleteMin(p.l, min)
   return tree.balanceDeleteL(p)
}

// Returns the result of deleting the right-most node of p.
func (tree *AVLBottomUp) deleteMax(p *Node, max **Node) *Node {
   tree.persist(&p)
   if p.r == nil {
      *max = p
      return p.l
   }
   p.r = tree.deleteMax(p.r, max)
   return tree.balanceDeleteR(p)
}

// Constructs a balanced tree with root `p` where all nodes in `l` are to the
// left of `p` and all nodes in `r` to the right of `p`.
//
// The rank of `r` is greater than or equal to the rank of `l`.
//
//                                    p
//                                    .       r
//                              l            ↙
//                             ↙            /\
//                            /\           /\ \
//                           /  \      ~l /  \ \
//                          /____\       /____\_\
//
//
// Follow the left spine of `r` to find a subtree that is similar in rank to `l`
// then build a new subtree with parent `p`, left subtree `l` and right `r`.
//
func (tree *AVLBottomUp) buildL(l, p, r *Node, sl list.Size) *Node {
   // assert(tree.rank(r) >= tree.rank(l))
   if tree.rankDifference(r, l) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(r) + 1)
      return p
   }
   tree.persist(&r)
   //
   // To update the size of `r`, which is the eventual size of its left subtree,
   // consider that the left subtree of `r` will consist of all the nodes in `l`,
   // then `p`,all the nodes currently in that subtree.
   //
   r.s = r.sizeL() + sl + 1
   r.l = tree.buildL(l, p, r.l, sl)
   return tree.balanceInsertL(r)
}

// Symmetric
func (tree *AVLBottomUp) buildR(l, p, r *Node, sl list.Size) *Node {
   // assert(tree.rank(l) >= tree.rank(r))
   if tree.rankDifference(l, r) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(l) + 1)
      return p
   }
   tree.persist(&l)
   l.r = tree.buildR(l.r, p, r, l.sizeR(sl))
   return tree.balanceInsertR(l)
}

// Constructs a balanced tree with root p where all nodes of l are to the left
// of p and all nodes in r are to the right of p.
func (tree *AVLBottomUp) build(l, p, r *Node, sl list.Size) *Node {
   // assert(sl == l.size())
   if tree.rank(l) < tree.rank(r) {
      return tree.buildL(l, p, r, sl)
   } else {
      return tree.buildR(l, p, r, sl)
   }
}

// Similar to buildL, but there is no `p` node yet to use for the local root.
//
// At some point we will need to delete the left-most node of `r` to use as `p`,
// but we delay that as long as possible to avoid descending all the way down to
// delete it, all the way back up as the recursion unwinds, then descend again
// along the same path anyway.
//
func (tree *AVLBottomUp) joinL(l, r *Node, sl list.Size) (p *Node) {
   if tree.rankDifference(r, l) <= 1 {
      return tree.build(l, p, tree.deleteMin(r, &p), sl)
   }
   tree.persist(&r)
   r.s = r.sizeL() + sl
   r.l = tree.joinL(l, r.l, sl)
   return tree.balanceInsertL(r)
}

// Symmetric
func (tree *AVLBottomUp) joinR(l, r *Node, sl list.Size) (p *Node) {
   if tree.rankDifference(l, r) <= 1 {
      return tree.build(tree.deleteMax(l, &p), p, r, sl - 1)
   }
   tree.persist(&l)
   l.r = tree.joinR(l.r, r, l.sizeR(sl))
   return tree.balanceInsertR(l)
}

// Constructs a balanced tree with root `p` where all nodes in `l` are to the
// left of `p` and all nodes in `r` to the right of `p`.
func (tree *AVLBottomUp) join(l, r *Node, sl list.Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) < tree.rank(r) {
      return tree.joinL(l, r, sl)
   } else {
      return tree.joinR(l, r, sl)
   }
}

func (tree *AVLBottomUp) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*AVLBottomUp).root)
   return &AVLBottomUp{
      Tree: Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, other.(*AVLBottomUp).root, tree.size),
         size:  tree.size + other.(*AVLBottomUp).size,
      },
   }
}

// Splits the tree of `p` into two trees `l` and `r` at position `i`, such that
// the resulting size of `l` is equal to `i`.
func (tree *AVLBottomUp) split(p *Node, i, s list.Size) (l, r *Node) {
   // assert(s == p.size())
   if p == nil {
      return
   }
   tree.persist(&p)
   if i <= (*p).sizeL() {
      l, r = tree.split(p.l, i, p.sizeL())
         r = tree.build(r, p, p.r, p.sizeL() - i)
   } else {
      l, r = tree.split(p.r, i - (p.sizeL() + 1), p.sizeR(s))
         l = tree.build(p.l, p, l, p.sizeL())
   }
   return l, r
}

func (tree *AVLBottomUp) Split(i list.Position) (list.List, list.List) {
   // assert(i <= tree.size)
   tree.share(tree.root)

   l, r := tree.split(tree.root, i, tree.size)

   return &AVLBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}},
          &AVLBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}
