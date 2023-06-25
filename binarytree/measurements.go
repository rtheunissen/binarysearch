package binarytree

import (
   "golang.org/x/exp/constraints"
)

func measurement[T constraints.Integer | constraints.Float](addr *T, delta T) {
   *addr = *addr + delta
}

type Measurement interface {
   Reset()
   Measure(BinaryTree) any
}

type PartitionCount struct {

}

var partitionCount uint64 = 0

func (PartitionCount) Reset()  {
   partitionCount = 0
}

func (PartitionCount) Measure(BinaryTree) any {
   return partitionCount
}

type PartitionDepth struct {

}
var partitionCost uint64 = 0

func (PartitionDepth) Reset()  {
   partitionCost = 0
}

func (PartitionDepth) Measure(BinaryTree) any {
   return partitionCost
}

type MaximumPathLength struct {

}

func (MaximumPathLength) Reset()  {
}

func (MaximumPathLength) Measure(tree BinaryTree) any {
   return tree.Root().MaximumPathLength()
}

type AveragePathLength struct {
}

func (accumulator *AveragePathLength) Measure(tree BinaryTree) any {
   return tree.Root().AveragePathLength()
}
func (AveragePathLength) Reset()  {
}
var allocations uint64 = 0

type Allocations struct {
}

func (Allocations) Reset()  {
   allocations = 0
}

func (Allocations) Measure(BinaryTree) any {
   return allocations
}

var rotations uint64 = 0

type Rotations struct {
}

func (accumulator *Rotations) Reset()  {
   rotations = 0
}

func (accumulator *Rotations) Measure(BinaryTree) any {
   return rotations
}


//
//
//
//
//type Allocations struct {
//   runtime.MemStats
//}
//func (measurement *Allocations) Reset()  {
//   runtime.ReadMemStats(&measurement.MemStats)
//}
//func (measurement *Allocations) Measure(result abstract.List) string {
//   var stats runtime.MemStats
//   runtime.ReadMemStats(&stats)
//   return strconv.FormatUint(stats.Mallocs - measurement.MemStats.Mallocs, 10)
//}
//
//
//
//
//
//
//
//type Measurements struct {
//   Operation operations.Operation
//   Lists []abstract.List // TODO: Measurable ??
//   Distributions []number.Distribution
//   Measurements []Measurement
//
//   Resolution uint64
//   Scale      uint64
//}
//
//func info(values ...any) {
//   fmt.Fprintln(os.Stderr, values...)
//}
//
//func line(values ...any) {
//   fmt.Fprintln(os.Stdout, values...)
//}
//
//func (ms Measurements) Run() {
//   if ms.Resolution > ms.Scale {
//      ms.Resolution = ms.Scale
//   }
//
//   //
//   header := []any{"#", "x"}
//   for _, measure := range ms.Measurements {
//      header = append(header, utility.TypeName(measure))
//   }
//   line(header...)
//
//   for _, impl := range ms.Lists {
//      line(utility.TypeName(impl))
//      info(utility.TypeName(impl))
//
//      for _, distribution := range ms.Distributions {
//         x := uint64(0)
//         if distribution != nil {
//            distribution = distribution.Seed(1234)
//         }
//         list := ms.Operation.Setup(impl, ms.Scale)
//         for _, measurement := range ms.Measurements {
//            measurement.Reset()
//         }
//         for ms.Operation.Valid(list, ms.Scale) {
//            list = ms.Operation.Update(list, distribution)
//
//            if x++; x % (ms.Scale / ms.Resolution) == 0 {
//               var row []any
//               row = append(row, x)
//               for _, measure := range ms.Measurements {
//                  row = append(row, measure.Measure(list))
//               }
//               line(row...)
//            }
//         }
//      }
//      line()
//      line()
//   }
//}