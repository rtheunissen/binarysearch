package binarytree

import (
   "binarysearch/abstract/list"
   "math"
)

type AVLBottomUp struct {
   Tree
   RankBalanced
}

func (AVLBottomUp) New() list.List {
   return &AVLBottomUp{}
}

func (tree *AVLBottomUp) Clone() list.List {
   return &AVLBottomUp{Tree: tree.Tree.Clone()}
}

func (tree *AVLBottomUp) Verify() {
   tree.verifySizes(tree.root, tree.size)
   tree.verifyHeight(tree.root, tree.size)
   tree.verifyRanks(tree.root)
}

func (tree *AVLBottomUp) verifySizes(p *Node, s list.Size) list.Size {
   if p == nil {
      return 0
   }
   sl := tree.verifySizes(p.l, p.sizeL())
   sr := tree.verifySizes(p.r, p.sizeR(s))

   invariant(s == sl + sr + 1)
   return s
}

func (tree *AVLBottomUp) verifyHeight(p *Node, s list.Size) {
   if p == nil {
      return
   }
   invariant(tree.rank(p) == tree.root.height())
   invariant(tree.rank(p) <= int(1.44 * math.Log2(float64(s + 2)) - 0.328))
}

func (tree *AVLBottomUp) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   // AVL rule: Every node is 1,1 or 1,2
   invariant(tree.isOneChild(p, p.l) || tree.isTwoChild(p, p.l))
   invariant(tree.isOneChild(p, p.r) || tree.isTwoChild(p, p.r))
   invariant(tree.isOneChild(p, p.l) || tree.isOneChild(p, p.r))
   invariant(tree.isOneChild(p, p.r) || tree.isOneChild(p, p.l))

   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
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
   if i <= p.s {
      p.s = p.s + 1
      p.l = tree.insert(p.l, i, x)
      return tree.balanceInsertL(p)
   } else {
      p.r = tree.insert(p.r, i - p.s - 1, x)
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
   // assert(tree.isZeroChild(p, p.l))
   // assert(tree.isOneChild(p, p.r) || tree.isTwoChild(p, p.r))
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
   // assert(tree.isZeroChild(p, p.l))
   // assert(tree.isTwoChild(p, p.r))
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
      //                          2               0   <-- 2-child
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
      //                              1       2   <-- should have rank 1
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
   // assert(tree.isZeroChild(p, p.l))
   // assert(tree.isTwoChild(p, p.r))
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

func (tree *AVLBottomUp) balanceDeleteL(p *Node) *Node {
   //
   //
   //
   if tree.isTwoTwo(p) {
      tree.demote(p)
      return p
   }
   //
   //
   //
   if tree.isThreeChild(p, p.l) {
      //
      //
      //
      if tree.isTwoChild(p.r, p.r.r) {
         //
         //
         //
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.r)
         tree.demote(p.l)
         tree.demote(p.l)
      } else {
         //
         //
         //
         if tree.isTwoChild(p.r, p.r.l) {
            //
            //
            //
            tree.rotateL(&p)
            tree.demote(p.l)
            tree.demote(p.l)
         } else {
            //
            //
            //
            tree.rotateL(&p)
            tree.promote(p)
            tree.demote(p.l)
         }
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

// Constructs a balanced tree with root `p` where all nodes in `l` are to the
// left of `p` and all nodes in `r` to the right of `p`.
//
// The rank of `r` is greater than or equal to the rank of `l`.
//
//                            l      p       r
//                                   .
//                                          /\
//                            /\           /  \
//                           /  \         /    \
//                          /____\       /______\
//
// Descend along the left spine of the right tree to find a subtree that is
// similar in rank to `l`, then construct a new tree there with `r` as the right
// subtree and `l` as the left subtree.
//
// The left subtree of `r` will eventually consist of all the nodes currently in
// that subtree, as well as `p`, as well as all the nodes in `l`.
//
// Balancing makes use of the same procedure as when inserting, however the
// double-rotate case is never needed and could be factored out.
//
func (tree *AVLBottomUp) buildL(l, p, r *Node, sl list.Size) *Node {
   if tree.rankDifference(r, l) <= 1 {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(tree.rank(r) + 1)
      return p
   }
   tree.persist(&r)
   r.s = r.s + sl + 1
   r.l = tree.buildL(l, p, r.l, sl)
   return tree.balanceInsertL(r)
}

// Symmetric.
func (tree *AVLBottomUp) buildR(l, p, r *Node, sl list.Size) *Node {
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

// Similar to buildL, but there is no `p` node yet to use for local root there.
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
   r.s = r.s + sl
   r.l = tree.joinL(l, r.l, sl)
   return tree.balanceInsertL(r)
}

// Symmetric.
func (tree *AVLBottomUp) joinR(l, r *Node, sl list.Size) (p *Node) {
   if tree.rankDifference(l, r) <= 1 {
      return tree.build(tree.deleteMax(l, &p), p, r, sl - 1)
   }
   tree.persist(&l)
   l.r = tree.joinR(l.r, r, l.sizeR(sl))
   return tree.balanceInsertR(l)
}

func (tree *AVLBottomUp) Split(i list.Position) (list.List, list.List) {
   // assert(i <= tree.size)
   tree.share(tree.root)

   l, r := tree.split(tree.root, i, tree.size)

   return &AVLBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}},
          &AVLBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
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
