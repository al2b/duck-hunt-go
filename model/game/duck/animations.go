package duck

import (
	"duck-hunt-go/engine"
)

const (
	animationWidth  = 32
	animationHeight = 32
)

const (
	animationFlyTop int = iota
	animationFlyTopRight
	animationFlyRight
	animationFlyBottomRight
	animationFlyBottom
	animationFlyBottomLeft
	animationFlyLeft
	animationFlyTopLeft
)

var animations = map[int]*engine.Animation{
	animationFlyTop: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 0 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 0 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyTopRight: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 1 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 1 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyRight: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 2 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 2 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyBottomRight: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 3 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 3 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyBottom: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 0 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
			{X: 0 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
		},
	),
	animationFlyBottomLeft: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 3 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 3 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
		},
	),
	animationFlyLeft: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 2 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 2 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
		},
	),
	animationFlyTopLeft: engine.NewAnimation(
		image,
		engine.AnimationFrames{
			{X: 1 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 1 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
		},
	),
}
