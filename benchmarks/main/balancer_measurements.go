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
   samples := flag.Int("samples", 10_000, "")
   scale := flag.Int("scale", 100_000, "")
   flag.Parse()

   BalancerMeasurement{
      Samples: *samples,
      Scale:   *scale,
      Measurements: []binarytree.Measurement{
         &binarytree.PartitionCount{},
         &binarytree.PartitionDepth{},
         &binarytree.AveragePathLength{},
         &binarytree.MaximumPathLength{},
         &binarytree.Rotations{},
      },
      Distributions: []distribution.Distribution{
         &distribution.Uniform{},
      },
      Strategies: []binarytree.Balancer{
         //&binarytree.Median{},
         //&binarytree.Height{},
         //&binarytree.Weight{},
         &binarytree.Constant{},
         //&binarytree.Cost{},
         &binarytree.Log{},
         //&binarytree.DSW{},
      },
   }.Run()
}


type BalancerMeasurement struct {
   Scale         int
   Samples       int
   Strategies    []binarytree.Balancer
   Measurements  []binarytree.Measurement
   Distributions []distribution.Distribution
}

func (measurement BalancerMeasurement) Run() {

   //
   for _, strategy := range measurement.Strategies {

      path := fmt.Sprintf(
         "benchmarks/csv/balancers/measurements/%s",
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

         fmt.Fprintln(os.Stderr, []any{
            time.Now().Format(time.RFC822),
            utility.NameOf(strategy),
            fmt.Sprintf("%d / %d", position, measurement.Scale),
         }...)

         for _, random := range measurement.Distributions {

            // Randomize the tree.
            instance.Tree = instance.Tree.Randomize(random.New(uint64(position)))

            // Reset measurements.
            for _, measurement := range measurement.Measurements {
               measurement.Reset()
            }

            instance.Tree = strategy.Restore(instance.Tree)

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
}