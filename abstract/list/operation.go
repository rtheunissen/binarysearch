package list

import (
	"trees/distribution"
)

type Operation interface {
	Setup(List, Size) List
	Valid(List, Size) bool // TODO: I think redundant

	// TODO: kinda liked the index() call because everything just uses one i
	// then we don't need to return position here because the index is passed in.
	Update(List, distribution.Distribution) (List, Position)
}
