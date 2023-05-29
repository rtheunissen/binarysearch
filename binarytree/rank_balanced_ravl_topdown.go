package binarytree

type AVLRelaxedTopDown struct {
	AVLWeakTopDown
}

func (AVLRelaxedTopDown) New() List {
	return &AVLRelaxedTopDown{
		AVLWeakTopDown: *AVLWeakTopDown{}.New().(*AVLWeakTopDown),
	}
}

func (tree *AVLRelaxedTopDown) Clone() List {
	return &AVLRelaxedTopDown{
		AVLWeakTopDown: AVLWeakTopDown{
			WAVL: WAVL{
				Tree: tree.Tree.Clone(),
			},
		},
	}
}

func (tree AVLRelaxedTopDown) Verify() {
	tree.Tree.verifySizes()
	tree.verifyValidRankDiffs(tree.root)
}

func (tree AVLRelaxedTopDown) verifyValidRankDiffs(p *Node) {
	if p == nil {
		return
	}
	invariant(tree.rank(p) >= p.height())
	invariant(tree.rank(p) > tree.rank(p.l))
	invariant(tree.rank(p) > tree.rank(p.r))

	tree.verifyValidRankDiffs(p.l)
	tree.verifyValidRankDiffs(p.r)
}

func (tree *AVLRelaxedTopDown) Select(i Size) Data {
	assert(i < tree.Size())
	return tree.lookup(tree.root, i)
}

func (tree *AVLRelaxedTopDown) Update(i Size, x Data) {
	assert(i < tree.Size())
	tree.copy(&tree.root)
	tree.update(tree.root, i, x)
}

func (tree *AVLRelaxedTopDown) Delete(i Position) Data {
	assert(i < tree.Size())
	x := tree.Tree.delete(&tree.root, tree.size, i)
	tree.size--
	return x
}

func (tree AVLRelaxedTopDown) join(l, r *Node, sl Size) (p *Node) {
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

func (tree AVLRelaxedTopDown) Join(other List) List {
	l := tree
	r := other.(*AVLRelaxedTopDown)
	tree.share(l.root)
	tree.share(r.root)

	p := tree.join(l.root, r.root, l.size)

	return &AVLRelaxedTopDown{AVLWeakTopDown: AVLWeakTopDown{WAVL{Tree: Tree{arena: tree.arena, root: p, size: l.size + r.size}}}}
}

func (tree AVLRelaxedTopDown) Split(i Position) (List, List) {
	assert(i <= tree.Size())
	tree.share(tree.root)
	l, r := JoinBased{Tree: tree.Tree, Joiner: tree}.split(tree.root, i, tree.size)

	return &AVLRelaxedTopDown{AVLWeakTopDown: AVLWeakTopDown{WAVL: WAVL{Tree: Tree{arena: tree.arena, root: l, size: i}}}},
		&AVLRelaxedTopDown{AVLWeakTopDown: AVLWeakTopDown{WAVL: WAVL{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}}}
}
