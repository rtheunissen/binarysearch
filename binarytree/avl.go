package binarytree

import (
   "binarysearch/abstract/list"
   "binarysearch/utility"
)

type AVL struct {
   Tree
}

func (AVL) height(p *Node) int {
   if p == nil {
      return -1
   } else {
      return int(p.y)
   }
}

func (tree *AVL) calculateHeight(p *Node) {
   p.y = uint64(utility.Max(tree.height(p.l), tree.height(p.r)) + 1)
}

//func (tree *AVL) calculateHeights(p *Node) {
//   if p == nil {
//      return
//   }
//   tree.calculateHeights(p.l)
//   tree.calculateHeights(p.r)
//   tree.calculateHeight(p)
//}

func (tree *AVL) verifyHeight(p *Node) {
   if p == nil {
      return
   }
   invariant(utility.Difference(tree.height(p.l), tree.height(p.r)) <= 1)

   invariant(tree.height(p) > tree.height(p.l))
   invariant(tree.height(p) > tree.height(p.r))

   tree.verifyHeight(p.l)
   tree.verifyHeight(p.r)
}

func (tree *AVL) verifySize(p *Node, s list.Size) list.Size {
   if p == nil {
      return 0
   }
   sl := tree.verifySize(p.l, p.s)
   sr := tree.verifySize(p.r, s-p.s-1)

   invariant(s == sl+sr+1)
   return s
}

func (tree *AVL) Verify() {
   tree.verifySize(tree.root, tree.size)
   tree.verifyHeight(tree.root)
}

func (tree *AVL) fix(p *Node) *Node {
   if tree.height(p.r) > tree.height(p.l)+1 {
   	if tree.height(p.r.l) > tree.height(p.r.r) {
   		tree.rotateRL(&p)
   		tree.calculateHeight(p.l)
   		tree.calculateHeight(p.r)
   		tree.calculateHeight(p)
   	} else {
   		tree.rotateL(&p)
   		tree.calculateHeight(p.l)
   		tree.calculateHeight(p)
   	}
   } else if tree.height(p.l) > tree.height(p.r)+1 {
   	if tree.height(p.l.r) > tree.height(p.l.l) {
   		tree.rotateLR(&p)
   		tree.calculateHeight(p.l)
   		tree.calculateHeight(p.r)
   		tree.calculateHeight(p)
   	} else {
   		tree.rotateR(&p)
   		tree.calculateHeight(p.r)
   		tree.calculateHeight(p)
   	}
   } else {
   	tree.calculateHeight(p)
   }
   return p
}

func (tree *AVL) Select(i list.Size) list.Data {
   // assert(i < tree.size)
   return tree.lookup(tree.root, i)
}

func (tree *AVL) Update(i list.Size, x list.Data) {
   // assert(i < tree.size)
   tree.copy(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *AVL) deleteMin(p *Node, min **Node) *Node {
   tree.copy(&p)
   if p.l == nil {
   	*min = p
   	return p.r
   }
   p.s = p.s - 1
   p.l = tree.deleteMin(p.l, min)
   return tree.fix(p)
}

func (tree *AVL) deleteMax(p *Node, max **Node) *Node {
   tree.copy(&p)
   if p.r == nil {
   	*max = p
   	return p.l
   }
   p.r = tree.deleteMax(p.r, max)
   return tree.fix(p)
}

func (tree *AVL) buildL(l, p, r *Node, sl list.Size) *Node {
   if tree.height(l)-tree.height(r) <= 1 {
   	p.l = l
   	p.r = r
   	p.s = sl
   	tree.calculateHeight(p)
   	return p
   }
   tree.copy(&l)
   l.r = tree.buildL(l.r, p, r, sl-l.s-1)
   return tree.fix(l)
}

func (tree *AVL) buildR(l, p, r *Node, sl list.Size) *Node {
   if tree.height(r)-tree.height(l) <= 1 {
   	p.l = l
   	p.r = r
   	p.s = sl
   	tree.calculateHeight(p)
   	return p
   }
   tree.copy(&r)
   r.s = 1 + sl + r.s
   r.l = tree.buildR(l, p, r.l, sl)
   return tree.fix(r)
}

func (tree *AVL) build(l, p, r *Node, sl list.Size) *Node {
   if tree.height(l) > tree.height(r) {
   	return tree.buildL(l, p, r, sl)
   } else {
   	return tree.buildR(l, p, r, sl)
   }
}

func (tree *AVL) joinL(l, r *Node, sl list.Size) (p *Node) {
   if tree.height(l)-tree.height(r) <= 1 {
   	return tree.build(tree.deleteMax(l, &p), p, r, sl-1)
   }
   tree.copy(&l)
   l.r = tree.joinL(l.r, r, sl-l.s-1)
   return tree.fix(l)
}

func (tree *AVL) joinR(l, r *Node, sl list.Size) (p *Node) {
   if tree.height(r)-tree.height(l) <= 1 {
   	return tree.build(l, p, tree.deleteMin(r, &p), sl)
   }
   tree.copy(&r)
   r.s = sl + r.s
   r.l = tree.joinR(l, r.l, sl)
   return tree.fix(r)
}

func (tree *AVL) join(l, r *Node, sl list.Size) (p *Node) {
   if r == nil {
   	return l
   }
   if l == nil {
   	return r
   }
   if tree.height(l) > tree.height(r) {
   	return tree.joinL(l, r, sl)
   } else {
   	return tree.joinR(l, r, sl)
   }
}

func (tree *AVL) Join(other AVL) AVL {
   tree.share(tree.root)
   tree.share(other.root)
   return AVL{
   	Tree{
   		arena: tree.arena,
   		root:  tree.join(tree.root, other.root, tree.size),
   		size:  tree.size + other.size,
   	},
   }
}

func (tree *AVL) split(p *Node, i, s list.Size) (l, r *Node) {
   if p == nil {
   	return
   }
   tree.copy(&p)
   if i <= (*p).s {
   	l, r = tree.split(p.l, i, p.s)
   	r = tree.build(r, p, p.r, p.s-i)
   } else {
   	l, r = tree.split(p.r, i-p.s-1, s-p.s-1)
   	l = tree.build(p.l, p, l, p.s)
   }
   return l, r
}

func (tree *AVL) Split(i list.Position) (AVL, AVL) {
   // assert(i <= tree.size)
   tree.share(tree.root)

   l, r := tree.split(tree.root, i, tree.size)

   return AVL{Tree{arena: tree.arena, root: l, size: i}},
   	AVL{Tree{arena: tree.arena, root: r, size: tree.size - i}}
}
