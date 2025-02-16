package draw

import (
	"duck-hunt-go/engine"
	"image"
	"image/color"
)

func Circle(point image.Point, radius int, color color.Color) CircleDrawer {
	return CircleDrawer{
		point:  point,
		radius: radius,
		color:  color,
	}
}

type CircleDrawer struct {
	point  image.Point
	radius int
	color  color.Color
}

func (drawer CircleDrawer) Draw(image *engine.Image) {
	x, y := drawer.radius, 0
	d := 1 - drawer.radius

	for x >= y {
		image.Set(drawer.point.X+x, drawer.point.Y+y, drawer.color)
		image.Set(drawer.point.X-x, drawer.point.Y+y, drawer.color)
		image.Set(drawer.point.X+x, drawer.point.Y-y, drawer.color)
		image.Set(drawer.point.X-x, drawer.point.Y-y, drawer.color)
		image.Set(drawer.point.X+y, drawer.point.Y+x, drawer.color)
		image.Set(drawer.point.X-y, drawer.point.Y+x, drawer.color)
		image.Set(drawer.point.X+y, drawer.point.Y-x, drawer.color)
		image.Set(drawer.point.X-y, drawer.point.Y-x, drawer.color)
		y++
		if d < 0 {
			d += 2*y + 1
		} else {
			x--
			d += 2*(y-x) + 1
		}
	}
}
