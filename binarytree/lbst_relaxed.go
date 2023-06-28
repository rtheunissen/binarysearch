package binarytree

import (
   . "binarysearch/abstract/list"
)

// LBSTRelaxed / Relaxed Weight-Balanced Tree (Experimental)
//
// Recap:
//    The _height_ of a tree is the longest path from the root to any other node.
//    The _height_ of a node is the number of links to follow to reach the root.
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
// if the path exceeded the height upper-bound; if the height exceeds the tree's
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
//       Check: size < (1 << ((height + 1) >> 1))
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
   Log
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
//
func (tree *LBSTRelaxed) delete(p **Node, s Size, i Size) (x Data) {
   for {
      tree.persist(p)
      sl := (*p).sizeL()
      sr := (*p).sizeR(s)
      if i < sl {
         s = sl
    (*p).s = sl - 1
         p = &(*p).l
         continue
      }
      if i > sl {
         s = sr
         i = i - sl - 1
         p = &(*p).r
         continue
      }
      x := (*p).x
      *p = tree.join((*p).l, (*p).r, sl, sr)
      return x
   }
}

// Deletes the node at position `i` from the tree.
// Returns the data that was in the deleted value.
func (tree *LBSTRelaxed) Delete(i Position) Data {
 assert(i < tree.size)
 x := tree.delete(&tree.root, tree.size, i)
 tree.size = tree.size - 1
 return x
}
func (tree *LBSTRelaxed) insert(p **Node, s Size, i Position, x Data) {
   var unbalancedNode **Node // An unbalanced node along the insertion path.
   var unbalancedSize Size   // The size of the unbalanced node.
   var height uint64         // The height of the insertion so far.
   // Search with increasing height until the end of the path is reached.
   for {
      // Attach a new node at the end of the path.
      if *p == nil {
         *p = tree.allocate(Node{x: x})

         // Check if a rebuild is required.
         if tree.size < 1 << (height >> 1) {
            tree.rebuild(unbalancedNode, unbalancedSize)
         }
         return
      }
      tree.persist(p)
      height++

      sl := (*p).sizeL()
      sr := (*p).sizeR(s)
      if i <= sl {
         if unbalancedNode == nil && !tree.isBalanced(sr, sl + 1) {
            unbalancedNode = p
            unbalancedSize = s + 1
         }
         p = insertL(*p)
         s = sl

      } else {
         if unbalancedNode == nil && !tree.isBalanced(sl, sr + 1) {
            unbalancedNode = p
            unbalancedSize = s + 1
         }
         p = insertR(*p, &i)
         s = sr
      }
   }
}
// Inserts a value `s` at position `i` in the tree.
func (tree *LBSTRelaxed) Insert(i Position, x Data) {
   assert(i <= tree.size)
   tree.size = tree.size + 1
   tree.insert(&tree.root, tree.size, i, x)
}

// Determines if the height at which a new node was inserted was too deep, i.e.
// if the new height of the tree exceeds the upper-bound for its size.
//
//        tooDeep := height > 2 * ⌊log₂(size)⌋
//
//     because root.height() <= MaximumPathLength(2 * math.FloorLog2(size))
//
//   height is height + 1 ?
func (tree *LBSTRelaxed) tooDeep(size Size, height uint64) bool {
   return (1 << ((height + 1) >> 1)) > size
}

func (tree *LBSTRelaxed) balance(p *Node, s Size) *Node {
   if s <= 3 {
      return p
   }
   if !(tree.balanced(p, s)) {
      p = tree.partition(p, s >> 1)
   }
   p.l = tree.balance(p.l, p.sizeL())
   p.r = tree.balance(p.r, p.sizeR(s))
   return p
}

func (tree *LBSTRelaxed) rebuild(p **Node, s Size) {
   *p = tree.balance(*p, s)
}


func (tree *LBSTRelaxed) Split(i Size) (List, List) {
   assert(i <= tree.size)

   tree.share(tree.root)
   l,r := tree.split(tree.root, i)

   return &LBSTRelaxed{Tree: Tree{arena: tree.arena, root: l, size: i}},
          &LBSTRelaxed{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
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
   r.l = tree.joinLr(l, r.l, sl, r.s) // wtf shouldn't this be joinlR
   r.s = sl + r.s
   return r
}

func (tree *LBSTRelaxed) join(l, r *Node, sl, sr Size) *Node {
   if sl > sr {
      return tree.joinLr(l, r, sl, sr)
   } else {
      return tree.joinlR(l, r, sl, sr)
   }
}

func (tree *LBSTRelaxed) Join(other List) List {
   tree.share(tree.root)
   tree.share(other.(*LBSTRelaxed).root)
   return &LBSTRelaxed{Tree: Tree{
      arena: tree.arena,
      root:  tree.join(tree.root, other.(*LBSTRelaxed).root, tree.size, other.(*LBSTRelaxed).size),
      size:  tree.size + other.(*LBSTRelaxed).size,
   }}
}
