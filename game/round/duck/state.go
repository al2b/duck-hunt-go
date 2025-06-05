package duck

type state int

const (
	stateIdle state = iota
	stateFly
	stateShot
	stateFall
	stateDown
)
