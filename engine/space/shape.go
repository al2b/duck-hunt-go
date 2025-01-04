package space

import (
	"github.com/jakecoffman/cp/v2"
)

type Shape struct {
	cpShape *cp.Shape
}

func (s Shape) SetElasticity(elasticity float64) Shape {
	s.cpShape.SetElasticity(elasticity)
	return s
}

func (s Shape) SetFriction(friction float64) Shape {
	s.cpShape.SetFriction(friction)
	return s
}
