package engine

import "image"

type Body interface {
	Positionable
	Shape() Shape
}

type Bodies []Body

func (s Bodies) Append(bodies ...Body) Bodies {
	return append(s, bodies...)
}

func (s Bodies) Appends(bodies ...Bodies) Bodies {
	for _, bodies := range bodies {
		s = append(s, bodies...)
	}
	return s
}

func NewRectangleBody(coordinates Coordinates, shape RectangleShape) RectangleBody {
	return RectangleBody{
		Coordinates:    coordinates,
		RectangleShape: shape,
	}
}

type RectangleBody struct {
	Coordinates
	RectangleShape
}

func NewPolygonBody(coordinates Coordinates, shape PolygonShape) PolygonBody {
	return PolygonBody{
		Coordinates:  coordinates,
		PolygonShape: shape,
	}
}

type PolygonBody struct {
	Coordinates
	PolygonShape
}

type BodySprite struct {
	Body
}

func (s BodySprite) Image() image.Image {
	shape := s.Body.Shape()

	img := image.NewNRGBA(image.Rect(
		0,
		0,
		shape.Width(),
		shape.Height(),
	))

	for i, _ := range shape {
		for j := 0; j < len(shape[i])-2; j += 2 {
			ImageLine(img,
				int(shape[i][j]), int(shape[i][j+1]),
				int(shape[i][j+2]), int(shape[i][j+3]),
				ColorGreen,
			)
		}
	}

	return img
}
