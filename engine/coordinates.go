package engine

func NewCoordinates(x, y, z float64) Coordinates {
	return Coordinates{
		x: x,
		y: y,
		z: z,
	}
}

type Coordinates struct {
	x float64
	y float64
	z float64
}

func (c Coordinates) Set(x, y, z float64) Coordinates {
	c.x = x
	c.y = y
	c.z = z
	return c
}

func (c Coordinates) Add(x, y, z float64) Coordinates {
	c.x += x
	c.y += y
	c.z += z
	return c
}

func (c Coordinates) X() float64 {
	return c.x
}

func (c Coordinates) SetX(x float64) Coordinates {
	c.x = x
	return c
}

func (c Coordinates) Y() float64 {
	return c.y
}

func (c Coordinates) SetY(y float64) Coordinates {
	c.y = y
	return c
}

func (c Coordinates) Z() float64 {
	return c.z
}

func (c Coordinates) SetZ(z float64) Coordinates {
	c.z = z
	return c
}

func (c Coordinates) AddZ(z float64) Coordinates {
	c.z += z
	return c
}

func (c Coordinates) SubZ(z float64) Coordinates {
	c.z -= z
	return c
}

func (c Coordinates) Move(vector Vector) Coordinates {
	c.x += vector.X
	c.y += vector.Y
	return c
}
