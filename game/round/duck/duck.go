package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"duck-hunt-go/game/round/layout"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
	"math/rand/v2"
)

func New(space *space.Space) *Duck {
	m := &Duck{
		space: space,
	}

	m.animation = engine.AnimationPlayer{
		Animation: NewAnimation(&m.body),
		OnEnd:     engine.PlayerOnEndLoop,
	}

	m.ImageDrawer = engine.ImageDrawer{
		engine.PointAdder{
			engine.Position2DPointer{&m.body},
			engine.Pt(-19, -19),
		},
		&m.animation,
	}

	return m
}

type Duck struct {
	space     *space.Space
	body      space.Body
	animation engine.AnimationPlayer
	engine.ImageDrawer
}

func (m *Duck) Init() tea.Cmd {
	m.animation.Play()

	// Init space body
	m.body = m.space.AddNewBody(1.0).
		SetPosition(engine.Vec2D(
			85+math.Round(rand.Float64()*85),
			layout.Ground-120,
		)).
		SetVelocity(engine.Vector2D{}.
			FromAngle(235 + (rand.Float64() * 90)).
			Scale(1),
		)

	m.body.AddNewCircle(18).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Duck) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		// Follow the mouse position
		m.body.SetPosition(engine.Vec2D(
			float64(msg.X), float64(msg.Y),
		))
		return engine.ConsoleLog("Go!")
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			m.body.SetVelocity(m.body.Velocity().Rotate(10))
		case tea.KeyLeft:
			m.body.SetVelocity(m.body.Velocity().Rotate(-10))
		}
	case engine.TickMsg:
		// Step animation
		m.animation.Step(msg.Duration)
	}

	return nil
}
