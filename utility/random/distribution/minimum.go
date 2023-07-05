package distribution

import "bst/utility/random"

type Minimum struct {
}

func (Minimum) New(uint64) random.Distribution {
   return Minimum{}
}

func (Minimum) LessThan(uint64) uint64 {
   return 0
}
