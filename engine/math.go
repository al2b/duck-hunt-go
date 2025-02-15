package engine

import (
	"math"
)

type Positionable interface {
	X() float64
	Y() float64
	Z() float64
}

func Round(value float64) int {
	iv := int(value)
	if value-float64(iv) < 0.5 {
		return iv
	}
	return iv + 1
}

// Degrees convert radians to degrees
func Degrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

// Radians convert degrees to radians
func Radians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func Abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
