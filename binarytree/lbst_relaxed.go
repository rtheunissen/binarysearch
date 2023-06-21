package binarytree

import (
   . "binarysearch/abstract/list"
)

// LBSTRelaxed / Relaxed Weight-Balanced Tree (Experimental)
//
// Recap:
//    The _height_ of a tree is the longest path from the root to any other node.
//    The _depth_ of a node is the number of links to follow to reach the root.
//    The _weight_ of a tree is equal to its size; the number of reachable nodes.
//
//    To guarantee a tree height upper-bound, a weight-balanced tree maintains
//    strict weight-balance at _every_ node in the tree.
//
// This balancing strategy combines the ideas of weight- and height-balance by
// allowing some nodes to not be weight-balanced, as long as the height of the
// tree is less than or equal to some upper-bound.
//
// This strategy can therefore be considered a "relaxed weight-balanced tree",
// because the height can be valid even when some nodes are not weight-balanced.
//
// When the height of the tree exceeds the upper-bound for its size, any node
// along the tree's longest path which is _not_ weight-balanced can be balanced
// to reduce its height by at least 1, thereby restoring the height bound.
//
// When inserting a new node, we count the number of nodes top-down to determine
// if the path exceeded the height upper-bound; if the depth exceeds the tree's
// height upper-bound, then at least one node along the insertion path must not
// be weight-balanced - we call this node the "scapegoat".
//
// ---
//
// Usually, a scapegoat tree is implemented quite differently:
//
// An "alpha" parameter `α` is chosen between 0.5 and 1.0:
//
//       Balanced:
//
//          sl <= α * s
//          sr <= α * s
//
//       MaximumPathLength upper-bound:
//
//          ⌊log1/α(size)⌋
//
//       Where `s` is the size of the parent, `sl` and `sr` are subtree sizes.
//
// This parameter imposes the computation of a logarithm with a fractional base.
// Choosing a=0.5 would be log₂, however the balance is then so strict that the
// amount of rebalancing makes the strategy non-viable.
//
// In comparison, the bitwise logarithmic weight-balance rule is much simpler,
// avoiding the logarithm, integer multiplication, and potential for overflow.
//
//       sl >= ~sl & (sr >> 1)
//       sr >= ~sr & (sl >> 1)
//
//       Check: size < (1 << ((depth + 1) >> 1))
//
//       The maximum height is 2 * ⌊log₂(m)⌋, where `m` is the maximum size
//       that the tree has reached. Deletions do not consider balance at all.
//
//       This height upper-bound is similar to the worst-case height of a
//       red-black tree of size `m`. TODO: what is it exactly?
//
// Usually, a scapegoat tree explicitly tracks the ceiling size `m` to determine
// when to rebuild after a delete: when the distance between the current size
// and the ceiling is too great, rebuild the entire tree and reset the ceiling
// to equal the size. This implementation forgoes this step, taking inspiration
// from the relaxed AVL and relaxed red-black trees.
//
// Another difference is that the textbook scapegoat tree chooses the deepest
// weight-unbalanced ancestor, i.e. the one furthest away from the root. This
// is done _after_ a height violation is detected, bottom-up towards the root
// until the first weight-unbalanced node is found. This implementation instead
// looks for a scapegoat top-down, choosing the first weight-unbalanced node
// encountered and therefore the highest weight-unbalanced ancestor.
//
// Most insertions, however, do not require a rebuild.
//
// Split is implemented using a similar approach to treaps and zip trees:
// one pass top-down, without considering balance. Both sides of the split
// then inherit the implicit ceiling of the tree that was split.
//
// Join simply removes either the maximum node of the left tree or the minimum
// node of the right tree for the joining node, whichever is larger. It is not
// clear what effect this has over time, but successive insertions will again
// improve the balance. This implementation includes it for completeness.
//
// ---
//
// References:
//    - Andersson, A. (1989). Improving Partial Rebuilding by Using Simple Balance Criteria. WADS.
//    - Andersson, A. (1999). General Balanced Trees. J. Algorithms, 30, 1-18.
//    - Galperin, I., & Rivest, R.L. (1993). LBSTRelaxed trees. SODA '93.
//    - Roura, S. (2001). A New Method for Balancing Binary Search Trees. ICALP.
//    - Muusse, I.J. (2017). An algorithm for balancing a binary search tree.
//

type LBSTRelaxed struct {
   Tree
}

