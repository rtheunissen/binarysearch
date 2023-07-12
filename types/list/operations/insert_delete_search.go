package operations

import (
   "bst/types/list"
   "bst/utility/random"
)

type InsertDeleteSearch struct {
   random.Source
}

func (operation *InsertDeleteSearch) Setup(strategy list.List, scale list.Size) list.List {
   operation.Source = random.New(scale)
   return strategy.New()
}

func (InsertDeleteSearch) Valid(instance list.List, scale list.Size) bool {
   return true
}

func (operation *InsertDeleteSearch) Update(instance list.List, number random.Distribution) (list.List, list.Position) {
   var i list.Size
   if instance.Size() == 0 {
      instance.Insert(0, 0)
      return instance, 0
   }
   switch random.Uint64() % 5 {
   case 0: fallthrough
   case 1:
      i = number.LessThan(instance.Size()+1)
      instance.Insert(i, 0)
   case 2:
      i = number.LessThan(instance.Size())
      instance.Delete(i)
   case 3:
      i = number.LessThan(instance.Size())
      instance.Select(i)
   case 4:
      i = number.LessThan(instance.Size())
      instance.Update(i, 0)
   }
   return instance, i
}
type InsertDeleteSearchPersistent struct {
   InsertDeleteSearch
}

func (operation InsertDeleteSearchPersistent) Update(instance list.List, dist random.Distribution) (list.List, list.Position) {
   return operation.InsertDeleteSearch.Update(instance.Clone(), dist)
}
