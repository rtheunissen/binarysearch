package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/distribution"
   . "binarysearch/utility"
)

type Balancer interface {
   Restore(Tree) Tree
   Verify(Tree)
}

type Balance interface {
   isBalanced(x, y Size) bool
}

type Partition struct {
   Balance
}

func (strategy Partition) Restore(tree Tree) Tree {
   tree.root = strategy.balance(&tree, tree.root, tree.size)
   return tree
}

func (strategy Partition) balance(tree *Tree, p *Node, s Size) *Node {
   if s < 3 {
      return p
   }
   sl := p.sizeL()
   sr := p.sizeR(s)

   assert(strategy.isBalanced(sl, sr) || sl < sr)
   assert(strategy.isBalanced(sr, sl) || sr < sl)

   // Replace `p` by its underlying median if not balanced.
   if !strategy.isBalanced(sl, sr) || !strategy.isBalanced(sr, sl) {
      p = tree.partition(p, s >> 1)
   }
   // Recursively balance the left and right subtrees.
   p.l = strategy.balance(tree, p.l, p.sizeL())
   p.r = strategy.balance(tree, p.r, p.sizeR(s))
   return p
}

func (strategy Partition) verify(p *Node, s Size) {
   if p == nil {
      return
   }
   sl := p.sizeL()
   sr := p.sizeR(s)

   invariant(strategy.isBalanced(sl, sr))
   invariant(strategy.isBalanced(sr, sl))

   strategy.verify(p.l, sl)
   strategy.verify(p.r, sr)

}

func (strategy Partition) Verify(tree Tree) {
   strategy.verify(tree.root, tree.size)

}

type Weight struct{}

func (balancer Weight) Restore(tree Tree) Tree {
   return Partition{balancer}.Restore(tree)
}

func (balancer Weight) Verify(tree Tree) {
   balancer.verify(tree.root, tree.size)
}

// -1 <= ⌊log₂(L)⌋ - ⌊log₂(R)⌋ <= 1
func (balancer Weight) verify(p *Node, s Size) {
   if p == nil {
      return
   }
   invariant(Difference(Log2(p.sizeL() + 1), Log2(p.sizeR(s) + 1)) <= 1)

   balancer.verify(p.l, p.sizeL())
   balancer.verify(p.r, p.sizeR(s))
}

func (Weight) isBalanced(x, y Size) bool {
   return !SmallerLog2(x+1, (y+1)>>1)
}




type Half struct{}

func (balancer Half) Restore(tree Tree) Tree {
   return Partition{balancer}.Restore(tree)
}

func (Half) isBalanced(x, y Size) bool {
   return !(x < (y >> 1))
}

func (balancer Half) Verify(tree Tree) {
   balancer.verify(tree.root, tree.size)
}

func (balancer Half) verify(p *Node, s Size) {
   if p == nil {
      return
   }
   //invariant(Difference(p.l.height(), p.r.height()) <= 1)
   //
   //balancer.verify(p.l, p.sizeL())
   //balancer.verify(p.r, p.sizeR(s))
}



type Half2 struct{}

func (balancer Half2) Restore(tree Tree) Tree {
   return Partition{balancer}.Restore(tree)
}

func (Half2) isBalanced(x, y Size) bool {
   return !((x + 1) < ((y + 1) >> 1))
}

func (balancer Half2) Verify(tree Tree) {
   balancer.verify(tree.root, tree.size)
}

func (balancer Half2) verify(p *Node, s Size) {
   if p == nil {
      return
   }
   //invariant(Difference(p.l.height(), p.r.height()) <= 1)
   //
   //balancer.verify(p.l, p.sizeL())
   //balancer.verify(p.r, p.sizeR(s))
}



type Log struct{}

func (balancer Log) Restore(tree Tree) Tree {
   return Partition{balancer}.Restore(tree)
}

func (balancer Log) Verify(tree Tree) {
   balancer.verify(tree.root, tree.size)
}

