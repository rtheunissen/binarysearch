package operations

import (
   "bst/abstract/list"
   "bst/utility/random/distribution"
)

type Insert struct {
}

func (operation *Insert) Setup(strategy list.List, scale list.Size) list.List {
   return strategy.New()
}

func (Insert) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
   i := dist.LessThan(instance.Size() + 1)
   x := list.Data(0)
   instance.Insert(i, x)
   return instance, i
}

func (Insert) Valid(instance list.List, scale list.Size) bool {
   return instance.Size() < scale
}

type InsertPersistent struct {
   Insert
}

func (operation InsertPersistent) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
   return operation.Insert.Update(instance.Clone(), dist)
}
