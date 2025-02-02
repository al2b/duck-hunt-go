package engine

type Path struct {
	Coordinates
	vector Vector
	frames int
	frame  int
}

func (p Path) To(x, y float64, frames int) Path {
	p.vector = Vector{
		x - p.X(),
		y - p.Y(),
	}.Divide(float64(frames))

	p.frames = frames
	p.frame = 0

	return p
}

func (p Path) Update() Path {
	if p.frame >= p.frames {
		return p
	}

	p.frame++
	p.Coordinates = p.Coordinates.Move(p.vector)

	return p
}
