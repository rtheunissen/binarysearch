package main

import (
   "bst/abstract/list"
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
   strategy := flag.String("strategy", "", "")
   flag.Parse()

   BalancerMeasurement{
      Iterations:        1,
      Samples:      10_000,
      Scale:     1_000_000,
      Measurements: []trees.Measurement{
         &trees.PartitionCount{},
         &trees.PartitionDepth{},
         &trees.AveragePathLength{},
         &trees.MaximumPathLength{},
         &trees.Rotations{},
      },
      Distributions: []random.Distribution{
         &distribution.Uniform{},
      },
      Strategy: utility.Resolve[trees.Balancer](*strategy, []trees.Balancer{
         &trees.Median{},
         &trees.Height{},
         &trees.Weight{},
         &trees.Log{},
         &trees.Cost{},
         &trees.DSW{},
      }),
   }.Run()
}


type BalancerMeasurement struct {
   Iterations    int
   Scale         int
   Samples       int
   Strategy      trees.Balancer
   Measurements  []trees.Measurement
   Distributions []random.Distribution
}

func (measurement BalancerMeasurement) Run() {
   if measurement.Strategy == nil {
      return
   }
   path := fmt.Sprintf(
      "docs/benchmarks/data/balancers/measurements/%s",
      utility.NameOf(measurement.Strategy),
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
   }

   //
   for _, measurement := range measurement.Measurements {
      header = append(header, utility.NameOf(measurement))
   }
   fmt.Fprintln(file, header...)

   instance := trees.Splay{}.New().(*trees.Splay)

   step := measurement.Scale / measurement.Samples

   for position := step; position <= measurement.Scale; position += step {
      //
      // Grow the tree if needed.
      //
      for instance.Size() < list.Size(position) {
          instance.Insert(0, 0)
      }
      //
      //
      //
      if position % (measurement.Scale / step) == 0 {
         fmt.Printf("%s %-10s %10d/%d\n",
            time.Now().Format(time.TimeOnly),
            utility.NameOf(measurement.Strategy),
            position,
            measurement.Scale)
      }
      //
      //
      //
      for _, random := range measurement.Distributions {

         source := random.New(uint64(position))

         for iteration := 1; iteration <= measurement.Iterations; iteration++ {
            //
            // Randomize the tree.
            //
            instance.Tree = instance.Tree.Randomize(source)
            //
            // Reset measurements.
            //
            for _, measurement := range measurement.Measurements {
               measurement.Reset()
            }
            //
            //
            //
            instance.Tree = measurement.Strategy.Restore(instance.Tree)
            //
            //
            //
            row := []any{
               utility.NameOf(random),
               measurement.Scale,
               position,
            }
            for _, measurement := range measurement.Measurements {
               row = append(row, fmt.Sprint(measurement.Measure(instance.Tree)))
            }
            fmt.Fprintln(file, row...)
         }
      }
   }
}