package duck

import (
	"duck-hunt-go/engine"
)

const (
	width  = 32
	height = 32
)

var (
	imageDuck = engine.MustLoadImage(assets, "assets/duck.png")
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

func (a *Animation) Step(angle float64) {
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
			{X: 0 * width, Y: 0 * width, Width: width, Height: height},
			{X: 0 * width, Y: 1 * width, Width: width, Height: height},
			{X: 0 * width, Y: 2 * width, Width: width, Height: height},
			{X: 0 * width, Y: 1 * width, Width: width, Height: height},
		},
	),
	animationFlyTopRight: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 1 * width, Y: 0 * width, Width: width, Height: height},
			{X: 1 * width, Y: 1 * width, Width: width, Height: height},
			{X: 1 * width, Y: 2 * width, Width: width, Height: height},
			{X: 1 * width, Y: 1 * width, Width: width, Height: height},
		},
	),
	animationFlyRight: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 2 * width, Y: 0 * width, Width: width, Height: height},
			{X: 2 * width, Y: 1 * width, Width: width, Height: height},
			{X: 2 * width, Y: 2 * width, Width: width, Height: height},
			{X: 2 * width, Y: 1 * width, Width: width, Height: height},
		},
	),
	animationFlyBottomRight: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 3 * width, Y: 0 * width, Width: width, Height: height},
			{X: 3 * width, Y: 1 * width, Width: width, Height: height},
			{X: 3 * width, Y: 2 * width, Width: width, Height: height},
			{X: 3 * width, Y: 1 * width, Width: width, Height: height},
		},
	),
	animationFlyBottom: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 0 * width, Y: 0 * width, Width: width, Height: height, FlipV: true},
			{X: 0 * width, Y: 1 * width, Width: width, Height: height, FlipV: true},
			{X: 0 * width, Y: 2 * width, Width: width, Height: height, FlipV: true},
			{X: 0 * width, Y: 1 * width, Width: width, Height: height, FlipV: true},
		},
	),
	animationFlyBottomLeft: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 3 * width, Y: 0 * width, Width: width, Height: height, FlipH: true},
			{X: 3 * width, Y: 1 * width, Width: width, Height: height, FlipH: true},
			{X: 3 * width, Y: 2 * width, Width: width, Height: height, FlipH: true},
			{X: 3 * width, Y: 1 * width, Width: width, Height: height, FlipH: true},
		},
	),
	animationFlyLeft: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 2 * width, Y: 0 * width, Width: width, Height: height, FlipH: true},
			{X: 2 * width, Y: 1 * width, Width: width, Height: height, FlipH: true},
			{X: 2 * width, Y: 2 * width, Width: width, Height: height, FlipH: true},
			{X: 2 * width, Y: 1 * width, Width: width, Height: height, FlipH: true},
		},
	),
	animationFlyTopLeft: engine.NewAnimation(
		imageDuck,
		engine.AnimationFrames{
			{X: 1 * width, Y: 0 * width, Width: width, Height: height, FlipH: true},
			{X: 1 * width, Y: 1 * width, Width: width, Height: height, FlipH: true},
			{X: 1 * width, Y: 2 * width, Width: width, Height: height, FlipH: true},
			{X: 1 * width, Y: 1 * width, Width: width, Height: height, FlipH: true},
		},
	),
}
