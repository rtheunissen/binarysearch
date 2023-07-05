package distribution

import "bst/utility/random"

type Median struct {
}

func (Median) New(uint64) random.Distribution {
   return Median{}
}

func (Median) LessThan(n uint64) uint64 {
   return n / 2
}
