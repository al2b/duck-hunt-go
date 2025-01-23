package engine

import (
	"github.com/solarlune/resolv"
	"image"
)

func NewBody(coordinates Point, shape BodyShape) *Body {
	return &Body{
		Point: coordinates,
		shape: shape,
	}
}

type Body struct {
	Point
	shape         BodyShape
	Intersections Intersections
}

func (b *Body) ResolvShape() *resolv.ConvexPolygon {
	var points []float64

	for _, sp := range b.shape {
		points = append(points, sp.X, sp.Y)
	}

	return resolv.NewConvexPolygon(
		b.X(),
		b.Y(),
		points,
	)
}

func (b *Body) Image() Image {
	return &BodyImage{
		body: b,
	}
}

type Bodies []*Body

func (b Bodies) Append(bodies ...*Body) Bodies {
	b = append(b, bodies...)
	return b
}

func (b Bodies) Appends(bodies ...Bodies) Bodies {
	for _, bodies := range bodies {
		b = append(b, bodies...)
	}
	return b
}

type BodyIntersection func()

type BodyShapePoint struct {
	X, Y float64
}

type BodyShape []BodyShapePoint

func (s *BodyShape) Width() int {
	var width int
	for _, p := range *s {
		x := int(p.X)
		if x > width {
			width = x
		}
	}
	return width + 1
}

func (s *BodyShape) Height() int {
	var height int
	for _, p := range *s {
		y := int(p.Y)
		if y > height {
			height = y
		}
	}
	return height + 1
}

type BodyImage struct {
	body *Body
}

func (s *BodyImage) Image8() Image8 {
	img := NewImage8(image.Rect(
		0,
		0,
		s.body.shape.Width(),
		s.body.shape.Height(),
	))

	s.image(
		image8Set(img, Color8Green),
		image8Set(img, Color8Red),
		image8Set(img, Color8Blue),
	)

	return img
}

func (s *BodyImage) Image24() Image24 {
	img := NewImage24(image.Rect(
		0,
		0,
		s.body.shape.Width(),
		s.body.shape.Height(),
	))

	s.image(
		image24Set(img, Color24Green),
		image24Set(img, Color24Red),
		image24Set(img, Color24Blue),
	)

	return img
}

func (s *BodyImage) image(set, setIntersected, setIntersection imageSet) {
	if len(s.body.Intersections) > 0 {
		set = setIntersected
	}

	for i := 0; i < len(s.body.shape)-1; i++ {
		imageLine(
			int(s.body.shape[i].X), int(s.body.shape[i].Y),
			int(s.body.shape[i+1].X), int(s.body.shape[i+1].Y),
			set,
		)
	}

	imageLine(
		int(s.body.shape[len(s.body.shape)-1].X), int(s.body.shape[len(s.body.shape)-1].Y),
		int(s.body.shape[0].X), int(s.body.shape[0].Y),
		set,
	)

	// Intersections
	for _, i := range s.body.Intersections {
		for _, it := range i.IntersectionSet.Intersections {
			imageLine(
				Round(it.Point.X-s.body.X()), Round(it.Point.Y-s.body.Y()),
				Round(it.Point.X+(it.Normal.X*20)-s.body.X()), Round(it.Point.Y+(it.Normal.Y*20)-s.body.Y()),
				setIntersection,
			)
		}
	}
}
