package animations

import (
   "github.com/eiannone/keyboard"
   "io"
   "math"
   "trees/abstract/list"
   "trees/distribution"
)

type BinaryTreeAnimation struct {
	Height int
	Offset int
	Frame  uint64
	list.Position
	io.Writer
	list.List
	list.Operation
	distribution.Distribution // remove
	list.Size
}

func (animation *BinaryTreeAnimation) Setup() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	animation.Frame = 0
	animation.List = animation.Operation.Setup(animation.List, animation.Size)
	animation.New(123)
}

func (animation *BinaryTreeAnimation) Close() {
	if err := keyboard.Close(); err != nil {
		panic(err)
	}
	//os.Exit(0)
}

func (animation *BinaryTreeAnimation) waitForKeyPress() keyboard.Key {
	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	return key
}

// This is like steps in the measure, because this is just measure!
//
// Every _step_ number of ticks along the operation, we do a thing, like measure
// and add to a central hub, .. however here the pace increases logarithmically,
// not linearly. Maybe then a "Samples" function also.
//
//func (a *Animation) seekFrame() {
//   //for a.Operation.Valid() {
//   //    a.Operation.Update()
//
//      if nextFrame == 0 || a.Tree.Size() % nextFrame == 0 {
//        break
//      }
//   //}
//}

func (animation *BinaryTreeAnimation) Valid() bool {
	return animation.Operation.Valid(animation.List, animation.Size)
}

func (animation *BinaryTreeAnimation) Update() {
	switch key := animation.waitForKeyPress(); key {

	// " ← " decreases the draw offset of the image within the viewport.
	case keyboard.KeyArrowLeft:
		animation.Offset = int(math.Max(float64(-animation.Height/2), float64(animation.Offset-1)))

	// " → " increases the draw offset of the image within the viewport.
	case keyboard.KeyArrowRight:
		animation.Offset = int(math.Min(float64(+animation.Height/2), float64(animation.Offset+1)))

	// " ↑ " increases the height of the viewport.
	case keyboard.KeyArrowUp:
		animation.Height = animation.Height + 1

	// " ↓ " Decreases the height of the viewport.
	case keyboard.KeyArrowDown:
		animation.Height = int(math.Max(0, float64(animation.Height-1)))

		// Exit
	case keyboard.KeyCtrlC:
		fallthrough
	case keyboard.KeyCtrlD:
		fallthrough
	case keyboard.KeyCtrlZ:
		fallthrough
	case keyboard.KeyCtrlQ:
		fallthrough
	case keyboard.KeyEsc:
		animation.Close()
	default:
		for animation.Valid() {
			animation.List, animation.Position = animation.Operation.Update(animation.List, animation.Distribution)
			if animation.shouldRenderFrame() {
				animation.Frame++
				return
			}
		}
	}
}

//	  1 to   100 : Render every page
//	100 to  1000 : Render every 10th
//
// 1000 to 10000 : Render every 100th etc.
func (animation *BinaryTreeAnimation) shouldRenderFrame() bool {
	nextLog10 := math.Ceil(math.Log10(float64(animation.List.Size() + 1)))
	nextPow10 := math.Pow10(int(nextLog10))
	frameSkip := uint64(nextPow10 / 100)
	return frameSkip == 0 || animation.List.Size()%frameSkip == 0
}

//
//type Animation struct {
//   io.Writer      // Drawing buffer
//   Context
//   Updater
//   Title        string
//   Introduction string
//   Graphics     []console.Graphic //
//   MaximumPathLength       int       // Viewport height
//   Offset       int       // Viewport offset
//}
//
//func (a *Animation) Reset() {
//   //debug.SetGCPercent(-1)
//
//
//   //a.Distribution = a.Distribution.Seed(a.Seed)
//
//   // Set up the BST instance for the animation
//   //a.BST = a.Operation.Reset(a.BST, a.Size)
//}
//
//func (a Animation) Teardown() {
//   os.Exit(0)
//}
//
//func (a *Animation) Valid() bool {
//   return true //a.Operation.Valid(a.BST, a.Size)
//}
//
//func (a *Animation) NextFrame() {
//   //for a.Valid() {
//   //   a.Position = a.Operation.Index(a.BST, a.Distribution)
//   //   a.BST = a.Operation.Apply(a.BST, a.Position)
//   //   //
//   //   //       1 to   100 : Render every page
//   //   //     100 to  1000 : Render every 10th
//   //   //    1000 to 10000 : Render every 100th etc.
//   //   //
//   //   nextLog10 := math.Ceil(math.Log10(float64(a.BST.Size() + 1)))
//   //   nextPow10 := math.Pow10(int(nextLog10))
//   //   nextFrame := uint64(nextPow10 / 200)
//   //
//   //   if nextFrame == 0 || a.BST.Size()%nextFrame == 0 {
//   //      break
//   //   }
//   //}
//}

