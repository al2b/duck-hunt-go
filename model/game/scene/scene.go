package scene

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	x      = 0
	y      = 0
	z      = 100
	width  = engine.Width
	ground = 184
)

var position = &engine.Position{X: x, Y: y, Z: z}

func New() *Scene {
	// Model
	m := &Scene{}

	// Body
	m.body = engine.NewBody(
		position,
		m.Intersect,
	).Shape(bodyShape)

	return m
}

type Scene struct {
	// Body
	body *engine.Body
}

func (m *Scene) Init() {}

func (m *Scene) Update(_ []tea.Msg) {}

func (m *Scene) Bodies() (bodies engine.Bodies) {
	return bodies.Append(m.body)
}

func (m *Scene) Intersect() {}

func (m *Scene) Sprites8() (sprites engine.Sprites8) {
	sprites.Append(sprite8)

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite8())
	}

	return sprites
}

func (m *Scene) Sprites24() (sprites engine.Sprites24) {
	sprites.Append(sprite24)

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite24())
	}

	return sprites
}
