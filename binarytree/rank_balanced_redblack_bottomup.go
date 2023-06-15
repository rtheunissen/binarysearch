//
//In a ranked binary tree obeying the red-black rule,
//the 0-children are the red nodes, the 1-children are the black nodes. All missing nodes have rank
//difference 1 and are black. The rank of a node is the number of black nodes on a path from the node
//to a leaf, not counting the node itself: this number is independent of the path. Some authors require
//that the root of a red-black tree be black, others allow it to be either red or black. In our formulation,
//the root has no rank difference, and hence no color. Since all rank differences are 0 or 1, we can
//store the balance information in one bit per node, indicating whether its rank difference is zero (it is
//red) or one (it is black).
//
//[Guibas and Sedgewick 1978],
//
//Red-Black Rule: All rank differences are 0 or 1, and no parent of a 0-child is a 0-child

package binarytree

import (
   . "binarysearch/abstract/list"
   "math"
)

type RedBlackBottomUp struct {
   Tree
}

func (RedBlackBottomUp) New() List {
   return &RedBlackBottomUp{}
}

func (tree *RedBlackBottomUp) Clone() List {
   return &RedBlackBottomUp{
      Tree: tree.Tree.Clone(),
   }
}

func (tree *RedBlackBottomUp) Select(i Size) Data {
   assert(i < tree.Size())
   return tree.lookup(tree.root, i)
}

func (tree *RedBlackBottomUp) Update(i Size, x Data) {
   assert(i < tree.Size())
   tree.copy(&tree.root)
   tree.update(tree.root, i, x)
}

func (tree *RedBlackBottomUp) rank(p *Node) int {
   return p.rank()
}

func (tree *RedBlackBottomUp) Delete(i Position) (x Data) {
   assert(i < tree.size)
   tree.size = tree.size - 1
   tree.root = tree.delete(tree.root, i, &x)
   return x
}


