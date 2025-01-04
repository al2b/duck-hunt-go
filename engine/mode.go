package engine

type Mode int

const (
	Mode8 Mode = iota
	Mode24
)

var mode = Mode24
