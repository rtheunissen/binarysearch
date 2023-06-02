package console

import (
   "io"
   "math"
   "strconv"
)

type StackedHistogram struct {
   Title  string
   Series [2][]uint64
   Width  int
   Height int
}

func (g StackedHistogram) Print(page io.Writer) {
   L := g.Series[0]
   R := g.Series[1]

   Println(page, " ╭ ", g.Title)
   Println(page, " │")

   for row := 0; row < g.Height; row++ {

      // Determine the width of the left and right bars for this row.
      barWidthL := 0
      barWidthR := 0
      if row < len(L) {
         barWidthL = int(math.Log2(float64(L[row])) + 1)
      }
      if row < len(R) {
         barWidthR = int(math.Log2(float64(R[row])) + 1)
      }

      Print(page, " │", PadLeft(strconv.Itoa(row), 4))
      Print(page, Repeat(" ", g.Width-barWidthL+1))
      if barWidthL > 0 {
         Print(page, "▕", Repeat("░", barWidthL))
      }
      if barWidthR > 0 {
         Print(page, Repeat("▓", barWidthR), "▏")
      }
      Println(page)
   }
   Println(page, " ╰")
}
