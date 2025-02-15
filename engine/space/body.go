package space

import (
	"duck-hunt-go/engine"
	"github.com/jakecoffman/cp/v2"
)

type Body struct {
	cpBody  *cp.Body
	cpSpace *cp.Space
}

func (b Body) Position() engine.Position {
	position := b.cpBody.Position()

	return engine.Position{
		X: position.X,
		Y: position.Y,
	}
}

func (b Body) Velocity() engine.Vector {
	velocity := b.cpBody.Velocity()

	return engine.Vector{
		X: velocity.X,
		Y: velocity.Y,
	}
}

func (b Body) SetPosition(position engine.Position) Body {
	b.cpBody.SetPosition(cp.Vector{X: position.X, Y: position.Y})
	return b
}

func (b Body) SetVelocity(vector engine.Vector) Body {
	b.cpBody.SetVelocityVector(cp.Vector{X: vector.X, Y: vector.Y})
	return b
}

func (b Body) AddNewPolygon(vertices engine.Positions, radius float64) Shape {
	cpVertices := make([]cp.Vector, len(vertices))
	for i, vertex := range vertices {
		cpVertices[i] = cp.Vector{
			X: vertex.X,
			Y: vertex.Y,
		}
	}
	cpShape := b.cpSpace.AddShape(
		cp.NewPolyShapeRaw(b.cpBody, len(cpVertices), cpVertices, radius),
	)

	switch b.cpBody.GetType() {
	case cp.BODY_DYNAMIC:
		b.cpBody.SetMoment(
			cp.MomentForPoly(b.cpBody.Mass(), len(cpVertices), cpVertices, cp.Vector{}, 0),
		)
	}

	return Shape{
		cpShape: cpShape,
	}
}

func (b Body) AddNewCircle(radius float64) Shape {
	cpShape := b.cpSpace.AddShape(
		cp.NewCircle(b.cpBody, radius, cp.Vector{}),
	)

	switch b.cpBody.GetType() {
	case cp.BODY_DYNAMIC:
		b.cpBody.SetMoment(
			cp.MomentForCircle(b.cpBody.Mass(), 0, radius, cp.Vector{}),
		)
	}

	return Shape{
		cpShape: cpShape,
	}
}
