package engine

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (vec Vector) FromAngle(angle float64) Vector {
	r := Radians(angle)
	vec.X = math.Cos(r)
	vec.Y = math.Sin(r)
	return vec
}

// Rotate rotates the vector by the given angle
func (vec Vector) Rotate(angle float64) Vector {
	a := Radians(angle)
	x := vec.X
	y := vec.Y
	vec.X = x*math.Cos(a) - y*math.Sin(a)
	vec.Y = x*math.Sin(a) + y*math.Cos(a)
	return vec
}

// magnitude returns the length of the Vector
func (vec Vector) magnitude() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

// unit normalizes the Vector (set to be of unit length)
func (vec Vector) unit() Vector {
	l := vec.magnitude()
	if l < 1e-8 || l == 1 {
		// If it's 0, then don't modify the vector
		return vec
	}
	vec.X, vec.Y = vec.X/l, vec.Y/l
	return vec
}

// Add adds the Vector with the given Vector
func (vec Vector) Add(other Vector) Vector {
	vec.X += other.X
	vec.Y += other.Y
	return vec
}

// Sub subtracts the Vector with the given Vector
func (vec Vector) Sub(other Vector) Vector {
	vec.X -= other.X
	vec.Y -= other.Y
	return vec
}

// Scale scales a Vector by the given scalar
func (vec Vector) Scale(scalar float64) Vector {
	vec.X *= scalar
	vec.Y *= scalar
	return vec
}

// Divide divides a Vector by the given scalar
func (vec Vector) Divide(scalar float64) Vector {
	vec.X /= scalar
	vec.Y /= scalar
	return vec
}

// dot returns the dot product of the Vector and the given Vector
func (vec Vector) dot(other Vector) float64 {
	return vec.X*other.X + vec.Y*other.Y
}

// Reflect reflects the Vector against the given normal
func (vec Vector) Reflect(normal Vector) Vector {
	n := normal.unit()
	return vec.Sub(n.Scale(2 * n.dot(vec)))
}

// Angle returns the angle
func (vec Vector) Angle() float64 {
	a := math.Atan2(vec.Y, vec.X)
	if a < 0 {
		a += 2 * math.Pi
	}
	return Degrees(a)
}

// Vec is shorthand for [Vector]{X, Y}.
func Vec(X, Y float64) Vector {
	return Vector{X, Y}
}

type Vectors []Vector
