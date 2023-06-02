package distribution

import (
   "binarysearch/random"
)

type Uniform struct {
   random.Source
}

func (uniform Uniform) New(seed uint64) Distribution {
   uniform.Source = random.New(seed)
   return &uniform
}

func (uniform *Uniform) LessThan(n uint64) uint64 {
   if n == 0 {
      panic("n must be > 0")
   }
   return random.LessThan(n, uniform.Source)
}
