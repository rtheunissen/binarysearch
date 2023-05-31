package operations

import (
   "binarysearch/abstract/list"
   "binarysearch/distribution"
   "binarysearch/random"
)

type SplitJoin struct {
}

func (operation *SplitJoin) Setup(strategy list.List, scale list.Size) list.List {
	instance := strategy.New()
	for instance.Size() < scale {
		instance.Insert(random.LessThan(instance.Size()+1, random.Uniform()), 0)
	}
	return instance
}

func (operation *SplitJoin) Valid(instance list.List, scale list.Size) bool {
	return true
}

func (operation *SplitJoin) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
	i := dist.LessThan(instance.Size() + 1)
	l, r := instance.Split(i)
	return r.Join(l), i
}
