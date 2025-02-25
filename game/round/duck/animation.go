package duck

import (
	"duck-hunt-go/engine"
	"time"
)

func NewAnimationPlayer(velociter engine.Velociter) AnimationPlayer {
	return AnimationPlayer{
		velociter: velociter,
		horizontalLeft: engine.NewAnimationPlayer(
			engine.Must(engine.LoadAnimation(assets, "assets/duck.horizontal.left.apng")),
		),
		horizontalRight: engine.NewAnimationPlayer(
			engine.Must(engine.LoadAnimation(assets, "assets/duck.horizontal.right.apng")),
		),
		angledLeft: engine.NewAnimationPlayer(
			engine.Must(engine.LoadAnimation(assets, "assets/duck.angled.left.apng")),
		),
		angledRight: engine.NewAnimationPlayer(
			engine.Must(engine.LoadAnimation(assets, "assets/duck.angled.right.apng")),
		),
	}
}

type AnimationPlayer struct {
	velociter       engine.Velociter
	horizontalLeft  *engine.AnimationPlayer
	horizontalRight *engine.AnimationPlayer
	angledLeft      *engine.AnimationPlayer
	angledRight     *engine.AnimationPlayer
}

func (animation AnimationPlayer) player() *engine.AnimationPlayer {
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

func (animation AnimationPlayer) Step(delta time.Duration) {
	animation.player().Step(delta)
}

func (animation AnimationPlayer) Image() *engine.Image {
	return animation.player().Image()
}
