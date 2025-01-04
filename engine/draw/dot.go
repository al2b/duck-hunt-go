package draw

import (
	"duck-hunt-go/engine"
	"image"
	"image/color"
)

func Dot(point image.Point, color color.Color) DotDrawer {
	return DotDrawer{
		point: point,
		color: color,
	}
}

type DotDrawer struct {
	point image.Point
	color color.Color
}

func (drawer DotDrawer) Draw(image *engine.Image) {
	image.Set(drawer.point.X, drawer.point.Y, drawer.color)
}
