package distribution

import "trees/random"

type Queue struct {
	random.Source
}

func (dist Queue) New(seed uint64) Distribution {
	dist.Source = random.New(seed)
	return &dist
}

func (dist *Queue) LessThan(n uint64) uint64 {
	if n == 0 {
		panic("n must be > 0") // TODO: can we maybe remove these?
	}
	return (n - 1) * (dist.Uint64() & 1 /* 0 or 1 */)
}
