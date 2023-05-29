package utility

import (
   "golang.org/x/exp/constraints"
)

func Difference[T constraints.Integer](a, b T) T {
   if a > b {
      return a - b
   } else {
      return b - a
   }
}
