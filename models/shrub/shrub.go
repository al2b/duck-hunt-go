package shrub

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	x = 193
	y = 122
	z = 20
)

var position = &engine.Position{X: x, Y: y, Z: z}

func New() *Shrub {
	// Model
	m := &Shrub{}

	// Body
	m.body = engine.NewBody(
		position,
		m.Intersect,
	).Shape(bodyShape)

	return m
}

type Shrub struct {
	// Body
	body *engine.Body
}

func (m *Shrub) Init() {}

func (m *Shrub) Update(_ []tea.Msg) {}

func (m *Shrub) Bodies() (bodies engine.Bodies) {
	return bodies.Append(m.body)
}

func (m *Shrub) Intersect() {}

func (m *Shrub) Sprites8() (sprites engine.Sprites8) {
	sprites.Append(sprite8)

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite8())
	}

	return sprites
}

func (m *Shrub) Sprites24() (sprites engine.Sprites24) {
	sprites.Append(sprite24)

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite24())
	}

	return sprites
}
