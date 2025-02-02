package layout

import "duck-hunt-go/engine"

var Tree = Element{
	Coordinates: engine.NewCoordinates(6, 32, 10),
	StaticImage: engine.NewStaticImage(imageTree),
	PolygonShape: engine.NewPolygonShape(
		0, 0,
		68, 0,
		68, 150,
		0, 150,
	),
}
