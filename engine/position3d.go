package engine

type Positioner3D interface {
	Position() Vector3D
}

type Position3DProjector struct {
	Positioner3D
	Projector
}

func (pp Position3DProjector) Position() Vector2D {
	return pp.Projector.Project(pp.Positioner3D.Position())
}

type Position3DOrderer struct {
	Positioner3D
}

func (po Position3DOrderer) Order() Order {
	return Order(po.Positioner3D.Position().Z)
}
