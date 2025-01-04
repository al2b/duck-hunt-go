package draw

import (
	"duck-hunt-go/engine"
	"image"
	"image/color"
)

func Segment(point0, point1 image.Point, color color.Color) SegmentDrawer {
	return SegmentDrawer{
		point0: point0,
		point1: point1,
		color:  color,
	}
}

type SegmentDrawer struct {
	point0, point1 image.Point
	color          color.Color
}

func (drawer SegmentDrawer) Draw(image *engine.Image) {
	dx := engine.Abs(drawer.point1.X - drawer.point0.X)
	dy := engine.Abs(drawer.point1.Y - drawer.point0.Y)
	sx, sy := 1, 1
	if drawer.point0.X >= drawer.point1.X {
		sx = -1
	}
	if drawer.point0.Y >= drawer.point1.Y {
		sy = -1
	}
	err := dx - dy
	for {
		image.Set(drawer.point0.X, drawer.point0.Y, drawer.color)
		if drawer.point0.X == drawer.point1.X && drawer.point0.Y == drawer.point1.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			drawer.point0.X += sx
		}
		if e2 < dx {
			err += dx
			drawer.point0.Y += sy
		}
	}
}