func (tree RedBlackBottomUp) fixDeleteL(p *Node) *Node {
   tree.copy(&p)

   //fmt.Println(" fixDeleteL where p is: ---------------")
   //p.Draw(os.Stdout)
   if p.l == nil && p.r == nil {
      return p
   }
   //fmt.Println("a")
   if tree.isBlack(p, p.l) {
      //fmt.Println("b")

      if isTwoChild(p, p.l) {
         //fmt.Println("c")

         if tree.isBlack(p, p.r) && tree.isRed(p.r, p.r.r) {
            //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS LEFT CHILD, AND ITS RIGHT NEPHEW IS RED")
            //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
            tree.rotateL(&p)
            demote(p.l)
            promote(p)

            if tree.isRed(p, p.r) {
               //fmt.Println("DOUBLE BLACK NODE HAS RED SIBLING - ROTATE TREE TO MAKE SIBLING BLACK")
               tree.rotateL(&p)

               if tree.isBlack(p.l, p.l.r) && tree.isBlack(p.l.r, p.l.r.r) {
                  //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS LEFT CHILD, AND ITS RIGHT NEPHEW IS BLACK")
                  //fmt.Println("ROTATE TREE TO MAKE OPPOSITE NEPHEW RED")
                  tree.copy(&p.l)
                  tree.rotateR(&p.l.r)
               }

               if tree.isBlack(p.l, p.l.r) && tree.isRed(p.l.r, p.l.r.r) {
                  //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS LEFT CHILD, AND ITS RIGHT NEPHEW IS RED")
                  //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
                  tree.copy(&p.l)
                  tree.rotateL(&p.l)
                  demote(p.l.l)
                  promote(p.l)
               }

            }

         } else {
            //fmt.Println("c2")



            if tree.isBlack(p, p.r) && tree.isBlack(p.r, p.r.r) && tree.isBlack(p.r, p.r.l) {
               //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING AND TWO BLACK NEPHEWS - PUSH UP BLACK LEVEL")
               demote(p)

            } else {

               if tree.isRed(p, p.r) {
                  //fmt.Println("DOUBLE BLACK NODE HAS RED SIBLING - ROTATE TO MAKE SIBLING BLACK")
                  tree.rotateL(&p)
                  //demote(p.r)
                  //promote(p)

                  if tree.isBlack(p.l, p.l.r) && tree.isBlack(p.l.r, p.l.r.r) && tree.isBlack(p.l.r, p.l.r.l) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING AND TWO BLACK NEPHEWS - PUSH UP BLACK LEVEL")
                     tree.copy(&p.l)
                     demote(p.l)

                  } else {

                     if tree.isBlack(p.l, p.l.r) && tree.isBlack(p.l.r, p.l.r.r) {
                        //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS LEFT CHILD, AND ITS RIGHT NEPHEW IS BLACK")
                        //fmt.Println("ROTATE TREE TO MAKE OPPOSITE NEPHEW RED")
                        tree.copy(&p.l)
                        tree.copy(&p.l.r)
                        tree.rotateR(&p.l.r)
                     }

                     if tree.isBlack(p.l, p.l.r) && tree.isRed(p.l.r, p.l.r.r) {
                        //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS LEFT CHILD, AND ITS RIGHT NEPHEW IS RED")
                        //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
                        tree.copy(&p.l)
                        tree.rotateL(&p.l)
                        promote(p.l)
                        demote(p.l.l)
                     }
                  }

               } else {

                  if tree.isBlack(p, p.r) && tree.isBlack(p.r, p.r.r) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS LEFT CHILD, AND ITS RIGHT NEPHEW IS BLACK")
                     //fmt.Println("ROTATE TREE TO MAKE OPPOSITE NEPHEW RED")
                     tree.copy(&p.r)
                     tree.rotateR(&p.r)
                  }

                  if tree.isBlack(p, p.r) && tree.isRed(p.r, p.r.r) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS LEFT CHILD, AND ITS RIGHT NEPHEW IS RED")
                     //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
                     tree.rotateL(&p)
                     promote(p)
                     demote(p.l)
                  }
               }
            }
         }

      } else {
         //fmt.Println("d")

         if p.l != nil && isTwoChild(p.l, p.l.l) {
            //fmt.Println("e")

            if tree.isRed(p.l, p.l.r) {
               //fmt.Println("DOUBLE BLACK NODE HAS RED SIBLING - ROTATE TO MAKE SIBLING BLACK")
               tree.copy(&p.l)
               tree.rotateL(&p.l)
               demote(p.l.l)
               promote(p.l)

               if tree.isBlack(p.l, p.l.r) {
                  //fmt.Println("f")

                  if tree.isBlack(p.l.r, p.l.r.r) && tree.isBlack(p.l.r, p.l.r.l) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING AND TWO BLACK NEPHEWS - PUSH UP BLACK LEVEL")
                     tree.copy(&p.r)
                     tree.copy(&p.r.r)
                     demote(p.r.r)
                     demote(p.r)
                  }
               }
            }
         }
      }
   } else {
      //fmt.Println("?")
      //fmt.Println("LEFT CHILD IS RED")
      if tree.isBlack(p.l, p.l.r) && tree.isBlack(p.l, p.l.l) {
         if isOneOne(p.l.r) && isOneOne(p.l.l) {
            demote(p.l)
         }
      }
   }
   return p
}

// // Symmetric cases for right child
//            sibling = node->parent->left;
//
//            // Case 1: Sibling is red
//            if (sibling->color == 1) {
//                sibling->color = 0;
//                node->parent->color = 1;
//                root = rightRotate(root, node->parent);
//                sibling = node->parent->left;
//            }
//
//            // Case 2: Sibling's children are both black
//            if (sibling->right->color == 0 && sibling->left->color == 0) {
//                sibling->color = 1;
//                node = node->parent;
//            } else {
//                // Case 3: Sibling's left child is black
//                if (sibling->left->color == 0) {
//                    sibling->right->color = 0;
//                    sibling->color = 1;
//                    root = leftRotate(root, sibling);
//                    sibling = node->parent->left;
//                }
//
//                // Case 4: Sibling's left child is red
//                sibling->color = node->parent->color;
//                node->parent->color = 0;
//                sibling->left->color = 0;
//                root = rightRotate(root, node->parent);
//                node = root;
//            }

func (tree RedBlackBottomUp) isBlack(parent, child *Node) bool {
   return !isZeroChild(parent, child)
}
func (tree RedBlackBottomUp) isRed(parent, child *Node) bool {
   return isZeroChild(parent, child)
}

