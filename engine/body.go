package engine

import (
	"github.com/solarlune/resolv"
	"image"
)

func NewBody(position *Position, onIntersect BodyIntersection) *Body {
	return &Body{
		position:    position,
		onIntersect: onIntersect,
	}
}

type Body struct {
	position      *Position
	shape         BodyShape
	Intersections Intersections
	onIntersect   BodyIntersection
}

func (b *Body) Shape(shape BodyShape) *Body {
	b.shape = shape
	return b
}

func (b *Body) ResolvShape() *resolv.ConvexPolygon {
	var points []float64

	for _, sp := range b.shape {
		points = append(points, sp.X, sp.Y)
	}

	return resolv.NewConvexPolygon(
		b.position.X,
		b.position.Y,
		points,
	)
}

func (b *Body) Sprite8() *Sprite8 {
	img := image.NewPaletted(image.Rect(
		0,
		0,
		b.shape.Width(),
		b.shape.Height(),
	), nil)

	c := Color8Green
	if len(b.Intersections) > 0 {
		c = Color8Red
	}

	for i := 0; i < len(b.shape)-1; i++ {
		ImageLine8(img,
			int(b.shape[i].X), int(b.shape[i].Y),
			int(b.shape[i+1].X), int(b.shape[i+1].Y),
			c,
		)
	}

	ImageLine8(img,
		int(b.shape[len(b.shape)-1].X), int(b.shape[len(b.shape)-1].Y),
		int(b.shape[0].X), int(b.shape[0].Y),
		c,
	)

	// Intersections
	for _, i := range b.Intersections {
		for _, it := range i.IntersectionSet.Intersections {
			ImageLine8(img,
				Round(it.Point.X-b.position.X), Round(it.Point.Y-b.position.Y),
				Round(it.Point.X+(it.Normal.X*20)-b.position.X), Round(it.Point.Y+(it.Normal.Y*20)-b.position.Y),
				Color8Blue,
			)
		}
	}

	return &Sprite8{
		Position: b.position,
		Image:    img,
	}
}

func (b *Body) Sprite24() *Sprite24 {
	img := image.NewNRGBA(image.Rect(
		0,
		0,
		b.shape.Width(),
		b.shape.Height(),
	))

	c := Color24Green
	if len(b.Intersections) > 0 {
		c = Color24Red
	}

	for i := 0; i < len(b.shape)-1; i++ {
		ImageLine24(img,
			int(b.shape[i].X), int(b.shape[i].Y),
			int(b.shape[i+1].X), int(b.shape[i+1].Y),
			c,
		)
	}

	ImageLine24(img,
		int(b.shape[len(b.shape)-1].X), int(b.shape[len(b.shape)-1].Y),
		int(b.shape[0].X), int(b.shape[0].Y),
		c,
	)

	// Intersections
	for _, i := range b.Intersections {
		for _, it := range i.IntersectionSet.Intersections {
			ImageLine24(img,
				Round(it.Point.X-b.position.X), Round(it.Point.Y-b.position.Y),
				Round(it.Point.X+(it.Normal.X*20)-b.position.X), Round(it.Point.Y+(it.Normal.Y*20)-b.position.Y),
				Color24Blue,
			)
		}
	}

	return &Sprite24{
		Position: b.position,
		Image:    img,
	}
}

type BodyIntersection func()

type Bodies []*Body

func (b *Bodies) Append(bodies ...*Body) Bodies {
	*b = append(*b, bodies...)
	return *b
}

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
