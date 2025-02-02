package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/scene/game/layout"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
	"math/rand/v2"
)

type Duck struct {
	engine.Motion
	Animation
	engine.RectangleShape
}

func (m *Duck) Init() tea.Cmd {
	// Init coordinates
	m.Coordinates = m.Coordinates.Set(
		85+math.Round(rand.Float64()*85),
		layout.Ground-height,
		5+math.Round(rand.Float64()*20),
	)

	// Init motion
	m.Motion = m.Motion.
		SetAngle(235 + (rand.Float64() * 90)).
		Scale(1)

	// Init animation
	m.Animation.Update(m.Angle())

	// Init shape
	m.RectangleShape = engine.NewRectangleShape(
		0, 0,
		width-1, height-1,
	)

	return nil
}

func (m *Duck) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		// Set coordinates
		m.Coordinates = m.Coordinates.SetXY(
			float64(msg.X),
			float64(msg.Y),
		)
		return engine.ConsoleLog("Duck!")
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			m.Motion = m.Motion.Rotate(10)
		case tea.KeyLeft:
			m.Motion = m.Motion.Rotate(-10)
		case tea.KeyUp:
			m.Coordinates = m.Coordinates.SubZ(10)
		case tea.KeyDown:
			m.Coordinates = m.Coordinates.AddZ(10)
		}
	case engine.TickMsg:
		// Update motion
		m.Motion = m.Motion.Update()
		// Update animation
		m.Animation.Update(m.Angle())
	}

	return nil
}
