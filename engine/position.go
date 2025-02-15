package engine

type Position struct {
	X, Y float64
}

func (p Position) Add(q Position) Position {
	return Position{p.X + q.X, p.Y + q.Y}
}

type Positions []Position

type Positioner interface {
	Position() Position
}

func NewAbsolutePosition(x, y float64) *AbsolutePosition {
	return &AbsolutePosition{
		position: Position{X: x, Y: y},
	}
}

type AbsolutePosition struct {
	position Position
}

func (position *AbsolutePosition) Position() Position {
	return position.position
}

func (position *AbsolutePosition) Move(x, y float64) {
	position.position.X, position.position.Y = x, y
}
