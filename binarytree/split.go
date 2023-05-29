package binarytree


func (tree *Tree) split(p *Node, i uint64) (*Node, *Node) {
   n := Node{}
   l := &n
   r := &n
   for p != nil {
      tree.copy(&p)
      if i <= p.s {
         p.s = p.s - i
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
   l.r = nil
   r.l = nil
   return n.r, n.l
}
