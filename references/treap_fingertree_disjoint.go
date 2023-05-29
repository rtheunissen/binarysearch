package references
//
//import (
//   
//   "math"
//)
//
//type FingerTreeDisjointGeometric struct {
//   Zip
//   head BST
//   tail BST
//}
//
//func (tree FingerTreeDisjointGeometric) HeadSize() uint64 {
//   return tree.head.size
//}
//
//// TODO: This function may not allocate memory
////       Maybe instead the tree can expose specific metrics
////       Rather than a root?
//func (tree FingerTreeDisjointGeometric) Root() BinaryTreeNode {
//   p := &Node{
//      l: tree.reverseL(tree.head.root.share(), nil, tree.head.size),
//      r: tree.reverseR(tree.tail.root.share(), nil),
//      s: tree.head.size,
//      y: math.MaxUint64,
//   }
//   return p
//}
//
//func (tree FingerTreeDisjointGeometric) Size() Size {
//   return tree.head.size + tree.tail.size
//}
//
//// in place
//func (tree Zip) toFingerTreeUniformDisjoint() (result FingerTreeDisjointGeometric) {
//   result.Random = tree.Random
//   result.head.root = nil
//   result.head.size = 0
//   result.tail.root = reverseR(tree.root, nil)
//   result.tail.size = tree.size
//   result.distribute()
//   return result
//}
//
//func (FingerTreeDisjointGeometric) fromArray(x []Value) FingerTreeDisjointGeometric {
//   return Zip{}.fromArray(x).toFingerTreeUniformDisjoint()
//}
//
//func (FingerTreeDisjointGeometric) New(x ...Value) List {
//   tree := FingerTreeDisjointGeometric{}.fromArray(x)
//   return &tree
//}
//
//func (tree FingerTreeDisjointGeometric) Clone() List {
//   tree.head.root.share()
//   tree.tail.root.share()
//   return &tree
//}
//
//func (tree *FingerTreeDisjointGeometric) reverseL(p *Node, g *Node, s Size) *Node {
//   assert(s == p.count())
//   for {
//      if p == nil {
//         return g
//      }
//      p = p.Copy()
//      sl := p.s
//      p.s = p.sizeR(s)
//      l := p.l
//      p.l = g
//      g = p
//      p = l
//      s = sl
//   }
//}
//
//func (tree *FingerTreeDisjointGeometric) reverseR(p *Node, g *Node) *Node {
//   for {
//      if p == nil {
//         return g
//      }
//       p = p.Copy()
//       r := p.r
//       p.r = g
//       g = p
//       p = r
//   }
//}
//
//func (tree *FingerTreeDisjointGeometric) rotateUpL(p *Node) {
//   for p.l != nil && p.y > p.l.y {
//      p.l = p.l.Copy()
//      l := p.l
//      p.l = l.l
//      l.l = p.r
//      p.r = l
//      p.s = p.s + l.s + 1
//      l.s = p.s - l.s - 1
//   }
//}
//func (tree *FingerTreeDisjointGeometric) rotateDownL(p *Node) {
//   for p.r != nil && p.r.y > p.y {
//      p.r = p.r.Copy()
//      r := p.r
//      p.r = r.l
//      r.l = p.l
//      p.l = r
//      r.s = p.s - r.s - 1
//      p.s = p.s - r.s - 1
//   }
//}
//func (tree *FingerTreeDisjointGeometric) rotateUpR(p *Node) {
//   for p.r != nil && p.y > p.r.y {
//      p.r = p.r.Copy()
//      r := p.r                // parent on the spine
//      p.r = r.r
//      r.r = p.l
//      p.l = r
//      p.s = p.s + r.s + 1
//   }
//}
//func (tree *FingerTreeDisjointGeometric) rotateDownR(p *Node) {
//   for p.l != nil && p.l.y > p.y {
//      p.l = p.l.Copy()
//      l := p.l
//      p.l = l.r
//      l.r = p.r
//      p.r = l
//      p.s = p.s - l.s - 1
//   }
//}
//
//
////tree.flipSizesToRightDescendingLeft(&p,  tree.head.size)
////tree.head.root = reverseL(p, nil)
////
////func (tree *FingerTreap) rotateParentRightOnLeftSpine(p *Node) {
////
////
////}
////func (tree *FingerTreap) rotateParentLeftOnRightSpine(p *Node) {
////
////}
////
////func (tree *FingerTreap) rotateIntoRightSpine(p *Node) {
////
////
////}
////func (tree *FingerTreap) rotateIntoLeftSpine(p *Node) {
////   //
////   //
////
////}
////
//
//
//
//func (tree FingerTreeDisjointGeometric) Access(i Position) Value {
//   assert(i < tree.Size())
//   if i < tree.head.size {
//      return tree.accessFromHead(i)
//   } else {
//      return tree.accessFromTail(tree.head.size + tree.tail.size - i - 1)
//   }
//}
//func (tree FingerTreeDisjointGeometric) accessFromHead(i Position) Value {
//   p := tree.head.root
//   for i > p.s {
//       i = i - p.s - 1
//       p = p.l
//   }
//   return BST{}.access(p, p.s+ i)
//}
//func (tree FingerTreeDisjointGeometric) accessFromTail(i Position) Value {
//   p := tree.tail.root
//   for i > p.s {
//       i = i - p.s - 1
//       p = p.r
//   }
//   return BST{}.access(p, p.s- i)
//}
//
//
//
//
//func (tree *FingerTreeDisjointGeometric) Update(i Position, x Value) {
//   assert(i < tree.Size())
//   if i < tree.head.size {
//      tree.updateFromHead(x, i)
//   } else {
//      tree.updateFromTail(x, tree.head.size + tree.tail.size - i - 1)
//   }
//}
//func (tree *FingerTreeDisjointGeometric) updateFromHead(x Value, i Position) {
//   tree.head.root = pathcopy(tree.head.root)
//   p := tree.head.root
//   for {
//      if i == 0 {
//         p.x = x
//         return
//      }
//      if i > p.s {
//         p.l = p.l.Copy()
//         i = i - p.s - 1
//         p = p.l
//      } else {
//         p.r = p.r.Copy()
//         BST{}.update(p.r, i - 1, x) // TODO: consider p.r.update vs update(p.r, i - 1, s)
//         return
//      }
//   }
//}
//func (tree *FingerTreeDisjointGeometric) updateFromTail(x Value, i Position) {
//   tree.tail.root = pathcopy(tree.tail.root)
//   p := tree.tail.root
//   for {
//      if i == 0 {
//         p.x = x
//         return
//      }
//      if i > p.s {
//         p.r = p.r.Copy()
//         i = i - p.s - 1
//         p = p.r
//      } else {
//         p.l = p.l.Copy()
//         BST{}.update(p.l, p.s- i, x)
//         return
//      }
//   }
//}
//
//
//
//
//
//func (tree *FingerTreeDisjointGeometric) appendHead(x Value) {
//   p := tree.allocate(Node{x: x})
//   p.l = tree.head.root
//   tree.rotateUpL(p)
//   tree.head.root = p
//   tree.head.size++
//   tree.distribute()
//}
//func (tree *FingerTreeDisjointGeometric) appendTail(x Value) {
//   p := tree.allocate(Node{x: x})
//   p.r = tree.tail.root
//   tree.rotateUpR(p)
//   tree.tail.root = p
//   tree.tail.size++
//   tree.distribute()
//}
//func (tree *FingerTreeDisjointGeometric) insertFromHead(x Value, i Position) {
//   if i == 0 {
//     tree.appendHead(x)
//     return
//   }
//   assert(i > 0)
//   assert(i < tree.head.size)
//
//   tree.head.root = pathcopy(tree.head.root)
//   tree.head.size++
//
//   n := tree.allocate(Node{x: x})
//   p := tree.head.root
//   for i > p.s+ 1 {
//     p.l = p.l.Copy()
//     i = i - p.s - 1
//     p = p.l
//   }
//   if rank(n) <= rank(p) {
//     Treap{}.insert(&p.r, i - 1, n)
//     p.s++
//   } else {
//     partition(p.r, i - 1, &p.r, &n.r)
//     n.s = p.s - i + 1
//     p.s = i - 1
//     n.l = p.l
//     p.l = n
//     tree.rotateUpL(n)
//   }
//}
//
//func (tree *FingerTreeDisjointGeometric) splitFromHead(i Position) (List, List) {
//   assert(i > 0)
//   assert(i < tree.head.size)
//
//   //
//   L := &FingerTreeDisjointGeometric{
//      head: BST{
//         root: tree.head.root.copy(),
//         size: i,
//      },
//   }
//
//   p := L.head.root
//   for i > p.s + 1 {
//      i = i - (p.s + 1)
//      p.l = p.l.Copy()
//      p = p.l
//   }
//   sl := i - 1
//   sr := (p.s + 1) - i
//
//   l, r := split(p.r, i - 1)
//   assert(sl == l.count())
//   assert(sr == r.count())
//
//   R := &FingerTreeDisjointGeometric{
//      head: BST{
//         root: tree.reverseL(r, p.l, sr),
//         size: tree.head.size - L.head.size,
//      },
//      tail: BST{
//         root: tree.tail.root,
//         size: tree.tail.size,
//      },
//   }
//   p.l = nil
//   p.r = l
//   p.s = sl
//
//   L.distribute()
//   R.distribute()
//   return L, R
//}
//
//func (tree *FingerTreeDisjointGeometric) splitFromTail(i Position) (List, List) {
//   assert(i > 0)
//   assert(i <= tree.tail.size)
//
//   L := &FingerTreeDisjointGeometric{head: *tree.head.Clone().(*BST)}
//   R := &FingerTreeDisjointGeometric{tail: *tree.tail.Clone().(*BST)}
//
//   R.tail.size = i
//   L.tail.size = tree.tail.size - i
//
//   R.tail.root = R.tail.root.copy()
//   L.head.root = L.head.root.copy()
//
//   p := R.tail.root
//   for i > p.s + 1 {
//      i = i - (p.s + 1)
//      p.r = p.r.Copy()
//      p = p.r
//   }
//
//   if p.l != nil {
//      p.l = p.l.Copy()
//   }
//   l, r := split(p.l, p.s + 1 - i)
//
//   if p.r != nil {
//      p.r = p.r.Copy()
//   }
//   L.tail.root = tree.reverseR(l.share(), p.r.share())
//
//   p.r = nil
//   p.l = r.share()
//   p.s = i - 1
//
//   L.distribute()
//   R.distribute()
//
//   return L, R
//}
//
//
//
//func (tree *FingerTreeDisjointGeometric) Split(i Position) (List, List) {
//   if i == 0 {
//      return FingerTreeDisjointGeometric{}.New(), tree.Clone()
//   }
//   if i < tree.head.size {
//      return tree.splitFromHead(i)
//   } else {
//      return tree.splitFromTail(tree.Size() - i)
//   }
//}
//
//
//
//
//func (tree *FingerTreeDisjointGeometric) insertFromTail(x Value, i Position) {
//   if i == 0 {
//      tree.appendTail(x)
//      return
//   }
//   assert(i > 0)
//   assert(i <= tree.tail.size)
//
//   tree.tail.root = pathcopy(tree.tail.root)
//   tree.tail.size++
//
//   n := tree.allocate(Node{x: x})
//   p := tree.tail.root
//
//   // find the turning node
//
//   for i > p.s+ 1 {
//      p.r = p.r.Copy()
//      i = i - p.s - 1
//      p = p.r
//   }
//
//   // insert descending
//   if rank(n) <= rank(p) {
//      Treap{}.insert(&p.l, p.s+ 1 - i, n)
//      p.s++
//   } else {
//      // insert on the right spine.
//      partition(p.l, p.s+ 1 - i, &n.l, &p.l)
//      n.s = p.s - i + 1
//      p.s = i - 1
//      n.r = p.r
//      p.r = n
//      tree.rotateUpR(n)
//   }
//}
//
//
//
//func (tree *FingerTreeDisjointGeometric) Insert(i Position, x Value) {
//   if i < tree.head.size {
//      tree.insertFromHead(x, i)
//   } else {
//      tree.insertFromTail(x, tree.Size() - i) // TODO: maybe this can just take i?
//   }
//   tree.distribute()
//   return
//}
//
//func (tree *FingerTreeDisjointGeometric) Delete(i Position) (x Value) {
//   if i < tree.head.size {
//      if i == 0 {
//         tree.deleteFirst(&x)
//         return x
//      }
//      tree.deleteFromHead(i, &x)
//      tree.distribute()
//      return x
//   } else {
//      if i == tree.head.size + tree.tail.size - 1 {
//         tree.deleteLast(&x)
//         return x
//      }
//      tree.deleteFromTail(tree.head.size + tree.tail.size - i - 1, &x)
//      tree.distribute()
//      return x
//   }
//}
//
//func (tree *FingerTreeDisjointGeometric) deleteFromHead(i Position, x *Value) {
//   tree.head.root = pathcopy(tree.head.root)
//   tree.head.size--
//   p := tree.head.root
//   for i > p.s+ 1 {
//      p.l = p.l.Copy()
//      i = i - p.s - 1
//      p = p.l
//   }
//   if i < p.s+ 1 {
//      Treap{}.delete(&p.r, i - 1, x)
//      p.s--
//      return
//   }
//   p.l = p.l.Copy()
//   g := p.l
//   *x = g.x
//   p.r = Treap{}.join(p.r, g.r, p.s)
//   p.l = g.l
//   p.s = p.s + g.s
//   tree.rotateDownL(p)
//}
//
//func (tree *FingerTreeDisjointGeometric) deleteFromTail(i Position, x *Value) {
//   tree.tail.root = pathcopy(tree.tail.root)
//   tree.tail.size--
//   p := tree.tail.root
//   for i > p.s+ 1 {
//      p.r = p.r.Copy()
//      i = i - p.s - 1
//      p = p.r
//   }
//   if i < p.s+ 1 {
//      Treap{}.delete(&p.l, p.s- i, x)
//      p.s--
//      return
//   }
//   p.r = p.r.Copy()
//   g := p.r
//   *x = g.x
//   p.l = Treap{}.join(g.l, p.l, g.s)
//   p.r = g.r
//   p.s = p.s + g.s
//   tree.rotateDownR(p)
//}
//
//func (tree *FingerTreeDisjointGeometric) deleteLast(x *Value) {
//
//   *x = tree.tail.root.x
//
//   tree.tail.root = tree.tail.root.copy()
//   tree.tail.root = tree.reverseR(tree.tail.root.l, tree.tail.root.r)
//   tree.tail.size = tree.tail.size - 1
//   tree.distribute()
//}
//
//
//func (tree *FingerTreeDisjointGeometric) deleteFirst(x *Value) {
//
//   *x = tree.head.root.x
//
//   tree.head.root = tree.head.root.copy()
//   tree.head.root = tree.reverseL(tree.head.root.r, tree.head.root.l, tree.head.root.s)
//   tree.head.size = tree.head.size - 1
//   tree.distribute()
//}
//
//func (tree *FingerTreeDisjointGeometric) distribute() {
//   if tree.head.root == nil {
//      tree.transferFromTailToHead()
//      if tree.tail.root == nil {
//         tree.transferFromHeadToTail()
//      }
//   } else if tree.tail.root == nil {
//      tree.transferFromHeadToTail()
//      if tree.head.root == nil {
//         tree.transferFromTailToHead()
//      }
//   }
//}
//
//func (tree *FingerTreeDisjointGeometric) transferFromTailToHead() {
//   if tree.tail.root == nil {
//      return
//   }
//   p := truncateR(&tree.tail.root)
//   sl := p.s + 1
//
//   tree.head.size = sl
//   tree.tail.size = tree.tail.size - tree.head.size
//   tree.head.root = tree.reverseL(p, nil, tree.head.size)
//}
//
//func (tree *FingerTreeDisjointGeometric) transferFromHeadToTail() {
//   if tree.head.root == nil {
//      return
//   }
//   p := truncateL(&tree.head.root)
//   sr := p.s + 1
//   p.s = 0
//
//   tree.tail.size = sr
//   tree.head.size = tree.head.size - tree.tail.size
//   tree.tail.root = tree.reverseR(p, nil)
//}
//// create new node with random rank
//// if this rank is less than the root's rank, we need a new root node and the
////    old root node will
//
//
//
//
//
//
//
//
//
//
//
//
//func (tree FingerTreeDisjointGeometric) Join(that List) List {
//   return &tree
//}
//
//
//
//func (tree FingerTreeDisjointGeometric) eachFromHead(p *Node, visit func (Value)) {
//   if p == nil {
//      return
//   }
//   visit(p.x)
//   p.r.inorder(visit)
//   tree.eachFromHead(p.l, visit)
//}
//func (tree FingerTreeDisjointGeometric) eachFromTail(p *Node, visit func (Value)) {
//   if p == nil {
//      return
//   }
//   tree.eachFromTail(p.r, visit)
//   p.l.inorder(visit)
//   visit(p.x)
//}
//func (tree FingerTreeDisjointGeometric) Each(visit func (Value)) {
//   tree.eachFromHead(tree.head.root, visit)
//   tree.eachFromTail(tree.tail.root, visit)
//}
//
//
//
//// func (p *Node) preorder(visit func(Value)) {
////    if p == nil {
////       return
////    }
////    visit(p.s)
////    p.l.preorder(visit)
////    p.r.preorder(visit)
//// }
////
//// func (p *Node) postorder(visit func(Value)) {
////    if p == nil {
////       return
////    }
////    p.l.postorder(visit)
////    p.r.postorder(visit)
////    visit(p.s)
//// }
//
//
//func (tree FingerTreeDisjointGeometric) verifyRankMaxHeap(p *Node) {
//   if p == nil {
//      return
//   }
//   invariant(rank(p) >= rank(p.l))
//   invariant(rank(p) >= rank(p.r))
//
//   tree.verifyRankMaxHeap(p.l)
//   tree.verifyRankMaxHeap(p.r)
//}
//
//func (tree FingerTreeDisjointGeometric) Verify() {
//   head := tree.head.root
//   tail := tree.tail.root
//
//   // There should always be a head and a tail when the size is > 1
//   // to utility worst case O(1) access to both the head and the tail.
//   if tree.head.size + tree.tail.size > 1 {
//      invariant(head != nil)
//      invariant(tail != nil)
//   }
//   //The root's size must be equal to the size of the left subtree.
//   invariant(tree.head.size == head.count())
//   invariant(tree.tail.size == tail.count())
//   //
//   // Verify internal treap positions along the spines.
//   for l := head; l != nil; l = l.l { verifySizes(l.r) }
//   for r := tail; r != nil; r = r.r { verifySizes(r.l) }
//
//   // Verify internal node counts, positions along the spines.
//   for l := head; l != nil; l = l.l { invariant(l.s == l.r.count()) }
//   for r := tail; r != nil; r = r.r { invariant(r.s == r.l.count()) }
//
//   //
//   for l := head; l != nil; l = l.l { tree.verifyRankMaxHeap(l.r) }
//   for r := tail; r != nil; r = r.r { tree.verifyRankMaxHeap(r.l) }
//
//   // Verify ranks ascending from the head and tail.
//   for l := head; l != nil; l = l.l { invariant(l.l == nil || rank(l) <= rank(l.l)) }
//   for r := tail; r != nil; r = r.r { invariant(r.r == nil || rank(r) <= rank(r.r)) }
//
//   //
//   for l := head; l != nil; l = l.l { invariant(rank(l) >= rank(l.r)) }
//   for r := tail; r != nil; r = r.r { invariant(rank(r) >= rank(r.l)) }
//}
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
////package structures
////
////import (
////  . "trees/pkg/abstract/list"
////  "trees/pkg/utility"
////  "golang.org/x/exp/rand"
////)
////
////
////
////// An implementation of a persistent list using a finger search tree (TreapFingerTreeEmbed).
//////
////// The first node of the sequence is the head, the last node is the tail. The
////// head is therefore the leftmost node in the tree and the tail the rightmost.
//////
////// Starting at the head, nodes along the left spine of the tree have their left
////// pointers reversed to point to their parent instead. Symmetrically along the
////// right spine starting from the tail, nodes have their right pointers reversed
////// to point to their parent. This provides access to internal nodes of the tree
////// by ascending along the spine from the head or tail and then descend inwards.
//////
////// The root has no children. Values access to nodes other than the root must start
////// from the head or the tail. There is no need to ascend both spines because the
////// position to search for can be compared to the relative position of the root
////// which is an absolute position in the sequence.
//////
//////
//////                               ROOT
//////                                   ↘ +4
//////                                    (m)
//////                            -2               +2
//////                            (s)             (l)
//////                        -1 ↗   ↘ +1     -1 ↙   ↖ +1
//////                        (e)     (a)     (p)     (e)
//////                       ↗                           ↖
//////                   HEAD                             TAIL
//////
//////
////// Another way to visualize this structure is to use a horizontal spine layout.
////// Notice that the relative position of every node along the spine indicates how
////// many nodes will be skipped if a search continues along in the same direction:
//////
//////
//////   HEAD                              ROOT                              TAIL
//////
//////   -1  -2    -4            -8        +16        +8            +4    +2  +1
//////    ○ → ○  →  ○ --------- → ○         ◎         ○ ← --------- ○  ←  ○ ← ○
//////        ↓     ↓             ↓                   ↓             ↓     ↓
//////        ○     ○             ○                   ○             ○     ○
//////            ↙   ↘        ↙     ↘             ↙     ↘        ↙   ↘
//////           ○     ○     ○         ○         ○         ○     ○     ○
//////                     ↙   ↘     ↙   ↘     ↙   ↘     ↙   ↘
//////                    ○     ○   ○     ○   ○     ○   ○     ○
//////
//////
//////
////// Operations:
//////
//////    From        O(n)        Creates a new instance from existing values.
//////    Get         O(lg n)     Returns a value at position.
//////    Set         O(lg n)     Updates a value at position.
//////    Insert      O(lg n)     Inserts a value at position, increasing length.
//////    Remove      O(lg n)     Removes a value at position, reducing length.
//////    Split       O(lg n)     Partition by position, in-place on the left.
//////    Join        O(lg n)     Concatenate by appending another TreapFingerTreeEmbed at the end.
//////    Push        O(1)        Adds a value after the last value.
//////    Unshift     O(1)        Adds a value before the first value.
//////    Pop         O(1)        Removes the last value.
//////    Shift       O(1)        Removes the first value.
//////
//////
//////
////// Reading list:
//////
//////  - Implement an immutable deque as a balanced binary tree? (2010)
//////    https://stackoverflow.com/q/3271256
//////
//////  - Functional Set Operations with Treaps (2001)
//////    Dan Blandford, Guy Blelloch
//////    https://www.cs.cmu.edu/afs/cs/project/pscico/pscico/papers/fingertrees/
//////    https://www.cs.cmu.edu/afs/cs/project/pscico/pscico/src/fingertrees/
//////
//////  - An O(n log log n)-Time Algorithm for Triangulating a Simple Polygon (1986)
//////    Appendix: Finger search trees
//////    R. E. Tarjan, C. J. Van Wyk
//////    https://api.semanticscholar.org/CorpusID:4981331
//////  - Finger trees: a simple general-purpose data structure (2006)
//////    R. Hinze, R. Paterson
//////    https://api.semanticscholar.org/CorpusID:6881581
//////
//////  - Finger Search Trees (2005)
//////    G. Brodal
//////    https://api.semanticscholar.org/CorpusID:5694716
//////
//////
//////
////// Symbols:
//////
//////       t   tree
//////       o   other tree
//////       i   index, offset, 0-based
//////       d   relative position, distance
//////       s   data
//////
//////
////type TreapFingerTreeEmbed struct {
////  Tree
////  root *Node // s ?
////  head *Node // l
////  tail *Node // r
////  size int   // w
////  // z another int available here
////}
////
////func (t *TreapFingerTreeEmbed) New() List {
////  return &TreapFingerTreeEmbed{}
////}
////
////// Creates a new instance of TreapFingerTreeEmbed by building a Treap then converting it to an TreapFingerTreeEmbed.
//////
////// TODO: Build recursively from the head and tail towards the root.
////// TODO: Explore building by push without copy, aka transience.
//////
////func (t *TreapFingerTreeEmbed) From(arr []Value) List {
////  return Treap{}.From(arr).(*Treap).ToFST()
////}
////
////// Returns the size of the tree, the length of its sequence.
////func (t *TreapFingerTreeEmbed) Len() int {
////  return t.size
////}
////
////// Creates a shallow copy of an TreapFingerTreeEmbed.
////// TODO just use structs where it makes sense to.
////func (t TreapFingerTreeEmbed) Clone() List {
////  /// TODO extract this to a node method
////  if t.root != nil {
////     t.root.rc++
////  }
////  ///////
////  if t.head != nil {
////     t.head.rc++
////  }
////  if t.tail != nil {
////     t.tail.rc++
////  }
////  return &t
////}
////
////// Returns the value at i.
////func (t *TreapFingerTreeEmbed) Get(i int) Value {
////  assert(i < t.size)
////  return t.seekTo(i).s
////}
////
////// Replaces the value at i, returns t.
////func (t *TreapFingerTreeEmbed) Set(i int, s Value) {
////  assert(i < t.size) // TODO move these assertions to the list impl?
////  t.shadowTo(i).withData(s)
////}
////
//////
////func (t *TreapFingerTreeEmbed) seekTo(i int) *Node {
////  switch {
////     case i < t.root.sizeRL(): return t.seekFromHead(i)
////     case i > t.root.sizeRL(): return t.seekFromTail(i)
////  default:
////     return t.root
////  }
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) shadowTo(i int) *Node {
////  switch {
////     case i < t.root.sizeRL(): return t.shadowFromHead(i)
////     case i > t.root.sizeRL(): return t.shadowFromTail(i)
////  default:
////     return shadow(&t.root)
////  }
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) seekFromHead(i int) *Node {
////  return t.seekL(t.head, i)
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) seekFromTail(i int) *Node {
////  return t.seekR(t.tail, i - t.size + 1)
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) shadowFromHead(i int) *Node {
////  return t.seekShadowL(shadow(&t.head), i)
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) shadowFromTail(i int) *Node {
////  return t.seekShadowR(shadow(&t.tail), i - t.size + 1)
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
//// func (t *TreapFingerTreeEmbed) seekL(p *Node, i int) *Node {
////  for {
////     invariant(i >= 0)
////     invariant(p.isL())
////     //
////     if i == 0 {
////        return p
////     }
////     //
////     if i + p.w < 0 { // d - (sizeLR(p) + 1) < 0
////        return seekTo(i, p.r)
////     }
////     //
////     i = i + p.w
////     p = p.l
////  }
//// }
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) seekR(p *Node, d int) *Node {
////  for {
////     invariant(d <= 0)
////     invariant(p.isR())
////
////     //
////     if d == 0 {
////        return p
////     }
////     //
////     if d + p.w > 0 { // d - (sizeRL(p) + 1) < 0
////        return seekTo(d, p.l)
////     }
////     //
////     d = d + p.w
////     p = p.r
////  }
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) seekShadowL(p *Node, d int) *Node {
////  for {
////     invariant(d >= 0)
////     invariant(p.isL())
////     //
////     if d == 0 {
////        return p
////     }
////     invariant(d > 0)
////     //
////     if d + p.w < 0 { // d - (sizeLR(p) + 1) < 0
////        return shadowTo(d, &p.r)
////     }
////     //
////     d = d + p.w
////     p = shadow(&p.l)
////  }
////}
////
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) seekShadowR(p *Node, d int) *Node {
////  for {
////     invariant(d <= 0)
////     invariant(p.isR())
////
////     //
////     if d == 0 {
////        return p
////     }
////     //
////     if d + p.w > 0 { // d - (sizeRL(p) + 1) < 0
////        return shadowTo(d, &p.l)
////     }
////     //
////     d = d + p.w
////     p = shadow(&p.r)
////  }
////}
////
////func generateRandomRank() Rank {
////  return random.Uint64()
////  //var level Rank
////  //for random.Uint64() & 1 == 0 {
////  //   level++
////  //}
////  ////fmt.Println(level)
////  //return level
////}
////
////// Creates a new node containing a given data value with a fixed random rank.
////func (TreapFingerTreeEmbed) createNodeWithData(s Value) *Node {
////  return &Node{
////     s: s,
////     z: generateRandomRank(),
////  }
////}
////func (Treap) createNodeWithData(s Value) *Node {
////  return &Node{
////     s: s,
////     z: generateRandomRank(),
////  }
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) Push(s Value) {
////  if t.isEmpty() {
////     t.root = t.createNodeWithData(s).withRelativePosition(+1)
////  } else {
////     t.tail = t.rotateUpR(t.createNodeWithData(s).withRelativePosition(+1).withR(t.tail))
////  }
////  t.size++
////}
////
////
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) Unshift(s Value) {
////  if t.isEmpty() {
////     t.root = t.createNodeWithData(s).withRelativePosition(+1)
////  } else {
////     //
////     t.root = copyOf(t.root); t.root.increaseInternalWeightOfR()
////     t.head = t.rotateUpL(t.createNodeWithData(s).withRelativePosition(-1).withL(t.head))
////  }
////  t.size++
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) Pop() (s Value) {
////  //
////  if t.tail == nil {
////     return t.removeRoot()
////  }
////
////  //
////  s = t.tail.s
////  t.tail = copyOf(t.tail)
////
////  //
////  if t.tail.hasL() {
////     t.tail.l = t.tail.l.copy().toR(t.tail.sizeRL())
////  }
////  t.tail = reverseR(t.tail.l, t.tail.r)
////  t.size--
////  return
////}
////
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) Shift() (s Value) {
////  assert(!t.isEmpty())
////  if t.head != nil {
////     s = t.head.s
////     t.dissolvehead
////  } else {
////     s = t.root.s
////     t.removeRoot()
////  }
////  return s
////}
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) rotateUpR(p *Node) *Node {
////  invariant(p.isR())
////  //
////  //
////  for p.hasR() {
////     if rankOf(p) <= rankOf(p.r) {
////        return p
////     }
////     t.rotateLeftOnRightSpine(p)
////  }
////  //
////  //
////  if rankOf(p) <= rankOf(t.root) {
////     return p
////  }
////  t.rotateLeftIntoRoot(p)
////  return nil
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) rotateUpL(p *Node) *Node {
////  invariant(p.isL())
////  //
////  //
////  for p.hasL() {
////     if rankOf(p) <= rankOf(p.l) {
////        return p
////     }
////     t.rotateRightOnLeftSpine(p)
////  }
////  //
////  //
////  if rankOf(p) <= rankOf(t.root) {
////     return p
////  }
////  t.rotateRightIntoRoot(p)
////  return nil
////}
////
////
////
////
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////
////func (t *TreapFingerTreeEmbed) rotateRightIntoRoot(p *Node) {
////  //visualize("↻")
////  invariant(p.isL())
////  invariant(!p.hasL())
////  invariant(!t.isEmpty())
////  //
////  //
////  g := copyOf(t.root)
////  r := p.r
////  //
////  //
////  appendR(&t.tail, g)
////
////  g.l = r; t.root = p; p.r = nil
////  rotateRelativePositions(g, p, r)
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////
////func (t *TreapFingerTreeEmbed) rotateLeftIntoRoot(p *Node) {
////  //visualize("↺")
////  invariant(p.isR())
////  invariant(!p.hasR())
////  invariant(!t.isEmpty())
////  // invariant rank of p is greater than root
////  g := copyOf(t.root)
////  l := p.l
////
////  //
////  //
////  appendL(&t.head, g)
////
////  g.r = l
////  t.root = p
////  p.l = nil
////  rotateRelativePositions(g, p, l)
////}
////
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) rotateLeftOnRightSpine(p *Node)  {
////  invariant(p.isR())
////  //visualize("↺")
////  g := p.r.copy()
////  //
////  // as
////  p.r, p.l, g.r = g.r, g, p.l
////  rotateRelativePositions(g, p, g.r)
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) rotateRightOnLeftSpine(p *Node) {
////  invariant(p.isL())
////  //visualize("↻")
////  g := p.l.copy()
////  //
////  // as
////  p.l, p.r, g.l = g.l, g, p.r
////  rotateRelativePositions(g, p, g.l)
////}
////
////
////
//////
//////
////func (t *TreapFingerTreeEmbed) Insert(i int, s Value) {
////  assert(i <= t.size)
////
////  // takes care of the null root case also.
////  if i == 0 {
////     t.Unshift(s)
////     return
////  }
////  if i == t.size {
////     t.Push(s)
////     return
////  }
////  if i < t.root.w {
////     t.insertFromHead(s, i)
////  } else {
////     t.insertFromTail(s, i)
////  }
////}
////
//////
////func (t *TreapFingerTreeEmbed) insertOnRightSpine(p *Node, n *Node, d int) {
////  invariant(!n.hasL())
////  invariant(!n.hasR())
////  invariant(d < 0)
////
////  n.l, p.l = unzip(p.l, d, p.sizeRL())
////
////  n.w = d + p.w + 1 // TODO revise
////  p.w = -d
////
////  //
////  if n.hasL() {
////     n.l.toL(n.sizeRL())
////  }
////  n.r = p.r
////  p.r = t.rotateUpR(n)
////}
////
//////
////func (t *TreapFingerTreeEmbed) insertOnLeftSpine(p *Node, n *Node, d int) {
////  invariant(!n.hasL())
////  invariant(!n.hasR())
////  invariant(d >= 0)
////  //
////  //
////  p.r, n.r = unzip(p.r, d, p.sizeLR())
////
////  //
////  n.w = p.w + d // TODO revise
////  p.w = -d - 1
////
////  //
////  if n.hasR() {
////     n.r.toR(n.sizeLR())
////  }
////  n.l = p.l
////  p.l = t.rotateUpL(n)
////}
////
////
//////
//////
////func (t *TreapFingerTreeEmbed) insertFromTail(s Value, i int) {
////  //visualize("↖")
////
////  t.tail = copyOf(t.tail)
////
////  n := t.createNodeWithData(s)
////  p := t.tail
////  d := i - t.size
////
////  //
////  t.size = t.size + 1
////
////  for {
////     //visualize("↖")
////     invariant(d <= 0)
////     invariant(p.isR())
////
////     if d + p.w < 0 {
////        d = d + p.w
////        p = shadow(&p.r)
////        continue
////     }
////     if rankOf(p) < rankOf(n) {
////        t.insertOnRightSpine(p, n, d)
////        return
////     }
////     //visualize("|↙")
////     (&Treap{}).insert(&p.l, n, d, p.sizeRL()); p.increaseInternalWeightOfR()
////     return
////  }
////}
////
////func (t *TreapFingerTreeEmbed) insertFromHead(s Value, i int) {
////  //visualize("↗")
////  invariant(i >= 0)
////
////  //
////  t.head = copyOf(t.head)
////  t.root = copyOf(t.root)
////  t.size = t.size + 1
////
////  t.root.increaseInternalWeightOfR()
////  //
////  n := t.createNodeWithData(s)
////  p := t.head
////  d := i - 1
////
////  for {
////     //visualize("↗")
////     invariant(d >= 0)
////     invariant(p.isL())
////
////     //
////     //
////     if d + p.w >= 0  {
////        d = d + p.w
////        p = shadow(&p.l)
////        continue
////     }
////     //
////     //
////     if rankOf(p) < rankOf(n) {
////        t.insertOnLeftSpine(p, n, d)
////        return
////     }
////     //
////     //
////     //visualize("|↘")
////     (&Treap{}).insert(&p.r, n, d, p.sizeLR()); p.increaseInternalWeightOfL()
////     return
////  }
////}
////
////
//////
//////
////func (t *TreapFingerTreeEmbed) removeFromHead(i int) (s Value) {
////  invariant(i >= 0)
////
////  //
////  //
////  if i == 0 {
////     return t.Shift()
////  }
////
////  //
////  t.size = t.size - 1
////  t.root = t.root.copy().withRelativePosition(t.root.w - 1)
////  t.head = t.head.copy()
////
////  //
////  p := t.head
////  d := i
////
////  //
////  //
////  for {
////     //visualize("↗")
////     invariant(d > 0)
////     invariant(p.isL())
////
////     //
////     //
////     if d + p.w < 0 {
////        //visualize("|↘")
////        p.decreaseInternalWeightOfL()
////        return (&Treap{}).remove(&p.r, d, p.sizeLR()+1).s
////     }
////
////     //
////     //
////     if d + p.w == 0 {
////        dissolveOnLeftSpine(p, &s)
////        return
////     }
////
////     //
////     //
////     d = d + p.w
////     p = shadow(&p.l)
////  }
////}
////
//////
//////
////func (t *TreapFingerTreeEmbed) removeFromTail(i int) (s Value) {
////  invariant(i >= 0)
////
////  //
////  if i == t.size - 1 {
////     return t.Pop()
////  }
////
////  //
////  t.size = t.size - 1
////  t.tail = t.tail.copy()
////
////  //
////  p := t.tail
////  d := i - t.size
////
////  //
////  //
////  for {
////     //visualize("↖")
////     invariant(d < 0)
////     invariant(p.isR())
////
////     //
////     //
////     if d + p.w > 0 {
////        //visualize("|↙")
////        p.decreaseInternalWeightOfR()
////        return (&Treap{}).remove(&p.l, d, p.sizeRL()+1).s
////     }
////
////     //
////     //
////     if d + p.w == 0 {
////        dissolveOnRightSpine(p, &s)
////        return
////     }
////
////     //
////     //
////     d = d + p.w
////     p = shadow(&p.r)
////  }
////}
////
//////
//////
////func dissolveOnLeftSpine(l *Node, s *Value) {
////  //visualize("!")
////
////  p  := copyOf(l.l)
////  g  := p.l
////  p.l = l
////  l.l = nil
////  *s  = p.s
////  *l  = *reverseL((&Treap{}).dissolveWithoutSize(p), g)
////}
////
//////
//////
////func dissolveOnRightSpine(r *Node, s *Value) {
////  //visualize("!")
////  invariant(r.isR())
////
////  p  := r.r.copy()
////  g  := p.r
////  p.r = r
////  r.r = nil
////  *s  = p.s
////  *r  = *reverseR((&Treap{}).dissolveWithoutSize(p), g)
////}
////
////
////
//////
//////
////func (t *TreapFingerTreeEmbed) Remove(i int) (s Value){
////  switch {
////  case i < t.root.sizeRL(): return t.removeFromHead(i)
////  case i > t.root.sizeRL(): return t.removeFromTail(i)
////  default:
////     return t.removeRoot()
////  }
////}
////
////
////func (t *TreapFingerTreeEmbed) toBST() *Treap {
////  if t.isEmpty() {
////     return &Treap{}
////  }
////  return &Treap{
////     size: t.size,
////     root: copyOf(t.root).
////        withL(reverseL(t.head.copyOrNil(), nil)).
////        withR(reverseR(t.tail.copyOrNil(), nil)),
////  }
////}
////
////
////
//////
//////
////func (t *TreapFingerTreeEmbed) isEmpty() bool {
////  return t.root == nil
////}
////
//////
//////
////func inorderL(p *Node, fn func(*Node)) {
////  if p == nil {
////     return
////  }
////  fn(p)
////  inorder(p.r, fn)
////  inorderL(p.l, fn)
////}
////
//////
////func inorderR(p *Node, fn func(*Node)) {
////  if p == nil {
////     return
////  }
////  inorderR(p.r, fn)
////  inorder(p.l, fn)
////  fn(p)
////}
//////
////func (t *TreapFingerTreeEmbed) inorder(fn func(*Node)) {
////  if t.isEmpty() {
////     return
////  }
////  inorderL(t.head, fn)
////  fn(t.root)
////  inorderR(t.tail, fn)
////}
////
//////
////func (t *TreapFingerTreeEmbed) Each(fn func(i int, s Value)) {
////  i := 0
////  t.inorder(func(n *Node) {
////     fn(i, n.s)
////     i++
////  })
////}
////
//////
////func (t *TreapFingerTreeEmbed) Array() []Value {
////  a := make([]Value, t.size)
////  i := 0
////  t.inorder(func(n *Node) {
////     a[i] = n.s
////     i++
////  })
////  return a
////}
////
//////
//////
////func (t *TreapFingerTreeEmbed) Validate() {
////  if t.isEmpty() {
////     assert(t.head == nil)
////     assert(t.tail == nil)
////     assert(t.size == 0)
////     return
////  }
////  t.verifyRelationPositionInvariant()
////  t.verifyRankHeapInvariant()
////  t.verifyRoot()
////}
////
////
////
//////
//////
////func (t *TreapFingerTreeEmbed) verifyRanksAlongLeftSpine(p *Node) {
////  for ; p != nil; p = p.l {
////     if p.hasR() {
////        assert(rankOf(p) >= rankOf(p.r))
////        p.r.verifyRankHeapInvariant()
////     }
////     if p.hasL() {
////        assert(rankOf(p) <= rankOf(p.l))
////     }
////  }
////}
////
//////
//////
////func (t *TreapFingerTreeEmbed) verifyRanksAlongRightSpine(p *Node) {
////  for ; p != nil; p = p.r {
////     if p.hasL() {
////        assert(rankOf(p) >= rankOf(p.l))
////        p.l.verifyRankHeapInvariant()
////     }
////     if p.hasR() {
////        assert(rankOf(p) <= rankOf(p.r))
////     }
////  }
////}
////
//////
////func (t *TreapFingerTreeEmbed) verifyRankHeapInvariant() {
////  t.verifyRanksAlongLeftSpine(t.head)
////  t.verifyRanksAlongRightSpine(t.tail)
////
////  if t.root != nil {
////     //
////     //
////     if t.head != nil {
////        for p := t.head; p.hasL(); p = p.l {
////           assert(rankOf(p) <= rankOf(t.root))
////        }
////     }
////     //
////     //
////     if t.tail != nil {
////        for p := t.tail; p.hasR(); p = p.r {
////           assert(rankOf(p) <= rankOf(t.root))
////        }
////     }
////  }
////}
////
//////
////func (t *TreapFingerTreeEmbed) verifyRoot() {
////  assert(t.root.l == nil)
////  assert(t.root.r == nil)
////  assert(t.root.sizeRL() == t.head.size())
////}
////
//////
////func (t *TreapFingerTreeEmbed) verifyRelationPositionInvariant() {
////  if t.root != nil {
////     assert(t.root.isR())
////     assert(t.head.size() == t.root.sizeRL())
////  }
////  for p := t.head; p != nil; p = p.l {
////     assert(p.isL())
////     p.verifyRelativePositionInvariant()
////  }
////  for p := t.tail; p != nil; p = p.r {
////     assert(p.isR())
////     p.verifyRelativePositionInvariant()
////  }
////}
////
////
////func (t *TreapFingerTreeEmbed) hashead bool {
////  return t.head != nil
////}
////
////func (t *TreapFingerTreeEmbed) hastail bool {
////  return t.tail != nil
////}
////
////func (t *TreapFingerTreeEmbed) hasRoot() bool {
////  return t.root != nil
////}
////
//////
//////
//////
//////
//////
//////
//////
//////
//////
////func (t *TreapFingerTreeEmbed) dissolvehead {
////  invariant(t.hashead)
////
////  //
////  t.root = copyOf(t.root).withRelativePosition(t.root.sizeRL())
////  //
////  p := t.head
////  g := t.head.l
////  r := t.head.r
////  //
////  //
////  if t.head.hasR() {
////     t.head = tree.reverseL(copyOf(r).toL(p.sizeLR()), g)
////  } else {
////     t.head = g
////  }
////  t.size--
////}
////
////
//////
//////
//////
//////
//////
//////
//////
//////
////
////
//////
//////
////func (t *TreapFingerTreeEmbed) Split(i Index) (List, List) {
////  assert(i <= t.size)
////  //
////  //
////  if i == t.size { return t, &TreapFingerTreeEmbed{} }
////  if i == 0      { return &TreapFingerTreeEmbed{}, t }
////
////  tmp := t.Clone().(*TreapFingerTreeEmbed) // get this of this all
////
////  if i <= t.root.sizeRL() {
////     return tmp, tmp.splitFromHead(i)
////  } else {
////     return tmp, tmp.splitFromTail(i)
////  }
////}
////
//////
//////
////func (t *TreapFingerTreeEmbed) Join(o List) List {
////
////
////  if o.Len() == 0 { return t }
////  if t.Len() == 0 { return o }
////  //
////  //
////  l := t.tail
////  r := o.(*TreapFingerTreeEmbed).head
////  s := 0
////
////  var p *Node
////  //
////  //
////  //
////  for r != nil && l != nil {
////     if rankOf(l) < rankOf(r) {
////        //
////        //
////        if p != nil && p.isL() {
////           p.toR(s)
////        }
////        //
////        //
////        s = s + l.w
////        l = prependR(copyOf(l), &p)
////
////     } else {
////        //
////        //
////        if p != nil && p.isR() {
////           p.toL(s)
////        }
////        //
////        //
////        s = s - r.w
////        r = prependL(copyOf(r), &p)
////     }
////  }
////  //
////  //
////  //
////  //
////  if r == nil {
////     //
////     //
////     r := copyOf(o.(*TreapFingerTreeEmbed).root).
////        withL(p).
////        withR(l).
////        withW(0).toR(s)
////     //
////     //
////     //
////     t = &TreapFingerTreeEmbed{
////        root: t.root,
////        head: t.head,
////        tail: o.(*TreapFingerTreeEmbed).tail,
////        size: t.size + o.(*TreapFingerTreeEmbed).size,
////     }
////     //
////     appendR(&t.tail, t.rotateUpR(r))
////
////  } else {
////     //
////     //
////     l := copyOf(t.root).
////        withL(r).
////        withR(p).
////        withW(0).toL(s)
////     //
////     //
////     //
////     t = &TreapFingerTreeEmbed{
////        root: copyOf(o.(*TreapFingerTreeEmbed).root).withW(t.size + o.(*TreapFingerTreeEmbed).root.w),
////        head: t.head,
////        tail: o.(*TreapFingerTreeEmbed).tail,
////        size: t.size + o.(*TreapFingerTreeEmbed).size,
////     }
////     //
////     appendL(&t.head, t.rotateUpL(l))
////  }
////  return t
////}
////
////// Can we use _dissolve here?
//////
////func (t *TreapFingerTreeEmbed) removeRoot() (s Value) {
////  invariant(t.root != nil)
////  s = t.root.s
////  //
////  if t.size == 1 {
////     t.size = 0
////     t.root = nil
////     return
////  }
////  //
////  //
////  if t.head == nil {
////     // TODO use tree construction here?
////     pr := followR(&t.tail)
////     t.head = tree.reverseL((*pr).l.copyOrNil(), nil)
////     t.root = copyOf(*pr).withL(nil).withR(nil)
////     t.size = t.size - 1
////     *pr = nil
////     return
////  }
////  //
////  //
////  if t.tail == nil {
////     pl := followL(&t.head)
////     t.tail = reverseR((*pl).r.copyOrNil(), nil)
////     t.root = copyOf(*pl).withL(nil).withR(nil).toR(t.root.sizeRL())
////     t.size = t.size - 1
////     *pl = nil
////     return
////  }
////  //
////  //
////  pl := followL(&t.head)
////  pr := followR(&t.tail)
////  //
////  //
////  sl := t.root.sizeRL()
////  sr := t.root.sizeRR(t.size)
////  //
////
////  //
////  l := (*pl).copy().toR(sl)
////  r := (*pr).copy().toL(sr)
////
////  var p *Node
////
////  //
////  // TODO is it not possible to already have a p on hand here?
////  (&Treap{}).joinRL(&p, l, r, sl + sr)
////  //
////  //
////  if rankOf(l) >= rankOf(r) {
////     *pr = reverseR(p.r, nil)
////     *pl = nil
////  } else {
////     *pl = tree.reverseL(p.l, nil)
////     *pr = nil
////  }
////  //
////  //
////  t.root = p.withL(nil).withR(nil)
////  t.size = t.size - 1
////  return
////}
////
////
////// TODO: if instead we use the same, does there exist a general algorithm?
//////       if so or otherwise, does it simplify anything?
//////
//////
//////
////func (t *TreapFingerTreeEmbed) splitFromHead(i int) (o *TreapFingerTreeEmbed) {
////  //
////  //
////  //
////  g := &t.head
////  d := i
////  for {
////     invariant(d > 0)
////     //
////     //
////     if d - ((*g).sizeLR() + 1) > 0 {
////        d -= (*g).sizeLR() + 1
////        g = &(shadow(g).l)
////        continue
////     }
////     //
////     //
////     //
////     //
////     p := shadow(g)
////     l, r := unzip(p.r, d - 1, p.sizeLR())
////     *g = nil
////     //
////     //
////     o = &TreapFingerTreeEmbed{
////        root: copyOf(t.root).withRelativePosition(t.root.w - i),
////        head: tree.reverseL(r, p.l),
////        tail: t.tail,
////        size: t.size - i,
////     }
////     //
////     //
////     *t = TreapFingerTreeEmbed{
////        root: p.withRelativePosition(i - (d - 1)).withL(nil).withR(nil),
////        tail: reverseR(l, nil),
////        head: t.head,
////        size: i,
////     }
////     return
////  }
////}
////
//////
//////
////func (t *TreapFingerTreeEmbed) splitFromTail(i int) (o *TreapFingerTreeEmbed) {
////  //
////  //
////  g := &t.tail
////  d := i - t.size
////  for {
////     invariant(d < 0)
////     //
////     //
////     if d + 1 + (*g).sizeRL() < 0 {
////        d += (*g).sizeRL() + 1
////        g = &shadow(g).r
////        continue
////     }
////     //
////     //
////     p := shadow(g)
////     l, r := unzip(p.l, d, p.sizeRL())
////     *g = nil
////
////     o = &TreapFingerTreeEmbed{
////        root: copyOf(p).withW(-d).withL(nil).withR(nil),
////        head: tree.reverseL(r, nil),
////        tail: t.tail,
////        size: t.size - i,
////     }
////     //
////     //
////     *t = TreapFingerTreeEmbed{
////        root: t.root,
////        tail: reverseR(l, p.r),
////        head: t.head,
////        size: i,
////     }
////     return
////  }
////}
////
////
////func (t Treap) DepthAlongTheSpines() (depths [2][]int) {
////  if t.root == nil {
////     return
////  }
////  traverseL(t.root.l, func(p *Node) {
////     depths[0] = append(depths[0], 1 + p.r.depth())
////  })
////  traverseR(t.root.r, func(p *Node) {
////     depths[1] = append(depths[1], 1 + p.l.depth())
////  })
////  return
////}
////
////
//////
////func (t TreapFingerTreeEmbed) WeightAlongTheSpinesLog2() (weights [2][]int) {
////  var l []*Node
////  var r []*Node
////
////  traverseL(t.head, func(p *Node) {
////     l = append(l, p)
////  })
////  traverseR(t.tail, func(p *Node) {
////     r = append(r, p)
////  })
////  for i := len(l) - 1; i >= 0; i-- {
////     weights[0] = append(weights[0], 1 + utility.Log2(l[i].sizeLR() + 1))
////  }
////  for i := len(r) - 1; i >= 0; i-- {
////     weights[1] = append(weights[1], 1 + utility.Log2(r[i].sizeRL() + 1))
////  }
////  return
////}
////
////func (t TreapFingerTreeEmbed) DepthAlongTheSpines() (depths [2][]int) {
////  var l []*Node
////  var r []*Node
////  traverseL(t.head, func(p *Node) {
////     l = append(l, p)
////  })
////  traverseR(t.tail, func(p *Node) {
////     r = append(r, p)
////  })
////  for i := len(l) - 1; i >= 0; i-- {
////     depths[0] = append(depths[0], 1 + l[i].r.depth())
////  }
////  for i := len(r) - 1; i >= 0; i-- {
////     depths[1] = append(depths[1], 1 + r[i].l.depth())
////  }
////  return
////}