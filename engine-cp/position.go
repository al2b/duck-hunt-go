package cp

import (
	"duck-hunt-go/engine"
	"github.com/jakecoffman/cp/v2"
)

type Positioner interface {
	Position() cp.Vector
}

type PositionPointer struct {
	Positioner
}

func (pp PositionPointer) Point() engine.Point {
	position := pp.Positioner.Position()
	return engine.Point{X: int(position.X), Y: int(position.Y)}
}
