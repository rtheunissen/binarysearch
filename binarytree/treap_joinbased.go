package binarytree

import (
	"trees/random"
)

type TreapJoinBased struct {
	TreapTopDown // TODO: maybe all trees should have only BST at the base, expanded
}

func (TreapJoinBased) New() List {
	return &TreapJoinBased{
		TreapTopDown{
			Source: random.New(random.Uint64()),
		},
	}
}

func (tree *TreapJoinBased) Clone() List {
	return &TreapJoinBased{
		TreapTopDown{
			Tree:   tree.Tree.Clone(),
			Source: tree.Source,
		},
	}
}

func (tree *TreapJoinBased) Insert(i Position, x Data) {
	assert(i <= tree.Size())
	tree.root = JoinBased{Tree: tree.Tree, Joiner: tree}.insert(tree.root, i, tree.size, tree.allocate(Node{x: x, y: tree.Source.Uint64()}))
	tree.size++
}
func (tree *TreapJoinBased) Delete(i Position) (x Data) {
	assert(i < tree.Size())
	tree.root = JoinBased{Tree: tree.Tree, Joiner: tree}.delete(tree.root, i, tree.size, &x)
	tree.size--
	return
}

func (tree TreapJoinBased) Split(i Position) (List, List) {
	assert(i <= tree.Size())
	tree.share(tree.root)
	l, r := JoinBased{Tree: tree.Tree, Joiner: &tree}.splitToBST(tree.root, i, tree.size)

	return &TreapJoinBased{TreapTopDown{l, tree.Source}}, // TODO consider merging Treap and Treap??
		&TreapJoinBased{TreapTopDown{r, tree.Source}}
}

func (tree TreapJoinBased) Join(that List) List {
	l := tree
	r := that.(*TreapJoinBased)
	tree.share(l.root)
	tree.share(r.root)
	return &TreapJoinBased{
		TreapTopDown{
			Source: tree.Source,
			Tree: Tree{
				arena: tree.arena,
				root:  tree.join(l.root, r.root, l.size),
				size:  l.size + r.size,
			},
		},
	}
}

func (tree TreapJoinBased) Verify() {
	tree.Tree.verifySizes()
	tree.verifyMaxRankHeap(tree.root)
}
