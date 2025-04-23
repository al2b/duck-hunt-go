package duck

import (
	"duck-hunt-go/engine"
	"embed"
	"time"
)

const frameDuration = time.Second / 60

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageHorizontal1 = engine.Must(engine.LoadImage(assets, "assets/duck.horizontal.1.png"))
	imageHorizontal2 = engine.Must(engine.LoadImage(assets, "assets/duck.horizontal.2.png"))
	imageHorizontal3 = engine.Must(engine.LoadImage(assets, "assets/duck.horizontal.3.png"))
	imageAngled1     = engine.Must(engine.LoadImage(assets, "assets/duck.angled.1.png"))
	imageAngled2     = engine.Must(engine.LoadImage(assets, "assets/duck.angled.2.png"))
	imageAngled3     = engine.Must(engine.LoadImage(assets, "assets/duck.angled.3.png"))

	// Animations
	animationHorizontalRight = engine.Animation{
		{imageHorizontal1, 3 * frameDuration},
		{imageHorizontal2, 3 * frameDuration},
		{imageHorizontal3, 3 * frameDuration},
	}
	animationHorizontalLeft = engine.Animation{
		{imageHorizontal1.FlipHorizontal(), 3 * frameDuration},
		{imageHorizontal2.FlipHorizontal(), 3 * frameDuration},
		{imageHorizontal3.FlipHorizontal(), 3 * frameDuration},
	}
	animationAngledRight = engine.Animation{
		{imageAngled1, 3 * frameDuration},
		{imageAngled2, 3 * frameDuration},
		{imageAngled3, 3 * frameDuration},
	}
	animationAngledLeft = engine.Animation{
		{imageAngled1.FlipHorizontal(), 3 * frameDuration},
		{imageAngled2.FlipHorizontal(), 3 * frameDuration},
		{imageAngled3.FlipHorizontal(), 3 * frameDuration},
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
