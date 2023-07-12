package main

import (
   "bst/abstract/list"
   "bst/abstract/list/operations"
   "bst/trees"
   "bst/utility"
   "bst/utility/random"
   "bst/utility/random/distribution"
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
   Distributions []random.Distribution
   Strategies    []list.List
   Measurements  []trees.Measurement
}

func main() {
   operation := flag.String("operation", "", "")
   flag.Parse()

   OperationMeasurement{
      Scale:         1_000_000,
      Samples:           1_000,
      Operation: utility.Resolve(*operation, []list.Operation{
         &operations.InsertPersistent{},
         &operations.InsertDeleteCyclesPersistent{},
         //&operations.InsertDeletePersistent{},
         //&operations.InsertDeleteSearchPersistent{},
      }),
      Distributions: []random.Distribution{
         &distribution.Uniform{},
         &distribution.Normal{},
         &distribution.Skewed{},
         &distribution.Zipf{},
         &distribution.Maximum{},
      },
      Strategies: []list.List{
         &trees.AVLBottomUp{},
         &trees.AVLTopDown{},
         &trees.AVLWeakTopDown{},
         &trees.AVLWeakBottomUp{},
         &trees.AVLRelaxedTopDown{},
         &trees.AVLRelaxedBottomUp{},
         &trees.RedBlackBottomUp{},
         &trees.RedBlackTopDown{},
         &trees.RedBlackRelaxedBottomUp{},
         &trees.RedBlackRelaxedTopDown{},
         &trees.LBSTBottomUp{},
         &trees.LBSTTopDown{},
         &trees.LBSTRelaxed{},
         &trees.WBSTBottomUp{},
         &trees.WBSTTopDown{},
         &trees.WBSTRelaxed{},
         &trees.TreapTopDown{},
         &trees.TreapFingerTree{},
         &trees.Randomized{},
         &trees.Zip{},
         &trees.Splay{},
         &trees.Conc{},
      },
      Measurements: []trees.Measurement{
         &trees.PartitionCount{},
         &trees.PartitionDepth{},
         &trees.AveragePathLength{},
         &trees.MaximumPathLength{},
         &trees.Rotations{},
         &trees.Allocations{},
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
         "docs/benchmarks/data/operations/measurements/%s/%s",
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
               row = append(row, fmt.Sprint(measurement.Measure(instance.(trees.BinaryTree))))
            }
            fmt.Fprintln(file, row...)
         }
         instance.Free()
      }
   }
}
