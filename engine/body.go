package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

type Body interface {
	Point
	Shape() BodyShape
	Update(msg tea.Msg) tea.Cmd
	Intersections() Intersections
}

type Bodies []Body

func (b Bodies) Append(bodies ...Body) Bodies {
	return append(b, bodies...)
}

func (b Bodies) Appends(bodies ...Bodies) Bodies {
	for _, bodies := range bodies {
		b = append(b, bodies...)
	}
	return b
}

func NewCoordinatedBody(coordinates Coordinates, shape BodyShape) *CoordinatedBody {
	return &CoordinatedBody{
		Coordinates: coordinates,
		shape:       shape,
	}
}

type CoordinatedBody struct {
	Coordinates
	shape         BodyShape
	intersections Intersections
}

func (b *CoordinatedBody) Shape() BodyShape {
	return b.shape
}

func (b *CoordinatedBody) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case IntersectionMsg:
		b.intersections = append(b.intersections, msg.Intersection)
	}
	return nil
}

func (b *CoordinatedBody) Intersections() Intersections {
	return b.intersections
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
	body Body
}

func (s *BodyImage) Image() image.Image {
	shape := s.body.Shape()

	img := image.NewNRGBA(image.Rect(
		0,
		0,
		shape.Width(),
		shape.Height(),
	))

	c := ColorGreen

	intersections := s.body.Intersections()

	if len(intersections) > 0 {
		c = ColorRed
	}

	for i := 0; i < len(shape)-1; i++ {
		ImageLine(img,
			int(shape[i].X), int(shape[i].Y),
			int(shape[i+1].X), int(shape[i+1].Y),
			c,
		)
	}

	ImageLine(img,
		int(shape[len(shape)-1].X), int(shape[len(shape)-1].Y),
		int(shape[0].X), int(shape[0].Y),
		c,
	)

	// Intersections
	for _, i := range intersections {
		for _, it := range i.IntersectionSet.Intersections {
			ImageLine(img,
				Round(it.Point.X-s.body.X()), Round(it.Point.Y-s.body.Y()),
				Round(it.Point.X+(it.Normal.X*20)-s.body.X()), Round(it.Point.Y+(it.Normal.Y*20)-s.body.Y()),
				ColorBlue,
			)
		}
	}

	return img
}
