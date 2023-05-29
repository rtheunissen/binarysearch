package console

import (
	"io"
	"trees/utility"
)

type Details struct {
	Labels []string
	Values []string
}

func (g Details) Print(page io.Writer) {
	maxLength := utility.LengthOfLongestString(g.Labels)
	for i, label := range g.Labels {
		Print(page, "\n ", utility.PadRight(label, maxLength+1), "  ", g.Values[i])
	}
	Println(page)
}
