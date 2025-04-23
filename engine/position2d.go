package engine

type Positioner2D interface {
	Position() Vector2D
}

type Position2DPointer struct {
	Positioner2D
}

func (pp Position2DPointer) Point() Point {
	pos := pp.Positioner2D.Position()
	return Point{X: int(pos.X), Y: int(pos.Y)}
}