//
//func (a *Animation) waitForKeyPress() keyboard.Key {
//   _, key, err := keyboard.GetSingleKey()
//   if err != nil {
//      panic(err)
//   }
//   return key
//}
//
//func (a *Animation) getDetails() console.Details {
//   return console.Details{
//      Labels: []string{
//         "Animation",
//         "Focus",
//         "Strategy",
//         "Operation",
//         "Distribution",
//         "Size",
//      },
//      Values: []string{
//         //a.Title,
//         //FocusBar{
//         //   Focus: a.Position,
//         //   Total: a.BST.Size(),
//         //}.Print(),
//         //TypeName(a.BST),
//         //TypeName(a.Operation),
//         //TypeName(a.Distribution),
//         //CommaSep(a.BST.Size()),
//      },
//   }
//}
//
//func (a *Animation) Update() {
//   switch key := a.waitForKeyPress(); key {
//
//   // Applies the operation until the next page should be rendered.
//   case keyboard.KeySpace:
//      fallthrough
//   case keyboard.KeyEnter:
//      a.NextFrame()
//
//   // ← Decreases the draw offset of the image within the viewport.
//   case keyboard.KeyArrowLeft:
//      a.Offset = int(math.Max(float64(-a.MaximumPathLength/2), float64(a.Offset-1)))
//
//   // → Increases the draw offset of the image within the viewport.
//   case keyboard.KeyArrowRight:
//      a.Offset = int(math.Min(float64(+a.MaximumPathLength/2), float64(a.Offset+1)))
//
//   // ↑ Increases the height of the viewport.
//   case keyboard.KeyArrowUp:
//      a.MaximumPathLength = a.MaximumPathLength + 1
//
//   // ↓ Decreases the height of the viewport.
//   case keyboard.KeyArrowDown:
//      a.MaximumPathLength = int(math.Max(0, float64(a.MaximumPathLength-1)))
//
//      // Exit
//   case keyboard.KeyCtrlC:
//      fallthrough
//   case keyboard.KeyCtrlD:
//      fallthrough
//   case keyboard.KeyCtrlZ:
//      fallthrough
//   case keyboard.KeyCtrlQ:
//      fallthrough
//   case keyboard.KeyEsc:
//      a.Teardown()
//   }
//}
//
//func (a *Animation) printToBuffer(content string) {
//   strings.Builder{}.
//   a.Writer.Print(content)
//}
//
//func (a *Animation) resetBuffer() {
//   a.Reset()
//}
//
//func (a *Animation) clearScreen() {
//   a.Print(ClearScreen())
//}
//
//// Write the buffer to the console.
//func (a *Animation) writeBufferToScreen() {
//   _, err := os.Stdout.WriteString(a.String())
//   if err != nil {
//      panic(err)
//   }
//}
//
//func (a *Animation) drawGraphicsToBuffer() {
//   for _, graphic := range a.Graphics { // Print the graphics
//      a.Print(graphic.Print())
//   }
//}
//
//func (a *Animation) printIntro() {
//   a.Print(a.Introduction)
//}
//
//func (a *Animation) Intro() {
//   a.resetBuffer()
//   //a.clearScreen()
//   a.printIntro()
//   a.writeBufferToScreen()
//   a.waitForKeyPress()
//}
//
//func (a *Animation) Render() {
//   a.resetBuffer()
//   //a.clearScreen()
//   a.drawGraphicsToBuffer()
//   a.writeBufferToScreen()
//}
