package engine

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (vec Vector3D) Position() Vector3D {
	return vec
}
