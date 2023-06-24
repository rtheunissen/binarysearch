package main

import (
   "binarysearch/abstract/list"
   "binarysearch/binarytree"
   "binarysearch/distribution"
   "binarysearch/utility"
   "fmt"
   "os"
   "path/filepath"
   "time"
)

func main() {
  BalancerBenchmark{
     Iterations:        100,
     Samples:           100,
     Scale:      10_000_000,
     Distributions: []distribution.Distribution{
        &distribution.Uniform{},
     },
     Strategies: []binarytree.Balancer{
        &binarytree.Median{},
        &binarytree.Height{},
        &binarytree.Weight{},
        &binarytree.Log{},
        &binarytree.Cost{},
        &binarytree.DSW{},
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
        "benchmarks/data/balancers/benchmarks/%s",
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

     instance := binarytree.Splay{}.New().(*binarytree.Splay)


     for position := step; position <= benchmark.Scale; position += step {

        fmt.Printf("%s %-10s %10d/%d\n",
           time.Now().Format(time.TimeOnly),
           utility.NameOf(strategy),
           position,
           benchmark.Scale)

        // Grow the tree.
        for instance.Size() < list.Size(position) {
            instance.Insert(0, 0)
        }

        for _, random := range benchmark.Distributions {

           var duration time.Duration

           source := random.New(uint64(position))

           for iteration := 1; iteration <= benchmark.Iterations; iteration++ {

              // Randomize the tree.
              instance.Tree = instance.Tree.Randomize(source)

              start := time.Now()

              instance.Tree = strategy.Restore(instance.Tree)

              duration += time.Since(start)
           }

           row := []any{
              utility.NameOf(random),
              fmt.Sprint(benchmark.Scale),
              fmt.Sprint(instance.Size()),
              fmt.Sprint(duration.Nanoseconds() / int64(benchmark.Iterations)),
           }
           fmt.Fprintln(file, row...)
        }
     }
     instance.Free()
  }
}
