package console

import (
   "io"
   "math"
   "strconv"
   "bst/utility"
)

var maxHistogramBarWidth = 120

type Histogram struct {
   Series [2][]int
   Height int
   Offset int
}

func (g Histogram) Print(page io.Writer) {
   //
   //
   l := g.Series[0]
   r := g.Series[1]
   h := g.Height

   // The number of rows we would ideally like to draw for each side.
   numberOfRowsL := len(l)
   numberOfRowsR := len(r)

   // Because of the offset, from the middle, we calculate how many of those
   // rows we have room to draw. As the offset increases, the image moves to
   // the right (or down), increasing the capacity of the left side.
   //
   // However, that capacity might exceed the total height.

   capacityL := int(math.Floor(float64(h)/2)) + g.Offset
   capacityR := int(math.Ceil(float64(h)/2)) - g.Offset

   l = l[utility.Max(0, utility.Min(len(l), len(l)-capacityL)):]
   l = l[:utility.Min(len(l), utility.Max(0, len(l)+capacityR))]

   r = r[utility.Min(len(r), utility.Max(0, 0-capacityL)):]
   r = r[:utility.Min(len(r), utility.Max(0, 0+capacityR))]

   // Calculate padding to keep the graphic vertically centered.
   paddingTop := 0
   paddingBot := 0

   // Do we have some empty space?
   if emptySpace := h - (len(l) + len(r)); emptySpace > 0 {
      paddingTop = utility.Max(0, utility.Min(emptySpace, capacityL-numberOfRowsL))
      paddingBot = utility.Max(0, utility.Min(emptySpace, capacityR-numberOfRowsR))
   }
   //
   Println(page)
   Println(page, " ┌ ") // ╭

   //
   for ; paddingTop >= 0; paddingTop-- {
      Println(page, " │")
   }
   //
   max := 0
   for _, width := range l {
      max = utility.Max(max, width)
      Print(page, " │")
      //console.Print(page, PadLeft(strconv.Itoa(width), 4))
      //console.Print(page, " ")
      Print(page, truncatedBar("░", width), "▏")
      Println(page)
   }
   //
   for _, width := range r {
      max = utility.Max(max, width)
      Print(page, " │")
      //console.Print(page, PadLeft(strconv.Itoa(width), 4))
      //console.Print(page, " ") ░   ▒   ▓
      Print(page, truncatedBar("▓", width), "▏")
      Println(page)
   }
   //
   for ; paddingBot >= 0; paddingBot-- {
      Println(page, " │")
   }
   //
   Print(page, " └") // └╰
   Print(page, utility.Repeat("─", max), "┤ ", strconv.Itoa(max))
   Println(page)
}

func truncatedBar(char string, width int) string {
   if width > maxHistogramBarWidth {
      return utility.Repeat(char, maxHistogramBarWidth) + "···"
   } else {
      return utility.Repeat(char, width)
   }
}
