package engine

func NewCoordinates(x, y, z float64) *Coordinates {
	return &Coordinates{
		ox: x, x: x,
		oy: y, y: y,
		oz: z, z: z,
	}
}

type Coordinates struct {
	ox, x float64
	oy, y float64
	oz, z float64
}

func (c *Coordinates) X() float64 {
	return c.x
}

func (c *Coordinates) SetX(x float64) *Coordinates {
	c.x = x
	return c
}

func (c *Coordinates) Y() float64 {
	return c.y
}

func (c *Coordinates) SetY(y float64) *Coordinates {
	c.y = y
	return c
}

func (c *Coordinates) Z() float64 {
	return c.z
}

func (c *Coordinates) SetZ(z float64) *Coordinates {
	c.z = z
	return c
}

func (c *Coordinates) AddZ(z float64) *Coordinates {
	c.z += z
	return c
}

func (c *Coordinates) SubZ(z float64) *Coordinates {
	c.z -= z
	return c
}

func (c *Coordinates) Move(movement Vector) *Coordinates {
	c.x += movement.X
	c.y += movement.Y
	return c
}

func (c *Coordinates) Reset() *Coordinates {
	c.x, c.y, c.z = c.ox, c.oy, c.oz
	return c
}

func NewRelativeCoordinates(coordinates Point, x, y, z float64) *RelativeCoordinates {
	return &RelativeCoordinates{
		coordinates: coordinates,
		x:           x,
		y:           y,
		z:           z,
	}
}

type RelativeCoordinates struct {
	coordinates Point
	x, y, z     float64
}

func (c *RelativeCoordinates) X() float64 {
	return c.coordinates.X() + c.x
}

func (c *RelativeCoordinates) Y() float64 {
	return c.coordinates.Y() + c.y
}

func (c *RelativeCoordinates) Z() float64 {
	return c.coordinates.Z() + c.z
}
