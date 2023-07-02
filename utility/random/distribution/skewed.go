package distribution

type Skewed struct {
   Beta
}

func (Skewed) New(seed uint64) Distribution {
   return &Skewed{Beta{a: 100, b: 50}.Seed(seed)}
}