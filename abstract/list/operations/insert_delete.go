package operations

import (
   "bst/abstract/list"
   "bst/utility/random/distribution"
)

type InsertDelete struct {
}

func (operation *InsertDelete) Setup(strategy list.List, scale list.Size) list.List {
   return strategy.New()
}

func (InsertDelete) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
   i := dist.LessThan(instance.Size() + 1)
   x := list.Data(0)
   instance.Insert(i, x)
   instance.Insert(i, x)
   instance.Delete(i)
   return instance, i
}

func (InsertDelete) Valid(instance list.List, scale list.Size) bool {
   return true
}

type InsertDeletePersistent struct {
   InsertDelete
}

func (operation InsertDeletePersistent) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
   return operation.InsertDelete.Update(instance.Clone(), dist)
}
