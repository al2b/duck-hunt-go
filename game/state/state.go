package state

type State int

const (
	StateIntro State = iota
	StateGame
)

var (
	Round = 1
	Score = 1337
)
