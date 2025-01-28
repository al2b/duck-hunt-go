package engine

import (
	"image"
	"image/draw"
	"slices"
)

type Sprite interface {
	Point
	Image
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
		img := sprite.Image24()
		if img != nil {
			draw.Draw(dst, dst.Bounds(), img, image.Point{
				X: -int(sprite.X()),
				Y: -int(sprite.Y()),
			}, draw.Over)
		}
	}

	return dst
}

type CoordinatedSprite struct {
	Coordinates
	Image
}
