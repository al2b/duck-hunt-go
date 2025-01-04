package engine

import (
	"image"
)

type Sprite8 struct {
	Position *Position
	Image    image.PalettedImage
}

type Sprites8 []*Sprite8

func (s *Sprites8) Append(sprites ...*Sprite8) Sprites8 {
	*s = append(*s, sprites...)
	return *s
}

type Sprite24 struct {
	Position *Position
	Image    image.Image
}

type Sprites24 []*Sprite24

func (s *Sprites24) Append(sprites ...*Sprite24) Sprites24 {
	*s = append(*s, sprites...)
	return *s
}
