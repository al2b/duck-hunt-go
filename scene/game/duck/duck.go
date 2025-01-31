package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/scene/game/layout"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
	"math/rand/v2"
)

const (
	maxHeight = 32
)

func New() *Duck {
	return &Duck{}
}

type Duck struct {
	engine.Coordinates
	*engine.Animation
	// Movement
	movement engine.Vector
}

func (m *Duck) Init() tea.Cmd {
	// Coordinates
	m.Coordinates = engine.NewCoordinates(
		85+math.Round(rand.Float64()*85),
		layout.Ground-maxHeight,
		5+math.Round(rand.Float64()*20),
	)

	// Movement
	m.movement = engine.
		VectorFromAngle(235 + (rand.Float64() * 90)).
		Scale(1)

	// Animation
	m.Animation = m.animation()

	return nil
}

func (m *Duck) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		// Coordinates
		m.Coordinates = m.Coordinates.SetXY(
			float64(msg.X),
			float64(msg.Y),
		)
		return engine.ConsoleLog("Duck!")
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			m.movement = m.movement.Rotate(10)
		case tea.KeyLeft:
			m.movement = m.movement.Rotate(-10)
		case tea.KeyUp:
			m.Coordinates = m.Coordinates.SubZ(10)
		case tea.KeyDown:
			m.Coordinates = m.Coordinates.AddZ(10)
		}
	case engine.TickMsg:
		// Coordinates
		m.Coordinates = m.Coordinates.Move(m.movement)
		// Animation
		animation := m.animation()
		if animation != m.Animation {
			m.Animation = animation
		} else {
			m.Animation.Update()
		}
	}

	return nil
}

func (m *Duck) Sprites() engine.Sprites {
	return engine.Sprites{m}
}

func (m *Duck) animation() *engine.Animation {
	angle := m.movement.Angle()

	animation := animationFlyRight

	switch true {
	case 23 <= angle && angle <= 67:
		animation = animationFlyBottomRight
	case 68 <= angle && angle <= 112:
		animation = animationFlyBottom
	case 113 <= angle && angle <= 157:
		animation = animationFlyBottomLeft
	case 158 <= angle && angle <= 202:
		animation = animationFlyLeft
	case 203 <= angle && angle <= 247:
		animation = animationFlyTopLeft
	case 248 <= angle && angle <= 292:
		animation = animationFlyTop
	case 293 <= angle && angle <= 337:
		animation = animationFlyTopRight
	}

	return animations[animation]
}

func (m *Duck) Bodies() (bodies engine.Bodies) {
	return bodies.Append(
		engine.NewCoordinatedBody(m.Coordinates,
			engine.BodyShape{
				{0, 0},
				{31, 0},
				{31, 31},
				{0, 31},
			},
		),
	)
}
