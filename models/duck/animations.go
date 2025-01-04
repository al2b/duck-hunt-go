package duck

import "duck-hunt-go/engine"

const (
	animationFlyTop engine.Animation = iota
	animationFlyTopRight
	animationFlyRight
	animationFlyBottomRight
	animationFlyBottom
	animationFlyBottomLeft
	animationFlyLeft
	animationFlyTopLeft
)

var animationFrames = engine.AnimationFrames{
	animationFlyTop: {
		{X: 0, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: 2},
		{X: 0, Y: 1},
	},
	animationFlyTopRight: {
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 1},
	},
	animationFlyRight: {
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 2, Y: 1},
	},
	animationFlyBottomRight: {
		{X: 3, Y: 0},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 3, Y: 1},
	},
	animationFlyBottom: {
		{X: 0, Y: 0, FlipV: true},
		{X: 0, Y: 1, FlipV: true},
		{X: 0, Y: 2, FlipV: true},
		{X: 0, Y: 1, FlipV: true},
	},
	animationFlyBottomLeft: {
		{X: 3, Y: 0, FlipH: true},
		{X: 3, Y: 1, FlipH: true},
		{X: 3, Y: 2, FlipH: true},
		{X: 3, Y: 1, FlipH: true},
	},
	animationFlyLeft: {
		{X: 2, Y: 0, FlipH: true},
		{X: 2, Y: 1, FlipH: true},
		{X: 2, Y: 2, FlipH: true},
		{X: 2, Y: 1, FlipH: true},
	},
	animationFlyTopLeft: {
		{X: 1, Y: 0, FlipH: true},
		{X: 1, Y: 1, FlipH: true},
		{X: 1, Y: 2, FlipH: true},
		{X: 1, Y: 1, FlipH: true},
	},
}
