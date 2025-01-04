package engine

func Round(value float64) int {
	iv := int(value)
	if value-float64(iv) < 0.5 {
		return iv
	}
	return iv + 1
}
