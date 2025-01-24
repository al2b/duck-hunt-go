package engine

type Scene interface {
	Width() int
	Height() int
	FPS() int
	Model
}
