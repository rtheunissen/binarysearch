package trees

func init() {
   //spew.Config.Indent = "       "
}

//
//func CompareBalancerPartitions(balancer Balancer, measurement Measurement) {
//      const scale = 1_000_000
//      const step = 10_000
//      const iterations = 10
//      tree := Tree{}.New()
//      for size := list.Size(step); size <= scale; size += step {
//         for tree.size < size {
//             tree.Insert(random.Uint64() % (size + 1), 0)
//         }
//         measurement.Reset()
//         for i := 1; i <= iterations; i++ {
//            tree = tree.Randomize(size)
//            tree = measurement.Measure(tree, func(tree Tree) Tree {
//               return balancer.Restore(tree)
//            })
//         }
//         fmt.Println(size, measurement.Result())
//      }
//      tree.Free()
//
//   //
//   //for name, rows := range results {
//   //   fmt.Println("#" + name)
//   //   for _, row := range rows {
//   //      fmt.Println(row...)
//   //   }
//   //   fmt.Println()
//   //   fmt.Println()
//   //}
//
//
//
//   //
//   ////
//   ////
//   ////
//   //fmt.Println("#Median")
//   //for _, size := range sizes {
//   //   tree = Tree{}.Vine(size)
//   //   tree = tree.Randomize(size)
//   //   tree.partitions = 0
//   //   tree.partitionSteps = 0
//   //   start := time.Now()
//   //   tree = Median{}.Restore(tree)
//   //   taken := time.Now().Sub(start)
//   //   fmt.Println("Median", size, tree.partitions, tree.partitionSteps, tree.root.AveragePathLength(), tree.root.MaximumPathLength(), taken.Nanoseconds())
//   //   tree.Free()
//   //}
//   //fmt.Println()
//   //fmt.Println()
//   ////
//   ////
//   ////
//   //fmt.Println("#Height")
//   //for _, size := range sizes {
//   //   tree = Tree{}.Vine(size)
//   //   tree = tree.Randomize(size)
//   //   tree.partitions = 0
//   //   tree.partitionSteps = 0
//   //   start := time.Now()
//   //   tree = Height{}.Restore(tree)
//   //   taken := time.Now().Sub(start)
//   //   fmt.Println("Height", size, tree.partitions, tree.partitionSteps, tree.root.AveragePathLength(), tree.root.MaximumPathLength(), taken.Nanoseconds())
//   //   tree.Free()
//   //}
//   //fmt.Println()
//   //fmt.Println()
//   ////
//   ////
//   ////
//   //fmt.Println("#Weight")
//   //for _, size := range sizes {
//   //   tree = Tree{}.Vine(size)
//   //   tree = tree.Randomize(size)
//   //   tree.partitions = 0
//   //   tree.partitionSteps = 0
//   //   start := time.Now()
//   //   tree = Weight{}.Restore(tree)
//   //   taken := time.Now().Sub(start)
//   //   fmt.Println("Weight", size, tree.partitions, tree.partitionSteps, tree.root.AveragePathLength(), tree.root.MaximumPathLength(), taken.Nanoseconds())
//   //   tree.Free()
//   //}
//   //
//   ////
//   //////
//   //////
//   //////
//   ////fmt.Println("#Median")
//   ////
//   ////   tree = Tree{}.Vine(size)
//   ////   tree = tree.Randomize(size)
//   ////   tree = Weight{}.Restore(tree)
//   ////   partitions = 0
//   ////   partition_steps = 0
//   ////   tree = Median{}.Restore(tree)
//   ////   fmt.Println("Median", size, partitions, partition_steps, tree.root.AveragePathLength(), tree.root.MaximumPathLength())
//   ////   tree.Free()
//   ////}
//   ////fmt.Println()
//   ////fmt.Println()
//   //////
//   //////
//   //////
//   ////fmt.Println("#Height")
//   ////for _, size := range sizes {
//   ////   tree = Tree{}.Vine(size)
//   ////   tree = tree.Randomize(size)
//   ////   tree = Weight{}.Restore(tree)
//   ////   partitions = 0
//   ////   partition_steps = 0
//   ////   tree = Height{}.Restore(tree)
//   ////   fmt.Println("Height", size, partitions, partition_steps, tree.root.AveragePathLength(), tree.root.MaximumPathLength())
//   ////   tree.Free()
//   ////}
//   ////fmt.Println()
//   ////fmt.Println()
//   //////
//   //////
//   //////
//   ////fmt.Println("#Weight")
//   ////for _, size := range sizes {
//   ////   tree = Tree{}.Vine(size)
//   ////   tree = tree.Randomize(size)
//   ////   partitions = 0
//   ////   partition_steps = 0
//   ////   tree = Weight{}.Restore(tree)
//   ////   fmt.Println("Weight", size, partitions, partition_steps, tree.root.AveragePathLength(), tree.root.MaximumPathLength())
//   ////   tree.Free()
//   ////}
//}

//func heightBound1(height int, size list.Size) bool {
//   return height > 2 * int(utility.Log2(size))
//}
//func heightBound2(height int, size list.Size) bool {
//   return size < (1 << ((height + 1) >> 1))
//}
//func heightBound3(height int, size list.Size) bool {
//   return (height + 1) / 2 > int(math.Log2(float64(size)))
//}
//func log2_1(x, y int) bool {
//   // assert(x <= y)
//   return (1 + int(math.Floor(math.Log2(float64(y))))) - (1 + int(math.Floor(math.Log2(float64(x))))) <= 1
//}
//func log2_2(x, y int) bool {
//   // assert(x <= y)
//   return x >= y / 2 && x <= y * 2
//}

func Beep() {
   //f, err := os.Open("dissolve.wav")
   //if err != nil {
   //   log.Fatal(err)
   //}
   //streamer, format, err := wav.Decode(f)
   //if err != nil {
   //   log.Fatal(err)
   //}
   //defer streamer.Close()
   //
   //speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
   //
   //done := make(chan bool)
   //speaker.Play(beep.Seq(streamer, beep.Callback(func() {
   //   done <- true
   //})))
   //<-done
}


