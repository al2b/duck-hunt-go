package engine

type Projector interface {
	Project(Vector3D) Vector2D
}

type OrthographicProjector struct{}

func (p OrthographicProjector) Project(pos Vector3D) Vector2D {
	return Vector2D{X: pos.X, Y: pos.Y}
}
