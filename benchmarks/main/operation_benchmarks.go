package main

import (
   "binarysearch/abstract/list"
   "binarysearch/binarytree"
   "binarysearch/distribution"
   "binarysearch/operations"
   "binarysearch/utility"
   "flag"
   "fmt"
   "time"
   "os"
   "path/filepath"
)

func main() {
   operation := flag.String("operation", "", "")
   flag.Parse()

   TreeBenchmark{
      Scale:      1_000_000,
      Samples:    1000,
      Iterations: 10,
      Operation: utility.Resolve(*operation, []list.Operation{
         &operations.Insert{},
         &operations.InsertPersistent{},
         &operations.InsertDelete{},
         &operations.InsertDeletePersistent{},
         &operations.InsertDeleteCycles{},
         &operations.InsertDeleteCyclesPersistent{},
         &operations.SplitJoin{},
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
         &binarytree.AVLJoinBased{},
         &binarytree.AVLWeakTopDown{},
         &binarytree.AVLWeakBottomUp{},
         &binarytree.AVLWeakJoinBased{},
         &binarytree.AVLRelaxedTopDown{},
         &binarytree.AVLRelaxedBottomUp{},
         &binarytree.RedBlackBottomUp{},
         &binarytree.RedBlackRelaxedBottomUp{},
         &binarytree.RedBlackRelaxedTopDown{},
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


type TreeBenchmark struct {
   Scale         int
   Samples       int
   Operation     list.Operation
   Distributions []distribution.Distribution
   Strategies    []list.List
   Iterations    int
}

func (benchmark TreeBenchmark) Run() {

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

         //
         stderr([]any{
            time.Now().Format(time.RFC822),
            utility.NameOf(benchmark.Operation),
            utility.NameOf(strategy),
            fmt.Sprint(iteration, "/", benchmark.Iterations),
         }...)

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