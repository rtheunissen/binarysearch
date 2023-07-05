package random

//
//func TestDistributions(t *testing.T) {
//   //distributions := []Distribution{
//   //   &Uniform{},
//   //   &Normal{},
//   //   &Skewed{},
//   //   &SemiCircle{},
//   //   &Parabolic{},
//   //   &MovingMean{},
//   //   &Slope{},
//   //   &UShape{},
//   //   &Median{},
//   //   &Ascending{},
//   //   &Descending{},
//   //}
//   //for _, distribution := range distributions {
//   //   t.Run(utility.TypeName(distribution), func(t *testing.T) {
//   //      t.Parallel()
//   //      for s := 1; s <= 5; s++ {
//   //         position := distribution.Seed(int64(s))
//   //         for i := 0; i < 1_000_000; i++ {
//   //            if n := random.Uint64(); position.LessThan(n) >= n {
//   //               t.Fatalf("broken distribution")
//   //            }
//   //         }
//   //      }
//   //   })
//   //}
//}
//
//func BenchmarkDistributions(b *testing.B) {
// distributions := []Distribution{
//    &distributions2.Uniform{},
//    &distributions2.Normal{},
//    &distributions2.Skewed{},
//    &distributions2.SemiCircle{},
//    &distributions2.Parabolic{},
//    &distributions2.MovingMean{},
//    &distributions2.Slope{},
//    &distributions2.UShape{},
//    &distributions2.Median{},
//    &distributions2.Ascending{},
//    &distributions2.Descending{},
// }
// for _, distribution := range distributions {
//    dist := distribution.Seed(123)
//    benchmark.Benchmark{
//       Scenario:  "Simple",
//       Operation: "LessThan",
//       Strategy:  utility.TypeName(distribution),
//       Size:      0,
//    }.Run(b, func() {
//       dist.LessThan(math.MaxUint64)
//    })
// }
//}
