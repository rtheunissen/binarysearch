package operations

import (
	"trees/abstract/list"
	"trees/distribution"
	"trees/random"
)

type InsertDelete struct {
	random.Source
}

func (operation *InsertDelete) Setup(strategy list.List, scale list.Size) list.List {
	operation.Source = random.New(scale)
	return strategy.New()
}

func (InsertDelete) Valid(instance list.List, scale list.Size) bool {
	return true
}

func (operation *InsertDelete) Update(instance list.List, number distribution.Distribution) (list.List, list.Position) {
	i := number.LessThan(instance.Size() + 1)
	instance.Insert(i, 0)
	instance.Insert(i, 0)
	instance.Delete(random.LessThan(instance.Size(), operation.Source))

	return instance, i
}

type InsertDeletePersistent struct {
	InsertDelete
}

func (operation InsertDeletePersistent) Update(instance list.List, dist distribution.Distribution) (list.List, list.Position) {
	return operation.InsertDelete.Update(instance.Clone(), dist)
}
