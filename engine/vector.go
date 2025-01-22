package engine

import "math"

func VectorFromAngle(angle float64) Vector {
	r := Radians(angle)
	return Vector{
		X: math.Cos(r),
		Y: math.Sin(r),
	}
}

type Vector struct {
	X float64
	Y float64
}

// Rotate rotates the vector by the given angle
func (v Vector) Rotate(angle float64) Vector {
	a := Radians(angle)
	x := v.X
	y := v.Y
	v.X = x*math.Cos(a) - y*math.Sin(a)
	v.Y = x*math.Sin(a) + y*math.Cos(a)
	return v
}

// magnitude returns the length of the Vector
func (v Vector) magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// unit normalizes the Vector (set to be of unit length)
func (v Vector) unit() Vector {
	l := v.magnitude()
	if l < 1e-8 || l == 1 {
		// If it's 0, then don't modify the vector
		return v
	}
	v.X, v.Y = v.X/l, v.Y/l
	return v
}

// sub subtracts the Vector with the given Vector
func (v Vector) sub(other Vector) Vector {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

// Scale scales a Vector by the given scalar
func (v Vector) Scale(scalar float64) Vector {
	v.X *= scalar
	v.Y *= scalar
	return v
}

// Divide divides a Vector by the given scalar
func (v Vector) Divide(scalar float64) Vector {
	v.X /= scalar
	v.Y /= scalar
	return v
}

// dot returns the dot product of the Vector and the given Vector
func (v Vector) dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Reflect reflects the Vector against the given normal
func (v Vector) Reflect(normal Vector) Vector {
	n := normal.unit()
	return v.sub(n.Scale(2 * n.dot(v)))
}

// Angle returns the angle
func (v Vector) Angle() float64 {
	a := math.Atan2(v.Y, v.X)
	if a < 0 {
		a += 2 * math.Pi
	}
	return Degrees(a)
}
