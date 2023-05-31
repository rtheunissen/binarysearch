package binarytree

import . "binarysearch/abstract/list"

type AVLJoinBased struct {
	AVL
}

func (AVLJoinBased) New() List {
	return &AVLJoinBased{}
}

func (tree *AVLJoinBased) Clone() List {
	return &AVLJoinBased{
		AVL: AVL{
			Tree: tree.Tree.Clone(),
		},
	}
}

func (tree *AVLJoinBased) insert(p *Node, i Position, s Size, x *Node) *Node {
	if p == nil {
		return x
	}
	tree.copy(&p)

	sl := p.s
	sr := s - p.s - 1

	if i <= p.s {
		p.s++
		return tree.build(tree.insert(p.l, i, sl, x), p, p.r, sl+1)
	} else {
		return tree.build(p.l, p, tree.insert(p.r, i-sl-1, sr, x), sl)
	}
}

func (tree *AVLJoinBased) Insert(i Position, x Data) {
	assert(i <= tree.size)
	tree.root = tree.insert(tree.root, i, tree.size, tree.allocate(Node{x: x}))
	tree.size++
}

func (tree AVLJoinBased) delete(p *Node, i Position, s Size, x *Data) *Node {
	tree.copy(&p)

	sl := p.s
	sr := s - p.s - 1

	if i == p.s {
		*x = p.x
		defer tree.release(p)
		return tree.join(p.l, p.r, sl)
	}
	if i < p.s {
		p.s--
		return tree.build(tree.delete(p.l, i, sl, x), p, p.r, sl-1)
	} else {
		return tree.build(p.l, p, tree.delete(p.r, i-sl-1, sr, x), sl)
	}
}

func (tree *AVLJoinBased) Delete(i Position) (x Data) {
	assert(i < tree.size)
	tree.root = tree.delete(tree.root, i, tree.size, &x)
	tree.size--
	return
}

func (tree *AVLJoinBased) Join(other List) List {
	return &AVLJoinBased{tree.AVL.Join(other.(*AVLJoinBased).AVL)}
}

func (tree *AVLJoinBased) Split(i Position) (List, List) {
	l, r := tree.AVL.Split(i)
	return &AVLJoinBased{l},
		&AVLJoinBased{r}
}
