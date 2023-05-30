package binarytree

import . "trees/abstract/list"

type RedBlackRelaxedTopDown struct {
	RedBlackRelaxed
}

func (RedBlackRelaxedTopDown) New() List {
	return &RedBlackRelaxedTopDown{}
}

func (tree *RedBlackRelaxedTopDown) Clone() List {
	return &RedBlackRelaxedTopDown{
		RedBlackRelaxed: RedBlackRelaxed{
			Tree: tree.Tree.Clone(),
		},
	}
}

// This top-down insertion algorithm was translated and paraphrased from the
// _Deletion Without Rebalancing in Binary Search Trees_ paper referenced above.
func (tree *RedBlackRelaxedTopDown) insert(p **Node, i Position, x Data) {
	//
	// "If the tree is empty, create a new node of rank zero containing the item
	//  to be inserted and make it the root, completing the insertion."
	//
	if *p == nil {
		tree.attach(p, x)
		return
	}
	tree.copy(p)
	//
	// "Otherwise, promote the root if 0,0."
	//
	if isZeroZero(*p) {
		promote(*p)
	}
	//
	// "This establishes the invariant for the main loop of the algorithm:
	//  *p is a non-nil node that is not a 0,0-node and not a 0-child.
	for {
		assert(!isZeroZero(*p))
		//
		// "From *p, take one step down the search path..."
		//
		if i <= (*p).s {
			//
			// LEFT
			//
			// "If the next node on the search path is nil, replace it by a new
			//  node of rank 0 containing the item to be inserted. This completes
			//  the insertion: the new node may be a 0-child, but *p is not."
			//
			if (*p).l == nil {
				tree.attachL(*p, x)
				return
			}
			if !isZeroZero((*p).l) && !isZeroChild(*p, (*p).l) {
				tree.pathLeft(&p)
				continue
			}
			if isZeroZero((*p).l) {
				tree.pathLeft(&p)
				promote(*p)
				continue
			}
			// In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child
			assert(isZeroChild(*p, (*p).l))
			assert(!isZeroChild((*p).l, (*p).l.l))
			assert(!isZeroChild((*p).l, (*p).l.l))

			// From y, take one step down the
			// search path to z. If z is null, replace z by a new node of rank 0 containing the item to
			// be inserted; if the new node is a 0-child, do a rotate or double rotate step to restore
			// the rank rule (Figure 6(d), (e), and (f)). This completes the insertion.
			if i <= (*p).l.s {
				//
				// LEFT LEFT
				//
				// "If this node is nil, replace it by a new node containing the
				//  item to be inserted. "
				//
				if (*p).l.l == nil {
					tree.attachLL(*p, x)
					if isZeroChild((*p).l, (*p).l.l) { // or is p.l.rank == 0 ?
						tree.rotateR(p)
					}
					return
				}
				//If z is not a 0,0-node, replace w by y and x by z, completing the
				// step (Figure 6(g)).
				if !isZeroZero((*p).l.l) {
					tree.pathLeft(&p)
					tree.pathLeft(&p)
					continue
				}
				// If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
				// x by z, completing the step (Figure 6(g)).
				if !isOneChild((*p).l, (*p).l.l) {
					tree.pathLeft(&p)
					tree.pathLeft(&p)
					tree.promote(*p)
					continue
				}
				// Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).
				//  (i) Node z is a 0,0-node and a 1-child, and y and z are both left or
				//both right children: promote z, do a rotate step, and replace w by z and x by the child of z along the search
				//path.
				tree.rotateR(p)
				tree.pathLeft(&p)
				tree.promote(*p)
				continue

			} else {
				//
				// LEFT RIGHT
				//
				if (*p).l.r == nil {
					tree.attachLR(*p, x)
					if isZeroChild((*p).l, (*p).l.r) { // or is p.l.rank == 0 ?
						tree.rotateLR(p)
					}
					return
				}
				// In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child.

				//If z is not a 0,0-node, replace w by y and x by z, completing the
				// step (Figure 6(g)).
				if !isZeroZero((*p).l.r) {
					tree.pathLeft(&p)
					tree.pathRight(&p, &i)
					continue
				}
				// If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
				// x by z, completing the step (Figure 6(g)).
				if !isOneChild((*p).l, (*p).l.r) {
					tree.pathLeft(&p)
					tree.pathRight(&p, &i)
					tree.promote(*p)
					continue
				}
				// Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).

				// (j) Node z is a 0,0-node and a 1-child, and exactly one of y and z is a left child: promote z, do a double
				//            //rotate step, replace w by whichever of x and y is on the search path from z after the rotations, and replace x
				//            //by the child of the new w on the search path.
				tree.rotateLR(p)
				tree.promote(*p)
				//
				// "If a double rotation is done, take one further step down the
				//  search path after the rotation. Ths completes the step."
				//
				if i <= (*p).s {
					tree.pathLeft(&p) // LRL
				} else {
					tree.pathRight(&p, &i) // LRR
				}
			}
		} else {
			//
			// RIGHT
			//
			// "If the next node on the search path is nil, replace it by a new
			//  node of rank 0 containing the item to be inserted. This completes
			//  the insertion: the new node may be a 0-child, but *p is not."
			//
			if (*p).r == nil {
				tree.attachR(*p, x)
				return
			}
			if !isZeroZero((*p).r) && !isZeroChild(*p, (*p).r) {
				tree.pathRight(&p, &i)
				continue
			}
			if isZeroZero((*p).r) {
				tree.pathRight(&p, &i)
				promote(*p)
				continue
			}
			// In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child

			// From y, take one step down the
			// search path to z. If z is null, replace z by a new node of rank 0 containing the item to
			// be inserted; if the new node is a 0-child, do a rotate or double rotate step to restore
			// the rank rule (Figure 6(d), (e), and (f)). This completes the insertion.
			if i > (*p).s+(*p).r.s+1 {
				//
				// RIGHT RIGHT
				//
				// "If this node is nil, replace it by a new node containing the
				//  item to be inserted. "
				//
				if (*p).r.r == nil {
					tree.attachRR(*p, x)
					if isZeroChild((*p).r, (*p).r.r) { // or is p.r.rank == 0 ?
						tree.rotateL(p)
					}
					return
				}
				// In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child.

				//If z is not a 0,0-node, replace w by y and x by z, completing the
				// step (Figure 6(g)).
				if !isZeroZero((*p).r.r) {
					tree.pathRight(&p, &i)
					tree.pathRight(&p, &i)
					continue
				}
				// If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
				// x by z, completing the step (Figure 6(g)).
				if !isOneChild((*p).r, (*p).r.r) {
					tree.pathRight(&p, &i)
					tree.pathRight(&p, &i)
					tree.promote(*p)
					continue
				}
				// Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).
				//  (i) Node z is a 0,0-node and a 1-child, and y and z are both left or
				//both right children: promote z, do a rotate step, and replace w by z and x by the child of z along the search
				//path.
				tree.rotateL(p)
				tree.pathRight(&p, &i)
				tree.promote(*p)
				continue

			} else {
				//
				// RIGHT LEFT
				//
				if (*p).r.l == nil {
					tree.attachRL(*p, x)
					if isZeroChild((*p).r, (*p).r.l) { // or is p.l.rank == 0 ?
						tree.rotateRL(p)
					}
					return
				}
				// In the remaining cases, y is a 0-child, and hence neither of its children is a 0-child.

				//If z is not a 0,0-node, replace w by y and x by z, completing the
				// step (Figure 6(g)).
				if !isZeroZero((*p).r.l) {
					tree.pathRight(&p, &i)
					tree.pathLeft(&p)
					continue
				}
				// If z is a 0,0-node but not a 1-child, promote z, and replace w by y and
				// x by z, completing the step (Figure 6(g)).
				if !isOneChild((*p).r, (*p).r.l) {
					tree.pathRight(&p, &i)
					tree.pathLeft(&p)
					tree.promote(*p)
					continue
				}
				// Otherwise (z is a 0,0-node and a 1-child), do a rotate or double rotate step to restore the rank rule (Figure 6(h) and (i), respectively).

				// (j) Node z is a 0,0-node and a 1-child, and exactly one of y and z is a left child: promote z, do a double
				//            //rotate step, replace w by whichever of x and y is on the search path from z after the rotations, and replace x
				//            //by the child of the new w on the search path.
				tree.rotateRL(p)
				tree.promote(*p)
				//
				// "If a double rotation is done, take one further step down the
				//  search path after the rotation. Ths completes the step."
				//
				if i > (*p).s {
					tree.pathRight(&p, &i) // RLR
				} else {
					tree.pathLeft(&p) // RLL
				}
			}
		}
	}
}

func (tree *RedBlackRelaxedTopDown) Select(i Size) Data {
	assert(i < tree.Size())
	return tree.lookup(tree.root, i)
}

func (tree *RedBlackRelaxedTopDown) Update(i Size, x Data) {
	assert(i < tree.Size())
	tree.copy(&tree.root)
	tree.update(tree.root, i, x)
}

func (tree *RedBlackRelaxedTopDown) Insert(i Position, x Data) {
	assert(i <= tree.Size())
	tree.size++
	tree.insert(&tree.root, i, x)
}

func (tree *RedBlackRelaxedTopDown) Join(other List) List {
	return &RedBlackRelaxedTopDown{
		tree.RedBlackRelaxed.Join(other.(*RedBlackRelaxedTopDown).RedBlackRelaxed),
	}
}

func (tree *RedBlackRelaxedTopDown) Split(i Position) (List, List) {
	l, r := tree.RedBlackRelaxed.Split(i)
	return &RedBlackRelaxedTopDown{l},
		&RedBlackRelaxedTopDown{r}
}
