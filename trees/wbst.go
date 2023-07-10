package trees

import (
   "bst/abstract/list"
   "math/big"
)

type WeightBalance interface {
   isBalanced(x, y list.Size) bool
   singleRotation(x, y list.Size) bool
}

type ThreeTwo struct {

}

func (ThreeTwo) isBalanced(x, y list.Size) bool {
   return 3 * (x + 1) >= (y + 1)
}

func (ThreeTwo) singleRotation(x, y list.Size) bool {
   return 2 * (x + 1) > (y + 1)
}

type Rational struct {
   Delta *big.Rat
   Gamma *big.Rat
   Cache map[[2]list.Size]bool
}

func (rat Rational) isBalanced(x, y list.Size) bool {
   if x >= y {
      return true
   }
   if rat.Cache == nil {
      rat.Cache = map[[2]list.Size]bool{}
   }
   key := [2]list.Size{x, y}
   if balanced, cached := rat.Cache[key]; cached {
      return balanced
   } else {
      var a big.Rat
      var b big.Rat
      a.SetUint64(x + 1)
      b.SetUint64(y + 1)
      balanced = a.Mul(rat.Delta, &a).Cmp(&b) >= 0
      rat.Cache[key] = balanced
      return balanced
   }
}

func (rat Rational) singleRotation(x, y list.Size) bool {
   if (x + 1) >= (y + 1) {
      return true
   }
   var a, b big.Rat
   a.SetUint64(x + 1)
   b.SetUint64(y + 1)
   single := a.Mul(rat.Gamma, &a).Cmp(&b) > 0
   return single
}

