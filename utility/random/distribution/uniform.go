package distribution

import (
   random2 "bst/utility/random"
)

type Uniform struct {
   random2.Source
}

func (uniform Uniform) New(seed uint64) random2.Distribution {
   uniform.Source = random2.New(seed)
   return &uniform
}

func (uniform *Uniform) LessThan(n uint64) uint64 {
   if n == 0 {
      panic("n must be > 0")
   }
   return random2.LessThan(n, uniform.Source)
}
