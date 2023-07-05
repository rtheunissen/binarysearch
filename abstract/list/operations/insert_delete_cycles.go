package operations

import (
   "bst/abstract/list"
   "bst/utility/random"
)

type InsertDeleteCycles struct {
   random.Source
   inserting bool
   scale     uint64
   cycles    uint64
}

func (operation *InsertDeleteCycles) Valid(instance list.List, scale list.Size) bool {
   return true
}

func (operation *InsertDeleteCycles) Setup(strategy list.List, scale list.Size) list.List {
   operation.Source = random.New(scale)
   operation.inserting = true
   operation.cycles = 10

   if operation.scale > operation.cycles {
      operation.scale = scale / operation.cycles
   } else {
      operation.scale = scale
   }
   return strategy.New() // TODO: do any operations actually use the strategy? Maybe new is worthless and we just seed random as needed in the lists
}

func (operation *InsertDeleteCycles) Update(instance list.List, dist random.Distribution) (list.List, list.Position) {
   //
   // Inserting when empty.
   //
   if instance.Size() == 0 {
      operation.inserting = true
   }
   //
   // Start deleting when at scale, which completes one iteration.
   //
   if instance.Size() != 0 && instance.Size() == operation.scale {
      operation.inserting = false
   }
   //
   // Start inserting again after deleting to half.
   //
   if operation.inserting == false && instance.Size() <= operation.scale/2 {
      operation.inserting = true
   }
   //
   //
   //
   var i list.Position
   if operation.inserting {
      //
      // Insert by the access distribution.
      //
      i = dist.LessThan(instance.Size() + 1)
      instance.Insert(i, 0)
   } else {
      //
      // Delete by reflection of the access distribution.
      //
      i = instance.Size() - dist.LessThan(instance.Size()) - 1
      instance.Delete(i)
   }
   return instance, i
}

type InsertDeleteCyclesPersistent struct {
   InsertDeleteCycles
}

func (operation *InsertDeleteCyclesPersistent) Update(list list.List, access random.Distribution) (list.List, list.Position) {
   return operation.InsertDeleteCycles.Update(list.Clone(), access)
}
