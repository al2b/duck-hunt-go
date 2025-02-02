package engine

func NewPath() *Path {
	return &Path{
		frames: 0,
		frame:  0,
	}
}

type Path struct {
	Coordinates
	vector Vector
	frames int
	frame  int
}

func (p *Path) To(x, y float64, frames int) {
	p.vector = Vector{
		x - p.X(),
		y - p.Y(),
	}.Divide(float64(frames))

	p.frames = frames
	p.frame = 0
}

func (p *Path) Update() {
	if p.frame >= p.frames {
		return
	}

	p.frame++

	p.Coordinates = p.Coordinates.Move(p.vector)
}
