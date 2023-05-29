package binarytree

type AVLRelaxedBottomUp struct {
	AVLWeakBottomUp
}

func (tree AVLRelaxedBottomUp) verifyValidRankDiffs(p *Node) {
	if p == nil {
		return
	}
	invariant(tree.rank(p) >= p.height())
	invariant(tree.rank(p) > tree.rank(p.l))
	invariant(tree.rank(p) > tree.rank(p.r))

	tree.verifyValidRankDiffs(p.l)
	tree.verifyValidRankDiffs(p.r)
}

func (tree AVLRelaxedBottomUp) Verify() {
	tree.verifySizes()
	tree.verifyValidRankDiffs(tree.root)
}

func (AVLRelaxedBottomUp) New() List {
	return &AVLRelaxedBottomUp{
		AVLWeakBottomUp: *AVLWeakBottomUp{}.New().(*AVLWeakBottomUp), // TODO: ew
	}
}

func (tree *AVLRelaxedBottomUp) Clone() List {
	return &AVLRelaxedBottomUp{
		AVLWeakBottomUp: AVLWeakBottomUp{
			WAVL: WAVL{
				Tree: tree.Tree.Clone(),
			},
		},
	}
}

func (tree *AVLRelaxedBottomUp) Delete(i Position) Data {
	assert(i < tree.Size())
	x := tree.Tree.delete(&tree.root, tree.size, i)
	tree.size--
	return x
}

func (tree AVLRelaxedBottomUp) join(l, r *Node, sl Size) (p *Node) {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	if tree.rank(l) <= tree.rank(r) {
		return tree.build(l, tree.deleteMin(&r), r, sl)
	} else {
		return tree.build(l, tree.deleteMax(&l), r, sl-1)
	}
}

func (tree AVLRelaxedBottomUp) Join(other List) List {
	l := tree
	r := other.(*AVLRelaxedBottomUp)

	tree.share(l.root)
	tree.share(r.root)

	return &AVLRelaxedBottomUp{
		AVLWeakBottomUp: AVLWeakBottomUp{
			WAVL: WAVL{
				Tree: Tree{
					arena: tree.arena,
					root:  tree.join(l.root, r.root, l.size),
					size:  l.size + r.size,
				},
			},
		},
	}
}

func (tree AVLRelaxedBottomUp) Split(i Position) (List, List) {
	assert(i <= tree.Size())
	tree.share(tree.root)
	l, r := JoinBased{Tree: tree.Tree, Joiner: tree}.split(tree.root, i, tree.size)

	return &AVLRelaxedBottomUp{AVLWeakBottomUp: AVLWeakBottomUp{WAVL: WAVL{Tree: Tree{arena: tree.arena, root: l, size: i}}}},
		&AVLRelaxedBottomUp{AVLWeakBottomUp: AVLWeakBottomUp{WAVL: WAVL{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}}
}
