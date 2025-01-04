package engine

import (
	"math"
)

func NewPosition() *Position {
	return &Position{}
}

type Position struct {
	X, Y, Z float64
}

func (p *Position) Move(direction *Direction) {
	rad := float64(direction.Angle) * math.Pi / 180
	p.X += math.Sin(rad) * direction.Velocity
	p.Y -= math.Cos(rad) * direction.Velocity
}

func (p *Position) DepthUp(z float64) {
	p.Z -= z
}

func (p *Position) DepthDown(z float64) {
	p.Z += z
}
