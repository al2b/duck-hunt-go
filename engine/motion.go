package engine

type Motion struct {
	Coordinates
	vector Vector
}

func (m Motion) Angle() float64 {
	return m.vector.Angle()
}

func (m Motion) SetAngle(angle float64) Motion {
	m.vector = m.vector.SetAngle(angle)
	return m
}

func (m Motion) Rotate(angle float64) Motion {
	m.vector = m.vector.Rotate(angle)
	return m
}

func (m Motion) Scale(scalar float64) Motion {
	m.vector = m.vector.Scale(scalar)
	return m
}

func (m Motion) Update() Motion {
	m.Coordinates = m.Coordinates.Move(m.vector)
	return m
}

func (m Motion) Reflect(normal Vector) Motion {
	m.vector = m.vector.Reflect(normal)
	return m
}
