package engine

type Scene interface {
	Size(windowSize Size) Size
	DrawModel
}
