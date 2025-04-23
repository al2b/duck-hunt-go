package engine

import (
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

func (vec Vector2D) FromAngle(angle float64) Vector2D {
	r := Radians(angle)
	vec.X = math.Cos(r)
	vec.Y = math.Sin(r)
	return vec
}

// Rotate rotates the vector by the given angle
func (vec Vector2D) Rotate(angle float64) Vector2D {
	a := Radians(angle)
	x := vec.X
	y := vec.Y
	vec.X = x*math.Cos(a) - y*math.Sin(a)
	vec.Y = x*math.Sin(a) + y*math.Cos(a)
	return vec
}

// magnitude returns the length of the Vector2D
func (vec Vector2D) magnitude() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

// unit normalizes the Vector2D (set to be of unit length)
func (vec Vector2D) unit() Vector2D {
	l := vec.magnitude()
	if l < 1e-8 || l == 1 {
		// If it's 0, then don't modify the vector
		return vec
	}
	vec.X, vec.Y = vec.X/l, vec.Y/l
	return vec
}

// Add adds the Vector2D with the given Vector2D
func (vec Vector2D) Add(other Vector2D) Vector2D {
	vec.X += other.X
	vec.Y += other.Y
	return vec
}

// Sub subtracts the Vector2D with the given Vector2D
func (vec Vector2D) Sub(other Vector2D) Vector2D {
	vec.X -= other.X
	vec.Y -= other.Y
	return vec
}

// Scale scales a Vector2D by the given scalar
func (vec Vector2D) Scale(scalar float64) Vector2D {
	vec.X *= scalar
	vec.Y *= scalar
	return vec
}

// Divide divides a Vector2D by the given scalar
func (vec Vector2D) Divide(scalar float64) Vector2D {
	vec.X /= scalar
	vec.Y /= scalar
	return vec
}

// dot returns the dot product of the Vector2D and the given Vector2D
func (vec Vector2D) dot(other Vector2D) float64 {
	return vec.X*other.X + vec.Y*other.Y
}

// Reflect reflects the Vector2D against the given normal
func (vec Vector2D) Reflect(normal Vector2D) Vector2D {
	n := normal.unit()
	return vec.Sub(n.Scale(2 * n.dot(vec)))
}

// Angle returns the angle
func (vec Vector2D) Angle() float64 {
	a := math.Atan2(vec.Y, vec.X)
	if a < 0 {
		a += 2 * math.Pi
	}
	return Degrees(a)
}

func (vec Vector2D) Position() Vector2D {
	return vec
}

// Vec2D is shorthand for [Vector2D]{x, y}.
func Vec2D(x, y float64) Vector2D {
	return Vector2D{X: x, Y: y}
}

type Vectors2D []Vector2D
