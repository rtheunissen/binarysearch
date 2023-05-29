package utility

import "golang.org/x/exp/constraints"

func Min[T constraints.Integer](x, y T) T {
   if x < y {
      return x
   }
   return y
}


func Max[T constraints.Integer](x, y T) T {
   if x < y {
      return y
   }
   return x
}