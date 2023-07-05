package distribution

import "bst/utility/random"

type Maximum struct {
}

func (Maximum) New(uint64) random.Distribution {
   return Maximum{}
}

func (Maximum) LessThan(n uint64) uint64 {
   if n == 0 {
      return n
   } else {
      return n - 1
   }
}
