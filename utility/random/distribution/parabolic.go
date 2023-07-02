package distribution

type Parabolic struct {
   Beta
}

func (Parabolic) New(seed uint64) Distribution {
   return &Parabolic{Beta{a: 2, b: 2}.Seed(seed)}
}
