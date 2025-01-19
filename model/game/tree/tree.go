package tree

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	x = 6
	y = 32
	z = 10
)

var position = &engine.Position{X: x, Y: y, Z: z}

func New() *Tree {
	// Model
	m := &Tree{}

	// Bodies
	m.bodies = append(m.bodies,
		engine.NewBody(
			position,
			m.Intersect,
		).Shape(bodiesShape[0]),
	)

	return m
}

type Tree struct {
	// Bodies
	bodies []*engine.Body
}

func (m *Tree) Init() {}

func (m *Tree) Update(_ []tea.Msg) {}

func (m *Tree) Bodies() (bodies engine.Bodies) {
	return bodies.Append(m.bodies...)
}

func (m *Tree) Intersect() {}

func (m *Tree) Sprites8() (sprites engine.Sprites8) {
	sprites.Append(sprite8)

	// Debug
	if engine.Debug() {
		for _, body := range m.bodies {
			sprites.Append(body.Sprite8())
		}
	}

	return sprites
}

func (m *Tree) Sprites24() (sprites engine.Sprites24) {
	sprites.Append(sprite24)

	// Debug
	if engine.Debug() {
		for _, body := range m.bodies {
			sprites.Append(body.Sprite24())
		}
	}

	return sprites
}
