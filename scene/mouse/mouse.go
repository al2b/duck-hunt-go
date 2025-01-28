package mouse

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Mouse {
	return &Mouse{}
}

type Mouse struct {
	coordinates engine.Coordinates
}

func (m *Mouse) Init() tea.Cmd {
	// Coordinates
	m.coordinates = engine.NewCoordinates(0, 0, 1000)

	return nil
}

func (m *Mouse) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		m.coordinates = m.coordinates.SetXY(
			float64(msg.X-(imageWidth/2)),
			float64(msg.Y-(imageHeight/2)),
		)
	}

	return nil
}

func (m *Mouse) Sprites() engine.Sprites {
	return engine.Sprites{
		engine.NewCoordinatedSprite(
			m.coordinates,
			imageMouse,
		),
	}
}

func (m *Mouse) Bodies() engine.Bodies {
	return nil
}
