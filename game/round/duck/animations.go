package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	"time"
)

var (
	// Animations
	animationFlyHorizontalRight = engine.Animation{
		{assets.DuckFlyHorizontalRight1, 3 * config.TickInterval},
		{assets.DuckFlyHorizontalRight2, 3 * config.TickInterval},
		{assets.DuckFlyHorizontalRight3, 3 * config.TickInterval},
	}
	animationFlyHorizontalLeft = engine.Animation{
		{assets.DuckFlyHorizontalLeft1, 3 * config.TickInterval},
		{assets.DuckFlyHorizontalLeft2, 3 * config.TickInterval},
		{assets.DuckFlyHorizontalLeft3, 3 * config.TickInterval},
	}
	animationFlyAngledRight = engine.Animation{
		{assets.DuckFlyAngledRight1, 3 * config.TickInterval},
		{assets.DuckFlyAngledRight2, 3 * config.TickInterval},
		{assets.DuckFlyAngledRight3, 3 * config.TickInterval},
	}
	animationFlyAngledLeft = engine.Animation{
		{assets.DuckFlyAngledLeft1, 3 * config.TickInterval},
		{assets.DuckFlyAngledLeft2, 3 * config.TickInterval},
		{assets.DuckFlyAngledLeft3, 3 * config.TickInterval},
	}
)

type animationFly struct {
	Velociter engine.Velociter
}

func (animation animationFly) animation() engine.AnimationInterface {
	directioner := engine.VerticalSextantDirectioner{animation.Velociter}

	switch directioner.Direction() {
	case engine.DirectionTopRight, engine.DirectionBottomRight:
		return animationFlyAngledRight
	case engine.DirectionTopLeft, engine.DirectionBottomLeft:
		return animationFlyAngledLeft
	case engine.DirectionLeft:
		return animationFlyHorizontalLeft
	default:
		return animationFlyHorizontalRight
	}
}

func (animation animationFly) Duration() time.Duration {
	return animation.animation().Duration()
}

func (animation animationFly) At(at time.Duration) *engine.Image {
	return animation.animation().At(at)
}
