package cp

import (
	"duck-hunt-go/engine"
	"github.com/jakecoffman/cp/v2"
)

type Velociter interface {
	Velocity() cp.Vector
}

type VelocityVelociter struct {
	Velociter
}

func (vv VelocityVelociter) Velocity() engine.Vector2D {
	velocity := vv.Velociter.Velocity()
	return engine.Vector2D{X: velocity.X, Y: velocity.Y}
}
