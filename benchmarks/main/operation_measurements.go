package main

import (
   "binarysearch/abstract/list"
   "binarysearch/binarytree"
   "binarysearch/distribution"
   "binarysearch/operations"
   "binarysearch/utility"
   "flag"
   "fmt"
   "os"
   "time"
   "path/filepath"
)


type OperationMeasurement struct {
   Scale         int
   Samples       int
   Operation     list.Operation
   Distributions []distribution.Distribution
   Strategies    []list.List
   Measurements  []binarytree.Measurement
}

func main() {
   operation := flag.String("operation", "", "")
   samples := flag.Int("samples", 1_000, "")
   scale := flag.Int("scale", 1_000_000, "")
   flag.Parse()

   OperationMeasurement{
      Scale:   *scale,
      Samples: *samples,
      Operation: utility.Resolve(*operation, []list.Operation{
         &operations.Insert{},
         &operations.InsertPersistent{},
         &operations.InsertDelete{},
         &operations.InsertDeletePersistent{},
         &operations.InsertDeleteCycles{},
         &operations.InsertDeleteCyclesPersistent{},
         &operations.SplitJoin{},
      }),
      Measurements: []binarytree.Measurement{
         &binarytree.PartitionCount{},
         &binarytree.PartitionDepth{},
         &binarytree.AveragePathLength{},
         &binarytree.MaximumPathLength{},
         &binarytree.Rotations{},
         &binarytree.Allocations{},
      },
      Distributions: []distribution.Distribution{
         &distribution.Uniform{},
         &distribution.Normal{},
         &distribution.Skewed{},
         &distribution.Zipf{},
         &distribution.Maximum{},
      },
      Strategies: []list.List{
         &binarytree.AVLBottomUp{},
         &binarytree.AVLJoinBased{},
         &binarytree.AVLWeakTopDown{},
         &binarytree.AVLWeakBottomUp{},
         &binarytree.AVLWeakJoinBased{},
         &binarytree.AVLRelaxedTopDown{},
         &binarytree.AVLRelaxedBottomUp{},
         &binarytree.RedBlackBottomUp{},
         &binarytree.RedBlackRelaxedTopDown{},
         &binarytree.RedBlackRelaxedBottomUp{},
         &binarytree.LBSTBottomUp{},
         &binarytree.LBSTTopDown{},
         &binarytree.LBSTJoinBased{},
         &binarytree.LBSTRelaxed{},
         &binarytree.TreapTopDown{},
         &binarytree.TreapJoinBased{},
         &binarytree.TreapFingerTree{},
         &binarytree.Randomized{},
         &binarytree.Zip{},
         &binarytree.Splay{},
         &binarytree.Conc{},
      },
   }.Run()
}


func (measurement OperationMeasurement) Run() {
   if measurement.Operation == nil {
      return
   }

   //
   for _, strategy := range measurement.Strategies {

      path := fmt.Sprintf(
         "benchmarks/data/operations/measurements/%s/%s",
         utility.NameOf(measurement.Operation),
         utility.NameOf(strategy),
      )
      err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
      if err != nil {
         panic(err)
      }
      file, err := os.Create(path)
      if err != nil {
         panic(err)
      }

      //
      header := []any{
         "Distribution",
         "Scale",
         "Size",
         "Step",
         "Position",
      }
      //
      for _, measurement := range measurement.Measurements {
         header = append(header, utility.NameOf(measurement))
      }
      //
      fmt.Fprintln(file, header...)

      //
      for _, distribution := range measurement.Distributions {
         //
         //
         fmt.Printf("%s %-32s %-32s %-32s\n",
            time.Now().Format(time.RFC822),
            utility.NameOf(measurement.Operation),
            utility.NameOf(strategy),
            utility.NameOf(distribution))

         //
         instance := measurement.Operation.Setup(strategy, list.Size(measurement.Scale))

         //
         access := distribution.New(1)

         //
         step := measurement.Scale / measurement.Samples

         //
         for position := step; position <= measurement.Scale; position = position + step {

            //
            for _, measurement := range measurement.Measurements {
               measurement.Reset()
            }
            //
            for i := 0; i < step; i++ {
               instance, _ = measurement.Operation.Update(instance, access)
            }
            //
            row := []any{
               utility.NameOf(access),
               fmt.Sprint(measurement.Scale),
               fmt.Sprint(instance.Size()),
               fmt.Sprint(step),
               fmt.Sprint(position),
            }
            //
            for _, measurement := range measurement.Measurements {
               row = append(row, fmt.Sprint(measurement.Measure(instance.(binarytree.BinaryTree))))
            }
            fmt.Fprintln(file, row...)
         }
         instance.Free()
      }
   }
}
