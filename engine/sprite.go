package engine

import (
	"image"
	"image/draw"
	"slices"
)

type Sprite interface {
	Point
	Image() image.Image
}

type Sprites []Sprite

func (s Sprites) Append(sprites Sprites) Sprites {
	return append(s, sprites...)
}

func (s Sprites) Flatten(width, height int) *image.NRGBA {
	dst := image.NewNRGBA(image.Rect(0, 0, width, height))

	// Sort by depth (z coordinate)
	slices.SortStableFunc(s, func(a, b Sprite) int {
		return int(a.Z() - b.Z())
	})

	for _, sprite := range s {
		img := sprite.Image()
		if img != nil {
			draw.Draw(dst, dst.Bounds(), img, image.Point{
				X: -int(sprite.X()),
				Y: -int(sprite.Y()),
			}, draw.Over)
		}
	}

	return dst
}

func NewCoordinatedSprite(coordinates Coordinates, image image.Image) *CoordinatedSprite {
	return &CoordinatedSprite{
		coordinates: coordinates,
		image:       image,
	}
}

type CoordinatedSprite struct {
	coordinates Coordinates
	image       image.Image
}

func (s *CoordinatedSprite) X() float64 {
	return s.coordinates.X()
}

func (s *CoordinatedSprite) Y() float64 {
	return s.coordinates.Y()
}

func (s *CoordinatedSprite) Z() float64 {
	return s.coordinates.Z()
}

func (s *CoordinatedSprite) Image() image.Image {
	return s.image
}
