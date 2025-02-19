package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/animation"
)

func NewAnimator(velociter engine.Velociter) Animator {
	return Animator{
		velociter:       velociter,
		horizontalLeft:  animation.MustLoad(animation.ApngFile(assets, "assets/duck.horizontal.left.apng")),
		horizontalRight: animation.MustLoad(animation.ApngFile(assets, "assets/duck.horizontal.right.apng")),
		angledLeft:      animation.MustLoad(animation.ApngFile(assets, "assets/duck.angled.left.apng")),
		angledRight:     animation.MustLoad(animation.ApngFile(assets, "assets/duck.angled.right.apng")),
	}
}

type Animator struct {
	velociter       engine.Velociter
	horizontalLeft  *animation.Animation
	horizontalRight *animation.Animation
	angledLeft      *animation.Animation
	angledRight     *animation.Animation
}

func (animator Animator) Animation() *animation.Animation {
	angle := animator.velociter.Velocity().Angle()
	switch true {
	case 30 <= angle && angle < 90:
		return animator.angledRight
	case 90 <= angle && angle < 150:
		return animator.angledLeft
	case 150 <= angle && angle < 210:
		return animator.horizontalLeft
	case 210 <= angle && angle < 270:
		return animator.angledLeft
	case 270 <= angle && angle < 330:
		return animator.angledRight
	default:
		return animator.horizontalRight
	}
}
