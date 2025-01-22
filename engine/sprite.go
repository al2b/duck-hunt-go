package engine

import (
	"image"
)

type Sprite interface {
	Point
	Image8() image.PalettedImage
	Image24() image.Image
}

type Sprites []Sprite

func (s Sprites) Append(sprites ...Sprite) Sprites {
	s = append(s, sprites...)
	return s
}

func (s Sprites) Appends(sprites ...Sprites) Sprites {
	for _, sprites := range sprites {
		s = append(s, sprites...)
	}
	return s
}

func NewUniformSprite(coordinates Point, color8 Color8, color24 Color24) *UniformSprite {
	return &UniformSprite{
		Point:   coordinates,
		color8:  color8,
		color24: color24,
	}
}

type UniformSprite struct {
	Point
	color8  Color8
	color24 Color24
}

func (s *UniformSprite) Image8() image.PalettedImage {
	return NewImageUniform8(s.color8)
}

func (s *UniformSprite) Image24() image.Image {
	return NewImageUniform24(s.color24)
}

func NewImageSprite(coordinates Point, image8 image.PalettedImage, image24 image.Image) *ImageSprite {
	return &ImageSprite{
		Point:   coordinates,
		image8:  image8,
		image24: image24,
	}
}

type ImageSprite struct {
	Point
	image8  image.PalettedImage
	image24 image.Image
}

func (s *ImageSprite) Image8() image.PalettedImage {
	return s.image8
}

func (s *ImageSprite) Image24() image.Image {
	return s.image24
}
