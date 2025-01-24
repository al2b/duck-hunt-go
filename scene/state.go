package scene

import "duck-hunt-go/engine"

type State int
type StateModels map[State]engine.Model

const (
	StateIntro State = iota
	StateGame
)
