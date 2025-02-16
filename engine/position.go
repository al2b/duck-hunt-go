package engine

type Positioner interface {
	Position() Vector
}

func NewAbsolutePosition(x, y float64) *AbsolutePosition {
	return &AbsolutePosition{
		position: Vec(x, y),
	}
}

type AbsolutePosition struct {
	position Vector
}

func (position *AbsolutePosition) Position() Vector {
	return position.position
}

func (position *AbsolutePosition) Move(x, y float64) {
	position.position.X, position.position.Y = x, y
}
