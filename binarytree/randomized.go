package binarytree

import (
   . "binarysearch/abstract/list"
   "binarysearch/random"
)

type Randomized struct {
	Tree
	random.Source // compare performance vs making this directly xoshiro
}

func (Randomized) New() List {
	return &Randomized{
		Source: random.New(random.Uint64()),
	}
}

//func (tree *Randomized) build(values []Data) *Node {
//   if len(values) == 0 {
//      return nil
//   }
//   m := random.LessThan(uint64(len(values)), tree.Source)
//
//   return tree.allocate(Node{
//      x: values[m],
//      s: m,
//      l: tree.build(values[:m]),
//      r: tree.build(values[m+1:]),
//   })
//}

func (tree *Randomized) Clone() List {
	return &Randomized{
		Tree:   tree.Tree.Clone(),
		Source: tree.Source, // TODO: a copy method? or clone?
	}
}

//  1. Node a new node for the value `s`.
//  2. Split the root into the left and right subtrees of the new node, such
//     that the first `i` nodes are on the left and the rest on the right.
//  3. Replace the previous root with the new node.
func (tree *Randomized) Insert(i Position, x Data) {
	p := &tree.root
	s := tree.size
	tree.size++
	for {
		if random.LessThan(s+1, tree.Source) == s {
			n := tree.allocate(Node{x: x})
			tree.splitInto(*p, i, &n.l, &n.r)
			n.s = i
			*p = n
			return
		}
		tree.copy(p)     //
		if i <= (*p).s { //
			s = (*p).s  //
			(*p).s++    //
			p = &(*p).l //
		} else {
			s -= (*p).s + 1 //
			i -= (*p).s + 1 //
			p = &(*p).r     //
		}
	}
}

func (tree *Randomized) Delete(i Position) Data {
	assert(i < tree.Size())
	p := &tree.root
	s := tree.size
	tree.size--
	for {
		tree.copy(p)
		if i == (*p).s {
			defer tree.release(*p)
			x := (*p).x
			*p = tree.join2((*p).l, (*p).r, (*p).s, s-(*p).s-1)
			return x
		}
		if i < (*p).s {
			s = (*p).s
			(*p).s--
			p = &(*p).l
		} else {
			s -= (*p).s + 1
			i -= (*p).s + 1
			p = &(*p).r
		}
	}
}

func (tree *Randomized) join2(l *Node, r *Node, sl, sr Size) (root *Node) {
	p := &root
	for {
		if l == nil {
			*p = r
			return
		}
		if r == nil {
			*p = l
			return
		}

		if random.LessThan(sl+sr, tree.Source) < sl {
			tree.copy(&l)
			sl = sl - l.s - 1
			*p = l
			p = &l.r
			l = *p
		} else {
			tree.copy(&r)
			sr = r.s
			r.s = r.s + sl
			*p = r
			p = &r.l
			r = *p
		}
	}
}

func (tree *Randomized) Select(i Size) Data {
	assert(i < tree.Size())
	return tree.lookup(tree.root, i)
}

func (tree *Randomized) Update(i Size, x Data) {
	assert(i < tree.Size())
	tree.copy(&tree.root)
	tree.update(tree.root, i, x)
}

func (tree *Randomized) splitInto(p *Node, i uint64, l, r **Node) {
	for p != nil {
		tree.copy(&p)
		if i <= p.s {
			*r = p
			p.s = p.s - i
			r = &p.l
			p = p.l
		} else {
			*l = p
			i = i - p.s - 1
			l = &p.r
			p = p.r
		}
	}
	*l = nil
	*r = nil
}

func (tree *Randomized) split(i Size) (Tree, Tree) {
	assert(i <= tree.Size())
	tree.share(tree.root)
	var l, r *Node
	tree.splitInto(tree.root, i, &l, &r)

	return Tree{arena: tree.arena, root: l, size: i},
		Tree{arena: tree.arena, root: r, size: tree.size - i}
}

func (tree *Randomized) Split(i Position) (List, List) {
	l, r := tree.split(i)
	return &Randomized{l, tree.Source},
		&Randomized{r, tree.Source}
}

// TODO check that the random source is independent
func (tree *Randomized) Join(that List) List { // TODO check if benchmarks are affected by poointer receivers here
	l := tree
	r := that.(*Randomized)
	tree.share(l.root)
	tree.share(r.root)
	return &Randomized{
		Tree{
			arena: tree.arena,
			root:  l.join2(l.root, r.root, l.size, r.size),
			size:  l.size + r.size,
		},
		tree.Source,
	}
}

func (tree *Randomized) Verify() {
	tree.verifySizes()
}
