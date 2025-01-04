package gun

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	width  = 37
	height = 37
)

func New() *Gun {
	// Model
	m := &Gun{
		position: engine.NewPosition(),
	}

	// Body
	m.body = engine.NewBody(
		m.position,
		m.Intersect,
	).Shape(bodyShape)

	return m
}

type Gun struct {
	// Position
	position *engine.Position
	// Body
	body *engine.Body
}

func (m *Gun) Init() {
	// Position
	m.position.Z = 1000
}

func (m *Gun) Update(msgs []tea.Msg) {
	// Messages
	for _, msg := range msgs {
		switch msg := msg.(type) {
		case tea.MouseMotionMsg:
			m.position.X = float64(msg.X - (width / 2))
			m.position.Y = float64(msg.Y - (height / 2))
		}
	}
}

func (m *Gun) Bodies() (bodies engine.Bodies) {
	return bodies.Append(m.body)
}

func (m *Gun) Intersect() {}

func (m *Gun) Sprites8() (sprites engine.Sprites8) {
	sprites.Append(&engine.Sprite8{
		Position: m.position,
		Image:    sprite8Image,
	})

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite8())
	}

	return sprites
}

func (m *Gun) Sprites24() (sprites engine.Sprites24) {
	sprites.Append(&engine.Sprite24{
		Position: m.position,
		Image:    sprite24Image,
	})

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite24())
	}

	return sprites
}
