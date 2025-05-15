package engine

import (
	"fmt"
)

type Size struct {
	Width, Height int
}

func (s Size) String() string {
	return fmt.Sprintf("%dx%d", s.Width, s.Height)
}

func (s Size) IsZero() bool {
	return s.Width == 0 && s.Height == 0
}

func (s Size) Add(size Size) Size {
	return Size{
		s.Width + size.Width,
		s.Height + size.Height,
	}
}

func (s Size) Sub(size Size) Size {
	return Size{
		s.Width - size.Width,
		s.Height - size.Height,
	}
}
