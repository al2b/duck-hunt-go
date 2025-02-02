package layout

import (
	"duck-hunt-go/engine"
)

var Shrub = Element{
	Coordinates: engine.NewCoordinates(193, 122, 20),
	StaticImage: engine.NewStaticImage(imageShrub),
	PolygonShape: engine.NewPolygonShape(
		0, 60,
		0, 29,
		1, 22,
		3, 16,
		7, 11,
		8, 7,
		9, 4,
		16, 0,
		23, 2,
		25, 4,
		29, 15,
		30, 25,
		30, 60,
	),
}
