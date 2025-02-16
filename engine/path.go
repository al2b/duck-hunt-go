package engine

import (
	"math"
	"time"
)

func NewPath() *Path {
	return &Path{}
}

type Path struct {
	position   Vector
	start, end Vector
	time       time.Duration
	duration   time.Duration
	easing     Easing
}

func (path *Path) Move(x, y float64) {
	path.position.X, path.position.Y = x, y
	path.time, path.duration = 0, 0
}

func (path *Path) Position() Vector {
	return path.position
}

func (path *Path) Step(delta time.Duration) {
	if path.time >= path.duration {
		return
	}
	path.time += delta
	progress := math.Min(
		float64(path.time)/float64(path.duration),
		1,
	)
	px, py := path.easing(progress)
	path.position.X = path.start.X + (path.end.X-path.start.X)*px
	path.position.Y = path.start.Y + (path.end.Y-path.start.Y)*py
}

func (path *Path) To(to Vector, easing Easing, duration time.Duration) {
	path.start, path.end = path.position, to
	path.easing = easing
	path.duration = duration
	path.time = 0
}

type Easing func(t float64) (float64, float64)

func LinearEasing() Easing {
	return func(t float64) (float64, float64) {
		return t, t
	}
}

func InQuadEasing() Easing {
	return func(t float64) (float64, float64) {
		progress := t * t
		return progress, progress
	}
}

func OutQuadEasing() Easing {
	return func(t float64) (float64, float64) {
		progress := 1 - (1-t)*(1-t)
		return progress, progress
	}
}

func InOutCubicEasing() Easing {
	return func(t float64) (float64, float64) {
		if t < 0.5 {
			return 4 * t * t * t, 4 * t * t * t
		}
		return 1 - math.Pow(-2*t+2, 3)/2,
			1 - math.Pow(-2*t+2, 3)/2
	}
}

func ElasticEasing(amplitude, period float64) Easing {
	return func(t float64) (float64, float64) {
		return math.Pow(2, -10*t)*math.Sin((t-period/4)*(2*math.Pi)/period)*amplitude + 1,
			math.Pow(2, -10*t)*math.Sin((t-period/4)*(2*math.Pi)/period)*amplitude + 1
	}
}

func BounceEasing(factor float64) Easing {
	return func(t float64) (float64, float64) {
		if t < 1/2.75 {
			return factor * t * t, factor * t * t
		} else if t < 2/2.75 {
			t -= 1.5 / 2.75
			return factor*t*t + 0.75, factor*t*t + 0.75
		} else if t < 2.5/2.75 {
			t -= 2.25 / 2.75
			return factor*t*t + 0.9375, factor*t*t + 0.9375
		}
		t -= 2.625 / 2.75
		return factor*t*t + 0.984375, factor*t*t + 0.984375
	}
}
