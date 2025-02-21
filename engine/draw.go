package engine

import (
	"image"
	"image/color"
	"image/draw"
)

type Drawer interface {
	Draw(*Image)
}

/*********/
/* Image */
/*********/

func DrawImage(point image.Point, img *Image) ImageDrawer {
	return ImageDrawer{
		point: point,
		image: img,
	}
}

func DrawCenteredImage(point image.Point, img *Image) ImageDrawer {
	size := img.Size()
	return DrawImage(
		point.Sub(image.Pt(
			(size.Width-1)/2,
			(size.Height-1)/2,
		)),
		img,
	)
}

type ImageDrawer struct {
	point image.Point
	image *Image
}

func (drawer ImageDrawer) Draw(img *Image) {
	draw.Draw(
		img.NRGBA,
		drawer.image.Bounds().Add(drawer.point),
		drawer.image.NRGBA,
		image.Point{},
		draw.Over,
	)
}

/*******/
/* Dot */
/*******/

func DrawDot(point image.Point, color color.Color) DotDrawer {
	return DotDrawer{
		point: point,
		color: color,
	}
}

type DotDrawer struct {
	point image.Point
	color color.Color
}

func (drawer DotDrawer) Draw(img *Image) {
	img.Set(drawer.point.X, drawer.point.Y, drawer.color)
}

/***********/
/* Segment */
/***********/

func DrawSegment(point0, point1 image.Point, color color.Color) SegmentDrawer {
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

func (drawer SegmentDrawer) Draw(img *Image) {
	dx := Abs(drawer.point1.X - drawer.point0.X)
	dy := Abs(drawer.point1.Y - drawer.point0.Y)
	sx, sy := 1, 1
	if drawer.point0.X >= drawer.point1.X {
		sx = -1
	}
	if drawer.point0.Y >= drawer.point1.Y {
		sy = -1
	}
	err := dx - dy
	for {
		img.Set(drawer.point0.X, drawer.point0.Y, drawer.color)
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

/**********/
/* Circle */
/**********/

func DrawCircle(point image.Point, radius int, color color.Color) CircleDrawer {
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

func (drawer CircleDrawer) Draw(img *Image) {
	x, y := drawer.radius, 0
	d := 1 - drawer.radius

	for x >= y {
		img.Set(drawer.point.X+x, drawer.point.Y+y, drawer.color)
		img.Set(drawer.point.X-x, drawer.point.Y+y, drawer.color)
		img.Set(drawer.point.X+x, drawer.point.Y-y, drawer.color)
		img.Set(drawer.point.X-x, drawer.point.Y-y, drawer.color)
		img.Set(drawer.point.X+y, drawer.point.Y+x, drawer.color)
		img.Set(drawer.point.X-y, drawer.point.Y+x, drawer.color)
		img.Set(drawer.point.X+y, drawer.point.Y-x, drawer.color)
		img.Set(drawer.point.X-y, drawer.point.Y-x, drawer.color)
		y++
		if d < 0 {
			d += 2*y + 1
		} else {
			x--
			d += 2*(y-x) + 1
		}
	}
}
