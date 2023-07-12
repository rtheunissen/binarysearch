package main

import (
   "bst/types/list"
   "bst/types/list/operations"
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

func main() {
   operation := flag.String("operation", "", "")
   flag.Parse()

   TreeBenchmark{
      Scale:       1_000_000,
      Samples:         1_000,
      Iterations:         10,
      Operation: utility.Resolve(*operation, []list.Operation{
         &operations.Insert{},
         &operations.InsertPersistent{},
         &operations.InsertDeleteCycles{},
         &operations.InsertDeleteCyclesPersistent{},
         //&operations.InsertDelete{},
         //&operations.InsertDeletePersistent{},
         //&operations.InsertDeleteSearch{},
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
   }.Run()
}


type TreeBenchmark struct {
   Scale         int
   Samples       int
   Operation     list.Operation
   Distributions []random.Distribution
   Strategies    []list.List
   Iterations    int
}

func (benchmark TreeBenchmark) Run() {
   if benchmark.Operation == nil {
      return
   }

   //
   for _, strategy := range benchmark.Strategies {

      path := fmt.Sprintf(
         "docs/benchmarks/data/operations/benchmarks/%s/%s",
         utility.NameOf(benchmark.Operation),
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
         "Iteration",
         "Duration",
      }
      //
      fmt.Fprintln(file, header...)

      //
      for iteration := 1; iteration <= benchmark.Iterations; iteration++ {

         //
         for _, random := range benchmark.Distributions {

            fmt.Printf("%s %-32s %-32s %-32s %10d/%d\n",
               time.Now().Format(time.RFC822),
               utility.NameOf(benchmark.Operation),
               utility.NameOf(strategy),
               utility.NameOf(random),
               iteration,
               benchmark.Iterations)

            //
            access := random.New(uint64(iteration))

            //
            instance := benchmark.Operation.Setup(strategy, list.Size(benchmark.Scale))

            //
            step := benchmark.Scale / benchmark.Samples

            //
            for position := step; position <= benchmark.Scale; position += step {

               start := time.Now()
               for i := 0; i < step; i++ {
                  instance, _ = benchmark.Operation.Update(instance, access)
               }
               duration := time.Since(start)

               fmt.Fprintln(file, []any{
                  utility.NameOf(access),
                  fmt.Sprint(benchmark.Scale),
                  fmt.Sprint(instance.Size()),
                  fmt.Sprint(step),
                  fmt.Sprint(position),
                  fmt.Sprint(iteration),
                  fmt.Sprint(duration.Nanoseconds()),
               }...)
            }
            instance.Free()
         }
      }
   }
}