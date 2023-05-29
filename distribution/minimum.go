package distribution

type Minimum struct {
}

func (Minimum) New(uint64) Distribution {
   return Minimum{}
}

func (Minimum) LessThan(uint64) uint64 {
   return 0
}
