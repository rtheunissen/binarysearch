package binarytree

// TODO: just inline all of this ? then remove the rank_balanced filename prefix?

type RankBalanced struct {
}

func (RankBalanced) rank(p *Node) int {
   return p.rank()
}

func (RankBalanced) promote(p *Node){
   p.y++
}

// Deprecated, why? use tree.rank
func (p *Node) rank() int {
   if p == nil {
      return -1
   } else {
      return int(p.y)
   }
}

func promote(p *Node) {
   p.y++
}

func demote(p *Node) {
   p.y--
}

func isZeroZero(p *Node) bool {
   return isZeroChild(p, p.l) && isZeroChild(p, p.r)
}

func isOneOne(p *Node) bool {
   return isOneChild(p, p.l) && isOneChild(p, p.r)
}

func isTwoTwo(p *Node) bool {
   return isTwoChild(p, p.l) && isTwoChild(p, p.r)
}

func rankDifference(parent, child *Node) int {
   assert(parent.rank() >= child.rank())
   return parent.rank() - child.rank()
}

func isZeroChild(parent, child *Node) bool {
   return parent.rank() == child.rank() //rankDifference(parent, child) == 0
}

func isOneChild(parent, child *Node) bool {
   return rankDifference(parent, child) == 1
}

func isTwoChild(parent, child *Node) bool {
   return rankDifference(parent, child) == 2
}

func isThreeChild(parent, child *Node) bool {
   return rankDifference(parent, child) == 3
}

