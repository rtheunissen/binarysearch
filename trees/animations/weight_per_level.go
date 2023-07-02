package animations

import (
   "bst/trees"
   "bst/utility"
   "bst/utility/console"
   "golang.org/x/text/language"
   "golang.org/x/text/message"
   "math"
)

type WeightsPerLevel struct {
   BinaryTreeAnimation
}

func (WeightsPerLevel) Introduction() string {
   return `
      
      This animation uses a log scale to draw the number of nodes per level on the
      left and right side of the root of a tree.
      
      ┌
      │              ▒▓
      │             ▒▒▓▓
      │            ▒▒▒▓▓▓
      │           ▒▒▒▒▓▓▓▓
      │          ▒▒▒▒▒▓▓▓▓▓
      │         ▒▒▒▒▒▒▓▓▓▓▓▓
      │        ▒▒▒▒▒▒▒▓▓▓▓▓▓▓
      │       ▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓
      │      ▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓
      │      ▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓
      │     ▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓
      │    ▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓
      │    ▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓
      │     ▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓
      │       ▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓
      │          ▒▒▒▒▒▓▓
      │
      │
      │
      │
      └`
}

func (animation WeightsPerLevel) Render() {
   animation.getGraphics().Print(animation)
}

func (animation WeightsPerLevel) getGraphics() console.Graphics {
   if animation.Frame == 0 {
      return console.Graphics{
         console.Clear,
         console.Text(animation.Introduction()),
      }
   } else {
      return console.Graphics{
         console.Clear,
         console.StackedHistogram{
            Title:  "Number of nodes per level, log2",
            Series: animation.List.(trees.BinaryTree).SymmetricWeightPerLevel(),
            Height: animation.Height,
            Width:  int(math.Log2(float64(animation.List.Size()))) + 1,
         },
         console.FocusBar{
            Focus: animation.Position,
            Total: animation.List.Size(),
         },
         console.Line,
         console.Details{
            Labels: []string{
               "Animation",
               "Strategy",
               "Operation",
               "Distribution",
               "Size",
            },
            Values: []string{
               utility.NameOf(animation),
               utility.NameOf(animation.List),
               utility.NameOf(animation.Operation),
               utility.NameOf(animation.Distribution),
               message.NewPrinter(language.English).Sprint(animation.List.Size()),
            },
         },
      }
   }
}
