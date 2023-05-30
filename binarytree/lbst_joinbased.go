package binarytree

import . "trees/abstract/list"

type LBSTJoinBased struct {
	LBST // TODO: Remove these, just use Tree at the base always
}

func (tree LBSTJoinBased) New() List {
	return &LBSTJoinBased{}
}

func (tree *LBSTJoinBased) Clone() List {
	return &LBSTJoinBased{
		LBST: LBST{
			Tree: tree.Tree.Clone(),
		},
	}
}

func (tree *LBSTJoinBased) Select(i Size) Data {
	assert(i < tree.Size())
	return tree.lookup(tree.root, i)
}

func (tree *LBSTJoinBased) Update(i Size, x Data) {
	assert(i < tree.Size())
	tree.copy(&tree.root)
	tree.update(tree.root, i, x)
}

func (tree *LBSTJoinBased) Insert(i Position, x Data) {
	assert(i <= tree.Size())
	tree.root = JoinBased{Tree: tree.Tree, Joiner: tree}.insert(tree.root, i, tree.size, tree.allocate(Node{x: x}))
	tree.size++
}

//
//func (tree LBSTBottomUp) Split(i Position) (List, List) {
//   l, r := tree.LBST.Split(i)
//
//   return &LBSTBottomUp{l},
//      &LBSTBottomUp{r}
//}
//
//func (tree LBSTBottomUp) Join(that List) List {
//   return &LBSTBottomUp{tree.LBST.Join(that.(*LBSTBottomUp).LBST)}
//}

func (tree *LBSTJoinBased) Delete(i Position) (x Data) {
	assert(i < tree.Size())
	tree.root = JoinBased{Tree: tree.Tree, Joiner: tree}.delete(tree.root, i, tree.size, &x)
	tree.size--
	return
}

func (tree LBSTJoinBased) Split(i Position) (List, List) {
	tree.share(tree.root)
	l, r := JoinBased{Tree: tree.Tree, Joiner: tree}.splitToBST(tree.root, i, tree.size)

	return &LBSTJoinBased{LBST{l}},
		&LBSTJoinBased{LBST{r}}
}

func (tree LBSTJoinBased) Join(that List) List {
	l := tree
	r := that.(*LBSTJoinBased)
	tree.share(l.root)
	tree.share(r.root)
	return &LBSTJoinBased{
		LBST{
			Tree: Tree{
				arena: tree.arena,
				root:  tree.join2(l.root, r.root, l.size, r.size),
				size:  l.size + r.size,
			},
		},
	}
}
