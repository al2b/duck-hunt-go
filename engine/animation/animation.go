package animation

import (
	"duck-hunt-go/engine"
	"time"
)

type Animation struct {
	size     engine.Size
	duration time.Duration
	frames   []Frame
	time     time.Duration
	current  int
}

func (a *Animation) Size() engine.Size {
	return a.size
}

func (a *Animation) Step(delta time.Duration) {
	a.time += delta
	if a.time >= a.duration {
		a.time -= a.duration
	}
}

func (a *Animation) Image() *engine.Image {
	var duration time.Duration
	for i, frame := range a.frames {
		duration += frame.Duration
		if a.time <= duration {
			return a.frames[i].Image
		}
	}
	return a.frames[len(a.frames)-1].Image
}

type Frame struct {
	Image    *engine.Image
	Duration time.Duration
}

type Animator interface {
	Animation() *Animation
}
