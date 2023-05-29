package binarytree

type AVLWeakBottomUp struct {
	WAVL // TODO: just use tree
}

func (AVLWeakBottomUp) New() List {
	return &AVLWeakBottomUp{}
}

func (tree *AVLWeakBottomUp) Clone() List {
	return &AVLWeakBottomUp{
		WAVL{
			Tree: tree.Tree.Clone(),
		},
	}
}

func (tree *AVLWeakBottomUp) insert(p *Node, i Position, x Data) *Node {
	if p == nil {
		return tree.allocate(Node{x: x})
	}
	tree.copy(&p)
	if i <= p.s {
		p.s = p.s + 1
		p.l = tree.insert(p.l, i, x)
		return tree.rebalanceBottomUpAfterInsertingLeft(p)
	} else {
		p.r = tree.insert(p.r, i-p.s-1, x)
		return tree.rebalanceBottomUpAfterInsertingRight(p)
	}
}

//
// "Deletion of a leaf may convert its parent, previously a 1,2 node
//  into a 2,2 leaf, violating the rank rule. In this case we begin
//  by demoting the parent, which may make it a 3-child."
//
// func (tree AVLWeakBottomUp) rebalanceBottomUpDeletingLeft(p *Node) *Node {
//    return tree.rebalanceDeletingLeft(p)
// }

//
// "Deletion of a leaf may convert its parent, previously a 1,2 node
//  into a 2,2 leaf, violating the rank rule. In this case we begin
//  by demoting the parent, which may make it a 3-child."
//
// func (tree AVLWeakBottomUp) rebalanceBottomUpDeletingRight(p *Node) *Node {
//    return tree.rebalanceDeletingRight(p)
// }

func (tree AVLWeakBottomUp) delete(p *Node, i Position, x *Data) *Node {
	tree.copy(&p)
	if i == p.s {
		*x = p.x
		defer tree.release(p)
		return tree.join(p.l, p.r, p.s)
	}
	if i < p.s {
		p.s = p.s - 1
		p.l = tree.delete(p.l, i, x)
	} else {
		p.r = tree.delete(p.r, i-p.s-1, x)
	}
	return tree.rebalanceOnDelete(p)
}

func (tree *AVLWeakBottomUp) Delete(i Position) (x Data) {
	assert(i < tree.Size())
	tree.root = tree.delete(tree.root, i, &x)
	tree.size = tree.size - 1
	return
}

func (tree *AVLWeakBottomUp) Insert(i Position, x Data) {
	assert(i <= tree.Size())
	tree.size = tree.size + 1
	tree.root = tree.insert(tree.root, i, x)
}

func (tree AVLWeakBottomUp) extractMin(p *Node, min **Node) *Node {
	if p.l == nil {
		*min = tree.replacedByRightSubtree(&p)
		return p
	}
	tree.copy(&p)
	p.s--
	p.l = tree.extractMin(p.l, min)
	return tree.rebalanceOnDelete(p)
}

func (tree AVLWeakBottomUp) extractMax(p *Node, max **Node) *Node {
	if p.r == nil {
		*max = tree.replacedByLeftSubtree(&p)
		return p
	}
	tree.copy(&p)
	p.r = tree.extractMax(p.r, max)
	return tree.rebalanceOnDelete(p)
}

func (tree AVLWeakBottomUp) join2(l, r *Node, sl, sr Size) (p *Node) {
	return nil // TODO: unused tree.join(l, r, sl)
}

func (tree AVLWeakBottomUp) join3(l, p, r *Node, sl, sr Size) *Node {
	return tree.build(l, p, r, sl)
}

func (tree AVLWeakBottomUp) join(l, r *Node, sl Size) (p *Node) {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	if tree.rank(l) <= tree.rank(r) {
		return tree.build(l, p, tree.extractMin(r, &p), sl)
	} else {
		return tree.build(tree.extractMax(l, &p), p, r, sl-1)
	}
}

func (tree *AVLWeakBottomUp) Select(i Size) Data {
	assert(i < tree.Size())
	return tree.lookup(tree.root, i)
}

func (tree *AVLWeakBottomUp) Update(i Size, x Data) {
	assert(i < tree.Size())
	tree.copy(&tree.root)
	tree.update(tree.root, i, x)
}

func (tree AVLWeakBottomUp) Join(that List) List {
	tree.share(tree.root)
	tree.share(that.(*AVLWeakBottomUp).root)
	return &AVLWeakBottomUp{
		WAVL{
			Tree: Tree{
				arena: tree.arena,
				root:  tree.join(tree.root, that.(*AVLWeakBottomUp).root, tree.size),
				size:  tree.size + that.(*AVLWeakBottomUp).size,
			},
		},
	}
}

func (tree AVLWeakBottomUp) Split(i Position) (List, List) {
	assert(i <= tree.Size())
	tree.share(tree.root)
	l, r := JoinBased{Tree: tree.Tree, Joiner: tree}.split(tree.root, i, tree.size)

	return &AVLWeakBottomUp{WAVL{Tree: Tree{arena: tree.arena, root: l, size: i}}},
		&AVLWeakBottomUp{WAVL{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}
}
