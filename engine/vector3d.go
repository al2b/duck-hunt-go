package engine

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (vec Vector3D) Position() Vector3D {
	return vec
}

// Vec3D is shorthand for [Vector3D]{x, y, z}.
func Vec3D(x, y, z float64) Vector3D {
	return Vector3D{X: x, Y: y, Z: z}
}