func (tree RedBlackBottomUp) fixDeleteR(p *Node) *Node {
   tree.copy(&p)

   //fmt.Println(" fixDeleteR where p is: ---------------")
   //p.Draw(os.Stdout)
   if p.l == nil && p.r == nil {
      return p
   }
   //fmt.Println("a")
   if tree.isBlack(p, p.r) {
      //fmt.Println("b")

      if isTwoChild(p, p.r) {
         //fmt.Println("c")

         if tree.isBlack(p, p.l) && tree.isRed(p.l, p.l.l) {
            //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS RIGHT CHILD, AND ITS LEFT NEPHEW IS RED")
            //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
            tree.rotateR(&p)
            demote(p.r)
            promote(p)

            if tree.isRed(p, p.l) {
               //fmt.Println("DOUBLE BLACK NODE HAS RED SIBLING - ROTATE TREE TO MAKE SIBLING BLACK")
               tree.rotateR(&p)

               if tree.isBlack(p.r, p.r.l) && tree.isBlack(p.r.l, p.r.l.l) {
                  //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS RIGHT CHILD, AND ITS LEFT NEPHEW IS BLACK")
                  //fmt.Println("ROTATE TREE TO MAKE OPPOSITE NEPHEW RED")
                  tree.copy(&p.r)
                  tree.copy(&p.r.l)
                  tree.rotateL(&p.r.l)
               }

               if tree.isBlack(p.r, p.r.l) && tree.isRed(p.r.l, p.r.l.l) {
                  //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS RIGHT CHILD, AND ITS LEFT NEPHEW IS RED")
                  //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
                  tree.copy(&p.r)
                  tree.rotateR(&p.r)
                  demote(p.r.r)
                  promote(p.r)
               }

            }

         } else {
            //fmt.Println("c2")



            if tree.isBlack(p, p.l) && tree.isBlack(p.l, p.l.l) && tree.isBlack(p.l, p.l.r) {
              //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING AND TWO BLACK NEPHEWS - PUSH UP BLACK LEVEL")
              demote(p)

            } else {

               if tree.isRed(p, p.l) {
                  //fmt.Println("DOUBLE BLACK NODE HAS RED SIBLING - ROTATE TO MAKE SIBLING BLACK")
                  tree.rotateR(&p)
                  //demote(p.r)
                  //promote(p)

                  if tree.isBlack(p.r, p.r.l) && tree.isBlack(p.r.l, p.r.l.l) && tree.isBlack(p.r.l, p.r.l.r) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING AND TWO BLACK NEPHEWS - PUSH UP BLACK LEVEL")
                     tree.copy(&p.r)
                     demote(p.r)

                  } else {

                     if tree.isBlack(p.r, p.r.l) && tree.isBlack(p.r.l, p.r.l.l) {
                        //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS RIGHT CHILD, AND ITS LEFT NEPHEW IS BLACK")
                        //fmt.Println("ROTATE TREE TO MAKE OPPOSITE NEPHEW RED")
                        tree.copy(&p.r)
                        tree.copy(&p.r.l)
                        tree.rotateL(&p.r.l)
                     }

                     if tree.isBlack(p.r, p.r.l) && tree.isRed(p.r.l, p.r.l.l) {
                        //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS RIGHT CHILD, AND ITS LEFT NEPHEW IS RED")
                        //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
                        tree.copy(&p.r)
                        tree.rotateR(&p.r)
                        promote(p.r)
                        demote(p.r.r)
                     }
                  }

               } else {

                  if tree.isBlack(p, p.l) && tree.isBlack(p.l, p.l.l) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS RIGHT CHILD, AND ITS LEFT NEPHEW IS BLACK")
                     //fmt.Println("ROTATE TREE TO MAKE OPPOSITE NEPHEW RED")
                     tree.copy(&p.l)
                     tree.rotateL(&p.l)
                  }

                  if tree.isBlack(p, p.l) && tree.isRed(p.l, p.l.l) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING, IS RIGHT CHILD, AND ITS LEFT NEPHEW IS RED")
                     //fmt.Println("ONE ROTATION CAN FIX DOUBLE BLACKNESS")
                     tree.rotateR(&p)
                     promote(p)
                     demote(p.r)
                  }
               }
            }
         }

      } else {
         //fmt.Println("d")

         if p.r != nil && isTwoChild(p.r, p.r.r) {
            //fmt.Println("e")

            if tree.isRed(p.r, p.r.l) {
               //fmt.Println("DOUBLE BLACK NODE HAS RED SIBLING - ROTATE TO MAKE SIBLING BLACK")
               tree.copy(&p.r)
               tree.rotateR(&p.r)
               demote(p.r.r)
               promote(p.r)

               if tree.isBlack(p.r, p.r.l) {
                  //fmt.Println("f")

                  if tree.isBlack(p.r.l, p.r.l.l) && tree.isBlack(p.r.l, p.r.l.r) {
                     //fmt.Println("DOUBLE BLACK NODE HAS BLACK SIBLING AND TWO BLACK NEPHEWS - PUSH UP BLACK LEVEL")
                     tree.copy(&p.r)
                     tree.copy(&p.r.r)
                     demote(p.r.r)
                     demote(p.r)
                  }
               }
            }
         }
      }
   } else {
      //fmt.Println("RIGHT CHILD IS RED")
      if tree.isBlack(p.r, p.r.l) && tree.isBlack(p.r, p.r.r) {
         if isOneOne(p.r.l) && isOneOne(p.r.r) {
            demote(p.r)
         }
      }
   }
   return p
}

