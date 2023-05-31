package binarytree


// Moves the node at position `i` in the tree of `p` to its root and returns the
// resulting root. This algorithm is identical to splay with no rotation steps.
//
func (tree Tree) partition(p *Node, i uint64) *Node {
   // assert(i < p.size())
   // measurement(&partitionCount, 1)

   n := Node{s: i}
   l := &n
   r := &n
   for i != p.s {
      // measurement(&partitionDepth, 1)
      //tree.copy(&p)
      if i < p.s {
         p.s = p.s - i - 1
         r.l = p
         r = r.l
         p = p.l
      } else {
         i = i - p.s - 1
         l.r = p
         l = l.r
         p = p.r
      }
   }
   //tree.copy(&p)
   l.r = p.l
   r.l = p.r
   p.l = n.r
   p.r = n.l
   p.s = n.s
   return p
}

// TODO: temp because randomize uses partition
func (tree Tree) partition2(p *Node, i uint64) *Node {
   // assert(i < p.size())

   n := Node{s: i}
   l := &n
   r := &n
   for i != p.s {
      //tree.copy(&p)
      if i < p.s {
         p.s = p.s - i - 1
         r.l = p
         r = r.l
         p = p.l
      } else {
         i = i - p.s - 1
         l.r = p
         l = l.r
         p = p.r
      }
   }
   //tree.copy(&p)
   r.l = p.r
   l.r = p.l
   p.l = n.r
   p.r = n.l
   p.s = n.s
   return p
}