// Creates a new LBSTRelaxed BST from existing values.
func (LBSTRelaxed) New() List {
   return &LBSTRelaxed{}
}

func (tree *LBSTRelaxed) Clone() List {
   return &LBSTRelaxed{
      Tree: tree.Tree.Clone(),
   }
}

func (tree *LBSTRelaxed) Verify() {
   tree.verifySizes()
}

// Inserts a value `s` at position `i` in the tree.
func (tree *LBSTRelaxed) Insert(i Position, x Data) {
   // assert(i <= tree.Size())

   var unbalancedNode **Node // An unbalanced node along the insertion path.
   var unbalancedSize Size   // The size of the unbalanced node.
   var depth uint64          // The depth of the insertion so far.

   p := &tree.root
   s := tree.size

   // Search with increasing depth until the end of the path is reached.
   //
   for *p != nil {
      tree.persist(p)
      depth++

      sl := (*p).s         // Size of the left subtree.  O(1)
      sr := s - (*p).s - 1 // Size of the right subtree. O(1)

      if i <= sl {
         //
         // LEFT
         //
         if unbalancedNode == nil && !tree.isBalanced(sr, sl+1) {
            unbalancedNode = p
            unbalancedSize = s + 1
         }
         p = insertL(*p)
         s = sl

      } else {
         //
         // RIGHT
         //
         if unbalancedNode == nil && !tree.isBalanced(sl, sr+1) {
            unbalancedNode = p
            unbalancedSize = s + 1
         }
         p = insertR(*p, &i)
         s = sr
      }
   }
   // Attach a new node at the end of the path.
   *p = tree.allocate(Node{x: x})
   tree.size++

   // Check if a rebuild is required.
   if tree.tooDeep(tree.size, depth) {
      tree.balance(unbalancedNode, unbalancedSize)
   }
}

// Determines if the depth at which a new node was inserted was too deep, i.e.
// if the new height of the tree exceeds the upper-bound for its size.
//
//        tooDeep := depth > 2 * ⌊log₂(size)⌋
//
//     because root.height() <= MaximumPathLength(2 * math.FloorLog2(size))
//
//   depth is height + 1 ?
func (tree *LBSTRelaxed) tooDeep(size Size, depth uint64) bool {
   return (1 << ((depth + 1) >> 1)) > size
}

// Determines if two sizes are balanced.
func (LBSTRelaxed) isBalanced(x, y Size) bool {
   return LogSize{}.isBalanced(x, y)
}

func (tree *LBSTRelaxed) balance(p **Node, s Size) {
   *p = Partition{LogSize{}}.balance(&tree.Tree, *p, s)
}

func (tree *LBSTRelaxed) Split(i Size) (List, List) {
   // assert(i <= tree.size)

   tree.share(tree.root)
   l,r := tree.split(tree.root, i)

   return &LBSTRelaxed{Tree{arena: tree.arena, root: l, size: i}},
          &LBSTRelaxed{Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree *LBSTRelaxed) joinLr(l, r *Node, sl, sr Size) *Node {
   if r == nil {
      return l
   }
   if tree.isBalanced(sr, sl) { // TODO: wrong way around?
      p := tree.deleteMax(&l)
      p.l = l
      p.r = r
      p.s = sl - 1
      return p
   }
   tree.persist(&l)
   l.r = tree.joinlR(l.r, r, sl-l.s-1, sr)
   return l
}

func (tree *LBSTRelaxed) joinlR(l, r *Node, sl, sr Size) *Node {
   if l == nil {
      return r
   }
   if tree.isBalanced(sl, sr) { // TODO: wrong way around?
      p := tree.deleteMin(&r)
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.persist(&r)
   r.l = tree.joinLr(l, r.l, sl, r.s)
   r.s = sl + r.s
   return r
}

func (tree *LBSTRelaxed) join4(l, r *Node, sl, sr Size) *Node {
   if sl > sr {
      return tree.joinLr(l, r, sl, sr)
   } else {
      return tree.joinlR(l, r, sl, sr)
   }
}

func (tree *LBSTRelaxed) join(l, r *LBSTRelaxed) *LBSTRelaxed {
   tree.share(l.root)
   tree.share(r.root)
   return &LBSTRelaxed{Tree{
      arena: tree.arena,
      root:  tree.join4(l.root, r.root, l.size, r.size),
      size:  l.size + r.size,
   }}
}

func (tree *LBSTRelaxed) Join(other List) List {
   return tree.join(tree, other.(*LBSTRelaxed))
}
