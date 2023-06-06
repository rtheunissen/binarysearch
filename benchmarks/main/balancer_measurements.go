package main

import (
   "binarysearch/binarytree"
   "binarysearch/distribution"
   "flag"
   "fmt"
   "binarysearch/utility"
   "os"
   "binarysearch/abstract/list"
   "time"
   "path/filepath"
)

func main() {
   strategy := flag.String("strategy", "", "")
   flag.Parse()

   BalancerMeasurement{
      Samples: 10_000,
      Scale:   100_000,
      Measurements: []binarytree.Measurement{
         &binarytree.PartitionCount{},
         &binarytree.PartitionDepth{},
         &binarytree.AveragePathLength{},
         &binarytree.MaximumPathLength{},
         &binarytree.Rotations{},
      },
      Distributions: []distribution.Distribution{
         &distribution.Uniform{},
         &distribution.Normal{},
         &distribution.Skewed{},
         &distribution.Zipf{},
         &distribution.Maximum{},
      },
      Strategy: utility.Resolve[binarytree.Balancer](*strategy, []binarytree.Balancer{
         &binarytree.Median{},
         &binarytree.Height{},
         &binarytree.HalfSize{},
         &binarytree.LogSize{},
         &binarytree.HalfWeight{},
         &binarytree.LogWeight{},
         &binarytree.Cost{},
         &binarytree.DSW{},
      }),
   }.Run()
}


type BalancerMeasurement struct {
   Scale         int
   Samples       int
   Strategy    binarytree.Balancer
   Measurements  []binarytree.Measurement
   Distributions []distribution.Distribution
}

func (measurement BalancerMeasurement) Run() {

   path := fmt.Sprintf(
      "benchmarks/csv/balancers/measurements/%s",
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

   instance := binarytree.Splay{}.New().(*binarytree.Splay)

   step := measurement.Scale / measurement.Samples

   for position := step; position <= measurement.Scale; position += step {

      // Grow the tree if needed.
      for instance.Size() < list.Size(position) {
          instance.Insert(0, 0)
      }

      fmt.Fprintf(os.Stderr, "%s %-10s %10d/%d\n",
         time.Now().Format(time.TimeOnly),
         utility.NameOf(measurement.Strategy),
         position,
         measurement.Scale)

      for _, random := range measurement.Distributions {

         // Randomize the tree.
         instance.Tree = instance.Tree.Randomize(random.New(uint64(position)))

         // Reset measurements.
         for _, measurement := range measurement.Measurements {
            measurement.Reset()
         }

         instance.Tree = measurement.Strategy.Restore(instance.Tree)

         row := []any{
            utility.NameOf(random),
            fmt.Sprint(measurement.Scale),
            fmt.Sprint(position),
         }
         for _, measurement := range measurement.Measurements {
            row = append(row, fmt.Sprint(measurement.Measure(instance.Tree)))
         }
         fmt.Fprintln(file, row...)
      }
   }
   instance.Free()
}