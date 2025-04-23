package engine

import (
	"image/color"
)

/*******/
/* Dot */
/*******/

type Dot struct {
	Point Point
	Color color.Color
}

func (dot Dot) Draw(dst *Image) {
	dst.Set(dot.Point, dot.Color)
}

/***********/
/* Segment */
/***********/

type Segment struct {
	Point0, Point1 Point
	Color          color.Color
}

func (segment Segment) Draw(dst *Image) {
	dx := Abs(segment.Point1.X - segment.Point0.X)
	dy := Abs(segment.Point1.Y - segment.Point0.Y)
	sx, sy := 1, 1
	if segment.Point0.X >= segment.Point1.X {
		sx = -1
	}
	if segment.Point0.Y >= segment.Point1.Y {
		sy = -1
	}
	err := dx - dy
	for {
		dst.Set(segment.Point0, segment.Color)
		if segment.Point0.X == segment.Point1.X && segment.Point0.Y == segment.Point1.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			segment.Point0.X += sx
		}
		if e2 < dx {
			err += dx
			segment.Point0.Y += sy
		}
	}
}

/*************/
/* Rectangle */
/*************/

type Rectangle struct {
	Point Point
	Size  Size
	Color color.Color
}

func (rectangle Rectangle) Draw(dst *Image) {
	dst.Draw(
		Segment{rectangle.Point, rectangle.Point.Add(Pt(rectangle.Size.Width-1, 0)), rectangle.Color},
		Segment{rectangle.Point.Add(Pt(rectangle.Size.Width-1, 0)), rectangle.Point.Add(Pt(rectangle.Size.Width-1, rectangle.Size.Height-1)), rectangle.Color},
		Segment{rectangle.Point.Add(Pt(rectangle.Size.Width-1, rectangle.Size.Height-1)), rectangle.Point.Add(Pt(0, rectangle.Size.Height-1)), rectangle.Color},
		Segment{rectangle.Point.Add(Pt(0, rectangle.Size.Height-1)), rectangle.Point, rectangle.Color},
	)
}

/**********/
/* Circle */
/**********/

type Circle struct {
	Point  Point
	Radius int
	Color  color.Color
}

func (circle Circle) Draw(dst *Image) {
	x, y := circle.Radius, 0
	d := 1 - circle.Radius

	for x >= y {
		dst.Set(circle.Point.Add(Pt(x, y)), circle.Color)
		dst.Set(circle.Point.Add(Pt(-x, y)), circle.Color)
		dst.Set(circle.Point.Add(Pt(x, -y)), circle.Color)
		dst.Set(circle.Point.Add(Pt(-x, -y)), circle.Color)
		dst.Set(circle.Point.Add(Pt(y, x)), circle.Color)
		dst.Set(circle.Point.Add(Pt(-y, x)), circle.Color)
		dst.Set(circle.Point.Add(Pt(y, -x)), circle.Color)
		dst.Set(circle.Point.Add(Pt(-y, -x)), circle.Color)
		y++
		if d < 0 {
			d += 2*y + 1
		} else {
			x--
			d += 2*(y-x) + 1
		}
	}
}
