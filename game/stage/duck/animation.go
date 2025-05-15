package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	"time"
)

var (
	// Animations
	animationHorizontalRight = engine.Animation{
		{assets.DuckHorizontal1, 3 * config.TickInterval},
		{assets.DuckHorizontal2, 3 * config.TickInterval},
		{assets.DuckHorizontal3, 3 * config.TickInterval},
	}
	animationHorizontalLeft = engine.Animation{
		{assets.DuckHorizontal1.FlipHorizontal(), 3 * config.TickInterval},
		{assets.DuckHorizontal2.FlipHorizontal(), 3 * config.TickInterval},
		{assets.DuckHorizontal3.FlipHorizontal(), 3 * config.TickInterval},
	}
	animationAngledRight = engine.Animation{
		{assets.DuckAngled1, 3 * config.TickInterval},
		{assets.DuckAngled2, 3 * config.TickInterval},
		{assets.DuckAngled3, 3 * config.TickInterval},
	}
	animationAngledLeft = engine.Animation{
		{assets.DuckAngled1.FlipHorizontal(), 3 * config.TickInterval},
		{assets.DuckAngled2.FlipHorizontal(), 3 * config.TickInterval},
		{assets.DuckAngled3.FlipHorizontal(), 3 * config.TickInterval},
	}
)

func NewAnimation(velociter engine.Velociter) Animation {
	return Animation{
		velociter: velociter,
	}
}

type Animation struct {
	velociter engine.Velociter
}

func (animation Animation) animation() engine.AnimationInterface {
	angle := animation.velociter.Velocity().Angle()
	switch true {
	case 30 <= angle && angle < 90:
		return animationAngledRight
	case 90 <= angle && angle < 150:
		return animationAngledLeft
	case 150 <= angle && angle < 210:
		return animationHorizontalLeft
	case 210 <= angle && angle < 270:
		return animationAngledLeft
	case 270 <= angle && angle < 330:
		return animationAngledRight
	default:
		return animationHorizontalRight
	}
}

func (animation Animation) Duration() time.Duration {
	return animation.animation().Duration()
}

func (animation Animation) At(at time.Duration) *engine.Image {
	return animation.animation().At(at)
}
