package console

import (
   . "binarysearch/utility"
   "io"
)

type Table struct {
   Title   string
   Columns []string
   Labels  []string
   Values  [][]string
}

const Padding = 4

func (table Table) Print(page io.Writer) {
   //
   //
   columnWidth := LengthOfLongestString(table.Columns)

   for _, values := range table.Values {
      columnWidth = Max(columnWidth, LengthOfLongestString(values))
   }

   //
   maxLabelWidth := CharacterCount(table.Title)
   maxLabelWidth = Max(maxLabelWidth, LengthOfLongestString(table.Labels))

   //
   // TITLE
   //
   Print(page, Repeat(" ", Padding))
   Print(page, Bold(PadRight(table.Title, maxLabelWidth)))

   //
   // COLUMN HEADERS
   //
   for _, columnLabel := range table.Columns {
      Print(page, Repeat(" ", Padding))
      Print(page, PadLeft(columnLabel, columnWidth))
   }
   Println(page)
   Print(page, Repeat(" ", Padding), Repeat("─", maxLabelWidth))
   Print(page, Repeat("─", (columnWidth+Padding)*len(table.Columns)))
   Println(page)

   //
   // ROWS
   //
   for row, values := range table.Values {
      //
      // LABELS
      //
      Print(page, Repeat(" ", Padding))
      Print(page, Italic(PadRight(table.Labels[row], maxLabelWidth)))

      //
      // VALUES
      //
      for _, value := range values {
         Print(page, Repeat(" ", Padding))
         Print(page, PadLeft(value, columnWidth))
      }
      Println(page)
   }
   Println(page)
}
