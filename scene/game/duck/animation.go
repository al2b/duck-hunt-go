package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/animation"
	"time"
)

func NewAnimation(velociter engine.Velociter) Animation {
	return Animation{
		velociter:       velociter,
		horizontalLeft:  animation.MustLoad(animation.ApngFile(assets, "assets/duck.horizontal.left.apng")),
		horizontalRight: animation.MustLoad(animation.ApngFile(assets, "assets/duck.horizontal.right.apng")),
		angledLeft:      animation.MustLoad(animation.ApngFile(assets, "assets/duck.angled.left.apng")),
		angledRight:     animation.MustLoad(animation.ApngFile(assets, "assets/duck.angled.right.apng")),
	}
}

type Animation struct {
	velociter       engine.Velociter
	horizontalLeft  *animation.Animation
	horizontalRight *animation.Animation
	angledLeft      *animation.Animation
	angledRight     *animation.Animation
}

func (animation Animation) Animation() *animation.Animation {
	angle := animation.velociter.Velocity().Angle()
	switch true {
	case 30 <= angle && angle < 90:
		return animation.angledRight
	case 90 <= angle && angle < 150:
		return animation.angledLeft
	case 150 <= angle && angle < 210:
		return animation.horizontalLeft
	case 210 <= angle && angle < 270:
		return animation.angledLeft
	case 270 <= angle && angle < 330:
		return animation.angledRight
	default:
		return animation.horizontalRight
	}
}

func (animation Animation) Step(delta time.Duration) {
	animation.Animation().Step(delta)
}

func (animation Animation) Image() *engine.Image {
	return animation.Animation().Image()
}
