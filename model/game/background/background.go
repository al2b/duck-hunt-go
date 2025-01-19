package background

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	x = 0
	y = 0
	z = 0
)

var (
	position = &engine.Position{X: x, Y: y, Z: z}
	sprite8  = &engine.Sprite8{
		Position: position,
		Image:    engine.NewImageUniform8(engine.Color8(117)),
	}
	sprite24 = &engine.Sprite24{
		Position: position,
		Image:    engine.NewImageUniform24(engine.Color24{R: 143, G: 192, B: 255, A: 255}),
	}
)

func New() *Background {
	return &Background{}
}

type Background struct{}

func (m *Background) Init() {}

func (m *Background) Update(_ []tea.Msg) {}

func (m *Background) Bodies() (bodies engine.Bodies) {
	return bodies
}

func (m *Background) Sprites8() (sprites engine.Sprites8) {
	return sprites.Append(sprite8)
}

func (m *Background) Sprites24() (sprites engine.Sprites24) {
	return sprites.Append(sprite24)
}
