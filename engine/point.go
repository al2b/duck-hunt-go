package engine

import "image"

type Point image.Point

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Point() Point {
	return p
}

type Pointer interface {
	Point() Point
}

func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}

type PointAdder []Pointer

func (a PointAdder) Point() (point Point) {
	for _, pointer := range a {
		point = point.Add(pointer.Point())
	}
	return
}
