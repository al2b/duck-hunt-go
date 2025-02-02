package engine

type Scene interface {
	Size() (int, int)
	FPS() int
	Model
	Sprites() Sprites
	Bodies() Bodies
}
