package trees

import (
   "bst/types/list"
   "bst/utility/random"
)

type TreapTopDown struct {
   Tree
   random.Source
}

func (TreapTopDown) New() list.List {
   return &TreapTopDown{
      Source: random.New(random.Uint64()),
   }
}

//func (tree *TreapTopDown) randomRank() uint64 {
//   return tree.Uint64()
//}

func (tree *TreapTopDown) Clone() list.List {
   return &TreapTopDown{
      Tree:   tree.Tree.Clone(),
      Source: tree.Source,
   }
}

// def join(pseudo_root, T1, T2):
//    if not T1 or not T2:
//        return T1 if T1 else T2
//
//    pseudo_root.left = T1
//    pseudo_root.right = T2
//
//    curr = pseudo_root
//    while True:
//        if not curr.right:
//            break
//        if curr.right.priority > T2.priority:
//            T1 = curr.right
//            curr.right = T1.left
//            T1.left = None
//            curr = T1
//        else:
//            curr = curr.right
//
//    return pseudo_root

//       if i < p.s {
//         p.s = p.s - i - 1
//         r.l = p
//         r = r.l
//         p = p.l
//      } else {
//         i = i - p.s - 1
//         l.r = p
//         l = l.r
//         p = p.r
//      }
func (tree *TreapTopDown) join2(l, r *Node, sl list.Size) (root *Node) {
   if l == nil { return r }
   if r == nil { return l }

   tree.persist(&l)
   tree.persist(&r)

   n := Node{}
   p := &n
   for {
      if l.y >= r.y {
         tree.persist(&l)
         sl = sl - l.s - 1
         p.l = l
         p = p.l
         if l.r == nil {
            p.r = r
            break
         }
         l = l.r
      } else {
         tree.persist(&r)
         r.s = r.s + sl
         p.r = r
         p = p.r
         if r.l == nil {
            p.l = l
            break
         }
         //tree.persist(&r.l)
         r = r.l
      }
   }
   if n.l == nil {
      return n.r
   } else {
      return n.l
   }
}

func (tree *TreapTopDown) join(l, r *Node, sl list.Size) (root *Node) {
   p := &root
   for {
      if l == nil { *p = r; return }
      if r == nil { *p = l; return }

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
func (tree *TreapTopDown) build(l, p, r *Node, sl, sr list.Size) (root *Node) {
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

func (tree TreapTopDown) delete(p **Node, i list.Position, x *list.Data) {
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

func (tree *TreapTopDown) insert(p **Node, i list.Position, n *Node) {
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

func (tree *TreapTopDown) Insert(i list.Position, x list.Data) {
   assert(i <= tree.size)
   tree.size++
   tree.insert(&tree.root, i, tree.allocate(Node{x: x, y: tree.Source.Uint64() % 10}))
}

func (tree *TreapTopDown) Delete(i list.Position) (x list.Data) {
   assert(i < tree.size)
   tree.delete(&tree.root, i, &x)
   tree.size--
   return
}

func (tree TreapTopDown) split(i list.Size) (Tree, Tree) {
   assert(i <= tree.size)
   tree.share(tree.root)
   l, r := tree.Tree.split(tree.root, i)

   return Tree{arena: tree.arena, root: l, size: i},
      Tree{arena: tree.arena, root: r, size: tree.size - i}
}

func (tree TreapTopDown) Split(i list.Position) (list.List, list.List) {
   assert(i <= tree.size)
   l, r := tree.split(i)
   return &TreapTopDown{Tree: l, Source: tree.Source},
      &TreapTopDown{Tree: r, Source: tree.Source}
}

func (tree *TreapTopDown) Select(i list.Size) list.Data {
   assert(i < tree.size)
   return tree.lookup(tree.root, i)
}

func (tree *TreapTopDown) Update(i list.Size, x list.Data) {
   assert(i < tree.size)
   tree.persist(&tree.root)
   tree.update(tree.root, i, x)
}

// TODO: Figure out a standard naming for "that"

func (tree TreapTopDown) Join(that list.List) list.List {
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
