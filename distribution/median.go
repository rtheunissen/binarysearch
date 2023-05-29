package distribution

type Median struct {
}

func (Median) New(uint64) Distribution {
   return Median{}
}

func (Median) LessThan(n uint64) uint64 {
   return n / 2
}
