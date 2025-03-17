package duck

import (
	"duck-hunt-go/engine"
	"embed"
	"time"
)

//go:embed assets/*.apng
var assets embed.FS

var (
	animationHorizontalLeft  = engine.Must(engine.LoadAnimation(assets, "assets/duck.horizontal.left.apng"))
	animationHorizontalRight = engine.Must(engine.LoadAnimation(assets, "assets/duck.horizontal.right.apng"))
	animationAngledLeft      = engine.Must(engine.LoadAnimation(assets, "assets/duck.angled.left.apng"))
	animationAngledRight     = engine.Must(engine.LoadAnimation(assets, "assets/duck.angled.right.apng"))
)

func NewAnimation(velociter engine.Velociter) Animation {
	return Animation{
		velociter:       velociter,
		horizontalLeft:  animationHorizontalLeft,
		horizontalRight: animationHorizontalRight,
		angledLeft:      animationAngledLeft,
		angledRight:     animationAngledRight,
	}
}

type Animation struct {
	velociter       engine.Velociter
	horizontalLeft  engine.AnimationInterface
	horizontalRight engine.AnimationInterface
	angledLeft      engine.AnimationInterface
	angledRight     engine.AnimationInterface
}

func (animation Animation) animation() engine.AnimationInterface {
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

func (animation Animation) Duration() time.Duration {
	return animation.animation().Duration()
}

func (animation Animation) At(at time.Duration) *engine.Image {
	return animation.animation().At(at)
}
