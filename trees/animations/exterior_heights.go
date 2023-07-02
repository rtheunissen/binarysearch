package animations

import (
   "bst/trees"
   "bst/utility"
   "bst/utility/console"
   "golang.org/x/text/language"
   "golang.org/x/text/message"
)

type ExteriorHeights struct {
   BinaryTreeAnimation
   heights [2][]int
   frame   uint64
}

func (animation *ExteriorHeights) Introduction() string {
   return `
     This animation is an effective way to visualize the shape of the tree.
   
     The height of a node is the number of levels from top to bottom when drawn
     symmetrically. The height is also the length of the longest path, plus one.
   
     The height of nil is 0, and the height of a node with no subtrees is 1.
   
     The left spine of a binary tree is the left-most path from the root, and
     the right spine is the right-most path from the root.
   
     The _exterior height_ of a tree is defined here as the height of each node
     along the spine. The difference between exterior and interior height is that
     the exterior height includes all reachable nodes. Notice in the example below
     that the node to the left of the root has an exterior height of 6 because the
     path L R R R R accesses 6 nodes. The interior height of that node would be 3.
   
   
                                         ROOT
        1    2        5           6       ▿       4           3        2    1
        ● <─ ● <───── ● <──────── ● <──── ○ ────> ● ────────> ● ─────> ● ─> ●
              ╲        ╲           ╲             ╱           ╱
               ●        ●           ●           ●           ●
                      ╱   ╲       ╱   ╲       ╱   ╲
                     ●     ●     ●     ●     ●     ●
                         ╱   ╲
                        ●     ●
                                ╲
                                 ●
   
   
     Using the example above, this is what we would like to draw:
   
           { 1, 2, 5, 6, 4, 3, 2, 1 }
   
           ╭
           │  1 ▒
           │  2 ▒▒
           │  5 ▒▒▒▒▒
           │  6 ▒▒▒▒▒▒
           │  4 ▓▓▓▓
           │  3 ▓▓▓
           │  2 ▓▓
           │  1 ▓
           ╰`
}

func (animation *ExteriorHeights) Update() {
   animation.BinaryTreeAnimation.Update()
   //
   //
   //
   if animation.BinaryTreeAnimation.Frame != animation.frame {
      animation.heights = animation.List.(trees.BinaryTree).ExteriorHeightsAlongTheSpines()
      animation.frame = animation.BinaryTreeAnimation.Frame
   }
}

func (animation *ExteriorHeights) Render() {
   animation.getGraphics().Print(animation)
}

func (animation *ExteriorHeights) getGraphics() console.Graphics {
   if animation.Frame == 0 {
      return console.Graphics{
         console.Clear,
         console.Text(animation.Introduction()),
      }
   }
   return console.Graphics{
      console.Clear,
      console.Histogram{
         Series: animation.heights,
         Height: animation.Height,
         Offset: animation.Offset,
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
            "Heights along the spines",
            utility.NameOf(animation.List),
            utility.NameOf(animation.Operation),
            utility.NameOf(animation.Distribution),
            message.NewPrinter(language.English).Sprint(animation.List.Size()),
         },
      },
   }
}
