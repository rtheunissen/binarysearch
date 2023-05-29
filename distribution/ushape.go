package distribution

type UShape struct { // TODO: arcsine?
   Beta
}

func (UShape) New(seed uint64) Distribution {
   return &UShape{Beta{a: 0.5, b: 0.5}.Seed(seed)}
}