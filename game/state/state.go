package state

type State int

const (
	Title State = iota
	Play
)

var (
	Round = 1
	Score = 1337
)
