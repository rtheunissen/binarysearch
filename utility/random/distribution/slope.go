package distribution

import "bst/utility/random"

type Slope struct {
   Beta
}

func (Slope) New(seed uint64) random.Distribution {
   return &Slope{Beta{a: 5.0, b: 1.0}.Seed(seed)}
}