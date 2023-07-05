package distribution

import "bst/utility/random"

type BiModal struct {
   Beta
   flip bool
}

func (BiModal) New(seed uint64) random.Distribution {
   return &BiModal{
      Beta: Beta{a: 5, b: 15}.Seed(seed),
   }
}

func (dist *BiModal) LessThan(n uint64) uint64 {
   if n == 0 {
      panic("n must be > 0")
   }
   u := dist.Float64()
   if dist.flip = !dist.flip; dist.flip {
      u = 1.0 - u
   }
   return uint64(u * float64(n - 1))



   //if rand. & 1 == 0 {
   //   return uint64(dist.Beta.Float64(5, 15) * float64(n))
   //} else {
   //   return uint64(dist.Beta.Float64(15, 5) * float64(n))
   //}
}