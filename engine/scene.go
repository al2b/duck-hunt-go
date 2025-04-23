package engine

type Scene interface {
	Size(windowSize Size) Size
	TPS() int
	Model
	Drawer
}
