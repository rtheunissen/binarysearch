package list

//
type Size = uint64

type Position = uint64

type Data = uint64

type Traversal interface {

   Size() Size

   Each(func(Data))
}

// https://en.wikipedia.org/wiki/List_(abstract_data_type)
type List interface {
   Traversal

   Verify()

   New() List

   Clone() List

   Free()

   Size() Size

   Select(Position) Data

   Update(Position, Data)

   Insert(Position, Data)

   Delete(Position) Data

   Split(Position) (List, List)

   Join(List) List
}

func Sample(impl List, size Size) List {
   instance := impl.New()
   for i := Size(0); i < size; i++ {
      instance.Insert(instance.Size(), i)
   }
   return instance
}

