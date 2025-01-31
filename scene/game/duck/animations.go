package duck

import (
	"duck-hunt-go/engine"
)

const (
	animationWidth  = 32
	animationHeight = 32
)

type animationType int

const (
	animationFlyTop animationType = iota
	animationFlyTopRight
	animationFlyRight
	animationFlyBottomRight
	animationFlyBottom
	animationFlyBottomLeft
	animationFlyLeft
	animationFlyTopLeft
)

type Animation struct {
	*engine.Animation
}

func (a *Animation) Update(angle float64) {
	var t animationType

	switch true {
	case 23 <= angle && angle <= 67:
		t = animationFlyBottomRight
	case 68 <= angle && angle <= 112:
		t = animationFlyBottom
	case 113 <= angle && angle <= 157:
		t = animationFlyBottomLeft
	case 158 <= angle && angle <= 202:
		t = animationFlyLeft
	case 203 <= angle && angle <= 247:
		t = animationFlyTopLeft
	case 248 <= angle && angle <= 292:
		t = animationFlyTop
	case 293 <= angle && angle <= 337:
		t = animationFlyTopRight
	default:
		t = animationFlyRight
	}

	if a.Animation != animations[t] {
		a.Animation = animations[t]
	} else {
		a.Animation.Update()
	}
}

var animations = map[animationType]*engine.Animation{
	animationFlyTop: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 0 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 0 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyTopRight: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 1 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 1 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyRight: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 2 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 2 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyBottomRight: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 3 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 3 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight},
		},
	),
	animationFlyBottom: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 0 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
			{X: 0 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
			{X: 0 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipV: true},
		},
	),
	animationFlyBottomLeft: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 3 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 3 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 3 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
		},
	),
	animationFlyLeft: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 2 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 2 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 2 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
		},
	),
	animationFlyTopLeft: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 1 * animationWidth, Y: 0 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 1 * animationWidth, Y: 2 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
			{X: 1 * animationWidth, Y: 1 * animationWidth, Width: animationWidth, Height: animationHeight, FlipH: true},
		},
	),
}
