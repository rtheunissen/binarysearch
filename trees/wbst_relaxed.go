package trees

import "bst/abstract/list"

type WBSTRelaxed struct {
   Tree
   WeightBalance
}

// Creates a new WBSTRelaxed BST from existing values.
func (WBSTRelaxed) New() list.List {
   return &WBSTRelaxed{
      WeightBalance: ThreeTwo{},
   }
}

func (tree *WBSTRelaxed) Clone() list.List {
   return &WBSTRelaxed{
      WeightBalance: tree.WeightBalance,
      Tree: tree.Tree.Clone(),
   }
}

func (tree *WBSTRelaxed) Verify() {
   tree.verifySizes()
}

//
func (tree *WBSTRelaxed) delete(p **Node, s list.Size, i list.Size) (x list.Data) {
   for {
      tree.persist(p)
      sl := (*p).sizeL()
      sr := (*p).sizeR(s)
      if i < sl {
         s = sl
         (*p).s = sl - 1
         p = &(*p).l
         continue
      }
      if i > sl {
         s = sr
         i = i - sl - 1
         p = &(*p).r
         continue
      }
      x := (*p).x
      *p = tree.join((*p).l, (*p).r, sl, sr)
      return x
   }
}

// Deletes the node at position `i` from the tree.
// Returns the data that was in the deleted value.
func (tree *WBSTRelaxed) Delete(i list.Position) list.Data {
   assert(i < tree.size)
   x := tree.delete(&tree.root, tree.size, i)
   tree.size = tree.size - 1
   return x
}
func (tree *WBSTRelaxed) insert(p **Node, s list.Size, i list.Position, x list.Data) {
   var unbalancedNode **Node    // An unbalanced node along the insertion path.
   var unbalancedSize list.Size // The size of the unbalanced node.
   var height uint64            // The height of the insertion so far.
   // Search with increasing height until the end of the path is reached.
   for {
      // Attach a new node at the end of the path.
      if *p == nil {
         *p = tree.allocate(Node{x: x})

         // Check if a rebuild is required.
         if tree.size < 1 << (height >> 1) {
            tree.rebuild(unbalancedNode, unbalancedSize)
         }
         return
      }
      tree.persist(p)
      height++

      sl := (*p).sizeL()
      sr := (*p).sizeR(s)
      if i <= sl {
         if unbalancedNode == nil && !tree.isBalanced(sr, sl + 1) {
            unbalancedNode = p
            unbalancedSize = s + 1
         }
         p = insertL(*p)
         s = sl

      } else {
         if unbalancedNode == nil && !tree.isBalanced(sl, sr + 1) {
            unbalancedNode = p
            unbalancedSize = s + 1
         }
         p = insertR(*p, &i)
         s = sr
      }
   }
}
// Inserts a value `s` at position `i` in the tree.
func (tree *WBSTRelaxed) Insert(i list.Position, x list.Data) {
   assert(i <= tree.size)
   tree.size = tree.size + 1
   tree.insert(&tree.root, tree.size, i, x)
}

// Determines if the height at which a new node was inserted was too deep, i.e.
// if the new height of the tree exceeds the upper-bound for its size.
//
//        tooDeep := height > 2 * ⌊log₂(size)⌋
//
//     because root.height() <= MaximumPathLength(2 * math.FloorLog2(size))
//
//   height is height + 1 ?
func (tree *WBSTRelaxed) exceeds(size list.Size, height uint64) bool {
   return (1 << ((height + 1) >> 1)) > size
}

func (tree *WBSTRelaxed) balance(p *Node, s list.Size) *Node {
   if s < 4 {
      return p
   }
   if !tree.isBalanced(p.sizeL(), p.sizeR(s)) ||
      !tree.isBalanced(p.sizeR(s), p.sizeL()) {
      p = tree.partition(p, s >> 1)
   }
   p.l = tree.balance(p.l, p.sizeL())
   p.r = tree.balance(p.r, p.sizeR(s))
   return p
}

func (tree *WBSTRelaxed) rebuild(p **Node, s list.Size) {
   *p = tree.balance(*p, s)
}


func (tree *WBSTRelaxed) Split(i list.Size) (list.List, list.List) {
   assert(i <= tree.size)

   tree.share(tree.root)
   l,r := tree.split(tree.root, i)

   return &WBSTRelaxed{WeightBalance: tree.WeightBalance, Tree: Tree{arena: tree.arena, root: l, size: i}},
          &WBSTRelaxed{WeightBalance: tree.WeightBalance, Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree *WBSTRelaxed) joinLr(l, r *Node, sl, sr list.Size) *Node {
   if r == nil {
      return l
   }
   if tree.isBalanced(sr, sl) { // TODO: wrong way around?
      p := tree.deleteMax(&l)
      p.l = l
      p.r = r
      p.s = sl - 1
      return p
   }
   tree.persist(&l)
   l.r = tree.joinlR(l.r, r, sl-l.s-1, sr)
   return l
}

func (tree *WBSTRelaxed) joinlR(l, r *Node, sl, sr list.Size) *Node {
   if l == nil {
      return r
   }
   if tree.isBalanced(sl, sr) { // TODO: wrong way around?
      p := tree.deleteMin(&r)
      p.l = l
      p.r = r
      p.s = sl
      return p
   }
   tree.persist(&r)
   r.l = tree.joinLr(l, r.l, sl, r.s) // wtf shouldn't this be joinlR
   r.s = sl + r.s
   return r
}

func (tree *WBSTRelaxed) join(l, r *Node, sl, sr list.Size) *Node {
   if sl > sr {
      return tree.joinLr(l, r, sl, sr)
   } else {
      return tree.joinlR(l, r, sl, sr)
   }
}

func (tree *WBSTRelaxed) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*WBSTRelaxed).root)
   return &WBSTRelaxed{
      WeightBalance: tree.WeightBalance,
      Tree: Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, other.(*WBSTRelaxed).root, tree.size, other.(*WBSTRelaxed).size),
         size:  tree.size + other.(*WBSTRelaxed).size,
      },
   }
}
