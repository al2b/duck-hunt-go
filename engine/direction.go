package engine

import "math"

type Direction int

const (
	DirectionRight Direction = iota
	DirectionTopRight
	DirectionTop
	DirectionTopLeft
	DirectionLeft
	DirectionBottomLeft
	DirectionBottom
	DirectionBottomRight
)

func (o Direction) String() string {
	return [...]string{"right", "top right", "top", "top left", "left", "bottom left", "bottom", "bottom right"}[o]
}

type Directioner interface {
	Direction() Direction
}

type VerticalSemicircleDirectioner struct {
	Velociter
}

func (d VerticalSemicircleDirectioner) Direction() Direction {
	velocity := d.Velocity()
	if velocity.X >= 0 {
		return DirectionRight
	}
	return DirectionLeft
}

type HorizontalSemicircleDirectioner struct {
	Velociter
}

func (d HorizontalSemicircleDirectioner) Direction() Direction {
	velocity := d.Velocity()
	if velocity.Y >= 0 {
		return DirectionTop
	}
	return DirectionBottom
}

type QuadrantDirectioner struct {
	Velociter
}

func (d QuadrantDirectioner) Direction() Direction {
	velocity := d.Velocity()
	if Abs(velocity.X) > Abs(velocity.Y) {
		if velocity.X > 0 {
			return DirectionRight
		}
		return DirectionLeft
	}
	if velocity.Y > 0 {
		return DirectionBottom
	}
	return DirectionTop
}

type VerticalSextantDirectioner struct {
	Velociter
}

func (d VerticalSextantDirectioner) Direction() Direction {
	velocity := d.Velocity()

	angle := math.Atan2(velocity.Y, velocity.X)
	if angle < 0 {
		angle += 2 * math.Pi
	}

	const piOver3 = math.Pi / 3
	const piOver6 = math.Pi / 6

	index := int((angle+piOver6)/piOver3) % 6

	return [...]Direction{
		DirectionRight, DirectionTopRight, DirectionTopLeft,
		DirectionLeft, DirectionBottomLeft, DirectionBottomRight,
	}[index]
}

type OctantDirectioner struct {
	Velociter
}

func (d OctantDirectioner) Direction() Direction {
	velocity := d.Velocity()

	angle := math.Atan2(velocity.Y, velocity.X)
	if angle < 0 {
		angle += 2 * math.Pi
	}

	const piOver4 = math.Pi / 4
	const piOver8 = math.Pi / 8

	index := int((angle+piOver8)/piOver4) & 7

	return [...]Direction{
		DirectionRight, DirectionTopRight, DirectionTop, DirectionTopLeft,
		DirectionLeft, DirectionBottomLeft, DirectionBottom, DirectionBottomRight,
	}[index]
}
