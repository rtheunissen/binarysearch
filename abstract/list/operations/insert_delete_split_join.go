package operations

import (
   "bst/abstract/list"
   "bst/utility/random"
   "bst/utility/random/distribution"
)

type InsertDeleteSplitJoin struct {
}

func (operation *InsertDeleteSplitJoin) Setup(strategy list.List, scale list.Size) list.List {
   return strategy.New()
}

func (InsertDeleteSplitJoin) Valid(instance list.List, scale list.Size) bool {
   return true
}

func (operation *InsertDeleteSplitJoin) Update(instance list.List, number distribution.Distribution) (list.List, list.Position) {
   var i list.Size
   if instance.Size() == 0 {
      instance.Insert(0, 0)
      return instance, 0
   }
   switch random.Uint64() % 4 {
   case 0: fallthrough
   case 1:
      i = number.LessThan(instance.Size()+1)
      instance.Insert(i, 0)
   case 2:
      i = number.LessThan(instance.Size())
      instance.Delete(i)
   case 3:
      i = number.LessThan(instance.Size())
      l, r := instance.Split(i)
      instance = r.Join(l)
   }
   return instance, i
}

type InsertDeleteSplitJoinPersistent struct {
   InsertDeleteSplitJoin
}

func (operation InsertDeleteSplitJoinPersistent) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
   return operation.InsertDeleteSplitJoin.Update(instance.Clone(), dist)
}
