package console

import (
   "strconv"

   "io"
)

type VHistogram struct {
   Series [2][]int
   Height int
   Offset int
   Width  int
}

func (g *VHistogram) Print(page io.Writer) {
   //
   for _, height := range g.Series[0] {
      g.Height = Max(g.Height, height)
   }
   for _, height := range g.Series[1] {
      g.Height = Max(g.Height, height)
   }

   //
   g.Width = Max(g.Width, len(g.Series[0]))

   padding := Repeat(" ", g.Width-len(g.Series[0])+1)

   Println(page)
   Println(page, PadLeft(strconv.Itoa(g.Height), 4), " ┌ ")

   for h := g.Height; h >= 0; h-- {
      //if h == g.Height {
      //   console.Print(page, util.PadLeft(strconv.Itoa(h), 4), " ┬")
      //} else {
      Print(page, Repeat(" ", 4+1), "│")
      //}
      Print(page, padding)

      for _, height := range g.Series[0] {
         if height >= h {
            Print(page, "░")
         } else {
            Print(page, " ")
         }
      } // ░   ▒   ▓
      for _, height := range g.Series[1] {
         if height >= h {
            Print(page, "▓")
         } else {
            Print(page, " ")
         }
      }
      Println(page)
   }
   Println(page, Repeat(" ", 4+1), "└")
   Println(page)
}

//func truncatedBar(char string, width int) string {
//   if width > maxVHistogramBarSize {
//      return Repeat(char, maxVHistogramBarSize) + "···"
//   } else {
//      return Repeat(char, width)
//   }
//}