func Sandbox() {

   //for x := 0; x < 1_000; x++ {
   //   for y := 0; y < 1_000; y++ {
   //      if (x + 1) >= y && (y + 1) >= x {
   //         fmt.Println(x, y, 1)
   //      } else {
   //         fmt.Println(x, y, 0)
   //      }
   //   }
   //}

   //n := uint64(1_000_000)
   //
   //op := operations.InsertDeleteCycles{}
   //
   //strategies := []list.List{
   //   &RedBlackRelaxedBottomUp{},
   //   &RedBlackRelaxedTopDown{},
   //   &AVLRelaxedBottomUp{},
   //   &AVLRelaxedTopDown{},
   //   &WBSTRelaxed{},
   //   &LBSTRelaxed{},
   //}
   //for i, s := range strategies {
   //   strategies[i] = op.Setup(s, n)
   //}
   //for j := 0;; j++ {
   //   for i, s := range strategies {
   //      strategies[i], _ = op.Update(s, distribution.Uniform{}.New(uint64(j)))
   //   }
   //   if (j+1) % 1_000 == 0 {
   //      fmt.Println()
   //      var buf bytes.Buffer
   //      for _, s := range strategies {
   //         fmt.Fprintf(&buf, "%-32s apl: %.4f, h: %d, s: %d\n", utility.NameOf(s), s.(BinaryTree).Root().AveragePathLength(), s.(BinaryTree).Root().MaximumPathLength(), s.Size())
   //      }
   //      fmt.Println(buf.String())
   //   }
   //}


   //
   //
   ////
   //N := 1000
   //
   //random.Seed(uint64(time.Now().UnixNano()))
   //
   //distributions := []random.Distribution{
   //   //&distribution.Skewed{},
   //   &distribution.Uniform{},
   //   //&distribution.Normal{},
   //   &distribution.Zipf{},
   //   &distribution.RandomBeta{},
   //   //&distribution.BiModal{},
   //   //&distribution.Parabolic{},
   //   //&distribution.Slope{},
   //   //&distribution.Maximum{},
   //   //&distribution.Queue{},
   //}
   //
   //body, err := os.ReadFile("docs/polytope/bottomup.csv")
   //if err != nil {
   // panic(err)
   //}
   //reader := csv.NewReader(bytes.NewReader(body))
   //reader.Comma = ' '
   //
   //pairs, err := reader.ReadAll()
   //if err != nil {
   // panic(err)
   //}
   //
   ////fmt.Println(len(pairs), "total")
   //
   ////count := 0
   //
   //for _, pair := range pairs {
   // delta, _ := new(big.Rat).SetString(pair[0])
   // gamma, _ := new(big.Rat).SetString(pair[1])
   //
   // //delta + 1 / delta <= gamma is valid
   // if new(big.Rat).Quo(new(big.Rat).Add(delta, big.NewRat(1,1)), delta).Cmp(gamma) <= 0 {
   //   continue
   // }
   // //fmt.Println("Testing", delta.FloatString(2), gamma.FloatString(2))
   // //count++
   // var wg sync.WaitGroup
   // wg.Add(1)
   // go func() {
   //   tree := &WBSTBottomUp{
   //      Delta: delta,
   //      Gamma: gamma,
   //   }
   //   defer func() {
   //      //fmt.Fprintln(os.Stderr, delta.FloatString(2), gamma.FloatString(2))
   //      if err := recover(); err == nil {
   //         //fmt.Println()
   //         fmt.Println(delta.FloatString(2), gamma.FloatString(2))
   //         //os.Stdout.Write([]byte("\a"))
   //      } else {
   //         //fmt.Fprintln(os.Stderr, err)
   //      }
   //      tree.Free()
   //      wg.Done()
   //   }()
   //   for _, dist1 := range distributions {
   //      for _, dist2 := range distributions {
   //         dist1 := dist1.New(random.Uint64())
   //         dist2 := dist2.New(random.Uint64())
   //         //
   //         // INSERT
   //         //
   //         for tree.size < list.Size(N) {
   //             tree.Insert(dist1.LessThan(tree.size+1), 0)
   //         }
   //         tree.Verify()
   //         //
   //         // DELETE
   //         //
   //         for tree.size > 0 {
   //            tree.Delete(dist1.LessThan(tree.size))
   //         }
   //         //
   //         // INSERT DELETE INSERT
   //         //
   //         for tree.size < list.Size(N) {
   //             tree.Insert(dist1.LessThan(tree.size+1), 0)
   //             tree.Delete(dist2.LessThan(tree.size))
   //             tree.Insert(dist1.LessThan(tree.size+1), 0)
   //         }
   //         tree.Verify()
   //         //
   //         // DELETE
   //         //
   //         for tree.size > 0 {
   //             tree.Delete(dist2.LessThan(tree.size))
   //         }
   //         tree.Free()
   //      }
   //  }
   // }()
   //   wg.Wait()
   //}
   //fmt.Println(count, "to go")

   //
   //
   //
   //N := 100_000
   //
   //random.Seed(uint64(time.Now().UnixNano()))
   //
   //distributions := []random.Distribution{
   // //&distribution.Skewed{},
   // //&distribution.Skewed2{},
   // //&distribution.Skewed3{},
   // //&distribution.Skewed4{},
   // //&distribution.Skewed5{},
   // //&distribution.Skewed6{},
   // //&distribution.Skewed7{},
   // //&distribution.Skewed8{},
   // //&distribution.Skewed9{},
   // //&distribution.Skewed10{},
   // //
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // &distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //
   // //&distribution.Uniform{},
   // //&distribution.Normal{},
   // //&distribution.Zipf{},
   // //&distribution.BiModal{},
   // //&distribution.Parabolic{},
   // //&distribution.Slope{},
   // //&distribution.Maximum{},
   // //&distribution.Queue{},
   //}
   //
   //body, err := os.ReadFile("docs/polytope/topdown.csv")
   //if err != nil {
   //  panic(err)
   //}
   //reader := csv.NewReader(bytes.NewReader(body))
   //reader.Comma = ' '
   //
   //pairs, err := reader.ReadAll()
   //if err != nil {
   //  panic(err)
   //}
   //
   //fmt.Println(len(pairs), "total")
   //
   //count := 0
   //
   //var wg sync.WaitGroup
   //
   //for _, pair := range pairs {
   //  delta, _ := new(big.Rat).SetString(pair[0])
   //  gamma, _ := new(big.Rat).SetString(pair[1])
   //
   //  // delta + 1 / delta <= gamma is valid
   //  if new(big.Rat).Quo(new(big.Rat).Add(delta, big.NewRat(1,1)), delta).Cmp(gamma) <= 0 {
   //     continue
   //  }
   //  fmt.Println("Testing", delta.FloatString(2), gamma.FloatString(2))
   //  count++
   //
   //  wg.Add(1)
   //  go func() {
   //     time.Sleep(100 * time.Millisecond)
   //
   //    tree := &WBSTTopDown{
   //       WeightBalance: Rational{
   //          Delta: delta,
   //          Gamma: gamma,
   //       },
   //    }
   //    defer func() {
   //       if err := recover(); err != nil {
   //          fmt.Println()
   //          fmt.Println(delta.FloatString(2), gamma.FloatString(2))
   //          os.Stdout.Write([]byte("\a"))
   //       }
   //       tree.Free()
   //       wg.Done()
   //    }()
   //    for  {
   //       fmt.Print(".")
   //       for _, dist1 := range distributions {
   //          for _, dist2 := range distributions {
   //             dist1 := dist1.New(random.Uint64())
   //             dist2 := dist2.New(random.Uint64())
   //             //
   //             // INSERT
   //             //
   //             for tree.size < list.Size(N) {
   //                 tree.Insert(dist1.LessThan(tree.size+1), 0)
   //             }
   //             tree.Verify()
   //             //
   //             // DELETE
   //             //
   //             for tree.size > 0 {
   //                tree.Delete(dist1.LessThan(tree.size))
   //             }
   //             //
   //             // INSERT DELETE INSERT
   //             //
   //             for tree.size < list.Size(N) {
   //                 tree.Insert(dist1.LessThan(tree.size+1), 0)
   //                 tree.Delete(dist2.LessThan(tree.size))
   //                 tree.Insert(dist1.LessThan(tree.size+1), 0)
   //             }
   //             tree.Verify()
   //             //
   //             // DELETE
   //             //
   //             for tree.size > 0 {
   //                 tree.Delete(dist2.LessThan(tree.size))
   //             }
   //             tree.Free()
   //          }
   //      }
   //   }
   //  }()
   //}
   //fmt.Println(count, "to go")
   //wg.Wait()
   ////
   ////
   ////
   ////
   ////
   ////
   //
   //
   //N := 100_000
   //
   //random.Seed(uint64(time.Now().UnixNano()))
   //
   //distributions := []random.Distribution{
   // //&distribution.Skewed{},
   // //&distribution.Skewed2{},
   // //&distribution.Skewed3{},
   // //&distribution.Skewed4{},
   // //&distribution.Skewed5{},
   // //&distribution.Skewed6{},
   // //&distribution.Skewed7{},
   // //&distribution.Skewed8{},
   // //&distribution.Skewed9{},
   // //&distribution.Skewed10{},
   // //
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // &distribution.RandomBeta{},
   // //&distribution.RandomBeta{},
   // //
   // //&distribution.Uniform{},
   // //&distribution.Normal{},
   // //&distribution.Zipf{},
   // //&distribution.BiModal{},
   // //&distribution.Parabolic{},
   // //&distribution.Slope{},
   // //&distribution.Maximum{},
   // //&distribution.Queue{},
   //}
   //
   //body, err := os.ReadFile("docs/polytope/topdown.csv")
   //if err != nil {
   //  panic(err)
   //}
   //reader := csv.NewReader(bytes.NewReader(body))
   //reader.Comma = ' '
   //
   //pairs, err := reader.ReadAll()
   //if err != nil {
   //  panic(err)
   //}
   //
   //fmt.Println(len(pairs), "total")
   //
   //count := 0
   //
   //var wg sync.WaitGroup
   //
   //for _, pair := range pairs {
   //  delta, _ := new(big.Rat).SetString(pair[0])
   //  gamma, _ := new(big.Rat).SetString(pair[1])
   //
   //  // delta + 1 / delta <= gamma is valid
   //  if new(big.Rat).Quo(new(big.Rat).Add(delta, big.NewRat(1,1)), delta).Cmp(gamma) <= 0 {
   //     continue
   //  }
   //  fmt.Println("Testing", delta.FloatString(2), gamma.FloatString(2))
   //  count++
   //
   //  wg.Add(1)
   //  go func() {
   //     time.Sleep(100 * time.Millisecond)
   //
   //    tree := &WBSTTopDown{
   //       WeightBalance: Rational{
   //          Delta: delta,
   //          Gamma: gamma,
   //       },
   //    }
   //    defer func() {
   //       if err := recover(); err != nil {
   //          fmt.Println()
   //          fmt.Println(delta.FloatString(2), gamma.FloatString(2))
   //          os.Stdout.Write([]byte("\a"))
   //       }
   //       tree.Free()
   //       wg.Done()
   //    }()
   //    for  {
   //       fmt.Print(".")
   //       for _, dist1 := range distributions {
   //          for _, dist2 := range distributions {
   //             dist1 := dist1.New(random.Uint64())
   //             dist2 := dist2.New(random.Uint64())
   //             //
   //             // INSERT
   //             //
   //             for tree.size < list.Size(N) {
   //                 tree.Insert(dist1.LessThan(tree.size+1), 0)
   //             }
   //             tree.Verify()
   //             //
   //             // DELETE
   //             //
   //             for tree.size > 0 {
   //                tree.Delete(dist1.LessThan(tree.size))
   //             }
   //             //
   //             // INSERT DELETE INSERT
   //             //
   //             for tree.size < list.Size(N) {
   //                 tree.Insert(dist1.LessThan(tree.size+1), 0)
   //                 tree.Delete(dist2.LessThan(tree.size))
   //                 tree.Insert(dist1.LessThan(tree.size+1), 0)
   //             }
   //             tree.Verify()
   //             //
   //             // DELETE
   //             //
   //             for tree.size > 0 {
   //                 tree.Delete(dist2.LessThan(tree.size))
   //             }
   //             tree.Free()
   //          }
   //      }
   //   }
   //  }()
   //}
   //fmt.Println(count, "to go")
   //wg.Wait()
   ////
   ////
   ////
   ////
   ////
   ////



























  //
  //N := 100
  //
  //random.Seed(2)
  //
  //step := big.NewRat(1, 100)
  //
  //deltaMin := big.NewRat(0, 1)
  //deltaMax := big.NewRat(5, 1)
  //gammaMin := big.NewRat(0, 1)
  //gammaMax := big.NewRat(3, 1)
  //
  //var Delta, Gamma big.Rat
  //
  //delta, gamma := &Delta, &Gamma
  //
  //distributions := []random.Distribution{
  //  &distribution.RandomBeta{},
  //  &distribution.Uniform{},
  //  &distribution.Skewed{},
  //  &distribution.Normal{},
  //  &distribution.Zipf{},
  //  &distribution.Maximum{},
  //}
  //
  //for delta.Set(deltaMin); delta.Cmp(deltaMax) <= 0; delta.Add(delta, step) {
  //  for gamma.Set(gammaMin); gamma.Cmp(gammaMax) <= 0; gamma.Add(gamma, step) {
  //     tree := &WBSTBottomUp{
  //        Delta: delta,
  //        Gamma: gamma,
  //     }
  //     var wg sync.WaitGroup
  //     wg.Add(1)
  //
  //     go func() {
  //        defer func() {
  //           tree.Free()
  //           if err := recover(); err == nil {
  //              fmt.Println(delta.FloatString(2), gamma.FloatString(2))
  //           }
  //           wg.Done()
  //        }()
  //        for _, dist := range distributions {
  //           dist := dist.New(random.Uint64())
  //           //
  //           // INSERT
  //           //
  //           for tree.size < list.Size(N) {
  //               tree.Insert(dist.LessThan(tree.size+1), 0)
  //           }
  //           //tree.Verify()
  //           //
  //           // DELETE
  //           //
  //           for tree.size > 0 {
  //               tree.Delete(dist.LessThan(tree.size))
  //           }
  //           //tree.Verify()
  //        }
  //     }()
  //     wg.Wait()
  //  }
  //}

   //random.Seed(4)
   //for {
   //  //fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
   //  tree := &RedBlackTopDown{}
   //  for i, n := uint64(0), random.Uint64() % 100_000; i < n; i++ {
   //     tree.Insert(random.Uint64() % (tree.Size() + 1), i)
   //     tree.Verify()
   //  }
   //  tree.Free()
   //  fmt.Print(".")
   //}

   //
   //for i := 1; i < 1000000; i++ {
   //   for h := 1; h < 32; h++ {
   //      h1 := heightBound1(h, Size(i))
   //      h2 := heightBound2(h, Size(i))
   //      h3 := heightBound3(h, Size(i))
   //      if int(2 * math.Log2(float64(i))) != int(2 *utility.Log2(uint64(i))) {
   //      //if h1 != h2 || h1 != h3 || h2 != h3 {
   //         panic(fmt.Sprint(i, h, h1, h2, h3, int(2 * math.Log2(float64(i))), int(2 *utility.Log2(uint64(i)))))
   //      }
   //   }
   //}

   //
   //random.Seed(4)
   //var d time.Duration
   //var f float64
   //for k := 0; k < 100; k++ {
   //  tree := &LBSTRelaxed{}
   //  size := 10_000
   //
   //  for i := 0; i < size; i++ {
   //     tree.Insert(random.LessThan(tree.Size()+1, random.Uniform()), Data(i))
   //  }
   //  t := time.Now()
   //  for tree.root != nil {
   //     i := random.LessThan(tree.Size(), random.Uniform())
   //     tree.Delete(i)
   //     //f += tree.root.AveragePathLength()
   //  }
   //  d += time.Since(t)
   //  tree.Free()
   //}
   //fmt.Println(f / (3 * 100_000))
   //fmt.Println(d)

   //
   //t := time.Now()
   //for j := 0; j < 10; j++ {
   //  tree := &LBSTTopDown{}
   //  for i := 0; i < 1000000; i++ {
   //     tree.Insert(random.Uint64() % (tree.Size() + 1), list.Data(i))
   //  }
   //
   //   fmt.Println(tree.root.height())
   //   for i := 0; i < 1000000; i++ {
   //      tree.Insert(0, list.Data(i))
   //   }
   //   tree.Free()
   //}
   //fmt.Println(time.Since(t))
   //fmt.Println(rotations)
   //fmt.Println(rotations)


   //
   //n := 100000
   //for x := 1; x < n; x++ {
   //   for y := x; y < n; y++ {
   //      if log2_1(x, y) != log2_2(x, y) {
   //         panic(fmt.Sprint("x", x, "y", y))
   //      }
   //   }
   //}


   //r := random.New(4)
   //for {
   //   //fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
   //   tree := &AVLBottomUp{}
   //
   //   for i, n := uint64(0), random.LessThan(8, r) + 1; i < n; i++ {
   //      tree.Insert(random.LessThan(tree.Size()+1, random.Uniform()), i)
   //   }
   //   //tree.Draw(os.Stdout)
   //   tree.Free()
   //}

   //
   //for {
   // fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
   // a := &AVLBottomUp{}
   // b := &AVLBottomUp{}
   //
   // for i, n := uint64(0), random.LessThan(100, random.Uniform()) + 1; i < n; i++ {
   //    a.Insert(random.LessThan(a.Size()+1, random.Uniform()), i)
   // }
   // for i, n := uint64(0), random.LessThan(100, random.Uniform()) + 1; i < n; i++ {
   //    b.Insert(random.LessThan(b.Size()+1, random.Uniform()), i)
   // }
   //
   // a.Draw(os.Stdout)
   // b.Draw(os.Stdout)
   //
   // c := a.Join(b).(*AVLBottomUp)
   //
   // c.Draw(os.Stdout)
   // c.Verify()
   //
   // c.Free()
   //}


   //
   //
   //for j := 0; j < 100000; j++ {
   //   //fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
   //   tree := &AVLBottomUp{}
   //   size := random.LessThan(1000, random.Uniform()) + 1
   //   for i := uint64(0); i < size; i++ {
   //      tree.Insert(random.LessThan(tree.Size()+1, random.Uniform()), i)
   //   }
   //
   //   //tree.Draw(os.Stdout)
   //   for tree.root != nil {
   //      //fmt.Println()
   //      //fmt.Println()
   //      //fmt.Println("before deleting =====================")
   //
   //      //tree.Verify()
   //      //tree.Draw(os.Stdout)
   //
   //      //fmt.Println("deleting")
   //      i := random.LessThan(tree.Size(), random.Uniform())
   //      //fmt.Println(i)
   //      tree.Delete(i)
   //      //fmt.Println("after deleting: ---------------")
   //      //tree.Draw(os.Stdout)
   //      //tree.Verify()
   //   }
   //   tree.Free()
   //}
   //fmt.Println(rotations)


   //for {
   //  fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
   //  tree := &RedBlackBottomUp{}
   //  size := random.LessThan(34, random.Uniform()) + 1
   //  for i := uint64(0); i < size; i++ {
   //     tree.Insert(random.LessThan(tree.Size()+1, random.Uniform()), i)
   //  }
   //
   //  tree.Draw(os.Stdout)
   //  for tree.root != nil {
   //     fmt.Println()
   //     fmt.Println()
   //     fmt.Println("before deleting max =====================")
   //
   //     tree.Verify()
   //     tree.Draw(os.Stdout)
   //
   //     fmt.Println("deleting max")
   //     i := tree.size - 1
   //     fmt.Println(i)
   //     tree.Delete(i)
   //     fmt.Println("after deleting max: ---------------")
   //     tree.Draw(os.Stdout)
   //     tree.Verify()
   //  }
   //  tree.Free()
   //}


   //for {
   //   fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
   //   tree := &RedBlackBottomUp{}
   //   size := random.LessThan(32, random.Uniform()) + 1
   //   for i := uint64(0); i < size; i++ {
   //      tree.Insert(random.LessThan(tree.Size()+1, random.Uniform()), i)
   //   }
   //
   //   tree.Draw(os.Stdout)
   //   for tree.root != nil {
   //      fmt.Println()
   //      fmt.Println()
   //      fmt.Println("before deleting min =====================")
   //
   //      tree.Verify()
   //      tree.Draw(os.Stdout)
   //
   //      fmt.Println("deleting min")
   //      i := uint64(0)
   //      fmt.Println(i)
   //      tree.Delete(i)
   //      fmt.Println("after deleting min: ---------------")
   //      tree.Draw(os.Stdout)
   //      tree.Verify()
   //   }
   //   tree.Free()
   //}


   //for {
   //   fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
   //   tree := &RedBlackBottomUp{}
   //   size := random.LessThan(13, random.Uniform()) + 1
   //   for i := uint64(0); i < size; i++ {
   //      tree.Insert(random.LessThan(tree.Size()+1, random.Uniform()), i)
   //   }
   //   fmt.Println("start deleting =====================")
   //   tree.Draw(os.Stdout)
   //   for tree.root.r != nil {
   //      fmt.Println()
   //      fmt.Println(" deleting p.r, which is currently: ---------------")
   //      tree.Verify()
   //      tree.Draw(os.Stdout)
   //      i := tree.root.s + tree.root.r.s + 1
   //      fmt.Println(i)
   //      tree.Delete(i)
   //      fmt.Println(" after deleting p.r: ---------------")
   //      tree.Draw(os.Stdout)
   //      tree.Verify()
   //   }
   //   tree.Free()
   //}




   //t := time.Now()
   //for j := 0; j < 100; j++ {
   //   tree := &LBSTRelaxed{}
   //   for i := 0; i < 1000000; i++ {
   //      tree.Insert(random.Uint64() % (tree.Size() + 1), Data(i))
   //   }
   //   tree.Free()
   //}
   //fmt.Println(time.Since(t))
   //
   //

   //for seed := uint64(time.Now().UnixNano()); ; seed++ {
   //   t := LBSTTopDown{}.New()
   //   r := random.New(seed)
   //   fmt.Println(seed, "---")
   //   for i := uint64(1); i <= 16; i++ {
   //      j := random.LessThan(t.Size()+1, r)
   //      fmt.Println(j)
   //      t.(*LBSTTopDown).Draw(os.Stderr)
   //      t.Insert(j, 0)
   //      t.(*LBSTTopDown).Draw(os.Stderr)
   //      t.Verify()
   //   }
   //   t.Free()
   //}

   //
   //a := LBSTRelaxed{}.New().(*LBSTRelaxed)
   //b := AVLBottomUp{}.New().(*AVLBottomUp)
   //c := Splay{}.New().(*Splay)
   //
   //for i := 0; i < 7; i++ {
   //  a.Insert(random.Uint64() % (uint64(i) + 1), 0)
   //  b.Insert(random.Uint64() % (uint64(i) + 1), 0)
   //  c.Insert(random.Uint64() % (uint64(i) + 1), 0)
   //}
   //a.Tree = a.Tree.Randomize(distribution.Zipf{}.New(1))
   //b.Tree = b.Tree.Randomize(distribution.Zipf{}.New(2))
   //c.Tree = c.Tree.Randomize(distribution.Zipf{}.New(3))
   //
   //a.Draw()
   //b.Draw()
   //c.Draw()
   //
   //a.Tree = a.Tree.Randomize(distribution.Zipf{}.New(1))
   //b.Tree = b.Tree.Randomize(distribution.Zipf{}.New(2))
   //c.Tree = c.Tree.Randomize(distribution.Zipf{}.New(3))
   //
   //a.Draw()
   //b.Draw()
   //c.Draw()

   //CompareBalancerPartitions(Weight{}, &Duration{})
   //
   //r := random.New(123)
   //t := LBSTRelaxed{}.New()
   //start := time.Now()
   //for i := list.Size(0); i < 10_000_000; i++ {
   //  t.Insert(random.LessThan(i + 1, r), 0)
   //}
   //fmt.Println(time.Now().Sub(start))
   //t.Free()
   //tree := Tree{}.Vine(10)
   //tree = tree.Randomize(123)
   //tree.Draw()
   //
   //tree = tree.Randomize(345)
   //
   //tree = tree.Randomize(123)
   //tree.Draw()
   //
   //tree = DSW{}.toVine(tree)
   //tree = tree.Randomize(123)
   //tree.Draw()
   //
   //Tree{}.Vine(10).Randomize(123).Draw()

   //MicroBenchmarkSmallerLog2()
   //
   //tree := AVLBottomUp{}.New().(*AVLBottomUp)
   //
   //fmt.Println("i", "apl")
   //for i := 0; i <= 8; i++ {
   //  tree.Insert(Position(i), 0)
   //  tree.Draw()
   //}

   //Tree{}.Vine(5).Randomize(random.New(123)).Balance().Draw()
   ////Tree{}.Vine(7).Randomize(random.New(345)).Balance().Draw()
   ////
   ////tree := AVLBottomUp{}.New().(*AVLBottomUp)
   ////
   ////fmt.Println("i", "apl")
   ////for i := 0; i <= 1_000_000; i++ {
   ////   tree.Insert(random.Uint64() % uint64(i + 1), 0)
   ////   if i % 1000 == 0 {
   ////      t := math.Log(float64(i))
   ////      a := tree.root.AveragePathLength()
   ////      fmt.Println(i, a, t, a / t, t / a)
   ////   }
   ////}

   //Tree{}.New().WorstCaseMedian(7).Draw()

   //Tree{}.New().Vine(6).Balance().ToWorstCaseMedian().Draw()
   //Tree{}.New().Vine(7).Balance().Draw()

   //start := time.Now()
   //
   //r := random.New(123)
   //for i := 0; i < 100; i++ {
   //   Tree{}.New().Vine(1_000_000).Randomize(r).toVine().Free()
   //}
   //fmt.Println(time.Now().Sub(start))

   //Tree{}.New().WorstCaseMedian(1).Draw()
   //Tree{}.New().WorstCaseMedian(2).Draw()
   //Tree{}.New().WorstCaseMedian(3).Draw()
   //Tree{}.New().WorstCaseMedian(4).Draw()
   //Tree{}.New().WorstCaseMedian(5).Draw()
   //Tree{}.New().WorstCaseMedian(6).Draw()
   //Tree{}.New().WorstCaseMedian(7).Draw()

   //   balancers := []Balancer{
   //      Weight{},
   //      //Height{},
   //      //Median{},
   //   }
   //
   //
   //   for _, balancer := range balancers {
   //
   //      scale := uint64(10000)
   //
   //      tree := Tree{}.New()
   //
   //      dist := random.New(123)
   //
   //      for tree.size < scale {
   //
   //         tree.Insert(random.LessThan(tree.size+1, dist), 0)
   //
   //         //t.Run("already balanced", func(t *testing.T) {
   //         tree = tree.Clone()
   //         tree = balancer.Restore(tree)
   //         //balancer.Verify(tree)
   //         //})
   //
   //         //t.Run("randomized", func(t *testing.T) {
   //         tree = tree.Clone().Randomize()
   //         tree = balancer.Restore(tree)
   //         //balancer.Verify(tree)
   //         //})
   //
   //         //t.Run("linked list", func(t *testing.T) {
   //         tree = tree.Clone().toVine()
   //         tree = balancer.Restore(tree)
   //         //balancer.Verify(tree)
   //         //})
   //      }
   //      //tree.Verify()
   //      tree.Free()
   //   }
   //
   //   fmt.Println(__partitioning_counter, __copy_counter)
   //
   //   //
   //t := AVLBottomUp{}.New()
   //for i := uint64(0);  i < 1000000; i++ {
   //   t = t.Clone()
   //   t.Insert(random.LessThan(t.Size() + 1, random.Uniform()), 0)
   //}
   //fmt.Println(__copy_counter)
   //t.Free()

   //t := TreapFingerTree{}.New()
   //for i := uint64(0);  i < 1000000; i++ {
   //   t = t.Clone()
   //   t.Insert(random.LessThan(t.Size() + 1, random.Uniform()), 0)
   //}
   //fmt.Println(numberOfCopies)
   //t.Free()
   //
   //for {
   //   a := AVLBottomUp{}.New()
   //   b := AVLBottomUp{}.New()
   //
   //   n1 := random.Uint64() % 10
   //   for i := uint64(0); i < n1; i++ {
   //      b.Insert(random.Uint64()%(a.Size()+1), Value(i))
   //   }
   //   n2 := random.Uint64() % 10
   //   for i := uint64(0); i < n2; i++ {
   //      a.Insert(random.Uint64()%(a.Size()+1), Value(i))
   //   }
   //
   //   c := a.Join(b)
   //
   //   c.(*AVLBottomUp).Draw()
   //
   //   c.Verify()
   //
   //   a.Free()
   //   b.Free()
   //   c.Free()
   //}

   //
   //t := TreapFingerTree{Tree: Tree{}.New(), Source: random.New(123)}
   //
   //t.size = 8
   //t.root = &Node{
   //  x: 3,
   //  y: 7,
   //  i: 2,
   //  t: &Node{
   //     x: 1,
   //     y: 3,
   //     i: 0,
   //     t: &Node{
   //        x: 2,
   //        y: 5,
   //        i: 0,
   //     },
   //  },
   //  r: &Node{
   //     x: 8,
   //     y: 4,
   //     i: 3,
   //     t: &Node{
   //        x: 7,
   //        y: 2,
   //        i: 2,
   //        t: &Node{
   //           x: 5,
   //           y: 1,
   //           i: 0,
   //           r: &Node{
   //              x: 6,
   //              y: 0,
   //              i: 0,
   //           },
   //        },
   //     },
   //     r: &Node{
   //        x: 4,
   //        y: 6,
   //        i: 0,
   //     },
   //  },
   //}
   //t.Verify()
   //t.Draw()
   //
   //t, r := t.Split(5)
   //
   //t.(*TreapFingerTree).Draw()
   //r.(*TreapFingerTree).Draw()
   //
   //t.Verify()
   //r.Verify()
   //t.Verify()
   //
   //n := list.Size(8)
   //
   //for i := list.Size(0); i < n; i++ {
   //   t.Insert(i, i + 1)
   //}
   //for i := list.Size(0); i < n; i++ {
   //   fmt.Println("===")
   //   fmt.Println(i)
   //   t, r := t.Split(i)
   //   t.(*TreapFingerTree).Draw(os.Stdout)
   //   t.(*TreapFingerTree).Draw(os.Stdout)
   //   r.(*TreapFingerTree).Draw(os.Stdout)
   //   t.Verify()
   //   r.Verify()
   //   t.Verify()
   //
   //   t = t.Join(r)
   //   fmt.Println("---")
   //   t.(*TreapFingerTree).Draw(os.Stdout)
   //}
   //t.Free()
   //}

   //for {
   //   t := TreapFingerTree{}.New()
   //   n := list.Size(8)
   //
   //   for i := list.Size(0); i < n; i++ {
   //       t.Insert(i, i + 1)
   //   }
   //   for i := list.Size(0); i < n; i++ {
   //      fmt.Println("===")
   //      fmt.Println(i)
   //      t, r := t.Split(i)
   //      t.(*TreapFingerTree).Draw(os.Stdout)
   //      t.(*TreapFingerTree).Draw(os.Stdout)
   //      r.(*TreapFingerTree).Draw(os.Stdout)
   //      t.Verify()
   //      r.Verify()
   //      t.Verify()
   //
   //      t = t.Join(r)
   //      fmt.Println("---")
   //      t.(*TreapFingerTree).Draw(os.Stdout)
   //   }
   //   t.Free()
   //}

   //Tree{}.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 50, 100).Draw(os.Stdout)
   //
   //
   //  Tree{
   //     root: &Node{
   //        x: 5,
   //        r: &Node{
   //           x: 6,
   //        },
   //        t: &Node{
   //           x: 4,
   //           t: &Node{
   //              x: 3,
   //              t: &Node{
   //                 x: 2,
   //                 t: &Node{
   //                    x: 1,
   //                 },
   //              },
   //           },
   //        },
   //     }}.Draw(os.Stdout)
   //
   //t.Run("linked list right", func(t *testing.T) {
   //   expect(Draw(Tree{
   //      root: &Node{
   //         x: 2,
   //         t: &Node{x: 1},
   //         r: &Node{
   //            x: 3,
   //            r: &Node{
   //               x: 4,
   //               r: &Node{
   //                  x: 5,
   //                  r: &Node{
   //                     x: 6,
   //                  },
   //               },
   //            },
   //         },
   //      },
   //   }), `
   //                           123
   //                            ╰───────────────╮
   //                                           234
   //                                            ╰───────╮
   //                                                   345
   //                                                    ╰───╮
   //                                                       456
   //                                                        ╰─╮
   //                                                         567`,
   //   )
   //})

   //r := random.New(123)
   //i := 0
   //for {
   //  fmt.Println("-------------")
   //  t := TreapFingerTree{}.New().(*TreapFingerTree)
   //
   //  for n := random.LessThan(16, r); n > 0; n-- {
   //     t.Insert(random.LessThan(t.size, r), n)
   //  }
   //  x, y := t.Split(random.LessThan(t.size + 1, r))
   //  //x.(*TreapFingerTree).Draw(os.Stdout)
   //  //y.(*TreapFingerTree).Draw(os.Stdout)
   //  x.Verify()
   //  y.Verify()
   //  t.Verify()
   //  t.Free()
   //
   //  fmt.Println(i)
   //  i++
   //}

   //rand := distribution.Uniform{}.Seed(1223)
   //for {
   //   tree := RedBlackRankBalanced{}.New().(*RedBlackRankBalanced)
   //   for i, n := uint64(0), uint64(32); i < n; i++ {
   //      tree.Insert(rand.LessThan(tree.size+1), i)
   //   }
   //   draw(tree.root)
   //}

   //draw := Draw{}
   //rand := distribution.Uniform{}.Seed(123)
   //for {
   //   tree := RedBlackRelaxedBottomUp{}.New().(*RedBlackRelaxedBottomUp)
   //   fmt.Println("---------")
   //   for i, n := uint64(0), uint64(32); i < n; i++ {
   //      tree.Insert(rand.LessThan(tree.size+1), i)
   //   }
   //
   //   joined := tree.Join(tree).(*RedBlackRelaxedBottomUp)
   //
   //   draw.draw(tree.root)
   //   draw.draw(joined.root)
   //   joined.Verify()
   //}

   //for {
   //   draw := Draw{}
   //   spew.Dump("========================================")
   //
   //
   //   L := RedBlackRelaxedBottomUp{}.New().(*RedBlackRelaxedBottomUp)
   //   R := RedBlackRelaxedBottomUp{}.New().(*RedBlackRelaxedBottomUp)
   //
   //   for i, n := uint64(0), rand.LessThan(16); i < n; i++ { L.Insert(rand.LessThan(L.size + 1), i) }
   //   for i, n := uint64(0), rand.LessThan(16); i < n; i++ { R.Insert(rand.LessThan(R.size + 1), i) }
   //
   //   draw.draw(L.root)
   //   draw.draw(R.root)
   //
   //   draw.draw(L.Join(R).(*RedBlackRelaxedBottomUp).root)
   //   L.Join(R).Verify()
   //
   //   draw.draw(R.Join(L).(*RedBlackRelaxedBottomUp).root)
   //   R.Join(L).Verify()
   //}

   //treap := TreapTopDown{}.from([]abstract.Value{1,2, 3, 4,5,6,7,8,9,10})
   //
   //treap.BST.Draw()

   //treap.Verify()
   //
   //
   //n := 1 << 15
   //
   //strategies := []abstract.List{
   //  &Zip{},
   //  &Randomized{},
   //  &TreapTopDown{},
   //}
   //results := make([][]int, n)
   //
   //for _, strategy := range strategies {
   //  fmt.Print(utility.TypeName(strategy), " ")
   //}
   //fmt.Println()
   //
   //for _, strategy := range strategies {
   //
   //  operation := DeletePersistent{}
   //
   //  distribution := &Ascending{}
   //
   //  instance := operation.Setup(strategy, abstract.Size(n))
   //
   //  x := 0
   //  for operation.Valid(instance) {
   //     instance, _ = operation.Update(instance, distribution)
   //     results[x] = append(results[x], instance.(BinaryTree).Root().Height())
   //     x++
   //  }
   //}
   //for _, heights := range results {
   //  for _, h := range heights {
   //     fmt.Print(h, " ")
   //  }
   //  fmt.Println()
   //}
   //
   //

   //AVLBottomUp{}.New(make([]list.Value, 1 << 4 - 1)...).(*AVLBottomUp).Draw()

   //values := make([]abstract.Value, 15)
   //for i := 0; i < 15; i++ {
   //   values[i] = abstract.Value(i)
   //}
   ////tree := Zip{}.New(values...)
   //tree := TreapTopDown{}.New(values...)
   //
   ////spew.Dump(tree)
   //
   //dist := distribution.Ascending{}
   //
   //for i := 0; i < 3; i++ {
   //   tree = tree.Clone()
   //   tree.Delete(dist.LessThan(tree.Size()))
   //   spew.Dump(numberOfCopies)
   //}

   //
   //tree := BST{}.New(make([]abstract2.Value, 100)...)
   //
   //randomized := randomize2(tree.(*BST).root, tree.Size(), (&distributions.Skewed{}).Seed(1))
   //
   ////spew.Dump(randomized.ExteriorHeightsAlongTheSpines())
   //
   //console.Histogram{
   //   Series: randomized.ExteriorHeightsAlongTheSpines(),
   //   MaximumPathLength: 20,
   //   Offset: 0,
   //}.Print(os.Stdout)
   //
   //
   //dist := (&distributions.Skewed{}).Seed(1)
   //
   //tree = BST{}.New()
   //for i := 0; i < 100; i++ {
   //   tree.Insert(dist.LessThan(tree.Size() + 1), 0)
   //}
   //
   //console.Histogram{
   //   Series: tree.(*BST).root.ExteriorHeightsAlongTheSpines(),
   //   MaximumPathLength: 20,
   //   Offset: 0,
   //}.Print(os.Stdout)
   //

   //N := 2000
   //I := &LBSTBottomUp{}
   //tree := I.New()
   //for i := 0; i < N; i++ {
   //   tree.Insert(abstract2.Position(rand.Intn(int(tree.Size() + 1))), 0)
   //}
   //var heights []int
   //tree.(*LBSTBottomUp).Inorder(func(p Node) {
   //   heights = append(heights, p.MaximumPathLength())
   //})
   //for x, y := range heights {
   //   fmt.Println(x, y)
   //}
   //
   //
   //   j := 0
   //
   //   for _, t := range heights[0] {
   //      fmt.Println(j, t + 1)
   //      j++
   //   }
   //   for _, r := range heights[1] {
   //      fmt.Println(j, r + 1)
   //      j++
   //   }
   //   if i + 1 < N {
   //      fmt.Println()
   //      fmt.Println()
   //   }
   //}

   //N := 2000
   //tree := AVL{}.New()
   //for i := 0; i < N; i++ {
   //   heights := tree.(*AVL).Root().InteriorHeightsAlongTheSpines()
   //
   //   j := 0
   //   tree.Insert(random.Uint64(), 0)
   //   for _, t := range heights[0] {
   //      fmt.Println(j, t + 1)
   //      j++
   //   }
   //   for _, r := range heights[1] {
   //      fmt.Println(j, r + 1)
   //      j++
   //   }
   //   if i + 1 < N {
   //      fmt.Println()
   //      fmt.Println()
   //   }
   //}

   //n := 9000
   //
   //for P := 1; P < 1000; P++ {
   //   spew.Dump(P)
   //   rand.Seed(int64(P))
   //
   //   v := []abstract2.Value{}
   //   for i := 0; i < n; i++ {
   //      v = append(v, uint64(i))
   //   }
   //   t := FingerTreeDisjointGeometric{}.New(v...)
   //   for i := uint64(0); i < t.Size(); i++ {
   //      t, r := t.Split(i)
   //      t.Verify()
   //      t.Verify()
   //      r.Verify()
   //      //
   //      invariant(slices.Equal(abstract2.Values(t), v))
   //      invariant(slices.Equal(abstract2.Values(t), v[:i]))
   //      invariant(slices.Equal(abstract2.Values(r), v[i:]))
   //   }
   //}

   //
   //v := []abstract.Value{}
   //n := 1000
   //for i := 0; i < n; i++ {
   //   v = append(v, uint64(i))
   //   t := FingerTreeDisjointGeometric{}.New(v...)
   //   for i := uint64(0); i < t.(*FingerTreeDisjointGeometric).HeadSize(); i++ {
   //      t := FingerTreeDisjointGeometric{}.New(v...)
   //      t, r := t.Split(i)
   //      t.Verify()
   //      r.Verify()
   //   }
   //}

   //spew.Dump(t)
   //spew.Dump(t, r)
   //
   //spew.Dump(abstract.Values(t), abstract.Values(r))

   // for depth := uint64(0); depth < 120; depth++ {
   //    for size := uint64(1); size < 1000000; size++ {
   //       a := sg1(depth, size)
   //       b := sg2(depth, size)
   //       if a != b {
   //          panic(fmt.Sprintf("no good! %d %d", depth, size))
   //       }
   //    }
   // }

//    for _, a := range distributions.Distributions(123) {
//       for _, b := range distributions.Distributions(123) {
//          tree := LBSTRelaxed{}.New().(*LBSTRelaxed)
//          for tree.size < 1000000 {
//             tree.Insert(a.LessThan(tree.size+1), 0)
//          }
//          for tree.size > 0 {
//             tree.Delete(b.LessThan(tree.size))
//          }
//          for tree.size < 1000000 {
//             tree.Insert(a.LessThan(tree.size+1), 0)
//             tree.Insert(a.LessThan(tree.size+1), 0)
//             tree.Delete(b.LessThan(tree.size))
//          }
//          print(".")
//       }
//    }
//    println()
//
//    // println("mixed")
//    // mixed()
//    // println("insert")
//    // onlyInsert()
//    // println("delete")
//    // onlyDelete()
//    // C := map[int]uint64{}
//    // // S := random.Source(123123123)
//    // // R := &Geometric{}
//    // Q := TrailingZeroes{source: random.Source(123)}
//    // H := 64
//    // // R.Seed(123456345656)
//    // for i := 0; i < 1000000000; i++ {
//    //    // C[R.next()]++
//    //    C[Q.next()]++
//    // }
//    // for i := 0; i < H; i++ {
//    //    println(util.PadLeft(fmt.Sprint(i), 3), util.Repeat("#",
//    //       int(math2.Ceil(math2.Log2(float64(C[i]))))))
//    // }
// }
//
// func printStats(k int) {
//    println()
//    println("1 rotations", uint64(float64(SINGLE_ROTATIONS) / float64(k)))
//    println("2 rotations", uint64(float64(DOUBLE_ROTATIONS) / float64(k)))
//    println("s rotations", uint64(float64(SINGLE_ROTATIONS + DOUBLE_ROTATIONS * 2) / float64(k)))
//    // println("promotions", uint64(float64(PROMOTIONS) / float64(k)))
//    // println("demotions", uint64(float64(DEMOTIONS) / float64(k)))
//    println("comparisons", uint64(float64(COMPARISONS) / float64(k)))
// }
//
// func mixed() {
//    n := list.Size(1_000_000)
//    r := distributions.Normal{}
//    x := list.Value(1)
//    r.Seed(123)
//    k := 5
//    for z := 0; z < k; z++ {
//       t := LBST{}.New()
//       for t.Size() < n {
//           t.Insert(r.LessThan(t.Size()+1), x)
//           t.Insert(r.LessThan(t.Size()+1), x)
//           t.Delete(r.LessThan(t.Size()))
//       }
//       print(".")
//    }
//    printStats(k)
// }
//
// func onlyDelete() {
//    n := list.Size(1_000_000)
//    r := distributions.Normal{}
//    r.Seed(123)
//    k := 5
//    for z := 0; z < k; z++ {
//       t := LBST{}.New(make([]list.Value, n)...)
//       for t.Size() > 0 {
//           t.Delete(r.LessThan(t.Size()))
//       }
//       print(".")
//    }
//    printStats(k)
//
// }
//
// func onlyInsert() {
//    n := list.Size(1_000_000)
//    r := distributions.Normal{}
//    x := list.Value(1)
//    r.Seed(123)
//    k := 5
//    for z := 0; z < k; z++ {
//       t := LBST{}.New()
//       for t.Size() < n {
//           t.Insert(r.LessThan(t.Size()+1), x)
//           x++
//       }
//       print(".")
//    }
//    printStats(k)
// }
//
// func InsertDeleteCyclesTest2() {
//    n := list.Size(10000)
//    r := random.Source(123)
//    x := list.Value(1)
//    t := AVLWeakTopDown{}.New()
//
//    for {
//       for t.Size() < n {
//          t.Insert(random.LessThan(t.Size() + 1, r), x); x++
//       }
//       t.Verify()
//       for t.Size() > n / 2 {
//          i := random.LessThan(t.Size(), r)
//          // spew.Dump(t, i)
//          t.Delete(i)
//          t.Verify()
//       }
//       print(".")
//    }
// }
//
// func InsertDeleteCyclesTest() {
//    t := AVLWeakTopDown{}.New()
//    n := list.Size(10000)
//    r := random.Source(123)
//    for {
//       for t.Size() < n {
//          t.Insert(random.LessThan(t.Size() + 1, r), 0)
//          t.Verify()
//       }
//       m := random.LessThan(n / 2, r) + 1
//       for t.Size() > m {
//          t.Delete(random.LessThan(t.Size(), r))
//          t.Verify()
//       }
//       print(".")
//    }
// }

//
//func sandboxJoinL() {
//
//
//   root :=
//      &Node{rank: 6, s: 8,
//         l: &Node{rank: 5, s: 4,
//            l: &Node{rank: 3, s: 1,
//               l: &Node{rank: 1, s: 0},
//               r: &Node{rank: 2, s: 0,
//                  r: &Node{rank: 1, s: 0},
//               },
//            },
//            r: &Node{rank: 3, s: 1,
//               l: &Node{rank: 1, s: 0},
//               r: &Node{rank: 1, s: 0},
//            },
//         },
//         r: &Node{rank: 5, s: 5,
//            l: &Node{rank: 3, s: 1,
//               l: &Node{rank: 1, s: 0},
//               r: &Node{rank: 2, s: 1,
//                  l: &Node{rank: 1, s: 0},
//                  r: &Node{rank: 1, s: 0},
//               },
//            },
//            r: &Node{rank: 3, s: 3,
//               l: &Node{rank: 2, s: 1,
//                  l: &Node{rank: 1, s: 0},
//                  r: &Node{rank: 1, s: 0},
//               },
//               r: &Node{rank: 2, s: 0,
//                  r: &Node{rank: 1, s: 0},
//               },
//            },
//         },
//      }
//
//
//
//
//
//   // root :=
//   //    &Node{rank: 7, s: 13,
//   //       l: &Node{rank: 5, s: 7,
//   //          l: &Node{rank: 4, s: 3,
//   //             l: &Node{rank: 2, s: 1,
//   //                l: &Node{rank: 1, s: 0},
//   //                r: &Node{rank: 1, s: 0},
//   //             },
//   //             r: &Node{rank: 2, s: 1,
//   //                l: &Node{rank: 1, s: 0},
//   //                r: &Node{rank: 1, s: 0},
//   //             },
//   //          },
//   //          r: &Node{rank: 3, s: 2,
//   //             l: &Node{rank: 2, s: 1,
//   //                l: &Node{rank: 1, s: 0},
//   //             },
//   //             r: &Node{rank: 2, s: 0,
//   //                r: &Node{rank: 1, s: 0},
//   //             },
//   //          },
//   //       },
//   //       r: &Node{rank: 5, s: 6,
//   //          l: &Node{rank: 3, s: 3,
//   //             l: &Node{rank: 2, s: 1,
//   //                l: &Node{rank: 1, s: 0},
//   //                r: &Node{rank: 1, s: 0},
//   //             },
//   //             r: &Node{rank: 2, s: 0,
//   //                r: &Node{rank: 1, s: 0},
//   //             },
//   //          },
//   //          r: &Node{rank: 3, s: 1,
//   //             l: &Node{rank: 1, s: 0},
//   //             r: &Node{rank: 2, s: 1,
//   //                l: &Node{rank: 1, s: 0},
//   //                r: &Node{rank: 1, s: 0},
//   //             },
//   //          },
//   //       },
//   //    }
//
//
//      tree := WAVL{
//         root: root,
//         size: root.count(),
//      }
//      tree.Verify()
//
//      spew.Dump(tree.Delete(1))
//
//      // spew.Dump(tree1)
//      // spew.Dump(tree2)
//
//      tree.Verify()
//
//
//   // rand.Seed(1)
//   // N := 64
//   // x := uint64(1)
//   // for {
//   //    t := WAVL{}.New()
//   //    for t.Size() < Size(N) {
//   //       t = t.Clone()
//   //       t.Insert(Pos(rand.Intn(int(t.Size()+1))), x); x++
//   //    }
//   //    t.Verify()
//   //    println("=====")
//   //
//   //    for t.Size() > Size(N / 4) {
//   //        t = t.Clone()
//   //        i := Pos(rand.Intn(int(t.Size())))
//   //        println("------")
//   //        println("REMOVING", i)
//   //        spew.Dump(t)
//   //        s := t.(*WAVL).Delete(i)
//   //        println("after removal", s)
//   //        spew.Dump(t.(*WAVL))
//   //        t.Verify()
//   //    }
//   //
//   //    for t.Size() < Size(N) {
//   //       t = t.Clone()
//   //       t.Insert(Pos(rand.Intn(int(t.Size()+1))), x); x++
//   //    }
//   //    t.Verify()
//   //    println("=====")
//   //    for t.Size() > Size(N / 4) {
//   //        i := Pos(rand.Intn(int(t.Size())))
//   //        println("------")
//   //        println("REMOVING", i)
//   //        spew.Dump(t)
//   //        s := t.(*WAVL).Delete(i)
//   //        println("after removal", s)
//   //        spew.Dump(t.(*WAVL))
//   //        t.Verify()
//   //    }
//   // }
//
//   // for {
//   //    l := WAVL{}.New().(*WAVL)
//   //    r := WAVL{}.New().(*WAVL)
//   //
//   //    for n := rand.Intn(N); n >= 0; n-- {
//   //        l.Insert(abstract.Pos(rand.Intn(int(l.Size()+1))), x); x++
//   //    }
//   //    for n := rand.Intn(N); n >= 0; n-- {
//   //        r.Insert(abstract.Pos(rand.Intn(int(r.Size()+1))), x); x++
//   //    }
//   //    // if l.root.rank >= r.root.rank && l.root.rank - r.root.rank > 2 {
//   //       p := l.Join(r)
//   //       p.Verify()
//   //    // }
//   //    print(".")
//   // }
//
//
//   // for {
//   //    l := WAVL{}.New().(*WAVL)
//   //    r := WAVL{}.New().(*WAVL)
//   //
//   //    for n := rand.Intn(N); n >= 0; n-- {
//   //       l.Insert(abstract.Pos(rand.Intn(int(l.Size()+1))), x); x++
//   //    }
//   //    for n := rand.Intn(N); n >= 0; n-- {
//   //       r.Insert(abstract.Pos(rand.Intn(int(r.Size()+1))), x); x++
//   //    }
//   //    p := l.Join(r)
//   //    p.Verify()
//   //    // assert(p.Size() == l.Size() + r.Size())
//   //    // assert(p.Size() == p.(*WAVL).root.count())
//   //    print(".")
//   // }
//}
}