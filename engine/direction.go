package engine

func NewDirection() *Direction {
	return &Direction{}
}

type Direction struct {
	Angle    int
	Velocity float64
}

func (d *Direction) RotateUp(angle int) {
	d.Angle = (d.Angle + angle) % 360
}

func (d *Direction) RotateDown(angle int) {
	d.Angle = ((360 + d.Angle) - angle) % 360
}
