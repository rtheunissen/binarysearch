package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/utility"
)

type LBST struct {
   Tree
}

// Determines if two sizes are balanced.
func (LBST) isBalanced(x, y Size) bool {
   //
   // Given the sizes of two subtrees, an LBST is balanced when the difference
   // between the integer parts of log₂(x) and log₂(y) is no greater than one.
   //
   //                  -1 <= ⌊log₂(x)⌋ - ⌊log₂(y)⌋ <= 1
   //
   // This is similar to the height-balance rule, except instead of height we
   // consider the discrete the binary logarithm of the size of the subtree.
   //
   //    When `x` < `y`, is ⌊log₂(x)⌋ at most one less than ⌊log₂(y)⌋?
   //
   // Looking at the binary representation, the most significant bit or MSB is
   // the left-most bit set to 1, starting from the right. The bit position of
   // the MSB is equal to the floor of the binary logarithm of that integer.
   //
   //                                        00001101
   //                                             ↖
   //                                              MSB of 13 is at position 3
   //                                              log₂(13) = ~3.7
   //
   // Using this information, we can compare the bit position of the MSB of each
   // size without the need to calculate the logarithm itself.
   //
   // The MSB of `y` is within one step of the MSB of `x` (therefore balanced)
   // if after shifting `y` right, the MSB of `x` is less than the MSB of `y`.
   //
   //    Would it take more than one shift to align the MSBs of `y` and `x`?
   //
   // For example:
   //
   //              ↓                       ↓                    ↓
   //       s:   00100000              00001000              00010001
   //       y:   00111001              00010001              01001101
   //              ↑ BALANCED             ↑ BALANCED          ↑ NOT BALANCED
   //
   //
   // Returns true for s > y.
   //
   // Complexity: O(1)
   //
   return !utility.SmallerLog2(x, y>>1)
}

func (LBST) singleRotation(x, y Size) bool {
   return !utility.SmallerLog2(x, y)
}

func (tree LBST) join2(l *Node, r *Node, sl, sr Size) (k *Node) {
   if l == nil {
      return r
   }
   if r == nil {
      return l
   }
   if sl <= sr {
      r = tree.extractMin(r, sr, &k)
      return tree.join3(l, k, r, sl, sr-1)
   } else {
      l = tree.extractMax(l, sl, &k)
      return tree.join3(l, k, r, sl-1, sr)
   }
}

func (tree LBST) extractMin(p *Node, s Size, x **Node) *Node {
   tree.copy(&p)
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

func (tree LBST) extractMax(p *Node, s Size, x **Node) *Node {
   tree.copy(&p)
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

func (tree LBST) Join(that LBST) LBST {
   l := tree
   r := that
   tree.share(l.root)
   tree.share(r.root)
   return LBST{
      Tree{
         arena: tree.arena,
         root:  tree.join2(l.root, r.root, l.size, r.size),
         size:  l.size + r.size,
      },
   }
}

func (tree LBST) join3(l, k, r *Node, sl, sr Size) *Node {
   //tree.pathcopy(&k) // optional?
   if sl <= sr {
      return tree.assembleRL(k, l, r, sl, sr)
   } else {
      return tree.assembleLR(k, l, r, sl, sr)
   }
}

func (tree *LBST) assembleLR(p *Node, l, r *Node, sl, sr Size) *Node {
   if tree.isBalanced(sr, sl) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.copy(&l)

   sll := l.s
   slr := sl - l.s - 1

   l.r = tree.assembleLR(p, l.r, r, slr, sr)
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

func (tree *LBST) assembleRL(p *Node, l, r *Node, sl, sr Size) *Node {
   if tree.isBalanced(sl, sr) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.copy(&r)

   srl := r.s
   srr := sr - r.s - 1

   r.l = tree.assembleRL(p, l, r.l, sl, srl)
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

func (tree LBST) Split(i Position) (LBST, LBST) {
   tree.share(tree.root)
   l, r := JoinBased{Tree: tree.Tree, Joiner: tree}.split(tree.root, i, tree.size)

   return LBST{Tree{arena: tree.arena, root: l, size: i}},
      LBST{Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree LBST) verifyBalance(p *Node, s Size) {
   if p == nil {
      return
   }
   invariant(utility.Difference(utility.Log2(p.sizeL()), utility.Log2(p.sizeR(s))) <= 1)

   tree.verifyBalance(p.l, p.sizeL())
   tree.verifyBalance(p.r, p.sizeR(s))
}

func (tree LBST) verifyHeight(root *Node, size Size) {
   invariant(root.height() <= int(2*utility.Log2(size)))
}

func (tree LBST) verify(root *Node, size Size) {
   tree.verifySizes()
   tree.verifyBalance(root, size)
   tree.verifyHeight(root, size)
}

func (tree LBST) Verify() {
   tree.verify(tree.root, tree.size)
}
