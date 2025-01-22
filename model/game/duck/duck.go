package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/model/game/layout"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
	"math/rand/v2"
)

const (
	maxHeight = 32
)

var coordinates = engine.NewCoordinates(0, 0, 0)

func New() *Duck {
	return &Duck{}
}

type Duck struct {
	// Movement
	movement engine.Vector
}

func (m *Duck) Init() tea.Cmd {
	// Coordinates
	coordinates.
		SetX(85 + math.Round(rand.Float64()*85)).
		SetY(layout.Ground - maxHeight).
		SetZ(5 + math.Round(rand.Float64()*20))

	// Movement
	m.movement = engine.
		VectorFromAngle(235 + (rand.Float64() * 90)).
		Scale(1)

	return nil
}

func (m *Duck) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		coordinates.
			SetX(float64(msg.X)).
			SetY(float64(msg.Y))
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			m.movement = m.movement.Rotate(-10)
		case tea.KeyLeft:
			m.movement = m.movement.Rotate(10)
		case tea.KeyUp:
			coordinates.SubZ(10)
		case tea.KeyDown:
			coordinates.AddZ(10)
		}
	case engine.TickMsg:
		// Coordinates
		coordinates.Move(m.movement)
		// Animation
		m.animation().Update()
	}

	return nil
}

func (m *Duck) Bodies() (bodies engine.Bodies) {
	return bodies.Append(body)
}

func (m *Duck) Sprites() (sprites engine.Sprites) {
	return sprites.Append(
		m.animation().Sprite(coordinates),
	)
}

func (m *Duck) animation() *engine.Animation {
	angle := m.movement.Angle()

	animation := animationFlyRight

	switch true {
	case 23 <= angle && angle <= 67:
		animation = animationFlyTopRight
	case 68 <= angle && angle <= 112:
		animation = animationFlyTop
	case 113 <= angle && angle <= 157:
		animation = animationFlyTopLeft
	case 158 <= angle && angle <= 202:
		animation = animationFlyLeft
	case 203 <= angle && angle <= 247:
		animation = animationFlyBottomLeft
	case 248 <= angle && angle <= 292:
		animation = animationFlyBottom
	case 293 <= angle && angle <= 337:
		animation = animationFlyBottomRight
	}

	return animations[animation]
}
