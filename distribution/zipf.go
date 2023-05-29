package distribution

import (
   "golang.org/x/exp/rand"
   "trees/random"
)

type Zipf struct {
	zipf *rand.Zipf
}

func (dist Zipf) New(seed uint64) Distribution {
	dist.zipf = rand.NewZipf(rand.New(random.New(seed)), 1.25, 1, 100)
	return &dist
}

func (dist *Zipf) LessThan(n uint64) uint64 {
	if n == 0 {
		panic("n must be > 0")
	}
	return uint64((float64(dist.zipf.Uint64()) / float64(100)) * float64(n-1))
}
