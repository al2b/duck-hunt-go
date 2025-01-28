package gun

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Gun {
	return &Gun{
		motion: engine.NewMotion(),
	}
}

type Gun struct {
	coordinates engine.Coordinates
	motion      *engine.Motion
}

func (m *Gun) Init() tea.Cmd {
	// Coordinates
	m.coordinates = engine.NewCoordinates(0, 0, 1000)

	// Motion
	m.motion.Reset()

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		// Motion
		m.motion.MoveTo(m.coordinates,
			float64(msg.X-(imageWidth/2)),
			float64(msg.Y-(imageHeight/2)),
			10,
		)
	case engine.TickMsg:
		// Motion
		m.coordinates = m.motion.Update(m.coordinates)
	}

	return nil
}

func (m *Gun) Sprites() engine.Sprites {
	return engine.Sprites{
		engine.NewCoordinatedSprite(
			m.coordinates,
			imageGun,
		),
	}
}

func (m *Gun) Bodies() (bodies engine.Bodies) {
	return bodies.Append(
		engine.NewBody(m.coordinates,
			engine.BodyShape{
				{13, 0},
				{23, 0},
				{36, 13},
				{36, 23},
				{23, 36},
				{13, 36},
				{0, 23},
				{0, 13},
			},
		),
	)
}
