package engine

import "time"

// PlayerState represents the current state of a Player
type PlayerState int

// Predefined player states
const (
	PlayerStateStopped PlayerState = iota
	PlayerStatePlaying
	PlayerStatePaused
)

type PlayerOnEnd int

// Predefined player on end behavior
const (
	PlayerOnEndStopRewind PlayerOnEnd = iota
	PlayerOnEndLoop
	PlayerOnEndPause
)

type Player struct {
	state PlayerState
	time  time.Duration
}

// Step updates the player
func (p *Player) Step(delta time.Duration, duration time.Duration, onEnd PlayerOnEnd) {
	if p.state != PlayerStatePlaying {
		return
	}

	p.time += delta

	if p.time <= duration {
		return
	}

	// Reached the end
	switch onEnd {
	case PlayerOnEndStopRewind:
		p.Stop()
		p.Rewind()
	case PlayerOnEndLoop:
		p.time -= duration
	case PlayerOnEndPause:
		p.Pause()
	}
}

// Time returns the current time
func (p *Player) Time() time.Duration {
	return p.time
}

// Play sets the playing state
func (p *Player) Play() {
	p.state = PlayerStatePlaying
}

func (p *Player) Playing() bool {
	return p.state == PlayerStatePlaying
}

// Stop resets the time to 0 and sets the stopped state
func (p *Player) Stop() {
	p.state = PlayerStateStopped
}

func (p *Player) Stopped() bool {
	return p.state == PlayerStateStopped
}

func (p *Player) Pause() {
	p.state = PlayerStatePaused
}

func (p *Player) Paused() bool {
	return p.state == PlayerStatePaused
}

func (p *Player) Rewind() {
	p.time = 0
}
