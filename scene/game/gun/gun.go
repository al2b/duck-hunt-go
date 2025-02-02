package gun

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Gun struct {
	engine.Coordinates
	engine.StaticImage
	engine.PolygonShape
	motion *engine.Motion
}

func (m *Gun) Init() tea.Cmd {
	// Init coordinates
	m.Coordinates = engine.NewCoordinates(0, 0, 1000)

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

	// Init motion
	m.motion = engine.NewMotion()

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		// Mutate motion
		m.motion.MoveTo(m.Coordinates,
			float64(msg.X-(width/2)),
			float64(msg.Y-(height/2)),
			10,
		)
	case engine.TickMsg:
		// Update motion
		m.Coordinates = m.motion.Update(m.Coordinates)
	}

	return nil
}