// -1 <= ⌊log₂(L)⌋ - ⌊log₂(R)⌋ <= 1
func (balancer Log) verify(p *Node, s Size) {
   if p == nil {
      return
   }
   invariant(Difference(Log2(p.sizeL()), Log2(p.sizeR(s))) <= 1)

   balancer.verify(p.l, p.sizeL())
   balancer.verify(p.r, p.sizeR(s))
}

func (Log) isBalanced(x, y Size) bool {
   return !SmallerLog2(x, y>>1)
}

//!(MSB(y) < MSB(x >> 1))

type Median struct{}

func (balancer Median) Restore(tree Tree) Tree {
   return Partition{balancer}.Restore(tree)
}

func (balancer Median) Verify(tree Tree) {
   balancer.verify(tree.root, tree.size)
}

// -1 <= L - R <= 1
func (balancer Median) verify(p *Node, s Size) {
   if p == nil {
      return
   }
   invariant(Difference(p.sizeL(), p.sizeR(s)) <= 1)

   balancer.verify(p.l, p.sizeL())
   balancer.verify(p.r, p.sizeR(s))
}

func (Median) isBalanced(x, y Size) bool {
   return !(x+1 < y)
}

type Height struct{}

func (balancer Height) Restore(tree Tree) Tree {
   return Partition{balancer}.Restore(tree)
}

func (Height) isBalanced(x, y Size) bool {
   return !SmallerLog2(x+1, y)
}

func (balancer Height) Verify(tree Tree) {
   balancer.verify(tree.root, tree.size)
}

// A node is height-balanced when the difference between the height of its
// subtrees is no greater than 1, and both subtrees are also height-balanced.
//
// invariant(p.height() <= FloorLog2(s))
func (balancer Height) verify(p *Node, s Size) {
   if p == nil {
      return
   }
   invariant(Difference(p.l.height(), p.r.height()) <= 1)

   balancer.verify(p.l, p.sizeL())
   balancer.verify(p.r, p.sizeR(s))
}

//
//func (balancer Weight) Balance(p *Node, s Size) *Node {
//   return PartitionBalancer{balancer}.balance(p, s)
//}
//
//func (balancer Median) Balance(p *Node, s Size) *Node {
//   return PartitionBalancer{balancer}.balance(p, s)
//}
//
//func (balancer Height) Balance(p *Node, s Size) *Node {
//   return PartitionBalancer{balancer}.balance(p, s)
//}

//
//type ArrayRebuilder struct {
//}
//
//func (balancer ArrayRebuilder) Restore(tree Tree) Tree {
//   tree.root = balancer.fromArray(tree, treeToArray(tree.root, tree.size), 1, tree.size)
//   return tree
//}
//
//func (balancer ArrayRebuilder) fromArray(tree Tree, values []Data, i, j Position) *Node {
//   if i > j {
//      return nil
//   }
//   m := i + (j - i ) >> 1
//
//   return tree.allocate(Node{
//      x: values[m - 1],
//      s: m - i,
//      l: balancer.fromArray(tree, values, i, m - 1),
//      r: balancer.fromArray(tree, values, m + 1, j),
//   })
//}
//
//func (balancer ArrayRebuilder) Verify(tree Tree) {
//   invariant(tree.root.height() <= int(Log2(tree.size)))
//}
//
//func treeToArray(p *Node, s Size) []Data {
//  array := make([]Data, s)
//  stack := make([]*Node, 0)
//  index := 0
//  for {
//     if p != nil {
//        stack = append(stack, p)
//        p = p.l
//     } else {
//        n := len(stack)
//        if n == 0 {
//           return array
//        }
//        p = stack[n-1]
//        stack = stack[:n-1]
//        array[index] = p.x
//        index++
//        p = p.r
//     }
//  }
//}

// https://web.eecs.umich.edu/~qstout/pap/CACM86.pdf
type DSW struct {
}

