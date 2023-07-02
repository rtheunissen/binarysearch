package animations

import (
   "bst/trees"
   "bst/utility"
   "bst/utility/console"
)

type InteriorHeights struct {
   BinaryTreeAnimation
}

func (InteriorHeights) Introduction() string {
   return `

      This animation is an effective way to visualize the shape of the tree.
      
      The height of a node is the number of levels from top to bottom when drawn
      symmetrically. The height is also the length of the longest path, plus one.
      
      The height of nil is 0, and the height of a node with no subtrees is 1.
      
      The left spine of a binary tree is the left-most path from the root, and
      the right spine is the right-most path from the root.
      
      The _interior height_ of a tree is defined here as the height of each node
      along the spine branching only "inward" from the spine.
      
      This can be visualized by pretending to pull the left- and rightmost nodes
      upward, suspending horizontal spines. The interior heights are then the
      heights of all nodes that start on a down-link from the spine, plus one:
      
      
                                         ROOT
        1    2        3           4           ▿   5           3        2    1
        ● <─ ● <───── ● <──────── ● <──── ○ ────> ● ────────> ● ─────> ● ─> ● ┄┄ 1
             │        │           │               │           │        │
             ●        ●           ●               ●           ●        ● ┄┄┄┄┄┄┄ 2
                    ╱   ╲       ╱   ╲           ╱   ╲           ╲
                   ●     ●     ●     ●         ●     ●           ● ┄┄┄┄┄┄┄┄┄┄┄┄┄ 3
                                ╲   ╱ ╲       ╱ ╲   ╱ ╲
                                 ● ●   ●     ●   ● ●   ● ┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄ 4
                                                      ╱
                                                     ● ┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄ 5
      
      
      
      Using the example above, this is what we would like to draw:
      
            H := { 1, 2, 3, 4, 5, 3, 2, 1 }
      
            ╭
            │  1 ▒
            │  2 ▒▒
            │  3 ▒▒▒
            │  4 ▒▒▒▒
            │  5 ▓▓▓▓▓
            │  3 ▓▓▓
            │  2 ▓▓
            │  1 ▓
            ╰`
}

func (animation InteriorHeights) Render() {
   animation.getGraphics().Print(animation)
}

func (animation InteriorHeights) getGraphics() console.Graphics {
   if animation.Frame == 0 {
      return console.Graphics{
         console.Clear,
         console.Text(animation.Introduction()),
      }
   }
   return console.Graphics{
      console.Clear,
      console.Histogram{
         Series: animation.List.(trees.BinaryTree).InteriorHeightsAlongTheSpines(),
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
            utility.NameOf(animation),
            utility.NameOf(animation.List),
            utility.NameOf(animation.Operation),
            utility.NameOf(animation.Distribution),
            utility.CommaSep(animation.List.Size()),
         },
      },
   }
}
