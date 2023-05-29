package distribution

type Descending struct {
   i uint64
}

func (Descending) New(seed uint64) Distribution {
   return &Descending{i: seed}
}

func (dist *Descending) LessThan(n uint64) uint64 {
   if dist.i > n || dist.i == 0 {
      dist.i = n
   }
   dist.i--
   return dist.i
}
