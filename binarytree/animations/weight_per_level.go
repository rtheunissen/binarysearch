package animations

import (
	"binarysearch/binarytree"
	console2 "binarysearch/console"
	"binarysearch/utility"
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

func (animation WeightsPerLevel) getGraphics() console2.Graphics {
	if animation.Frame == 0 {
		return console2.Graphics{
			console2.Clear,
			console2.Text(animation.Introduction()),
		}
	} else {
		return console2.Graphics{
			console2.Clear,
			console2.StackedHistogram{
				Title:  "Number of nodes per level, log2",
				Series: animation.List.(binarytree.BinaryTree).SymmetricWeightPerLevel(),
				Height: animation.Height,
				Width:  int(math.Log2(float64(animation.List.Size()))) + 1,
			},
			console2.FocusBar{
				Focus: animation.Position,
				Total: animation.List.Size(),
			},
			console2.Line,
			console2.Details{
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
