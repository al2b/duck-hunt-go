package engine

func NewMotion() *Motion {
	return &Motion{
		frames: 0,
		frame:  0,
	}
}

type Motion struct {
	vector Vector
	frames int
	frame  int
}

func (m *Motion) MoveTo(coordinates Coordinates, x, y float64, frames int) {
	m.vector = Vector{
		x - coordinates.X(),
		y - coordinates.Y(),
	}.Divide(float64(frames))

	m.frames = frames
	m.frame = 0
}

func (m *Motion) Update(coordinates Coordinates) Coordinates {
	if m.frame >= m.frames {
		return coordinates
	}

	m.frame++

	return coordinates.Move(m.vector)
}
