package binarytree
import . "binarysearch/abstract/list"

import (
   "binarysearch/random"
)

type TreapTopDown struct {
   Tree
   random.Source
}

func (TreapTopDown) New() List {
   return &TreapTopDown{
      Source: random.New(random.Uint64()),
   }
}

//func (tree *TreapTopDown) randomRank() uint64 {
//   return tree.Uint64()
//}

func (tree *TreapTopDown) Clone() List {
   return &TreapTopDown{
      Tree:   tree.Tree.Clone(),
      Source: tree.Source,
   }
}

func (tree *TreapTopDown) join(l, r *Node, sl Size) (root *Node) {
   // assert(sl == l.size())
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
      if l.y >= r.y {
         tree.persist(&l)
         sl = sl - l.s - 1
         *p = l
         p = &l.r
         l = *p
      } else {
         tree.persist(&r)
         r.s = r.s + sl
         *p = r
         p = &r.l
         r = *p
      }
   }
}

//
//// TODO: this is identical to zip tree
//func (TreapTopDown) build(x []Data) *TreapTopDown {
//   tree := TreapTopDown{
//      Tree:   Tree{}.New(),
//      Source: random.New(random.Uint64()),
//   }
//   //
//   // Build bottom-up with a reversed right spine.
//   //
//   var p *Node
//   for _, x := range x {
//      //
//      // Node a new node with random rank.
//      //
//      p = tree.allocate(Node{
//         x: x,
//         y: tree.randomRank(),
//         r: p,
//      })
//      //
//      // Rotate up
//      //
//      for p.r != nil && rank(p) > rank(p.r) { // TODO: n.r nil check might be redundant because of the nil check in rank
//         g := p.r // parent on the spine
//         p.r = g.r
//         g.r = p.l
//         p.l = g
//         p.s = p.s + g.s + 1
//      }
//   }
//   //
//   // Reverse the right spine.
//   //
//   var r *Node
//   for p != nil {
//      g := p.r
//      p.r = r
//      r = p
//      p = g
//   }
//   tree.root = r
//   tree.size = Size(len(x))
//   return &tree
//}
func (tree *TreapTopDown) rank(p *Node) uint64 {
   if p == nil {
      return 0
   } else {
      return p.y
   }
}
func (tree *TreapTopDown) build(l, p, r *Node, sl, sr Size) (root *Node) {
   //tree.pathcopy(&p)
   if tree.rank(p) >= tree.rank(l) && tree.rank(p) >= tree.rank(r) {
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   if tree.rank(l) > tree.rank(r) {
      l.r = tree.build(l.r, p, r, sl-l.s-1, sr)
      return l
   } else {
      r.l = tree.build(l, p, r.l, sl, r.s)
      r.s = r.s + sl + 1
      return r
   }
}

func (tree TreapTopDown) delete(p **Node, i Position, x *Data) {
   for {
      if i == (*p).s {
         *x = (*p).x
         if (*p).l == nil && (*p).r == nil {
            *p = nil
         } else {
            tree.persist(p)
            defer tree.free(*p) // TODO what
            *p = tree.join((*p).l, (*p).r, (*p).s)
         }
         return
      }
      tree.persist(p)
      if i < (*p).s {
         p = deleteL(*p)
      } else {
         p = deleteR(*p, &i)
      }
   }
}

func (tree *TreapTopDown) insert(p **Node, i Position, n *Node) {
   for {
      if *p == nil {
         *p = n
         return
      }
      if (*p).y <= n.y {
         n.l, n.r = tree.Tree.split(*p, i)
         n.s = i
         *p = n
         return
      }
      tree.persist(p)
      if i <= (*p).s {
         p = insertL(*p)
      } else {
         p = insertR(*p, &i)
      }
   }
}

func (tree *TreapTopDown) Insert(i Position, x Data) {
   // assert(i <= tree.size)
   tree.size++
   tree.insert(&tree.root, i, tree.allocate(Node{x: x, y: tree.Source.Uint64()}))
}

func (tree *TreapTopDown) Delete(i Position) (x Data) {
   // assert(i < tree.Size())
   tree.delete(&tree.root, i, &x)
   tree.size--
   return
}

func (tree TreapTopDown) split(i Size) (Tree, Tree) {
   // assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := tree.Tree.split(tree.root, i)

   return Tree{arena: tree.arena, root: l, size: i},
      Tree{arena: tree.arena, root: r, size: tree.size - i}
}

func (tree TreapTopDown) Split(i Position) (List, List) {
   // assert(i <= tree.Size())
   l, r := tree.split(i)
   return &TreapTopDown{Tree: l, Source: tree.Source},
      &TreapTopDown{Tree: r, Source: tree.Source}
}

func (tree *TreapTopDown) Select(i Size) Data {
   // assert(i < tree.Size())
   return tree.lookup(tree.root, i)
}

func (tree *TreapTopDown) Update(i Size, x Data) {
   // assert(i < tree.Size())
   tree.persist(&tree.root)
   tree.update(tree.root, i, x)
}

// TODO: Figure out a standard naming for "that"

func (tree TreapTopDown) Join(that List) List {
   tree.share(tree.root)
   tree.share(that.(*TreapTopDown).root)
   return &TreapTopDown{
      Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, that.(*TreapTopDown).root, tree.size),
         size:  tree.size + that.(*TreapTopDown).size,
      },
      tree.Source,
   }
}

func (tree TreapTopDown) verifyMaxRankHeap(p *Node) {
   if p == nil {
      return
   }
   invariant(tree.rank(p) >= tree.rank(p.l))
   invariant(tree.rank(p) >= tree.rank(p.r))

   tree.verifyMaxRankHeap(p.l)
   tree.verifyMaxRankHeap(p.r)
}

func (tree TreapTopDown) Verify() {
   tree.verifySizes()
   tree.verifyMaxRankHeap(tree.root)
}
