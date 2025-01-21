package layout

import "duck-hunt-go/engine"

var bodies = engine.Bodies{
	engine.NewBody(coordinates,
		engine.BodyShape{
			{0, 0},
			{width - 1, 0},
			{width - 1, Ground - 1},
			{0, Ground - 1},
		},
	),
	// Tree
	engine.NewBody(treeCoordinates,
		engine.BodyShape{
			{0, 0},
			{68, 0},
			{68, 150},
			{0, 150},
		},
	),
	// Shrub
	engine.NewBody(shrubCoordinates,
		engine.BodyShape{
			{0, 60},
			{0, 29},
			{1, 22},
			{3, 16},
			{7, 11},
			{8, 7},
			{9, 4},
			{16, 0},
			{23, 2},
			{25, 4},
			{29, 15},
			{30, 25},
			{30, 60},
		},
	),
}
