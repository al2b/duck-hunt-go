package engine

func NewPosition() *Position {
	return &Position{}
}

type Position struct {
	X, Y, Z float64
}

func (p *Position) Move(movement Vector) {
	p.X += movement.X
	p.Y -= movement.Y
}

func (p *Position) DepthUp(z float64) {
	p.Z -= z
}

func (p *Position) DepthDown(z float64) {
	p.Z += z
}
