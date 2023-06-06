package main

import (
   "binarysearch/binarytree"
   "binarysearch/distribution"
   "flag"
   "fmt"
   "binarysearch/utility"
   "time"
   "os"
   "binarysearch/abstract/list"
   "path/filepath"
)

func main() {
   iterations := flag.Int("iterations", 100, "")
   samples := flag.Int("samples", 100, "")
   scale := flag.Int("scale", 1_000_000, "")
   flag.Parse()

   BalancerBenchmark{
      Samples:    *samples,
      Scale:      *scale,
      Iterations: *iterations,
      Distributions: []distribution.Distribution{
         &distribution.Uniform{},
      },
      Strategies: []binarytree.Balancer{
         &binarytree.Median{},
         &binarytree.Height{},
         &binarytree.Weight{},
         &binarytree.Cost{},
         &binarytree.DSW{},
         &binarytree.Log{},
         &binarytree.Constant{},
      },
   }.Run()
}


type BalancerBenchmark struct {
   Scale         int
   Samples       int
   Strategies    []binarytree.Balancer
   Distributions []distribution.Distribution
   Iterations    int
}

func (benchmark BalancerBenchmark) Run() {

   //
   for _, strategy := range benchmark.Strategies {

      path := fmt.Sprintf(
         "benchmarks/csv/balancers/benchmarks/%s",
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
      //
      header := []any{
         "Distribution",
         "Scale",
         "Size",
         "Duration",
      }
      fmt.Fprintln(file, header...)

      step := benchmark.Scale / benchmark.Samples

      //
      for iteration := 1; iteration <= benchmark.Iterations; iteration++ {

         instance := binarytree.Splay{}.New().(*binarytree.Splay)

         fmt.Fprintln(os.Stderr, []any{
            time.Now().Format(time.RFC822),
            utility.NameOf(strategy),
            fmt.Sprint(iteration, "/", benchmark.Iterations),
         }...)

         for position := step; position <= benchmark.Scale; position += step {

            // Grow the tree.
            for instance.Size() < list.Size(position) {
               instance.Insert(0, 0)
            }

            for _, random := range benchmark.Distributions {

               // Randomize the tree.
               instance.Tree = instance.Tree.Randomize(random.New(uint64(position)))

               start := time.Now()

               instance.Tree = strategy.Restore(instance.Tree)

               duration := time.Since(start)

               row := []any{
                  utility.NameOf(random),
                  fmt.Sprint(benchmark.Scale),
                  fmt.Sprint(instance.Size()),
                  fmt.Sprint(duration.Nanoseconds()),
               }
               fmt.Fprintln(file, row...)
            }
         }
         instance.Free()
      }
   }
}
