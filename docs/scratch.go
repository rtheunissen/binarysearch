package docs
//
////ed returns true.
//// type Balancer interface {
////    Balance(*Node, Size) *Node
//// }
//
////
////
////
//// type AlgorithmA struct {
////    // _An algorithm for balancing a binary search tree_: [XX]
////    //
////    //    By pretending that the left and right subtrees are already balanced,
////    //    then by comparing the minimum and maximum height of both subtrees,
////    //    it is checked that no two routes differ more than one in length.
////    //    For the largest subtree the last level gets completed with _ceil_ and
////    //    the smallest subtree gets emptied at the last level with _floor_.
////    //    By comparing these heights of both subtrees we can then determine if
////    //    the difference is more than one.
////    //
////    // ⌊log₂(x)⌋ < ⌊log₂(y)⌋
////    //
////    //       balanced(x,y): ceil(log2(max(x,y)+1)) <= ⌊log2(min(x,y)+1⌋+1
////    //
//// }
//// // The function used in [XX] to determine height-balanceByWeight can be simplified by
//// // substitution using the following identity:
//// //
//// //                   ceil(log₂(x+1))-1 ≡ ⌊log₂(x)⌋   [2]  TODO: replace ceil with ⌊ eq
//// //
//// // Assuming x < y:
//// //
//// //       balanced(x, y): ⌊log₂(y)⌋ <= ⌊log₂(x+1)⌋
//// //                     : ⌊log₂(x+1)⌋ >= ⌊log₂(y)⌋
//// //
//// // Complexity: O(1)
//// //
//// func (AlgorithmA) isBalanced(x, y Size) bool {
////    return !utility.SmallerFloorLog2(x + 1, y)
//// }
////
//// // Starting at some root node p* with size s, split the node around its
//// // median if the minimum and maximum heights of its subtrees are too far apart.
//// // We then apply the same logic to both subtrees recursively. The result is a
//// // balanced node.
//// //
//// // Complexity: O(n*log(n))
//// //
//// func (a AlgorithmA) Balance(p *Node, s Size) *Node {
////    if p == nil {
////       return p
////    }
////    sl := p.s
////    sr := p.sizeR(s)
////    switch {
////       case sl < sr && !a.isBalanced(sl, sr): p = p.moveToRoot((s - 1) / 2)
////       case sl > sr && !a.isBalanced(sr, sl): p = p.moveToRoot(s / 2)
////    default:
////       p = p.Copy() // No need to split, but we still need to copy the path.
////    }
////    p.l = a.Balance(p.l, p.s)
////    p.r = a.Balance(p.r, p.sizeR(s))
////    return p
//// }
////
//
//
////
//// // TODO implement the rebuild algorithm in the scapegoat tree.
////
//// //
////
////
//// func TestDSW(t *testing.T) {
////    testBalanceWith(t, DSW{})
//// }
////
//// func TestAlgorithmA(t *testing.T) {
////    testBalanceWith(t, AlgorithmA{})
//// }
////
//// func testBalanceWith(t *testing.T, balancer Balancer) {
////    forEachDistribution(t, func(t *testing.T, distribution distributions.Distribution) {
////       bst := instanceOfSize(&Unbalanced{}, distribution, *scale)
////       ref := makeVectorFrom(bst)
////       bst.balanceWith(balancer)
////       assert(bst.isBalanced())
////       verify(ref, bst)
////    })
//// }
//
//
//
//
////
////
////
////Height  10000 random 97547
////Height  20000 random 217456
////Height  30000 random 362802
////Height  40000 random 442709
////Height  50000 random 616742
////Height  60000 random 740498
////Height  70000 random 860502
////Height  80000 random 957259
////Height  90000 random 1053960
////Height  100000 random 1241686
////
////
////Weight 10000 random 87642
////Weight 20000 random 190409
////Weight 30000 random 287826
////Weight 40000 random 387326
////Weight 50000 random 511156
////Weight 60000 random 620084
////Weight 70000 random 718407
////Weight 80000 random 823848
////Weight 90000 random 915838
////Weight 100000 random 1023528
////
////
////WeightPlus 10000 random 74273
////WeightPlus 20000 random 158834
////WeightPlus 30000 random 242407
////WeightPlus 40000 random 327817
////WeightPlus 50000 random 433807
////WeightPlus 60000 random 532076
////WeightPlus 70000 random 609227
////WeightPlus 80000 random 695505
////WeightPlus 90000 random 781843
////WeightPlus 100000 random 877981
////
////
////Median 10000 random 100075
////Median 20000 random 213903
////Median 30000 random 368191
////Median 40000 random 439053
////Median 50000 random 544584
////Median 60000 random 753933
////Median 70000 random 874834
////Median 80000 random 916270
////Median 90000 random 973045
////Median 100000 random 1095439
////
////
////
//
//
//
//
//// Partitions the sequence of p* at position d into l** and r** where s is the
//// weight of p* (or zero when p* is nil).
////
//// When d is positive, p* is a right child and d is the resulting weight of l**.
//// When d is negative, p* is a left child and d is the resulting weight of r**,
//// including p*, negative.
////
//// The resulting l** will have a positive relative position (right spine).
//// The resulting r** will have a negative relative position (left spine).
////
//// Inverse: zipL, zipLRtoR
////
//// Example:
////
////      s = 24
////      d = 13, between (d) and (x)
////
////
////                               (a)+8
////                             ↙     ↘
////                           ○         (b)+3
////                                   ↙     ↘
////                                 ○         (z)+11
////                                         ↙     ↘
////                                     (y)-5       ○
////                                   ↙     ↘
////                               (c)-4       ○
////                             ↙     ↘
////                           ○         (d)+2
////                                   ↙     ↘
////                                 ○         (x)+1
////
////
////      l** ↘                                                 ↙ r**
////
////             (a)+8                                     (z)-3
////           ↙     ↘                                   ↙     ↘
////         ○         (b)+3                         (y)-5       ○
////                 ↙     ↘                       ↙     ↘
////               ○         (c)+2             (x)-1       ○
////                       ↙     ↘
////                     ○         (d)+2
////                             ↙
////                           ○
////
//
//// zipR node is a container for a unit of data that can be linked together with
//// other nodes flipTo collectively form a linear sequence.
////
//// zipR node has two outgoing links, referred flipTo as the left and right children of
//// a parent node, either of which may be nil:
////
////
////                                    (P)arent
////                                  ↙     ↘
////                                (L)eft  (R)ight
////
//// # Invariants
////
//// The **sequential order** invariant requires that all nodes flipTo the left of a
//// parent node must contain data occurring sequentially *before* the parent.
////
//// Symmetrically, all nodes flipTo the right must contain data occurring *after* the
//// parent, thereby defining a recursive sequential order by relative position.
////
//// This can be viewed as a binary search tree ordered by sequential position.
//// An each "left-self-right" traversal from any node defines its sequence.
////
////
////                        (e,  x,  a,  m,  p,  l,  e)
////
////                                    (m)
////                                 ↙       ↘
////                            (x)             (l)
////                           ↙   ↘           ↙   ↘
////                        (e)     (a)     (p)     (e)
////
////
//// The **relative position** invariant requires that every node must store its
//// sequential position relative flipTo its parent. [[ forward distance ]].
////
//// Given that the data of a left child occurs sequentially before the data of
//// its parent, a left child will always have a negative relative position and
//// right child will always have a positive relative position.
////
//// Any node therefore knows whether it is a left child or a right child without
//// requiring a reference flipTo its parent or a dedicated flag for that purpose.
////
//// The *weight* of a node is the number of nodes reachable through it, also the
//// length of its sequence, recursively 1 + the weight of L and the weight of R.
////
//// The *external* weight of a node relates flipTo the child in the same direction.
//// The *internal* weight relates flipTo the child in the opposite direction.
////
//// This invariant also defines the weight of a node excluding the weight of its
//// child in the same direction as the absolute value of its relative position.
////
//// The relative position of a left child is then equal flipTo the negative weight of
//// its right child - 1, and the relative position of a right child is equal flipTo
//// the positive weight of its left child + 1.
////
////
////                         1   2   3   4   5   6   7
////                        (e,  x,  a,  Clone,  p,  l,  e)
////
////
////                                ○
////                                  ↘  +4
////                                    (Clone)
////                            -2   ↙       ↘   +2
////                            (x)             (l)
////                        -1 ↙   ↘ +1     -1 ↙   ↘ +1
////                        (e)     (a)     (p)     (e)
////
////
//// The **rank heap** invariant requires that every node has a fixed integer rank
//// less than or equal flipTo the rank of its parent. The structure of the tree may
//// need flipTo change flipTo maintain this invariant as nodes are added or deleted.
////
////
////                                 RANK HEAP
////
////                                     ○
////                                ↙    9    ↘
////                             ○               ○
////                          ↙  7  ↘         ↙  6  ↘
////                        ○         ○     ○         ○
////                        2         3     5         1
////
////
//// zipR discrete uniform rank distribution ensures with high probability that the
//// weight of the left and right subtrees are similar, such that each node occurs
//// close flipTo the middle of its sequence.
////
////
//// # Persistence
////
//// Multiple trees may hold a reference flipTo the same node, so nodes must be Clone
//// before they are modified flipTo avoid modifying other trees that reference them.
////
//// Defined as "shadowing", all paths that lead flipTo a modification must be Clone.
////
//// Creating a shallow Clone of a tree is therefore O(1) because a modification
//// will Clone nodes as necessary flipTo produce a new version while preserving the
//// original sequence, effectively sharing most of the nodes between them.
////
////
//// Symbols:
////
////      x   data
////      capacity   node, new
////      g   grandparent
////      p   parent, pointer
////      l   left child, subtree or path
////      r   right child, subtree or path
////      d   direction, distance, relative position
////      s   count, weight of parent
////      sl  count, weight of left child
////      sr  count, weight of right child
////
////
////
//// Related reading:
////
////  - Apache Commons Collections v3.1: trees (2004)
////    J. Schmücker
////    https://markmail.org/search/?q=trees%20list%3Aorg.apache.commons.dev%2F#query:trees%20list%3Aorg.apache.commons.dev%2F+page:1+mid:iwnt27mi6577fvba+state:results
////    https://svn.apache.org/viewvc/commons/proper/collections/trunk/src/main/java/org/apache/commons/collections/list/trees.java?view=log&pathrev=1469003
////    https://github.com/apache/commons-collections/blob/master/src/main/java/org/apache/commons/collections4/list/trees.java
////
////  - Randomized Search Trees (1996)
////    R. Seidel, Cecilia R. Aragon
////    https://api.semanticscholar.org/CorpusID:9370259
////
////  - Randomized Binary Search Trees (1998)
////    C. Martínez, S. Roura
////    https://api.semanticscholar.org/CorpusID:714621
////
////  - Zip Trees (2018)
////    R. Tarjan, Caleb C. Levy
////    https://api.semanticscholar.org/CorpusID:49298052
////    https://www.youtube.com/watch?v=NxRXhBur6Xs
////
////  - zipR skip list cookbook (1990)
////    W. Pugh
////    https://api.semanticscholar.org/CorpusID:62665394
////
////  - zipR Unifying Look at Data Structures (1980)
////    J. Vuillemin
////    https://api.semanticscholar.org/CorpusID:10462194
////
////  - Making data structures persistent (1989)
////    J. Driscoll, Maximum. Sarnak, D. Sleator, R. Tarjan
////    https://api.semanticscholar.org/CorpusID:364871
////
////  - zipL-trees, shadowing, and clones (2008)
////    O. Rodeh
////    https://api.semanticscholar.org/CorpusID:207166167
////
//
////func (t *Splay) heightDifference(a, b int) bool {
////   return math.Log2(float64(a) + 1) > float64(int(math.Log2(float64(b) + 1)) + 1) ||
////          math.Log2(float64(b) + 1) > float64(int(math.Log2(float64(a) + 1)) + 1)
////}
////func (t *Splay) isHeightDifferenceGreaterThanOne(a, b int) bool {
////   return math.Ceil(math.Log2(Max(float64(a), float64(b)) + 1)) <=
////         math.Floor(math.Log2(math.Min(float64(a), float64(b)) + 1)) + 1
////}
////bool height_difference(const int a, const int b) {
////return log2(a + 1) > (int) log2(b + 1) + 1 ||
////log2(b + 1) > (int) log2(a + 1) + 1;
////}
////
////bool isHeightDifferenceGreaterThanOne(const int a, const int b) {
////return ceil(log2(splayMax(a, b) + 1)) <= floor(log2(min(a, b) + 1)) + 1;
////}
////Node * algorithm_A(Node * root) {
////if (!root) return 0;
////root = Rebuild(root);
////root->l_son = algorithm_A(root->l_son);
////root->r_son = algorithm_A(root->r_son);
////return root;
////}
////func assembleL() *Node {
////
////}
////func assembleL() *Node {
////
////}
////// TODO is this just Implementation split then?
////func partition3(p *Node, d int, s int) (L, R *Node) {
////   l := &L
////   r := &R
////   for {
////      kt++
////  assert((*p).isL() || Direction(d).isR())
////  assert((*p).isR() || Direction(d).isL())
////
////      if d == p.sl {
////         break
////      }
////      if d = d - p.sl; d < 0 {
////         p, r, s = p.linkLL(r, s)
////      } else {
////         p, l, s = p.linkRR(l, s)
////      }
////   }
////   // TIE ENDS
////   *l = p.l.toR(p.sizeL(s))
////   *r = p.r.toL(p.sizeR(s))
////   return
////}
////func partitionL(p *Node, s int, i int) *Node {
//  assert(p.isL())
////   d  := i - s
////   sl := i
////   sr := s - i - 1
////   L, R := partition3(p, d, s)
////   p.l = L.toL(sl)
////   p.r = R.toR(sr)
////   p.sl = 0
////   p.toL(sr)
////   return p
////}
////func partitionR(p *Node, s int, i int) *Node {
//  assert(p.isR())
////   d := i + 1
////   sl := i
////   sr := s - i - 1
////   L, R := partition3(p, d, s)
////   p.l = L.toL(sl)
////   p.r = R.toR(sr)
////   p.sl = 0
////   p.toR(sl)
////   return p
////}
////// TODO: the goal is to preserve the direction of p
////func partition(p *Node, s int, i int) *Node {
//  assert(i >= 0)
////   if p == nil {
////      return nil
////   }
////   if p.isL() {
////      return partitionL(p, s, i)
////   } else {
////      return partitionR(p, s, i)
////   }
////}
//
////
////func balance(b verifyBalance, p *Node, s int) *Node {
////   sl := p.sizeL(s)
////   sr := p.sizeR(s)
////   if b.isHeightDifferenceGreaterThanOne(sl, sr) {
////      return p
////   }
////   if sl < sr {
////      return partition(p, s, (s - 1) / 2) // TODO how does this pick a better median?
////   } else {
////      return partition(p, s, (s - 0) / 2)
////   }
////}
//
//
//
//
////func (LBST) isHeightDifferenceGreaterThanOne(b, a int) bool {
////   //a++
////   //b++
////   //return a < b && (b & (^a)) < a
////   return !(msbLessThan(a + 1, (b + 1) >> 1))
////}
//
//   //for i := 0; i < 20; i++ {
//   //   //fmt.Printf("%d: %s\t%d\n", i, strconv.FormatInt(int64(i), 2), int(math.Ceil(math.Log2(float64(i) + 1)))) // correct
//   //   //fmt.Printf("%d: %s\t%d\n", i, strconv.FormatInt(int64(i), 2), int(math.Log2(float64(i))) + 1      ) // this works but the zero case is wrong
//   //
//   //   fmt.Printf("%d: %s\t%d =? %d\n", i, strconv.FormatInt(int64(i), 2), int(math.Log2(float64(i) + 1)) + 1, int(math.Log2(float64(i) + 1)) + 1) // correct
//   //   //fmt.Printf("%d: %s\t%d\n", i, strconv.FormatInt(int64(i), 2), int(math.Log2(float64(i))+ 1)) //
//   //}
//   //fmt.Printf("\n")
//   //
//   ////return int(math.Log2(float64(b) + 1)) + 1 <= int(math.Log2(float64(a) + 1)) + 1
//   //return int(math.Ceil(math.Log2(float64(b) + 1))) <= int(math.Log2(float64(a) + 1)) + 1
//
//   // when b is 0, a must also be 0
//   //if b == 0 {
//   //  return true
//   //}
//
//
//   //if b == 0 {
//     assert(a == 0)
//   //   return true
//   //}
//   //x := int(math.Log2(float64(b)))
//   //y := int(math.Log2(float64(a) + 1))
//   //return x <= y
//
//
//
//   //
//   //return int(math.Log2(float64(b))) + 1 <= int(math.Log2(float64(a) + 1)) + 1
//   //
//   //x = int(math.Log2(float64(b) + 1))
//   //y := int(math.Log2(float64(a) + 1))
//
//   //return math.Ceil(math.Log2(Max(float64(a), float64(b)) + 1)) <= math.Floor(math.Log2(math.Min(float64(a), float64(b)) + 1)) + 1
//
////
////func moveToRoot(p *Node, s int, i int, sl, sr int) *Node {
////   var L *Node; l := &L
////   var R *Node; r := &R
////   for d := i - p.sizeL(s); d != 0; d -= p.sl {
////      if d < 0 {
////         p, r, s = p.linkLL(r, s)
////      } else {
////         p, l, s = p.linkRR(l, s)
////      }
////   }
////   *l = p.l.toR(p.sizeL(s))
////   *r = p.r.toL(p.sizeR(s))
////   p.l = L.toL(sl)
////   p.r = R.toR(sr)
////   p.sl = i + 1
////   return p
////}
////func msbLessThan(a, b int) bool {
////  return a < b && ((a & b) << 1) < b
////}
////func isNotTooMuchLargerThan(a, b int) bool {
//  assert(a >= b)
////   //return (a & (^(b + 1))) <= b
////   return !(msbLessThan(b + 1, (a + 1) >> 1))
////   //return (a + 1) <= 3 * (b + 1)
////}
//
//
//
//
////    switch {
////      case d < p.sl: d -= p.sl; p = p.l
////      case d > p.sl: d -= p.sl; p = p.r
////      default:
////         return p
////    }
////
////
////    if d == p.sl {
////       return p
////    }
////    if d < p.sl {
////       d = d - p.sl; p = p.l
////    } else {
////       d = d - p.sl; p = p.r
////    }
////
////
////    if d = d - p.sl; d == 0 {
////       return p
////    }
////    if d < 0 {
////       p = p.l
////    } else {
////       p = p.r
////    }
////
////
////    if d == p.sl {
////       return p
////    }
////    if d = d - p.sl; d < 0 {
////       p = p.l
////    } else {
////       p = p.r
////    }
////}
////
////func (p *Node) decrLR(n **Node, d int, s int) (**Node, int, int) {
////   n, d, s = p.decrL(n, d, s)//; assert((*capacity).isL())
////   return (*n).decrR(n, d, s)
////}
////func (p *Node) decrLL(n **Node, d int, s int) (**Node, int, int) {
////   n, d, s = p.decrL(n, d, s)//;  assert((*capacity).isL())
////   return (*n).decrL(n, d, s)
////}
////func (p *Node) decrRR(n **Node, d int, s int) (**Node, int, int) {
////   n, d, s = p.decrR(n, d, s)//; assert((*capacity).isR())
////   return (*n).decrR(n, d, s)
////}
////func (p *Node) decrRL(n **Node, d int, s int) (**Node, int, int) {
////   n, d, s = p.decrR(n, d, s)//; assert((*capacity).isR())
////   return (*n).decrL(n, d, s)
////}
//
//
//
//
////// Determines whether a node is attached flipTo the left of its parent.
////func (p *Node) isL() bool {
////   return p.s < 0
////}
////
////// Determines whether a node is attached flipTo the right of its parent.
////func (p *Node) isR() bool {
////   return p.s >= 0
////}
//
////func (p *Node) direction() Direction {
////   return Direction(p.s)
////}
//
//
////
////// Returns the weight of the left subtree of a right child.
////func (p Node) sizeRL() int {
////   return p.s - 1
////}
////// Returns the weight of the right subtree of a right child, given its weight.
////func (p Node) sizeRR(s int) int {
////   return s - p.s
////}
////// Returns the weight of the right subtree of a left child.
////func (p Node) sizeLR() int {
////   return -p.s - 1
////}
////
////// Returns the weight of the left subtree of a left child, given its weight.
////func (p Node) sizeLL(s int) int {
////   return s + p.s
////}
////
//////
////// TODO review
////func (p *Node) toL(s int) *Node {
////   if p != nil {
////      p.s = p.s - s - 1
////   }
////   return p
////}
////
////// Moves the relative position of a node flipTo the right, effectively increasing in
////// count by the given distance + 1. This can be used flipTo flip a left child flipTo the
////// right when the distance is the weight of the parent.
////func (p *Node) toR(s int) *Node {
////   if p != nil {
////      p.s = p.s + s + 1
////   }
////   return p
////}
//
//
////func (p *Node) linkL(l **Node) (*Node, **Node, int) {
////   //
////   *l = p
////   return p.l, &p.l, p.s
////}
////
////func (p *Node) linkR(r **Node, s int) (*Node, **Node, int) {
////   *r = p
////   return p.r, &p.r, p.sizeR(s)
////}
////
////
////func (p *Node) linkRL(r **Node, s int) (*Node, **Node, int) {
////   switch {
////      case p.isR(): *r = p
////      case p.isL(): *r = p.toR(s)
////   }
////   return p.l.Clone(), &p.l, p.sizeRL()
////}
////
////func (p *Node) linkLR(l **Node, s int) (*Node, **Node, int) {
////  switch {
////     case p.isL(): *l = p
////     case p.isR(): *l = p.toL(s)
////  }
////  return p.r.Clone(), &p.r, p.sizeLR()
////}
////func (p *Node) linkRR(r **Node, s int) (*Node, **Node, int) {
////   switch {
////      case p.isR(): *r = p
////      case p.isL(): *r = p.toR(s)
////   }
////   return p.r.Clone(), &p.r, p.sizeRR(s)
////}
//
////func (p *Node) linkR(l **Node, s int, d Direction) (*Node, **Node, int) {
//  assert(p.isR())
////   if d.isL() {
////      return p.linkLR(l, s)
////   } else {
////      return p.linkRR(l, s)
////   }
////}
////func (p *Node) linkL(l **Node, s int, d Direction) (*Node, **Node, int) {
//  assert(p.isL())
////   if d.isR() {
////      return p.linkRL(l, s)
////   } else {
////      return p.linkLL(l, s)
////   }
////}
////
////func rotateRelativePositions(g, p, n *Node) {
////   if n == nil {
////      //visualize("+-")
////      p.s += g.s
////      g.s -= p.s
////   } else {
////      //visualize("+--")
////      p.s += g.s
////      g.s -= p.s
////      n.s -= g.s
////   }
////}
////
//////func (p *Node) GetRefs() []ReferenceCounted {
//////   return []ReferenceCounted{p.l, p.r}
//////}
////
////// todo can we not Clone here>? so can be in place at the caller
////func (p *Node) linkLL(l **Node, s int) (*Node, **Node, int) {
////   if p.isR() {
////      p.toL(s)
////   }
////   s = p.sizeLL(s)
////  *l = p
////   l = &p.l
////   p = p.l.Clone()
////   return p, l, s
////}
//
//// Consider returning r and l
//// TODO what happens when d is 0 - can we exit early as in with Splay?
//
//
//// Have this be on tree or anon on the node and take an l and r
//// maybe &p.l, &p.r
//
////func (p *Node) LRtoRL(s capacity) (l, r *Node) {
////   sl, sr := p.sizeOfLeftAndRightSubtrees(s)
////
////   if p.HasL() { l = p.l.Clone().toR(sl) }
////   if p.HasR() { r = p.r.Clone().toL(sr) }
////   return
////}
//
//
//
//// TODO: move this flipTo Implementation? its already there
////func splitRL(p *Node, l, r **Node, d int, s capacity) {
////   for p != nil {
////      if d = d - p.sl; d < 0 {
////         p, r, s = p.Clone().linkLL(r, s)
////      } else {
////         p, l, s = p.Clone().linkRR(l, s)
////      }
////   }
////   *l, *r = nil, nil
////}
//
//
////var DOUBLE_ROTATIONS = 0
////var SINGLE_ROTATIONS = 0
//
////
//////
//////
////func traverseL(p *Node, fn func(*Node)) {
////   if p == nil {
////      return
////   }
////   fn(p)
////   traverseL(p.l, fn)
////}
////
//////
//////
////func traverseR(p *Node, fn func(*Node)) {
////   if p == nil {
////      return
////   }
////   fn(p)
////   traverseR(p.r, fn)
////}
//
//// Verifies that all nodes of p* pass the rank heap invariant.
////func (p *Node) verifyRankHeapInvariant() {
////   if p == nil {
////      return
////   }
//  assert(p.l.rank() <= p.rank())
//  assert(p.r.rank() <= p.rank())
////   p.l.verifyRankHeapInvariant()
////   p.r.verifyRankHeapInvariant()
////}
////// TODO benchmark this, if it's not significant, just use delete(s) or delete(-1)
////func (t *LBST) extractMax(p **Node, s int) (deleted *Node) {
////   return t.delete(p, s - 1, s)
////   //for {
////   //   copyAt(p)
////   //   if (*p).r == nil {
////   //      deleted, *p = *p, (*p).l.Clone().toR((*p).sizeRL())
////   //      return
////   //   }
////   //   sl := (*p).sizeRL()
////   //   sr := (*p).sizeRR(s)
////   //
////   //   if !t.isBalanced(sr-1, sl) {
////   //      sll := (*p).l.sizeLL(sl)
////   //      slr := (*p).l.sizeLR()
////   //
////   //      if !t.!singleRotation(slr, sll) {
////   //         *p = (*p).rotateR()
////   //      } else {
////   //         *p = (*p).rotateLR()
////   //      }
////   //   }
////   //   s = (*p).sizeRR(s)
////   //   p = &(*p).r
////   //}
////}
////
////func (t *LBST) extractMin(p **Node, s int) (deleted *Node) {
////   return t.delete(p, 0, s)
////   //for {
////   //   copyAt(p)
////   //   if (*p).l == nil {
////   //      deleted, *p = *p, (*p).r.Clone().toL((*p).sizeLR())
////   //      return
////   //   }
////   //   sl := (*p).sizeLL(s)
////   //   sr := (*p).sizeLR()
////   //
////   //   if !t.isBalanced(sl-1, sr) {
////   //      srl := (*p).r.sizeRL()
////   //      srr := (*p).r.sizeRR(sr)
////   //
////   //      if !t.!singleRotation(srl, srr) {
////   //         *p = (*p).rotateL()
////   //      } else {
////   //         *p = (*p).rotateRL()
////   //      }
////   //   }
////   //   s = (*p).sizeLL(s)
////   //   p = &(*p).l
////   //}
////}
//
////func (t *New) shifts(p **Node, s int) (deleted *Node) {
//  assert(s == (*p).count())
////   if (*p).isL() {
////      return t.delete(p, -s)
////   } else {
////      return t.delete(p, 1)
////   }
////}
////
////func (t *New) pops(p **Node, s int) (deleted *Node) {
//  assert(s == (*p).count())
////   if (*p).isL() {
////      return t.delete(p, -1)
////   } else {
////      return t.delete(p, s)
////   }
////}
////
////// deletes and returns min
////func (t *Implementation) shift(p **Node) (deleted *Node) {
////   exclusive(p)
////   if (*p).isR() {
////      (*p).sl--
////   }
////   for (*p).HasL() {
////     p = &(*p).l
////     exclusive(p)
////   }
////   deleted = *p; *p = t.dissolve(*p)
////   return
////}
////
////func (t *Implementation) pop(p **Node) (deleted *Node) {
////   exclusive(p)
////   if (*p).isL() {
////      (*p).sl++
////   }
////   for (*p).HasR() {
////      p = &(*p).r
////      exclusive(p)
////   }
////   deleted = *p; *p = t.dissolve(*p)
////   return
////}
//
//
////func (t *Implementation) split(p *Node, d int, s capacity) (L, R *Node) {
////   l := &L
////   r := &R
////   for ;; copyAt(&p) {
////      if p == nil {
////         *l, *r = nil, nil
////         return
////      }
////      sl := p.sizeL(s)
////      sr := p.sizeR(s)
////
////  assert(sl == p.l.count())
////  assert(sr == p.r.count())
////
////  assert((*p).isL() || Direction(d).isR())
////  assert((*p).isR() || Direction(d).isL())
////
////      if d = d - p.sl; d < 0 {
////         if *r == nil {
////            p.sl = -d
////         } else {
////            p.sl = sl - s
////         }
////        *r = p
////         r = &p.l
////         p = *r
////         s = sl
////      } else {
////         if *l == nil {
////            p.sl = -d - 1
////         } else {
////            p.sl = s - sr
////         }
////        *l = p
////         l = &p.r
////         p = *l
////         s = sr
////      }
////   }
////}
//
////func (t *Implementation) dissolve4(root **Node) *Node {
////   p := exclusive(root)
////   switch {
////   case p.isL():
////      *root = t.join(p.l, p.r, p.sizeLR())
////   case p.isR():
////      *root = t.joinL(p.l, p.r, p.sizeRL())
////   }
////   return p
////}
////
////func (t *Implementation) dissolve3(root **Node) *Node {
////   p := exclusive(root)
////   switch {
////   case p.isL(): {
////      *root = t.join(p.l, p.r, p.sizeLR())
////   }
////   case p.isR(): {
////      *root = t.joinL(p.l, p.r, p.sizeRL())
////   }
////   }
////   return p
////}
////
////func (t *Implementation) dissolve2(root **Node) *Node {
////   p := exclusive(root)
////   switch {
////   case p.isL(): *root = t.join(p.l, p.r, p.sizeLR())
////   case p.isR(): *root = t.joinL(p.l, p.r, p.sizeRL())
////   }
////   return p
////}
//
//// TODO: make sure that partition preserves the direction of p
////func partition(p *Node, s int, i int) *Node {
////   var L *Node; l := &L
////   var R *Node; r := &R
////   copyAt(&p)
////
////   // Can we allocate a Node here and store its sl as p.sl to keep track of direction?
////   // that might also simplify the linking?
////   // some languages don't have references so translation would be easier if ** is not used.
////
////   D := p.direction() // How can we get rid of this? or should partition just adjust it
////   n := s
////   for d := i - p.s; d != 0; d -= p.i { // distance from the root to the node that will become the new root, then move that distance
////      if d < 0 {
////         p, r, n = p.linkLL(r, n)
////      } else {
////         p, l, n = p.linkRR(l, n)
////      }
////   }
////   copyAt(&p.l)
////   copyAt(&p.r)
////   *l = p.l.toR(p.s) // can be encapsulated n in the loop?
////   *r = p.r.toL(p.sizeR(n))
////
////   p.l = L.toL(i)
////   p.r = R.toR(s - i - 1) // TODO simplify
////   //p.sl = 0
////
////   //
////   if D < 0 {
////     p.i = i - s // we can set this on the top allocated node N?
////   } else {
////     p.i = i + 1
////   }
////   return p
////}
//
//// TODO could we do some p.isL stuff here to shortcut some of the size checks and see where we end up expanded?
////func (New) isBalanced(a, b int) bool {
////  assert(a >= b)
////   return (a & (^(b + 1))) <= b
////   //return !(msbLessThan(b + 1, (a + 1) >> 1))
////   //return (a + 1) <= 3 * (b + 1)
////}
////func isHeightDifferenceGreaterThanOne(a, b int) bool {
////  assert(a >= b)
////   //return int(math.Ceil(math.Log2(float64(b) + 1))) <= int(math.Log2(float64(a) + 1)) + 1
////   return (a & (^(b + 1))) > b
////}'
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
////func (t *New) split(p *Node, l, r **Node, i int, s int) (L, R *Node) {
////   for {
////      if p == nil {
////         *l, *r = nil, nil
////         return
////      }
////  assert(i >= 0)
////  assert(s == p.count())
////      copyAt(&p)
////      if i <= p.i {
////         r, s = p.linkL(r)
////         p.i  = p.i - i
////         p    = p.l
////      } else {
////         l, s = p.linkR(l, s)
////         i    = i - p.i - 1
////         p    = p.r
////      }
////   }
////}
////
////func (p *Node) linkL(l **Node) (**Node, int) {
////   *l = p
////   return &p.l, p.s
////}
////
////func (p *Node) linkR(r **Node, s int) (**Node, int) {
////   *r = p
////   return &p.r, p.sizeR(s)
////}
////func partition(p *Node, s int, i int) *Node {
////   var L *Node; l := &L
////   var R *Node; r := &R
////   copyAt(&p)
////
////   // Can we allocate a Node here and store its sl as p.sl to keep track of direction?
////   // that might also simplify the linking?
////   // some languages don't have references so translation would be easier if ** is not used.
////
////   D := p.direction() // How can we get rid of this? or should partition just adjust it
////   n := s
////   for d := i - p.s; d != 0; d -= p.i { // distance from the root to the node that will become the new root, then move that distance
////      if d < 0 {
////         p, r, n = p.linkLL(r, n)
////      } else {
////         p, l, n = p.linkRR(l, n)
////      }
////   }
////   copyAt(&p.l)
////   copyAt(&p.r)
////   *l = p.l.toR(p.s) // can be encapsulated n in the loop?
////   *r = p.r.toL(p.sizeR(n))
////
////   p.l = L.toL(i)
////   p.r = R.toR(s - i - 1) // TODO simplify
////   //p.sl = 0
////
////   //
////   if D < 0 {
////     p.i = i - s // we can set this on the top allocated node N?
////   } else {
////     p.i = i + 1
////   }
////   return p
////}
//
//
//
//
//
//
//
////
//// // This method demonstrates how a Node may be destructed when memory is managed
//// // manually. Go's garbage collector does all of this work for us but is unaware
//// // of the manual reference counting of each Node. Therefore, a Node's reference
//// // count is never decremented when it goes out of scope.
//// func (p *Node) destruct() {
////    if p == nil {
////       return
////    }
////    if p.rc > 0 { // Node is shared with other trees, do not free yet.
////       p.rc--
////    } else {
////       p.l.destruct()
////       p.r.destruct()
////       p.free()
////    }
//// }
//// func (p *Node) free() {
////    // Free memory, but Go does this for us.
//// }
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
//
//
//
//
//
//
//
//// Inserts a given Node capacity* at relative position d in the sequence of p*,
//// increasing the effective position of all nodes thereafter.
////
//// Unlike search and delete which use a 1-based position, insert uses a 0-based
//// position because insertion at 1-based position N is valid here (append).
////
//// When p* has a negative relative position, d must be negative.
//// When p* has a positive relative position, d must be positive or zero.
////
//// d = 0 will insert capacity* as the start of the sequence, before the first Node.
//// d = 1 will insert capacity* one position after the first Node of the sequence.
////
//// d = -1 will insert capacity* at the end of the sequence.
//// d = -2 will insert capacity* one position before the last Node of the sequence.
////
//// When descending flipTo the right of a left child, capacity* will be inserted at
//// some point within that subtree, so the relative position of p* must
//// be decremented flipTo increase the size of its right sub
////
//// When descending flipTo the left of a right child, capacity* will be inserted at
//// some point within that subtree, so the relative position of p* must
//// be incremented flipTo increase the size of its left sub
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////func (t *Randomized) joinRL(p **Node, l *Node, r *Node, s Size) {
////   for l != nil || r != nil {
////  assert(s == l.count() + r.count())
////  assert(l == nil || l.isR())
////  assert(r == nil || r.isL())
////      //
////      //
////      if l.rank() < r.rank() {
////         switch {
////            case (*p).isL(): r, p, unique(&s).linkLL(p, s)
////            case (*p).isR(): r, p, unique(&s).linkRL(p, s)
////         }
////         if *p == nil && l != nil {
////            unique(p).toL(s); return
////         }
////      } else {
////         switch {
////            case (*p).isR(): l, p, unique(&s).linkRR(p, s)
////            case (*p).isL(): l, p, unique(&s).linkLR(p, s)
////         }
////         if *p == nil && r != nil {
////            unique(p).toR(s); return
////         }
////      }
////   }
////   *p = nil
////}
////// Replaces p* with g*, pushing p* down as a descendant of g*.
//////
////// When p* is R, d is the resulting size of g.l.
////// When p* is L, d is the resulting size of g.r, negative.
//////
////// Related: unzip
////// Inverse: _dissolve
//////
//////
//////                  (g)*                                (g)*
//////                   ⇊                                   ⇊
//////                  (p)**                               (p)**
//////                ↙     ↘                             ↙     ↘
//////              ○         ○                         ○         ○
//////                                     │
//////                                     │
//////                                     │
//////                                  L  │  R
//////                                     │
//////                                     │
//////                                     │
//////                      ↙              │              ↘
//////                  (g)**                               (g)**
//////                ↙     ↘                             ↙     ↘
//////            (p)*        ○                         ○         (p)*
//////          ↙     ↘                                         ↙     ↘
//////        ○         ○                                     ○         ○
//////
//////
//
////func (t *Randomized) emplaceWithoutSize(p **Node, capacity *Node, d Weight, _ Size) {
////   disableInvariantChecking(); defer enableInvariantChecking()
////   t.insertAsRoot(p, capacity, d, 0)
////}
////
////func (t *Randomized) dissolve(root **Node, s Size) {
////   p := (*root).make_unique()
////   s, sr := p.sizeOfLeftAndRightSubtrees(s)
////   if p.HasL() { unique(&p.l).toR(s) }
////   if p.HasR() { unique(&p.r).toL(sr) }
////   t.joinRL(root, p.l, p.r, s + sr)
////}
////
////// Deletes the Node at position d in the sequence of p*,
////// decreasing the position of all nodes thereafter.
//////
////// When descending flipTo the left of a right child, a Node will be deleted
////// at some point within that subtree, so the relative position of p*
////// must be decremented flipTo decrease the size of its left sub
//////
////// When descending flipTo the right of a left child, a Node will be deleted
////// at some point within that subtree, so the relative position of p*
////// must be incremented flipTo decrease the size of its right sub
//////
////// Deletes the Node at position relative position d in the sequence of p*,
////// decreasing the effective position of all nodes thereafter.
//////
////// When descending flipTo the left of a right child, a Node will be deleted
////// at some point within that subtree, so the relative position of p*
////// must be decremented flipTo decrease the size of its left sub
//////
////// When descending flipTo the right of a left child, a Node will be deleted
////// at some point within that subtree, so the relative position of p*
////// must be incremented flipTo decrease the size of its right sub
//////
////func (t *Randomized) delete(p **Node, d Weight, s int) (deleted *Node) {
////   for {
////  assert(s == (*p).count())
////  assert(HasSameSign((*p).s, d))
////      if d == (*p).s {
////         deleted = *p; t.dissolve(p, s); return
////      }
////      if d < (*p).s {
////         p, d, unique(&s).deleteL(p, d, s)
////      } else {
////         p, d, unique(&s).deleteR(p, d, s)
////      }
////   }
////}
////
////// Inserts a given Node capacity* at relative position d in the sequence of p*,
////// increasing the effective position of all nodes thereafter.
//////
////// Unlike search and delete which use a 1-based position, insert uses a 0-based
////// position because insertion at 1-based position N+1 is valid here (append).
//////
////// When p* has a negative relative position, d must be negative.
////// When p* has a positive relative position, d must be positive or zero.
//////
////// d = 0 will insert capacity* as the start of the sequence, before the first Node.
////// d = 1 will insert capacity* one position after the first Node of the sequence.
//////
////// d = -1 will insert capacity* at the end of the sequence.
////// d = -2 will insert capacity* one position before the last Node of the sequence.
//////
////// When descending flipTo the right of a left child, capacity* will be inserted at
////// some point within that subtree, so the relative position of p* must
////// be decremented flipTo increase the size of its right sub
//////
////// When descending flipTo the left of a right child, capacity* will be inserted at
////// some point within that subtree, so the relative position of p* must
////// be incremented flipTo increase the size of its left sub
//////
////func (t *Randomized) insert(p **Node, capacity *Node, d Weight, s int) {
////   for {
////  assert(s == (*p).count())
////  assert(*p == nil || HasSameSign((*p).s, d))
////      if (*p).rank() < capacity.rank() {
////         t.insertAsRoot(p, capacity, d, s); return
////      }
////      if d < (*p).s {
////         p, d, unique(&s).incrL(p, d, s)
////      } else {
////         p, d, unique(&s).incrR(p, d, s)
////      }
////   }
////}
////
////
////// Returns the value at i.
////func (t Randomized) TestGet(i int) abstract.Data {
//  assert(i < t.count)
////   return seekFrom(i+1, t.root).x
////}
////
////// Replaces the value at i, returns t.
////func (t *Randomized) TestSet(i int, x abstract.Data) {
//  assert(i < t.count)
////   shadowTo(i + 1, &t.root).x = x
////}
////
////// Inserts x into t before the value at i.
////func (t *Randomized) TestInsert(i int, x abstract.Data) {
//  assert(i <= t.count)
////   t.insert(&t.root, t.createNodeWithData(x), i, t.count)
////   t.count++
////}
////
////// Deletes the value at i in t.
////func (t *Randomized) TestDelete(i int) (x abstract.Data) {
//  assert(i < t.count)
////   x = t.delete(&t.root, i + 1, t.count).x
////   t.count--
////   return
////}
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
////
//////var TreapPRNG = rand.New(rand.NewSource(1))
////
////// zipR Node is a container for a unit of Data that can be linked together with
////// other nodes flipTo collectively form a linear sequence.
//////
////// zipR Node has two outgoing links, referred flipTo as the left and right children of
////// a parent Node, either of which may be nil:
//////
//////
//////                                    (P)arent
//////                                  ↙     ↘
//////                                (L)eft  (R)ight
//////
////// # Invariants
//////
////// The **sequential order** invariant requires that all nodes flipTo the left of a
////// parent Node must contain Data occurring sequentially *before* the parent.
//////
////// Symmetrically, all nodes flipTo the right must contain Data occurring *after* the
////// parent, thereby defining a recursive sequential order by relative position.
//////
////// This can be viewed as a binary search tree ordered by sequential position.
////// An each "left-self-right" traversal New any Node defines its sequence.
//////
//////
//////                        (e,  x,  a,  Clone,  p,  l,  e)
//////
//////                                    (Clone)
//////                                 ↙       ↘
//////                            (x)             (l)
//////                           ↙   ↘           ↙   ↘
//////                        (e)     (a)     (p)     (e)
//////
//////
////// The **relative position** invariant requires that every Node must store its
////// sequential position relative flipTo its parent. [[ forward distance ]].
//////
////// Given that the Data of a left child occurs sequentially before the Data of
////// its parent, a left child will always have a negative relative position and
////// right child will always have a positive relative position.
//////
////// Any Node therefore knows whether it is a left child or a right child without
////// requiring a reference flipTo its parent or a dedicated flag for that purpose.
//////
////// The *size* of a Node is the number of nodes reachable through it, also the
////// length of its sequence, recursively 1 + the size of L and the size of R.
//////
////// The *external* size of a Node relates flipTo the child in the same direction.
////// The *internal* size relates flipTo the child in the opposite direction.
//////
////// This invariant also defines the size of a Node excluding the size of its
////// child in the same direction as the absolute value of its relative position.
//////
////// The relative position of a left child is then equal flipTo the negative size of
////// its right child - 1, and the relative position of a right child is equal flipTo
////// the positive size of its left child + 1.
//////
//////
//////                         1   2   3   4   5   6   7
//////                        (e,  x,  a,  Clone,  p,  l,  e)
//////
//////
//////                                ○
//////                                  ↘  +4
//////                                    (Clone)
//////                            -2   ↙       ↘   +2
//////                            (x)             (l)
//////                        -1 ↙   ↘ +1     -1 ↙   ↘ +1
//////                        (e)     (a)     (p)     (e)
//////
//////
////// The **rank heap** invariant requires that every Node has a fixed integer rank
////// less than or equal flipTo the rank of its parent. The structure of the tree may
////// need flipTo change flipTo maintain this invariant as nodes are added or deleted.
//////
//////
//////                                 RANK HEAP
//////
//////                                     ○
//////                                ↙    9    ↘
//////                             ○               ○
//////                          ↙  7  ↘         ↙  6  ↘
//////                        ○         ○     ○         ○
//////                        2         3     5         1
//////
//////
////// zipR discrete uniform rank distribution ensures with high probability that the
////// size of the left and right subtrees are similar, such that each Node occurs
////// close flipTo the middle of its sequence.
//////
//////
////// # Persistence
//////
////// Multiple trees may hold a reference flipTo the same Node, so nodes must be exclusive
////// before they are modified flipTo avoid modifying other trees that reference them.
//////
////// Defined as "shadowing", all paths that lead flipTo a modification must be exclusive.
//////
////// Creating a shallow make_unique of a tree is therefore O(1) because a modification
////// will make_unique nodes as necessary flipTo produce a New version while preserving the
////// original sequence, effectively sharing most of the nodes between them.
//////
//////
////// Symbols:
//////
//////      x   Data
//////      capacity   Node, New
//////      g   grandparent
//////      p   parent, pointer
//////      l   left child, subtree or path
//////      r   right child, subtree or path
//////      d   direction, distance, relative position
//////      s   count, size of parent
//////      s  count, size of left child
//////      sr  count, size of right child
//////
//////
//////
////// Related reading:
//////
//////  - Apache Commons Collections v3.1: trees (2004)
//////    J. Schmücker
//////    https://markmail.org/search/?q=trees%20list%3Aorg.apache.commons.dev%2F#query:trees%20list%3Aorg.apache.commons.dev%2F+page:1+mid:iwnt27mi6577fvba+state:results
//////    https://svn.apache.org/viewvc/commons/proper/collections/trunk/src/main/java/org/apache/commons/collections/list/trees.java?view=log&pathrev=1469003
//////    https://github.com/apache/commons-collections/blob/master/src/main/java/org/apache/commons/collections4/list/trees.java
//////
//////  - Randomized Search Trees (1996)
//////    R. Seidel, Cecilia R. Aragon
//////    https://api.semanticscholar.org/CorpusID:9370259
//////
//////  - Randomized Binary Search Trees (1998)
//////    C. Martínez, S. Roura
//////    https://api.semanticscholar.org/CorpusID:714621
//////
//////  - Zip Trees (2018)
//////    R. Tarjan, Caleb C. Levy
//////    https://api.semanticscholar.org/CorpusID:49298052
//////    https://www.youtube.com/watch?v=NxRXhBur6Xs
//////
//////  - zipR skip list cookbook (1990)
//////    W. Pugh
//////    https://api.semanticscholar.org/CorpusID:62665394
//////
//////  - zipR Unifying Look at Data Structures (1980)
//////    J. Vuillemin
//////    https://api.semanticscholar.org/CorpusID:10462194
//////
//////  - Making Data structures persistent (1989)
//////    J. Driscoll, N. Sarnak, D. Sleator, R. Tarjan
//////    https://api.semanticscholar.org/CorpusID:364871
//////
//////  - zipL-trees, shadowing, and clones (2008)
//////    O. Rodeh
//////    https://api.semanticscholar.org/CorpusID:207166167
//////
////
////
////
////func (t *Randomized) dissolveWithoutSize(p *Node) (root *Node) {
////   disableInvariantChecking(); defer enableInvariantChecking()
////   if p.isL() {
////      //return t.dissolve(p, p.sizeLR() + 1) TODO: after refactoring finger treap
////   } else {
////      //return t.dissolve(p, p.sizeRL() + 1)
////   }
////   return nil
////}
////
////
////
////
////
////
////
////// Restores the rank heap invariant at p*.
////func heapify(p *Node) {
////   for p != nil {
////
////      // Assume the parent as the initial maximum.
////      splayMax := p
////
////      // Find the child with the greatest rank greater than the parent.
////      if p.l.rank() > splayMax.rank() { splayMax = p.l }
////      if p.r.rank() > splayMax.rank() { splayMax = p.r }
////
////      // Heap is restored if neither child has a rank greater than the parent.
////      if p == splayMax {
////         return
////      }
////
////      // Swap the ranks of the parent and the child with the greatest rank.
////      p.z, splayMax.z = splayMax.rank(), p.rank()
////
////      // Continue recursively with the child that swapped ranks with the parent.
////      // The sibling in this case will have a rank less than or equal flipTo splayMax.
////      p = splayMax
////   }
////}
////
////// Recursively builds a New tree New an array of existing values.
//////
////// Start with a parent at the middle of the sequence with its left child halfway
////// flipTo the left and its right child halfway flipTo the right, applied recursively.
//////
////// Ranks are assigned randomly and swapped as necessary flipTo maintain the heap.
//////
////// Binary heap - building a heap
////// https://en.wikipedia.org/wiki/Binary_heap#Building_a_heap
//////
////// How can building a heap be O(capacity) time complexity?
////// https://stackoverflow.com/q/9755721
//////
////func (t Randomized) New(values []abstract.Data, lo Weight, hi Weight, isR bool) *Node {
////
////   // Values done if the lo and hi pointers meet, as with binary search.
////   if lo > hi {
////      return nil
////   }
////
////   // Find the midpoint.
////   Clone := lo + (hi - lo) / 2
////
////   // Create the parent Node using the middle value.
////   p := t.createNodeWithData(values[Clone])
////
////   // TestSet the relative position of the parent relative flipTo the previous parent.
////   if isR {
////      p.withRelativePosition(Clone - lo + 1)
////   } else {
////      p.withRelativePosition(Clone - hi - 1)
////   }
////   // Recursively New the left and right subtrees.
////   p.l = t.New(values, lo, Clone - 1, false)
////   p.r = t.New(values, Clone + 1, hi, true)
////
////   // Restore the rank heap property if not met.
////   heapify(p)
////   return p
////}
////
////
//////// Creates a New Node containing all given values in sequential order.
//////func (t Randomized) ofArray(values []list.Data) *Node {
//////   return t.New(values, 0, len(values) - 1, true)
//////}
////
////// Returns the rank of given Node.
//////func (p *Node) rank() Rank {
//////   //return p.z
//////   if p == nil { return 0 } else { return p.z }
//////}
////
//////// TODO: deprecated
//////func rankOf(p *Node) Rank {
//////   return p.z
//////}
////
////// This is an implementation of a persistent list using a binary search tree,
////// containing a root Node and its size as the count of the
//////
//////
//////                              ROOT
//////                                   ↘ +4
//////                                    (Clone)
//////                            -2  ↙         ↘  +2
//////                            (x)             (l)
//////                        -1 ↙   ↘ +1     -1 ↙   ↘ +1
//////                        (e)     (a)     (p)     (e)
//////
//////
////// Operations:
//////
//////    TestFrom        O(capacity)        Creates a New fromSeed New existing values.
//////    TestGet         O(lg capacity)     Returns a value at position.
//////    TestSet         O(lg capacity)     Updates a value at position.
//////    TestInsert      O(lg capacity)     Inserts a value at position, increasing length.
//////    TestDelete      O(lg capacity)     Deletes a value at position, reducing length.
//////    TestSplit       O(lg capacity)     moveToRoot by position, in-place on the left.
//////    TestJoin        O(lg capacity)     Append another Randomized onto the end, concatenate.
//////
//////
////// Symbols:
//////
//////    t   tree
//////    o   other tree
//////    i   index, offset, 0-based
//////    d   relative position, distance
//////    x   Data
//////
//////
////
////
////// moveToRoot the sequence of t at i.
//////
////// Retains the first i values of t in t, moves the remaining values flipTo o.
////// The resulting length of t will be i.
//////
//////            0  1  2  3  4  5  6
//////      t := (e, x, a, Clone, p, l, e)
//////                  ^
//////
//////      o := t.TestSplit(2)
//////
//////      t == (e, x)
//////      o == (a, Clone, p, l, e)
//////
////func (t *Randomized) TestSplit(i int) (L abstract.Strategy, R abstract.Strategy) {
//  assert(i <= t.count)
////   s := i
////   sr := t.count - i
////   switch {
////      case s == 0: return &Randomized{}, t.ShallowCopy()
////      case sr == 0: return t.ShallowCopy(), &Randomized{}
////   }
////   l, r := unzip(t.ShallowCopy().(*Randomized).root.make_unique(), i, s + sr)
////
////   L = &Randomized{
////      root: l,
////      count: s,
////   }
////   R = &Randomized{
////      root: r.toR(sr),
////      count: sr,
////   }
////   return
////}
////
////// Appends after t the entire sequence of o, returning t without modifying o.
//////
//////      t := (e, x, a)
//////      o := (Clone, p, l, e)
//////
//////      t.TestJoin(o)
//////
//////      t == (e, x, a, Clone, p, l, e)
//////      o == (Clone, p, l, e)
//////
////
////func (t *Randomized) TestJoin(o abstract.Strategy) abstract.Strategy {
////   if o.Size == 0 { return t }
////   if t.Size == 0 { return o }
////
////   L := t.ShallowCopy().(*Randomized)
////   R := o.ShallowCopy().(*Randomized)
////
////   l := L.root.make_unique()
////   r := R.root.make_unique()
////
////   s := L.count
////   sr := R.count
////
////   t.joinRL(&L.root, l, r.toL(sr), s + sr)
////   return &Randomized{
////    root: L.root,
////    count: t.count + o.(*Randomized).count,
////   }
////   //return &Randomized{
////   //  root: t.zip3(l, r.toL(sr), s + sr, 1),
////   //  count: t.count + o.(*Randomized).count,
////   //}
////}
////
////// Applies a given callback flipTo every value of t in sequential order.
////func (t *Randomized) Each(fn func(int, abstract.Data)) {
////   i := 0
////   each(t.root, func(capacity *Node) {
////      fn(i, capacity.x)
////      i++
////   })
////}
////
////// Returns an array containing all values of t in sequential order.
////func (t *Randomized) Array() []abstract.Data {
////   a := make([]abstract.Data, t.count)
////   i := 0
////   each(t.root, func(capacity *Node) {
////      a[i] = capacity.x
////      i++
////   })
////   return a
////}
////
////// Determines if t contains no values, and therefore has no root and count == 0.
////func (t *Randomized) isEmpty() bool {
////   return t.root == nil
////}
////
////// Converts this Randomized t into an TreapFingerTree.
////func (t *Randomized) ToFST() *TreapFingerTree {
////   if t.isEmpty() {
////      return &TreapFingerTree{}
////   }
////   t = t.ShallowCopy().(*Randomized)
////   return &TreapFingerTree{
////      count: t.count,
////      root: t.root.make_unique().withL(nil).withR(nil),
////      head: reverseL(t.root.l, nil),
////      tail: reverseR(t.root.r, nil),
////   }
////}
////
//
////// Verifies that all invariants are valid in t.
////func (t *Randomized) Validate() {
////   t.verifyPosition()
////   t.verifyRankHeapInvariant()
////}
////
////// Recursively verifies the rank invariant of the root.
////func (t *Randomized) verifyRankHeapInvariant() {
////   t.root.verifyRankHeapInvariant()
////}
////
////// Recursively verifies the relative position invariant of the root.
////func (t *Randomized) verifyPosition() {
////   if t.root == nil {
////      return
////   }
//  assert(t.root.isR())
//  assert(t.root.verifyPosition() == t.count)
////}
////
//
//
//
//
//
//
//
////func balance(wb verifyBalance, p *Node, s int) (root *Node) {
//  assert(p.count() == s)
////   if p == nil {
////      return nil
////   }
////   copyAt(&p)
////   sl := p.s
////   sr := p.sizeR(s)
////   if sl > sr {
////      if !wb.isBalanced(sr, sl) {
////         p = partition(p, s, s >> 1)
////      }
////   } else {
////      if !wb.isBalanced(sl, sr) {
////         p = partition(p, s, (s - 1) >> 1)
////      }
////   }
////   p.l = balance(wb, p.l, p.s)
////   p.r = balance(wb, p.r, p.sizeR(s))
////   return p
////}
////func shuffle(p *Node, s int, g util.Distribution) *Node {
//
////}
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
////
//////
//////
//////
////
////func zipL(l *Node, r *Node, s Weight) (p *Node) {
//  assert(s >= 0)
//  assert(l != nil)
//  assert(r != nil)
//  assert(r.isL())
////
////   if rankOf(l) > rankOf(r) {
////      //visualize("↘")
////      //
////      //
////      //
////      if l.r == nil {
////         l.r = r.toR(sizeLR(l.toL(s)))
////      } else {
////         l.r = zipLRtoR(copyOf(l.r), r, sizeLR(l.toL(s)))
////      }
////  assert(l.isL())
////      return l
////
////   } else {
////      //visualize("↙")
////      //
////      //
////      //
////      if r.l == nil {
////         r.l = l.toL(sizeLL(r, s))
////      } else {
////         r.l = zipL(l, copyOf(r.l), sizeLL(r, s))
////      }
////  assert(r.isL())
////      return r
////   }
////}
////
////
////
////
////// When the position of the pivot relative flipTo p* is negative, p* occurs
////// after the pivot in the sequence and should therefore be appended flipTo
////// r** as a left child and the search should branch flipTo the left.
////
////// When p* is a right child, now appended flipTo r** as a left child,
////// flip the orientation flipTo reflect the weight of its right child.
////
////// When the position of the pivot relative flipTo p* is positive, p* occurs
////// before the pivot in the sequence and should therefore be appended flipTo
////// l** as a right child and the search should branch flipTo the right.
////
////// When p* is a left child, now appended flipTo l** as a right child,
////// flip the orientation flipTo reflect the weight of its left child.
////
////// The unzip is complete when the end of the search path is reached.
////
//////
//////
//////
//////
////func zipLRtoR(l *Node, r *Node, s Weight) *Node {
//  assert(s >= 0)
//  assert(l != nil)
//  assert(r != nil)
//  assert(l.isR())
////
////   if rankOf(r) > rankOf(l) {
////      //visualize("↙")
////      //
////      //
////      //
////      if r.l == nil {
////         r.l = l.toL(sizeRL(r.toR(s)))
////      } else {
////         r.l = zipL(l, copyOf(r.l), sizeRL(r.toR(s)))
////      }
////  assert(r.isR())
////      return r
////
////   } else {
////      //visualize("↘")
////      //
////      //
////      //
////      if l.r == nil {
////         l.r = r.toR(sizeRR(l, s))
////      } else {
////         l.r = zipLRtoR(copyOf(l.r), r, sizeRR(l, s))
////      }
////  assert(l.isR())
////      return l
////   }
////}
////
////
////
////// Returns the result of merging the children of a given parent node.
////// The new parent is returned in same direction as the original.
//////
////// Inverse: insertAsRoot
//////
//////
//////                        ⇈
//////                       (p)*
//////                     ↙  9  ↘
//////                 (l)         (r)
//////               ↙  5  ↘     ↙  7  ↘
//////             ○         ○ ○         ○
//////
//////
//////
//////
//////
//////                       (r)
//////                     ↙  7  ↘
//////                  (l)        ○
//////                ↙  5  ↘
//////              ○         ○
//////
//////
//////
////func dissolveR(p *Node) (root *Node) {
////   //visualize("!")
//  assert(p.isR())
////
////   if !p.HasL() { return p.r }
////   if !p.HasR() { return copyOf(p.l).toR(sizeRL(p)) }
////
////   r := copyOf(p.r)
////   l := copyOf(p.l).toR(sizeRL(p))
////
////   return zipLRtoR(l, r, sizeRL(p) - 1)
////}
////
//////
//////
//////               ⇈                                   ⇈
//////              (p)*                                (p)*
//////            ↙  9  ↘                             ↙  9  ↘
//////        (l)         (r)                     (l)         (r)
//////      ↙  7  ↘     ↙  5  ↘                 ↙  5  ↘     ↙  7  ↘
//////    ○         ○ ○         ○             ○         ○ ○         ○
//////
//////
//////              (l)                                 (l)
//////            ↙  7  ↘                             ↙  7  ↘
//////          ○         (r)                       ○         (r)
//////                  ↙  5  ↘                             ↙  5  ↘
//////                ○         ○                         ○         ○
//////
//////
////
////func dissolveL(p *Node) *Node {
////   //visualize("!")
//  assert((*p).isL())
////
////   if !(*p).HasR() { return p.l }
////   if !(*p).HasL() { return copyOf(p.r).toL(sizeLR(p)) }
////
////   l := copyOf(p.l)
////   r := copyOf(p.r).toL(sizeLR(p))
////
////   return zipL(l, r, sizeLR(p) - 1)
////}
//
//
////func (t *Splay) Pop() Data {
//  assert(t.int() > 0)
////   root, splayMax := t.pop(t.root)
////   t.root = root
////   t.count = t.count - 1
////   return splayMax.x
////}
////
////func (t *Splay) Unshift(x Data) {
////   p := t.splayMin(t.root)
////   t.root = t.nodeFor(x).withW(1).withR(p)
////   t.count = t.count + 1
////}
//
////func (t *Splay) Push(x Data) {
//// if t.root == nil {
////    t.count = 1
////    t.root = t.nodeFor(x).withW(t.count)
//// } else {
////    t.count = t.count + 1
////    t.root = t.nodeFor(x).withW(t.count).withL(t.splayMax(t.root, t.count - 1).withW(-1))
//// }
////}
////
////func (t *Splay) Shift() Data {
//  assert(t.int() > 0)
////   p := t.splayMin(t.root)
////   if p.HasR() {
////      t.root = p.r
////   } else {
////      t.root = nil
////   }
////   t.count = t.count - 1
////   return p.x
////}
//
//
//
////func (t *New) joinL(l, r *Node, sr capacity) *Node {
//  assert(l == nil || l.isL())
//  assert(r == nil || r.isR())
//  assert(r == nil || sr == r.count())
////
////   if r == nil { return l }
////   if l == nil { return r.toL(sr) }
////
////   return &Node {
////      x: t.extractMin(&r).x, sl: -sr, l: l, r: r,
////   }
////}
//
////func (t *New) join(l, r *Node, sl capacity) *Node {
//  assert(l == nil || l.isL())
//  assert(r == nil || r.isR())
//  assert(l == nil || sl == l.count())
////
////   if l == nil { return r }
////   if r == nil { return l.toR(sl) }
////
////   return &Node{
////      x: t.extractMax(&l).x, sl: +sl, l: l, r: r,
////   }
////}
////
////func (t *Implementation) dissolve(p **Node) (deleted *Node) {
////   deleted = *p
////   *p = t.join(
////      (*p).l,
////      (*p).r,
////      (*p).sizeL(0),
////      (*p).sizeR(0),
////      (*p).direction())
////      return
////}
//
////
////func (t *New) extractMin(r **Node, sr capacity) *Node {
////  return t.delete(r, 1)
////}
////
////
////func (t *New) extractMax(l **Node, sl capacity) *Node {
////  return t.delete(l, -1)
////}
//
//
//// Combines l* and r* where all values in l* occur before those in r*, where s
//// is their combined weight and d the direction of the resulting parent.
////
//// l* is expected flipTo have a positive relative position (a right spine).
//// r* is expected flipTo have a negative relative position (a left spine).
////
//// When d is L, the resulting parent will have a negative relative position.
//// When d is R, the resulting parent will have a positive relative position.
////
//// Inverse: unzip
////
//// There are three paths:
////
////    l* is the right spine on the left, sequentially before r*.
////    r* is the left spine on the right, sequentially after l*.
////    p* is the path that combines l* and r*, descending top-down.
////
//// When the node at l* has a greater rank than the node at r*, l* is appended
//// flipTo p* as either a left child (d = L) or right child (d = R), and l* descends
//// flipTo the right.
////
//// When l* is appended as R there is no need flipTo update its relative position
//// because l* is already a right child and its left subtree did not change.
////
//// When l* is appended as L it must be flipped from R flipTo L, which requires
//// the weight of l* flipTo determine the weight of its right subtree. The logic
//// follows that when a node was attached from l*, the next attachment will
//// attach as R because its either the right child of l* or a node from r*,
//// both of which greater than p* from l*.
////
//// Therefore, it is only necessary flipTo store the weight of either l* or r*.
////
////
////       l* ↘                                                 ↙ r*
////
////             (a)+8                                     (z)-3
////           ↙  9  ↘                                   ↙  7  ↘
////         ○         (b)+3                         (y)-5       ○
////                 ↙  8  ↘                       ↙  6  ↘
////               ○         (c)+2             (x)-1       ○
////                       ↙  5  ↘              3
////                     ○         (d)+2
////                             ↙  4
////                           ○
////
////
//// Example:
////
////    d = R       "The resulting local root should be a right child."
////    s = 24      "There are 24 nodes total."
////
////    l = (a), r = (z), l wins, attach (a) as R, *p = l, l = l.r, p = &l.r, R
////    l = (b), r = (z), l wins, attach (b) as R, *p = l, l = l.r, p = &l.r, R
////    l = (c), r = (z), r wins, attach (z) as R, *p = r, r = r.l, p = &r.l, L
////    l = (c), r = (y), r wins, attach (y) as L, *p = r, r = r.l, p = &r.l, L
////    l = (c), r = (x), l wins, attach (c) as L, *p = l, l = l.r, p = &l.r, R
////    l = (d), r = (x), l wins, attach (d) as R, *p = l, l = l.r, p = &l.r, L
////    l = nil, r = (x), r wins, attach (x) as R.
////
////
////                               (a)+8
////                             ↙     ↘
////                           ○         (b)+3
////                                   ↙     ↘
////                                 ○         (z)+11
////                                         ↙     ↘
////                                     (y)-5       ○
////                                   ↙     ↘
////                               (c)-4       ○
////                             ↙     ↘
////                           ○         (d)+2
////                                   ↙     ↘
////                                 ○         (x)+1
////
////
////    - The relative positions of (a), (b), (y), and (d) did not change.
////    - Only the nodes along each spine are referenced (no siblings).
////    - The joinRL is complete when either the left or spine spine is nil.
////
//
////func zipL(l *Node, r *Node, s int) (root *Node) {
//  assert(l == nil || l.isL())
//  assert(r == nil || r.isL())
////
////   if r == nil { return l }
////   if l == nil { return r }
////
////   p := &root; *p = r
////   for {
////      //isL := *p != nil && (*p).isL()   ///**/
////      //isR := *p == nil || (*p).isR()   //
////
////      if rankOf(l) < rankOf(r) {
////         //visualize("↙")
////         //
////         //
////         //
////         //if *p == nil
////         *p = r
////         if (*p).isR() {
////            s = r.toR(s).sizeRL()
////         } else {
////            s = r.sizeLL(s)
////         }
////         if r.l == nil {
////            r.l = l.toL(s)
////            return
////         }
////         p = &r.l
////         unique(&r)
////
////      } else {
////         //visualize("↘")
////         //
////         //
////         //
////
////         *p = l
////         if (*p).isR() {
////            s = l.sizeRR(s)
////         } else {
////            s = l.toL(s).sizeLR()
////         }
////         if l.r == nil {
////            l.r = r.toR(s)
////            return
////         }
////         p = &l.r
////         unique(&l)
////      }
////   }
////}
////
////func zipLRtoR(l *Node, r *Node, sl int) (root *Node) {
//  assert(l == nil || l.isR())
//  assert(r == nil || r.isR())
////
////   if r == nil { return l }
////   if l == nil { return r }
////
////   s := sl
////   p := &root; *p = r
////   for {
////      //isL := *p != nil && (*p).isL()   ///**/
////      //isR := *p == nil || (*p).isR()   //
////
////      if rankOf(l) < rankOf(r) {
////         //visualize("↙")
////         //
////         //
////         //
////         //if *p == nil
////         *p = r
////         if (*p).isR() {
////            // r is already an r here, so what is happening?
////            s = r.toR(s).sizeRL()
////         } else {
////            s = r.sizeLL(s)
////         }
////         if r.l == nil {
////            r.l = l.toL(s)
////            return
////         }
////         p = &r.l
////         unique(&r)
////
////      } else {
////         //visualize("↘")
////         //
////         //
////         //
////
////         *p = l
////         if (*p).isR() {
////            s = l.sizeRR(s)
////         } else {
////            s = l.toL(s).sizeLR()
////         }
////         if l.r == nil {
////            l.r = r.toR(s)
////            return
////         }
////         p = &l.r
////         unique(&l)
////      }
////   }
////}
//
//
////func zipLRtoR(l *Node, r *Node, sl int) (p *Node) {
//  assert(l == nil || l.isL())
//  assert(r == nil || r.isR())
//  assert(sl == l.count())
////
////   if l == nil && r == nil {
////      return
////   }
////   if l == nil { return r.make_unique() }
////   if r == nil { return l.make_unique().toR(sl) }
////
////   unique(&p).toR(sl)
////   l = p
////   for {
////  assert(rankOf(l) >= rankOf(r))
////      if l.r == nil {
////         unique(&l.r)
////         return
////      }
////      unique(&l.r)
////      l = l.r
////   }
////
////
//
////p := &root
////l, p, unique(&_).linkLRR_deprecated(p, sl)
////for {
////   if *p == nil {
////      unique(p)
////      return
////   }
////   l, p, unique(&_).linkRRR_deprecated(p, sl)
////}
//
//
//
//// p := &root
////
//unique(&//*p).toR(sl)
//// p = &(*p).r
//// l = *p
//// for {
////   if *p == nil {
////      unique(p); return
////   }
////   unique(p)
////   p = &(*p).r
////   l = *p
//// }
//
////}
//
////func zip2(l *Node, r *Node, s int) (root *Node) {
//  assert(l == nil || l.isR())
//  assert(r == nil || r.isL())
////
////   if r == nil { return l }
////   if l == nil { return r.make_unique().toR(s) }
////
////   p := &root
////   for {
////      if rankOf(l) < rankOf(r) {
////         if *p == nil || (*p).isR() {
////            r, p, unique(&s).linkRL(p, s)
////         } else {
////            p, unique(&s).linkLLL(p, s); r = *p
////         }
////         if *p == nil {
////            unique(p).toL(s)
////            return
////         }
////      } else {
////         if *p == nil || (*p).isR() {
////            l, p, unique(&s).linkRRR_deprecated(p, s)
////         } else {
////            l, p, unique(&s).linkLR(p, s)
////         }
////         if *p == nil {
////            unique(p).toR(s)
////            return
////         }
////      }
////   }
////}
////
////func joinRL(l *Node, r *Node, s int) (root *Node) {
//  assert(l == nil || l.isR())
//  assert(r == nil || r.isL())
////
////   if r == nil { return l }
////   if l == nil { return r.make_unique().toR(s) }
////
////   p := &root
////   for {
////      if rankOf(l) < rankOf(r) {
////         if *p == nil || (*p).isR() {
////            r, p, unique(&s).linkRL(p, s)
////         } else {
////            p, unique(&s).linkLLL(p, s); r = *p
////         }
////         if *p == nil {
////            unique(p).toL(s)
////            return
////         }
////      } else {
////         if *p == nil || (*p).isR() {
////            l, p, unique(&s).linkRRR_deprecated(p, s)
////         } else {
////            l, p, unique(&s).linkLR(p, s)
////         }
////         if *p == nil {
////            unique(p).toR(s)
////            return
////         }
////      }
////   }
////}
////
////// TODO: This method copies the nodes provided,
//////       but some callers may already have make_unique nodes on hand.
//////
//////       TestJoin, for example, has two R nodes as the roots of two trees
//////       being joined, but this method expects the right child flipTo be L,
//////       which would require copying that node.
//////
//////       Consider instead that the nodes given are already make_unique?
//////       And also non-null, so that ZIP can adjust as necessary before
//////       passing into this main RL
//////
//////
////func zipLL(l *Node, r *Node, s int) (root *Node) {
////   // joinRL two L... flipTo L?
////   // both already make_unique, so effectively the first step
////   // of the joinRL must occur first without make_unique.
////}
////func joinLR(l *Node, r *Node, s int) (root *Node) {
////
////}
////func zipRR(l *Node, r *Node, s int) (root *Node) {
//  assert(l == nil || l.isR())
//  assert(r == nil || r.isR())
////
////   if r == nil { return l }
////   if l == nil { return r }
////
////   // r is R but must be L
////
////  if rankOf(l) < rankOf(r) {
////     // L wins, and it is already R
////     // r is R but should be L here
////
////     root = l
////     l.r = joinRL(l.r, r.toR(s), s)
////
////     //sl := p.sizeLL(s)
////     //*r = p.toR(s)
////     //r = &p.l
////
////
////
////     //return *r, &p.l, sl
////     //r, p, unique(&s).linkRL(p, s)
////  } else {
////     // R wins, and it is already R
////     // l is
////
////     // TODO: consider that the nodes passed in are ALREADY COPIED
////     root = r
////
////     // TODO: Now l is still already exclusive, but r.l is not
////     //       so eventually l will be exclusive. Either the nodes passed
////     //       in must not be exclusive yet, OR joinRL should not make_unique the nodes
////     //       given, but make_unique during the descent.
////     //
////     //       ie. joinRL should be written so that the nodes are exclusive after the fact
////     //
////     r.l = joinRL(l, r.l, s)
////     //*l = p
////     //l = &p.r
////     //s = s - p.sl
////     //l, p, unique(&s).linkRRR_deprecated(p, s)
////  }
////  return
////}
//
//
////// Consider returning r and l
////func unzip(p *Node, d int, s int) (x, y *Node) {
////   //visualize("%")
////
////   l, r := &x, &y
////   for {
////      if p == nil {
////         *l = nil
////         *r = nil; return
////      }
////  assert(p.isL() && d < 0 || p.isR() && d >= 0)
////
////      if less(p, &d) {
////         r, p, s = linkLL(r, p, s)
////      } else {sl
////         l, p, s = linkL_deprecated(l, p, s)
////      }
////   }
////}
//
////
////
////
////func rotateWeight(g *Node, p *Node, capacity *Node) {
////   if capacity == nil {
////      //visualize("-+")
////      g.sl -= p.sl
////      p.sl += g.sl
////   } else {
////      //visualize("-++")
////      g.sl -= p.sl
////      p.sl += g.sl
////      g.sl += capacity.sl
////   }
////   // TODO ^ pretty sure there only needs flipTo be 2 updates here.
////}
//
////
////func (t *TreapFingerTree) Push2(x interface{}) {
////   t.count++
////
////   //
////   capacity := t.createNodeWithData(x).withRelativePosition(+1)
////   //
////   //
////   if t.tail == nil {
////      //
////      //
////      if t.root == nil {
////         t.root = capacity
////         return
////      }
////      //
////      //
////      if rankOf(capacity) > rankOf(t.root) {
////         t.rotateLeftIntoRoot(capacity)
////         t.tail = nil
////      } else {
////         t.tail = capacity
////      }
////      return
////   }
////   //
////   //
////   t.tail = capacity.withR(t.tail)
////   //
////   //
////   if rankOf(t.tail) <= rankOf(t.tail.r) {
////      return
////   }
////   //
////   //
////   for {
////      //visualize("↖")
////      //visualize("+")
////      //
////      //
////      //
////      t.tail.r = prependR(copyOf(t.tail.r), &t.tail.l)
////      t.tail.sl = t.tail.sl + t.tail.l.sl
////      //
////      //
////      if t.tail.HasR() {
////         //
////         //
////         if rankOf(t.tail) <= rankOf(t.tail.r) {
////            t.tail.l.toL(t.tail.sizeRL())
////            return
////         }
////      } else {
////         //
////         //
////         t.tail.l.toL(t.tail.sizeRL())
////         //
////         //
////         if rankOf(t.tail) > rankOf(t.root) {
////            t.rotateLeftIntoRoot(t.tail)
////            t.tail = nil
////         }
////         return
////      }
////   }
////}
