package distribution

import (
	"binarysearch/random"
	"golang.org/x/exp/rand"
)

type Normal struct {
	rand *rand.Rand
	mean float64
	sdev float64
}

func (dist Normal) New(seed uint64) Distribution {
	dist.rand = rand.New(random.New(seed))
	dist.mean = 0.50
	dist.sdev = 0.15
	return &dist
}

func (dist *Normal) LessThan(n uint64) uint64 {
	if n == 0 {
		panic("n must be > 0")
	}
	mean := dist.mean
	sdev := dist.sdev
	for {
		if v := dist.rand.NormFloat64()*sdev + mean; v >= 0 && v < 1.0 {
			return uint64(v * float64(n))
		}
	}
}
