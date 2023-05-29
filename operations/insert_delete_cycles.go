package operations

import (
	"trees/abstract/list"
	"trees/distribution"
	"trees/random"
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
	operation.inserting = true
	operation.cycles = 10
	operation.scale = scale / operation.cycles
	operation.Source = random.New(scale)

	// Insert to half.
	instance := strategy.New()
	for instance.Size() < (operation.scale / 2) {
		instance.Insert(random.LessThan(instance.Size()+1, operation.Source), 0)
	}
	return instance
}

func (operation *InsertDeleteCycles) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
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
	var i list.Position
	if operation.inserting {
		//
		// Insert by the access distribution.
		//
		i = dist.LessThan(instance.Size() + 1)
		instance.Insert(i, 0)
	} else {
		//
		// Delete uniformly.
		//
		i = random.LessThan(instance.Size(), operation.Source)
		instance.Delete(i)
	}
	return instance, i
}

type InsertDeleteCyclesPersistent struct {
	InsertDeleteCycles
}

func (operation *InsertDeleteCyclesPersistent) Update(list list.List, access distribution.Distribution) (list.List, list.Position) {
	return operation.InsertDeleteCycles.Update(list.Clone(), access)
}
