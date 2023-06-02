package benchmarks

import (
   "binarysearch/abstract/list"
   "binarysearch/distribution"
   "binarysearch/utility"
   "fmt"
   "time"
)

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
   	stdout("#" + utility.NameOf(strategy))
   	stdout(header...)

   	//
   	for iteration := 1; iteration <= benchmark.Iterations; iteration++ {

   		//
   		stderr([]any{
   			time.Now().Format(time.RFC822),
   			alignL(utility.NameOf(benchmark.Operation)),
   			alignL(utility.NameOf(strategy)),
   			alignR(fmt.Sprint(iteration, "/", benchmark.Iterations)),
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

   				stdout([]any{
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
   	stdout()
   	stdout()
   }
}
