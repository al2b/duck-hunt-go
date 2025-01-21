package gun

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	width  = 37
	height = 37
)

var coordinates = engine.NewCoordinates(0, 0, 1000)

func New() *Gun {
	return &Gun{}
}

type Gun struct{}

func (m *Gun) Init() tea.Cmd {
	// Coordinates
	coordinates.Reset()

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		coordinates.
			SetX(float64(msg.X - (width / 2))).
			SetY(float64(msg.Y - (height / 2)))
	}

	return nil
}

func (m *Gun) Bodies() (bodies engine.Bodies) {
	return bodies.Append(body)
}

func (m *Gun) Sprites() (sprites engine.Sprites) {
	return sprites.Append(sprite)
}
