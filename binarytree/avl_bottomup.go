package binarytree

type AVLBottomUp struct {
	AVL
}

func (AVLBottomUp) New() List {
	return &AVLBottomUp{}
}

func (tree *AVLBottomUp) Clone() List {
	return &AVLBottomUp{AVL{Tree: tree.Tree.Clone()}}
}

func (tree *AVLBottomUp) insert(p *Node, i Position, x Data) *Node {
	if p == nil {
		return tree.allocate(Node{x: x})
	}
	tree.copy(&p)
	if i <= p.s {
		p.s = p.s + 1
		p.l = tree.insert(p.l, i, x)
	} else {
		p.r = tree.insert(p.r, i-p.s-1, x)
	}
	return tree.fix(p)
}

func (tree *AVLBottomUp) Insert(i Position, x Data) {
	tree.root = tree.insert(tree.root, i, x)
	tree.size++
}

func (tree *AVLBottomUp) delete(p *Node, i Position, x *Data) *Node {
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
	return tree.fix(p)
}

func (tree *AVLBottomUp) Delete(i Position) (x Data) {
	assert(i < tree.size)
	tree.root = tree.delete(tree.root, i, &x)
	tree.size--
	return
}

func (tree *AVLBottomUp) Join(other List) List {
	return &AVLBottomUp{tree.AVL.Join(other.(*AVLBottomUp).AVL)}
}

func (tree *AVLBottomUp) Split(i Position) (List, List) {
	l, r := tree.AVL.Split(i)
	return &AVLBottomUp{l},
		&AVLBottomUp{r}
}
