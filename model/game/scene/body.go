package scene

import "duck-hunt-go/engine"

var bodyShape = engine.BodyShape{
	{0, 0},
	{width - 1, 0},
	{width - 1, ground - 1},
	{0, ground - 1},
}
