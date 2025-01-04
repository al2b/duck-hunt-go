package engine

type Positioner interface {
	Position() Vector
}

type AbsolutePosition Vector

func (position AbsolutePosition) Position() Vector {
	return Vector(position)
}

func (position *AbsolutePosition) SetPosition(vector Vector) {
	*position = AbsolutePosition(vector)
}