func (tree *RedBlackBottomUp) delete(p *Node, i Position, x *Data) *Node {
   tree.copy(&p)
   if i == p.s {
      *x = p.x
      defer tree.release(p)
      return tree.join(p.l, p.r, p.s)
   }
   if i < p.s {
      p.s = p.s - 1
      p.l = tree.delete(p.l, i, x)
      return tree.fixDeleteL(p)
   } else {
      p.r = tree.delete(p.r, i-p.s-1, x)
      return tree.fixDeleteR(p)
   }
   //tree.Draw(os.Stdout)




   //if p.isLeaf() && isTwoTwo(p) {
   //   demote(p)
   //   return p
   //}
   //if isThreeChild(p, p.r) {
   //   if isTwoChild(p, p.l) {
   //      demote(p)
   //
   //   } else if isTwoTwo(p.l) {
   //      demote(p.l)
   //      demote(p)
   //
   //   } else if isOneChild(p.l, p.l.l) {
   //      tree.rotateR(&p)
   //      promote(p)
   //      demote(p.r)
   //
   //      assert(isTwoChild(p, p.l))
   //      assert(isOneChild(p, p.r))
   //
   //      if p.r.l == nil {
   //         assert(isTwoTwo(p.r))
   //         demote(p.r)
   //      }
   //   } else {
   //      tree.rotateLR(&p)
   //      promote(p)
   //      promote(p)
   //      demote(p.l)
   //      demote(p.r)
   //      demote(p.r)
   //
   //      assert(isTwoChild(p, p.l))
   //      assert(isTwoChild(p, p.r))
   //   }
   //} else if isThreeChild(p, p.l) {
   //   if isTwoChild(p, p.r) {
   //      demote(p)
   //
   //   } else if isTwoTwo(p.r) {
   //      demote(p.r)
   //      demote(p)
   //
   //   } else if isOneChild(p.r, p.r.r) {
   //      tree.rotateL(&p)
   //      promote(p)
   //      demote(p.l)
   //
   //      assert(isOneChild(p, p.l))
   //      assert(isTwoChild(p, p.r))
   //
   //      if p.l.r == nil {
   //         assert(isTwoTwo(p.l))
   //         demote(p.l)
   //      }
   //   } else {
   //      tree.rotateRL(&p)
   //      promote(p)
   //      promote(p)
   //      demote(p.l)
   //      demote(p.l)
   //      demote(p.r)
   //
   //      assert(isTwoChild(p, p.l))
   //      assert(isTwoChild(p, p.r))
   //   }
   //}
   //return p
}

func (tree *RedBlackBottomUp) insert(p *Node, i Position, x Data) *Node {
   if p == nil {
      return tree.allocate(Node{x: x})
   }
   tree.copy(&p)
   if i <= p.s {
      p.s = p.s + 1
      p.l = tree.insert(p.l, i, x)
      return tree.fixInsertL(p)
   } else {
      p.r = tree.insert(p.r, i-p.s-1, x)
      return tree.fixInsertR(p)
   }
}

func (tree *RedBlackBottomUp) Insert(i Position, x Data) {
   assert(i <= tree.size)
   tree.size = tree.size + 1
   tree.root = tree.insert(tree.root, i, x)
   return
}

func (tree RedBlackBottomUp) split(p *Node, i, s Size) (l, r *Node) {
   if p == nil {
      return
   }
   tree.copy(&p)

   sl := p.s
   sr := s - p.s - 1

   if i <= (*p).s {
      l, r = tree.split(p.l, i, sl)
      r = tree.build(r, p, p.r, sl-i)
   } else {
      l, r = tree.split(p.r, i-sl-1, sr)
      l = tree.build(p.l, p, l, sl)
   }
   return l, r
}

func (tree RedBlackBottomUp) Split(i Position) (List, List) {
   assert(i <= tree.Size())
   tree.share(tree.root)
   l, r := tree.split(tree.root, i, tree.size)
   return &RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: l, size: i}},
          &RedBlackBottomUp{Tree: Tree{arena: tree.arena, root: r, size: tree.size - i}}
}

