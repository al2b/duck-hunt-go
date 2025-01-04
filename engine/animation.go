package engine

type Animation int

type AnimationFrame struct {
	X, Y         int
	FlipH, FlipV bool
}

type AnimationFrames map[Animation][]AnimationFrame
