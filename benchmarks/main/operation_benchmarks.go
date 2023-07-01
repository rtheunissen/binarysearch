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

func main() {
   operation := flag.String("operation", "", "")
   flag.Parse()

   TreeBenchmark{
      Scale:       1_000_000,
      Samples:          1000,
      Iterations:         10,
      Operation: utility.Resolve(*operation, []list.Operation{
         &operations2.Insert{},
         //&operations2.InsertPersistent{},
         &operations2.InsertDelete{},
         //&operations2.InsertDeletePersistent{},
         &operations2.InsertDeleteCycles{},
         //&operations2.InsertDeleteCyclesPersistent{},
         //&operations2.InsertDeleteSearch{},
         //&operations2.InsertDeleteSearchPersistent{},
      }),
      Distributions: []distribution.Distribution{
         &distribution.Uniform{},
         //&distribution.Normal{},
         //&distribution.Skewed{},
         &distribution.Zipf{},
         //&distribution.Maximum{},
      },
      Strategies: []list.List{
         &binarytree.AVLBottomUp{},
         &binarytree.AVLTopDown{},
         //&binarytree.AVLWeakTopDown{},
         //&binarytree.AVLWeakBottomUp{},
         //&binarytree.AVLRelaxedTopDown{},
         //&binarytree.AVLRelaxedBottomUp{},
         //&binarytree.RedBlackBottomUp{},
         //&binarytree.RedBlackTopDown{},
         //&binarytree.RedBlackRelaxedBottomUp{},
         //&binarytree.RedBlackRelaxedTopDown{},
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
   }.Run()
}


type TreeBenchmark struct {
   Scale         int
   Samples       int
   Operation     list.Operation
   Distributions []distribution.Distribution
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
         "benchmarks/data/operations/benchmarks/%s/%s",
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

         fmt.Printf("%s %-32s %-32s %10d/%d\n",
            time.Now().Format(time.RFC822),
            utility.NameOf(benchmark.Operation),
            utility.NameOf(strategy),
            iteration,
            benchmark.Iterations)

         //
         for _, random := range benchmark.Distributions {

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