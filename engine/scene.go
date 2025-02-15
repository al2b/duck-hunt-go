package engine

type Scene interface {
	Size(windowSize Size) Size
	FPS() int
	Model
	Drawer
}
