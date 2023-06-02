package console

import (
   "io"
   "math"
)

type FocusBar struct {
   Focus uint64
   Total uint64
}

func (f FocusBar) Print(page io.Writer) {
   if f.Total == 0 {
      return
   }
   w := int(math.Log2(float64(f.Total))) + 1
   x := int(float64(f.Focus) / float64(f.Total) * float64(w))

   if x > w {
      x = 0
   }
   Print(page, " ")
   Print(page, Repeat("░", x))
   Print(page, Repeat("▓", 1))
   Print(page, Repeat("░", w-x-1))
   Print(page, " ")
}
