package engine

import "fmt"

type Size struct {
	Width, Height int
}

func (s Size) String() string {
	return fmt.Sprintf("%dx%d", s.Width, s.Height)
}

type Sized interface {
	Size() Size
}

type AbsoluteSize Size

func (s AbsoluteSize) Size() Size {
	return Size(s)
}
