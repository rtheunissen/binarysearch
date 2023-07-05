package distribution

import "bst/utility/random"

type Skewed struct {
   Beta
}

func (Skewed) New(seed uint64) random.Distribution {
   return &Skewed{Beta{a: 100, b: 50}.Seed(seed)}
}