func (balancer DSW) Verify(tree Tree) {
   invariant(tree.root.height() <= int(Log2(tree.size)))
}

//func (tree Tree) fromVineToTree() Tree {
//   //leaves ← size + 1 − 2**⌊log2(size + 1))⌋
//
//   leaves := NextPowerOf2LessThanOrEqualTo(tree.size+1) - 1 //tree.size + 1 - (1 << FloorLog2(tree.size + 1))  // size + 1 - (1 << FloorLog2(size + 1))
//   tree.compress(tree.root, tree.size - leaves)
//   leaves = tree.size - leaves
//   for leaves > 1 {
//      tree.compress(tree.root, leaves / 2)
//      leaves = leaves / 2
//   }
//   return tree
//}

func (balancer DSW) Restore(tree Tree) Tree {
   return balancer.toTree(balancer.toVine(tree))
}

func (balancer DSW) toVine(tree Tree) Tree {
   n := Node{}
   l := &n
   p := tree.root
   s := tree.size
   for p != nil {
      //tree.copy(&p)
      for p.l != nil {
         p = p.rotateR()
         //tree.rotateR(&p)
      }
      l.r = p
      l = l.r
      p = p.r
   }
   tree.root = n.r
   tree.size = s
   return tree
}

func (balancer DSW) toTree(tree Tree) Tree {
   p := tree.root
   s := tree.size
   m := 1<<Log2(s+1) - 1

   p = balancer.compress(p, s-Size(m))
   for m > 1 {
      p = balancer.compress(p, Size(m)>>1)
      m = m >> 1
   }
   tree.root = p
   tree.size = s
   return tree
}

func (balancer DSW) compress(p *Node, c Size) *Node {
   n := Node{}
   l := &n
   n.r = p
   for ; c > 0; c-- {
      l.r = p.rotateL()
      l = l.r
      p = l.r
   }
   return n.r
}

//func (tree Tree) compress(count Size) Tree {
//   n := Node{}
//   l := &n
//   p := tree.root
//   for ; count > 0; count-- {
//      tree.copy(&p)
//      tree.rotateL(&p)
//      l.r = p
//      l = l.r
//      p = p.r
//   }
//   tree.root = n.r
//   return tree
//}
//
//
// routine compress(root, count)
//    scanner ← root
//    for i ← 1 to count
//        child ← scanner.right
//        scanner.right ← child.right
//        scanner ← scanner.right
//        child.right ← scanner.left
//        scanner.left ← child

//n := Node{}
//l := &n
//n.r = p
//for ; count > 0; count-- {
//   p = p.rotateL()
//   l.r = p
//   l = l.r
//   p = p.r
//}
//return n.r

//if count == 0 {
//   return v
//}
//tree.rotateL(&v)
//v.r = tree.compress(v.r, count - 1)
//return v
//

//for ; count > 0; count-- {
//  tree.rotateL(v)
//  v = &(*v).r
//}

//n := Node{}
//l := &n
//p := *v
//for ; count > 0; count-- {
//   p = p.rotateL()
//   l.r = p
//   l = l.r
//   p = p.r
//}
//*v = n.r

//
//func (tree Tree) fromVineToTree() Tree {
//   //p := tree.root
//   s := NextPowerOf2LessThanOrEqualTo(tree.size + 1) - 1
//   //var count = tree.size - s
//   for count := tree.size - s; count > 0; count-- {
//      tree.compress(&tree.root, count)
//   }
//   for s > 1 {
//      for count := s / 2; count > 0; count-- {
//         tree.compress(&tree.root, count)
//      }
//      s = s / 2
//   }
//   return tree
//}
//
//func (tree Tree) compress(p **Node, count Size) {
//   for ; count > 0; count-- {
//      tree.copy(p)
//      tree.rotateL(p)
//      p = &(*p).r
//   }
//}

