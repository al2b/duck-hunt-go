package animation

import (
	"duck-hunt-go/engine"
	"time"
)

type Animation struct {
	Size     engine.Size
	Duration time.Duration
	Frames   []Frame
	time     time.Duration
	current  int
}

func (a *Animation) Step(delta time.Duration) {
	a.time += delta
	if a.time >= a.Duration {
		a.time -= a.Duration
	}
}

func (a *Animation) Image() *engine.Image {
	var duration time.Duration
	for i, frame := range a.Frames {
		duration += frame.Duration
		if a.time <= duration {
			return a.Frames[i].Image
		}
	}
	return a.Frames[len(a.Frames)-1].Image
}

type Frame struct {
	Image    *engine.Image
	Duration time.Duration
}

type Animator interface {
	Animation() *Animation
}
