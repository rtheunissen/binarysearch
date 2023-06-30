package main

import (
   "binarysearch/abstract/list"
   operations2 "binarysearch/abstract/list/operations"
   "binarysearch/binarytree"
   "binarysearch/distribution"
   "binarysearch/utility"
   "flag"
   "fmt"
   "os"
   "path/filepath"
   "time"
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
   flag.Parse()

   OperationMeasurement{
      Scale:         1_000_000,
      Samples:           1_000,
      Operation: utility.Resolve(*operation, []list.Operation{
         &operations2.Insert{},
         &operations2.InsertPersistent{},
         &operations2.InsertDelete{},
         &operations2.InsertDeletePersistent{},
         &operations2.InsertDeleteCycles{},
         &operations2.InsertDeleteCyclesPersistent{},
         &operations2.InsertDeleteSearch{},
         &operations2.InsertDeleteSearchPersistent{},
      }),
      Distributions: []distribution.Distribution{
         &distribution.Uniform{},
         &distribution.Normal{},
         &distribution.Skewed{},
         &distribution.Zipf{},
         &distribution.Maximum{},
      },
      Strategies: []list.List{
         &binarytree.AVLBottomUp{},
         &binarytree.AVLWeakTopDown{},
         &binarytree.AVLWeakBottomUp{},
         &binarytree.AVLRelaxedTopDown{},
         &binarytree.AVLRelaxedBottomUp{},
         &binarytree.RedBlackBottomUp{},
         &binarytree.RedBlackTopDown{},
         &binarytree.RedBlackRelaxedBottomUp{},
         &binarytree.RedBlackRelaxedTopDown{},
         //&binarytree.LBSTBottomUp{},
         //&binarytree.LBSTTopDown{},
         //&binarytree.LBSTRelaxed{},
         //&binarytree.TreapTopDown{},
         //&binarytree.TreapFingerTree{},
         //&binarytree.Randomized{},
         //&binarytree.Zip{},
         //&binarytree.Splay{},
         //&binarytree.Conc{},
      },
      Measurements: []binarytree.Measurement{
         &binarytree.PartitionCount{},
         &binarytree.PartitionDepth{},
         &binarytree.AveragePathLength{},
         &binarytree.MaximumPathLength{},
         &binarytree.Rotations{},
         &binarytree.Allocations{},
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