//
//func (tree Tree) fromVineToTree() Tree {
//   //leaves ← size + 1 − 2**⌊log2(size + 1))⌋
//
//   leaves := NextPowerOf2LessThanOrEqualTo(tree.size+1) - 1 //tree.size + 1 - (1 << FloorLog2(tree.size + 1))  // size + 1 - (1 << FloorLog2(size + 1))
//   tree.compress(tree.root, tree.size - leaves)
//   leaves = tree.size - leaves
//   for leaves > 1 {
//      tree.compress(tree.root, leaves / 2)
//      leaves = leaves / 2
//   }
//   return tree
//}

//func (tree Tree) compress(p *Node, rotations Size) {
//   for ; rotations > 0; rotations-- {
//      tree.rotateL(&p)
//      tree.copy(&p.r)
//      p = p.r
//   }
//}
//
//

// There is a clever morris-based traversal version that benchmarked slower.
//
//     // Convert tree to a "vine", i.e., a sorted linked list,
//    // using the right pointers to point to the next node in the list
//    tail ← root
//    rest ← tail.right
//    while rest ≠ nil
//        if rest.left = nil
//            tail ← rest
//            rest ← rest.right
//        else
//            temp ← rest.left
//            rest.left ← temp.right
//            temp.right ← rest
//            rest ← temp
//            tail.right ← temp

//	func (tree Tree) Vine(size Size) (root *Node) {
//	 p := &root
//	 for ; size > 0; size-- {
//	   *p = tree.allocate(Node{})
//	    p = &(*p).r
//	 }
//	 return
//	}
func (Tree) Vine(size Size) Tree {
   t := Tree{}
   n := Node{}
   p := &n
   for t.size = 0; t.size < size; t.size++ {
      p.r = t.allocate(Node{})
      p = p.r
   }
   t.root = n.r
   return t
}

func (Tree) WorstCaseMedianVine(size Size) Tree {
   assert(size > 0)
   t := Tree{}
   n := Node{}
   p := &n
   for t.size = 0; t.size < (size-1)/2+1; t.size++ {
      p.r = t.allocate(Node{})
      p = p.r
   }
   for ; t.size < size; t.size++ {
      p.l = t.allocate(Node{})
      p.s = size - t.size
      p = p.l
   }
   t.root = n.r
   return t
}

// public void flatten(TreeNode root) {
//    if (root == null)
//        return;
//    flatten(root.right);
//    flatten(root.left);
//    root.right = prev;
//    root.left = null;
//    prev = root;
//}

// void flatten(TreeNode *root) {
//	while (root) {
//		if (root->left && root->right) {
//			TreeNode* t = root->left;
//			while (t->right)
//				t = t->right;
//			t->right = root->right;
//		}

//
//        if(root->left)
//		    root->right = root->left;
//		root->left = NULL;
//		root = root->right;
//	}
//}

//func (tree Tree) treeToVineAlt(p *Node) (vine *Node) {
//   if p == nil {
//      return nil
//   }
//   vine = p
//   for {
//      if p.l != nil && p.r != nil {
//         t := p.l
//         for t.r != nil {
//            t = t.r
//         }
//         if p.r != nil {
//         }
//         t.r = p.r
//      }
//      if p.l != nil {
//         p.r = p.l
//      }
//      p.l = nil
//      p.s = 0
//      if p.r == nil {
//         break
//      }
//      p = p.r
//   }
//   return vine
//}

// TODO: make pointer
func (tree Tree) Randomize(access distribution.Distribution) Tree {
   tree.root = tree.randomize(access, tree.root, tree.size)
   return tree
}

func (tree Tree) randomize(access distribution.Distribution, p *Node, s Size) *Node {
   assert(p.size() == s)
   if p == nil {
      return nil
   }
   p = tree.partition(p, access.LessThan(s))
   p.l = tree.randomize(access, p.l, p.sizeL())
   p.r = tree.randomize(access, p.r, p.sizeR(s))
   return p
}
