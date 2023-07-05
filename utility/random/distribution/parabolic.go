package distribution

import "bst/utility/random"

type Parabolic struct {
   Beta
}

func (Parabolic) New(seed uint64) random.Distribution {
   return &Parabolic{Beta{a: 2, b: 2}.Seed(seed)}
}
