package cp

import (
	"duck-hunt-go/engine"
	"github.com/jakecoffman/cp/v2"
)

func BodyPositioner2DFunc(positioner engine.Positioner2D) cp.BodyPositionFunc {
	return func(body *cp.Body, _ float64) {
		position := positioner.Position()
		body.SetPosition(
			cp.Vector{X: position.X, Y: position.Y},
		)
	}
}
