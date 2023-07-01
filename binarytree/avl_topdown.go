package binarytree

import (
   "binarysearch/abstract/list"
)

// AVLTopDown
//
// This a standard, recursive, bottom-up implementation of an AVL tree
// using the rank-balanced framework of Haeupler, Sen, and Tarjan [].
//
// Using ranks makes it easy to reason about the height of each subtree and
// provides an intuitive way to adjust ranks after rotations. Balancing is
// annotated in one direction only since the algorithms are symmetrical.
//
// A choice was made to not unify the symmetric cases using the direction-based
// technique of Ben Pfaff and others because it makes the logic more difficult
// to follow even though there would be less code overall.
//
// Storing ranks rather than rank differences takes up an entire integer field,
// but it makes `join` easier to implement and is consistent with the other
// rank-balanced implementations. It is possible to store only rank differences
// instead of ranks to use only 1 bit for the balancing information, if needed.
//
type AVLTopDown struct {
   AVLBottomUp
}

func (tree *AVLTopDown) New() list.List {
   return &AVLTopDown{}
}

func (tree *AVLTopDown) Clone() list.List {
   return &AVLTopDown{
      AVLBottomUp{
         Tree: tree.Tree.Clone(),
      },
   }
}

func (tree *AVLTopDown) Insert(i list.Position, x list.Data) {
   tree.insert(&tree.root, i, x)
   tree.size = tree.size + 1
}

func (tree *AVLTopDown) insert(p **Node, i list.Position, x list.Data) {
   //fmt.Println("======================================", i)
   //fmt.Println(i)
   //fmt.Println()

   if *p == nil {
      *p = tree.allocate(Node{x: x})
      return
   }
   tree.persist(p)

   g := p // safe node
   d := i

   //(*p).Draw(os.Stderr)
   for {
      if i <= (*p).sizeL() {
         //fmt.Println("LEFT")
         //(*p).Draw(os.Stderr)


         (*p).s++

         // LEFT
         if (*p).l == nil {
            //fmt.Println("Attach to the left of p at", (*p).y)
            //(*p).s++
            (*p).l = tree.allocate(Node{x: x})
            //(*p).Draw(os.Stderr)

            // TODO: Update balance factors
            //fmt.Println("Update balance factors from", (*g).y)

            for q := *g; q != (*p).l; {
               tree.promote(q)
               if d < q.sizeL() { // Not sure if <= or <
                  q = q.l
               } else {
                  d = d - q.s - 1 // Not sure if this should not have the -1
                  q = q.r
               }
            }

            //(*p).Draw(os.Stderr)
            // TODO: Rebalance at s
            //fmt.Println("rebalance at", (*g).y, "?")

            //*g = tree.balanceInsertL(*g)
            //*g = tree.balanceInsertR(*g)
            *g = tree.balance(*g)
            //*g = tree.balance(*g)

            //fmt.Println("done?")
            //(*p).Draw(os.Stderr)
            // break
            return

         }
         p = &(*p).l
         tree.persist(p)
         if !tree.isOneOne(*p) {
            if tree.isOneChild((*p), (*p).l) {
               //
            } else {
               //
            }
            if tree.isOneChild((*p), (*p).r) {
               //
            } else {
               //
            }
            //fmt.Println("p.l is 1,2 so set safe node to &p.l which has rank", (*p).y)
            g = p
            d = i
         } else {
            if tree.isOneChild((*p), (*p).l) {
               //
            } else {
               //
            }
            if tree.isOneChild((*p), (*p).r) {
               //
            } else {
               //
            }
         }
      } else {
         // RIGHT
         //(*p).Draw(os.Stderr)

         if (*p).r == nil {
            //fmt.Println("Attach to the right of p at", (*p).y)
            (*p).r = tree.allocate(Node{x: x})
            //(*p).Draw(os.Stderr)

            // TODO: Update balance factors
            //fmt.Println("Update balance factors from", (*g).y)

            for q := *g; q != (*p).r; {
               tree.promote(q)
               if d < q.sizeL() { // Not sure if <= or <
                  q = q.l
               } else {
                  d = d - q.s - 1 // Not sure if this should not have the -1
                  q = q.r
               }
            }

            //(*p).Draw(os.Stderr)
            // TODO: Rebalance at s
            //fmt.Println("rebalance at", (*g).y, "?")

            //*g = tree.balanceInsertL(*g)
            //*g = tree.balanceInsertR(*g)
            *g = tree.balance(*g)
            //*g = tree.balance(*g)

            //fmt.Println("done?")
            //(*p).Draw(os.Stderr)
            // break
            return

         }
         i = i - (*p).s - 1
         p = &(*p).r
         tree.persist(p)
         if !tree.isOneOne(*p) {
            //fmt.Println("p.r is 1,2 so set safe node to &p.r which has rank", (*p).y)
            g = p
            d = i
         }
      }
   }
}

func (tree *AVLTopDown) balance(p *Node) *Node {
   if tree.isTwoTwo(p) {
      tree.demote(p)
      return p
   }
   if tree.isThreeChild(p, p.l) {
      if tree.isTwoChild(p.r, p.r.r) {
         tree.rotateRL(&p)
         tree.promote(p)
         tree.demote(p.r)
         tree.demote(p.l)
         tree.demote(p.l)
      } else {
         tree.rotateL(&p)
         tree.demote(p.l)
         tree.demote(p.l)
      }
   }
   if tree.isThreeChild(p, p.r) {
      if tree.isTwoChild(p.l, p.l.l) {
         tree.rotateLR(&p)
         tree.promote(p)
         tree.demote(p.l)
         tree.demote(p.r)
         tree.demote(p.r)
      } else {
         tree.rotateR(&p)
         tree.demote(p.r)
         tree.demote(p.r)
      }
   }
   return p
}

func (tree *AVLTopDown) Join(other list.List) list.List {
   tree.share(tree.root)
   tree.share(other.(*AVLTopDown).root)
   return &AVLTopDown{
      AVLBottomUp{
         Tree: Tree{
            arena: tree.arena,
            root:  tree.join(tree.root, other.(*AVLTopDown).root, tree.size),
            size:  tree.size + other.(*AVLTopDown).size,
         },
      },
   }
}

func (tree *AVLTopDown) Split(i list.Position) (list.List, list.List) {
   l, r := tree.AVLBottomUp.Split(i)
   return &AVLTopDown{*l.(*AVLBottomUp)},
          &AVLTopDown{*r.(*AVLBottomUp)}
}
