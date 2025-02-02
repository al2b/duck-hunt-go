package engine

import (
	"image"
	"image/draw"
	"slices"
)

type Sprite interface {
	Positionable
	Image() image.Image
}

type Sprites []Sprite

func (s Sprites) Append(sprites ...Sprite) Sprites {
	return append(s, sprites...)
}

func (s Sprites) Appends(sprites ...Sprites) Sprites {
	for _, sprites := range sprites {
		s = append(s, sprites...)
	}
	return s
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
		Coordinates: coordinates,
		image:       image,
	}
}

type CoordinatedSprite struct {
	Coordinates
	image image.Image
}

func (s *CoordinatedSprite) Image() image.Image {
	return s.image
}
