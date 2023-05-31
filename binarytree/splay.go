package binarytree

import . "binarysearch/abstract/list"

type Splay struct {
	Tree
}

func (tree Splay) New() List {
	return &Splay{}
}

func (tree *Splay) Clone() List {
	return &Splay{Tree: tree.Tree.Clone()}
}

// `i` is the number of nodes that will be attached to the left still
// so we are not tracking the total size at any point.
// we reduce the size of the left subtree of p by i - 1
//
// size left reduces by (i + 1)
//
// the resulting size of p is then (s - p.s - 1 + i)
// the size before
//
//  1. Link `p` to the left of `r`
//
//  2. Set r
//
//     (p)            (r)
//     ↙
func (tree *Tree) linkL(p *Node, r *Node, i Position) (*Node, *Node, Size) {
	tree.copy(&p.l)
	p.s = p.s - i - 1
	r.l = p
	r = r.l
	p = p.l
	return p, r, i // for those that will not be attached.. some might
	// not be linked, but it is not aware of how it is linked
} // So maybe it is the number of nodes that must be taken away
//  p.s - i - 1 is the future size of the left subtree
//  p.s is the current size of the left subtree.
//  r.s is the future size of p
//  so in the future there are (i + 1) fewer nodes in the left subtree.

//

//  those nodes will NOT be linked.
//
//  So (i+i) is a

func (tree *Tree) linkR(p *Node, l *Node, i Position) (*Node, *Node, Size) {
	tree.copy(&p.r)
	i = i - p.s - 1
	l.r = p
	l = l.r
	p = p.r
	return p, l, i
}

func (tree *Splay) rotateL(p *Node) *Node {
	tree.copy(&p.r)
	return p.rotateL()
}

func (tree *Splay) rotateR(p *Node) *Node {
	tree.copy(&p.l)
	return p.rotateR()
}

func (tree *Splay) splay(p *Node, i Position) *Node {
	tree.copy(&p)
	n := Node{s: i}
	l := &n
	r := &n
	for i != p.s {
		if i < p.s {
			if i < p.l.s {
				p, r, i = tree.linkL(tree.rotateR(p), r, i)
			} else if i > p.l.s {
				p, r, i = tree.linkL(p, r, i)
				p, l, i = tree.linkR(p, l, i)
			} else {
				p, r, i = tree.linkL(p, r, i)
				break
			}
		} else {
			if i > p.s+p.r.s+1 {
				p, l, i = tree.linkR(tree.rotateL(p), l, i)
			} else if i < p.s+p.r.s+1 {
				p, l, i = tree.linkR(p, l, i)
				p, r, i = tree.linkL(p, r, i)
			} else {
				p, l, i = tree.linkR(p, l, i)
				break
			}
		}
	}
	l.r = p.l
	r.l = p.r
	p.r = n.l
	p.l = n.r
	p.s = n.s
	return p
}

func (tree *Splay) Splay(i Position) {
	tree.root = tree.splay(tree.root, i)
}

func (tree *Splay) Size() Size {
	return tree.size
}

// 1. Splay the node at `i`
// 2. Return the root.
func (tree *Splay) Select(i Position) (x Data) {
	assert(i < tree.Size())
	tree.Splay(i)
	return tree.root.x
}

// 1. Splay the node to be updated, at position `i`.
// 2. Update the root's data.
// 3. Return the root.
func (tree *Splay) Update(i Position, x Data) {
	assert(i < tree.Size())
	tree.Splay(i)
	tree.root.x = x
}

//  1. Node a new node for the Data `s`.
//  2. Split the root into the left and right subtrees of the new node, such
//     that the first `i` nodes are on the left and the rest on the right.
//  3. Replace the previous root with the new node.
func (tree *Splay) Insert(i Position, x Data) {
	assert(i <= tree.Size())
	//
	//
	if i == tree.size {
		tree.root = tree.allocate(Node{x: x, s: tree.size, l: tree.splayMax(tree.root)})
		tree.size++
		return
	}
	l, r := tree.split(tree.root, tree.size, i)
	tree.root = tree.allocate(Node{x: x, s: i, l: l, r: r})
	tree.size++
}

// 1. Splay the node to be deleted, making it the root.
// 2. Replace the root by a join of its left and right subtrees.
// 3. Return the deleted node.
func (tree *Splay) Delete(i Position) (x Data) {
	assert(i < tree.Size())

	tree.Splay(i)
	defer tree.release(tree.root)
	x = tree.root.x
	tree.root = tree.join(tree.root.l, tree.root.r)
	tree.size--
	return
}

func (tree *Splay) Split(i Position) (List, List) {
	assert(i <= tree.Size())
	tree.share(tree.root)

	if i == tree.size {
		return &Splay{Tree{arena: tree.arena, root: tree.root, size: tree.size}},
			&Splay{Tree{arena: tree.arena, root: nil, size: 0}}
	}
	//
	//
	l, r := tree.split(tree.root, tree.size, i)

	return &Splay{Tree{arena: tree.arena, root: l, size: i}},
		&Splay{Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

// 1. Splay the node at `i`.
// 2. Cut the left subtree of the root as l**, leaving a nil in its place.
// 3. The remaining root and right subtree is r**.
func (tree Splay) split(p *Node, s, i Position) (l, r *Node) {
	assert(i < s)
	p = tree.splay(p, i)
	l = p.l
	r = p
	r.l = nil
	r.s = 0
	return
}

func (tree *Splay) Join(that List) List { // TODO check if benchmarks are affected by poointer receivers here
	tree.share(tree.root)
	tree.share(that.(*Splay).root)
	return &Splay{Tree{arena: tree.arena, root: tree.join(tree.root, that.(*Splay).root), size: tree.Size() + that.Size()}}
}

func (tree *Splay) splayMax(l *Node) *Node {
	if l == nil { // TODO is this ever nil?
		return nil
	}
	tree.copy(&l)
	for l.r != nil {
		if l.r.r != nil {
			tree.copy(&l.r)
			tree.copy(&l.r.r)
			l.r = l.r.rotateL()
			l = l.rotateL()
		} else {
			tree.copy(&l.r)
			l = l.rotateL()
		}
	}
	return l
}

// 1. Splay the right-most node of l*, which wouldn't have a right subtree.
// 2. Set the right subtree of the splayed node to r*.
// 3. Return the splayed node.
func (tree *Splay) join(l *Node, r *Node) *Node {
	if l == nil {
		return r
	}
	l = tree.splayMax(l)
	l.r = r
	return l
}

func (tree *Splay) Verify() {
	tree.verifySizes()
}
