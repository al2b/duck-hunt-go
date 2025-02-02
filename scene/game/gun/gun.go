package gun

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Gun struct {
	*engine.Path
	engine.StaticImage
	engine.PolygonShape
}

func (m *Gun) Init() tea.Cmd {
	// Init path
	m.Path = engine.NewPath()

	// Init image
	m.StaticImage = engine.NewStaticImage(imageGun)

	// Init shape
	m.PolygonShape = engine.NewPolygonShape(
		13, 0,
		23, 0,
		36, 13,
		36, 23,
		23, 36,
		13, 36,
		0, 23,
		0, 13,
	)

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		// Change path
		m.Path.To(
			float64(msg.X-(width/2)),
			float64(msg.Y-(height/2)),
			10,
		)
	case engine.TickMsg:
		// Update path
		m.Path.Update()
	}

	return nil
}
