package distribution

type Ascending struct {
   i uint64
}

func (Ascending) New(seed uint64) Distribution {
   return &Ascending{i: 0}
}

func (dist *Ascending) LessThan(n uint64) uint64 {
   if dist.i >= n {
      dist.i = 0
      return 0
   }
   i := dist.i
   dist.i++
   return i
}