func (tree RedBlackBottomUp) build(l, p, r *Node, sl Size) *Node {
   //tree.copy(&l) // ??
   tree.copy(&p) // ??
   //tree.copy(&r) // ??
   if tree.rank(l) == tree.rank(r) {
      p.l = l
      p.r = r
      p.s = sl
      p.y = uint64(p.l.rank() + 1)
      return p
   }
   if tree.rank(l) < tree.rank(r) {
      tree.copy(&r)
      r.s = 1 + sl + r.s
      r.l = tree.build(l, p, r.l, sl)
      return tree.fixInsertL(r)
   } else {
      tree.copy(&l)
      l.r = tree.build(l.r, p, r, sl-l.s-1)
      return tree.fixInsertR(l)
   }
}


func (tree *RedBlackBottomUp) deleteMin(p *Node, min **Node) *Node {
   tree.copy(&p)
   if p.l == nil {
      *min = p
      return p.r
   }
   p.s = p.s - 1
   p.l = tree.deleteMin(p.l, min)
   return tree.fixDeleteL(p)
}

func (tree *RedBlackBottomUp) deleteMax(p *Node, max **Node) *Node {
   tree.copy(&p)
   if p.r == nil {
      *max = p
      return p.l
   }
   p.r = tree.deleteMax(p.r, max)
   return tree.fixDeleteR(p)
}

func (tree RedBlackBottomUp) join(l, r *Node, sl Size) (p *Node) {
   if l == nil { return r }
   if r == nil { return l }
   if tree.rank(l) < tree.rank(r) {
      var min *Node
      r = tree.deleteMin(r, &min)
      return tree.build(l, min, r, sl)
   } else {
      var max *Node
      l = tree.deleteMax(l, &max)
      return tree.build(l, max, r, sl-1)
   }
}

func (tree RedBlackBottomUp) Join(other List) List {
   tree.share(tree.root)
   tree.share(other.(*RedBlackBottomUp).root)
   return &RedBlackBottomUp{
      Tree: Tree{
         arena: tree.arena,
         root:  tree.join(tree.root, other.(*RedBlackBottomUp).root, tree.size),
         size:  tree.size + other.(*RedBlackBottomUp).size,
      },
   }
}

func (tree RedBlackBottomUp) fixInsertL(p *Node) *Node {
   if isZeroChild(p, p.l) {
      if isZeroChild(p, p.r) {
         if isZeroChild(p.l, p.l.l) || isZeroChild(p.l, p.l.r) {
            promote(p)
         }
      } else {
         if isZeroChild(p.l, p.l.l) {
            return p.rotateR()
         }
         if isZeroChild(p.l, p.l.r) {
            return p.rotateLR()
         }
      }
   }
   return p
}

func (tree RedBlackBottomUp) fixInsertR(p *Node) *Node {
   if isZeroChild(p, p.r) {
      if isZeroChild(p, p.l) {
         if isZeroChild(p.r, p.r.r) || isZeroChild(p.r, p.r.l) {
            promote(p)
         }
      } else {
         if isZeroChild(p.r, p.r.r) {
            return p.rotateL()
         }
         if isZeroChild(p.r, p.r.l) {
            return p.rotateRL()
         }
      }
   }
   return p
}

func (tree RedBlackBottomUp) verifyHeight(p *Node, s Size) {
   ////fmt.Println(p.height(), 2 * math.Floor(math.Log2(float64(s + 1))))
   invariant(float64(p.height()) <= 2 * math.Floor(math.Log2(float64(s + 1))))
}

func (tree RedBlackBottomUp) verifyRanks(p *Node) {
   if p == nil {
      return
   }
   invariant(tree.rank(p) >= tree.rank(p.l))
   invariant(tree.rank(p) >= tree.rank(p.r))

   // All rank differences are 0 or 1.
   invariant(isZeroChild(p, p.l) || isOneChild(p, p.l))
   invariant(isZeroChild(p, p.r) || isOneChild(p, p.r))

   // No parent of a 0-child is a 0-child.
   if isZeroChild(p, p.l) {
      invariant(!isZeroChild(p.l, p.l.l))
      invariant(!isZeroChild(p.l, p.l.r))
   }
   if isZeroChild(p, p.r) {
      invariant(!isZeroChild(p.r, p.r.l))
      invariant(!isZeroChild(p.r, p.r.r))
   }
   tree.verifyRanks(p.l)
   tree.verifyRanks(p.r)
}

func (tree RedBlackBottomUp) Verify() {
   tree.Tree.Verify()
   tree.verifyRanks(tree.root)
   tree.verifyHeight(tree.root, tree.size)
}
