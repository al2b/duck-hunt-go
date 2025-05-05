package engine

import "time"

type TimerInterface interface {
	Duration() time.Duration
}

type Timer struct {
	Span time.Duration
}

func (t Timer) Duration() time.Duration {
	return t.Span
}

/**********/
/* Player */
/**********/

type TimerPlayer struct {
	Timer TimerInterface
	OnEnd PlayerOnEnd
	Player
}

func (p *TimerPlayer) Step(delta time.Duration) {
	p.Player.Step(delta, p.Timer.Duration(), p.OnEnd)
}
