package gun

import "duck-hunt-go/engine"

var body = engine.NewBody(coordinates,
	engine.BodyShape{
		{13, 0},
		{23, 0},
		{36, 13},
		{36, 23},
		{23, 36},
		{13, 36},
		{0, 23},
		{0, 13},
	},
)
