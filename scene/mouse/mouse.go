package mouse

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Mouse struct {
	engine.Coordinates
	engine.StaticImage
}

func (m *Mouse) Init() tea.Cmd {
	// Init coordinates
	m.Coordinates = engine.NewCoordinates(0, 0, 1000)

	// Init image
	m.StaticImage = engine.NewStaticImage(imageMouse)

	return nil
}

func (m *Mouse) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		// Update coordinates
		m.Coordinates = m.Coordinates.SetXY(
			float64(msg.X-(width/2)),
			float64(msg.Y-(height/2)),
		)
	}

	return nil
}
