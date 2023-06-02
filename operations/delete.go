package operations

import (
   "binarysearch/abstract/list"
   "binarysearch/distribution"
   "binarysearch/random"
)

type Delete struct {
}

func (Delete) Setup(strategy list.List, size list.Size) list.List {
   instance := strategy.New()
   for instance.Size() < size {
      instance.Insert(random.LessThan(instance.Size()+1, random.Uniform()), 0)
   }
   return instance
}

func (Delete) Valid(instance list.List, scale list.Size) bool {
   return instance.Size() > 0
}

func (Delete) Update(instance list.List, number distribution.Distribution) (list.List, list.Position) {
   i := number.LessThan(instance.Size())
   instance.Delete(i)
   return instance, i
}

type DeletePersistent struct {
   Delete
}

func (operation *DeletePersistent) Update(instance list.List, number distribution.Distribution) (list.List, list.Position) {
   return operation.Delete.Update(instance.Clone(), number)
}
