package duck

import "duck-hunt-go/engine"

var body = engine.NewBody(coordinates,
	engine.BodyShape{
		{0, 0},
		{31, 0},
		{31, 31},
		{0, 31},
	},
)
