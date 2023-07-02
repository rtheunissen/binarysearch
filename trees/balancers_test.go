package trees

import (
   "bst/abstract/list"
   "bst/utility"
   "bst/utility/random"
   "bst/utility/random/distribution"
   "testing"
)

func TestBalancers(t *testing.T) {
   balancers := []Balancer{
      &Median{},
      &Height{},
      &Weight{},
      &Log{},
      &Cost{},
      &DSW{},
   }
   distributions := []distribution.Distribution{
      &distribution.Uniform{},
      &distribution.Normal{},
      &distribution.Skewed{},
      &distribution.Zipf{},
      &distribution.Maximum{},
   }
   testBalancers(t, 1000, balancers, distributions) // TODO: make consistent with test suites and benchmarks patterns exactly
}

func testBalancers(t *testing.T, scale list.Size, balancers []Balancer, distributions []distribution.Distribution) {
   for _, balancer := range balancers {

      t.Run(utility.NameOf(balancer), func(t *testing.T) {

         for _, distribution := range distributions {

            t.Run(utility.NameOf(distribution), func(t *testing.T) {

               tree := Tree{}
               reference := list.Reference{}
               dist := distribution.New(1)

               for tree.size < scale {

                  //
                  i := dist.LessThan(tree.size + 1)
                  x := random.Uint64()

                  //
                  tree.Insert(i, x)
                  reference.Insert(i, x)

                  //
                  tree = balancer.Restore(tree.Clone())
                  tree = balancer.Restore(tree.Clone())

                  balancer.Verify(tree)
               }
               //
               reference.Assert(t, tree)
               tree.Free()
            })
         }
      })
   }
}
