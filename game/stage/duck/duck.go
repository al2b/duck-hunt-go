package duck

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
	"duck-hunt-go/game/config"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
	"math"
	"math/rand/v2"
)

func New(space *cp.Space) *Duck {
	// Model
	m := &Duck{}

	// Space body
	bodyMass := 1.0
	bodyShapeRadius := 18.0
	m.body = space.AddBody(cp.NewBody(bodyMass, cp.MomentForCircle(bodyMass, 0, bodyShapeRadius, cp.Vector{})))
	bodyShape := space.AddShape(cp.NewCircle(m.body, bodyShapeRadius, cp.Vector{}))
	bodyShape.SetElasticity(1)
	bodyShape.SetFriction(0)

	// Animation
	m.animation = engine.AnimationPlayer{
		Animation: NewAnimation(enginecp.VelocityVelociter{m.body}),
		OnEnd:     engine.PlayerOnEndLoop,
	}

	// Drawer
	m.ImageDrawer = engine.ImageDrawer{
		engine.PointAdder{
			enginecp.PositionPointer{m.body},
			engine.Pt(-19, -19),
		},
		&m.animation,
	}

	return m
}

type Duck struct {
	body      *cp.Body
	animation engine.AnimationPlayer
	engine.ImageDrawer
}

func (m *Duck) Init() tea.Cmd {
	m.animation.Play()

	// Init space body
	m.body.SetPosition(cp.Vector{
		85 + math.Round(rand.Float64()*85),
		config.Ground - 120,
	})
	m.body.SetVelocityVector(cp.
		ForAngle(engine.Radians(235 + (rand.Float64() * 90))).
		Mult(1),
	)

	return nil
}

func (m *Duck) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		// Follow the mouse position
		m.body.SetPosition(cp.Vector{
			float64(msg.X), float64(msg.Y),
		})
		return engine.ConsoleLog("Go!")
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			m.body.SetVelocityVector(m.body.Velocity().Rotate(cp.ForAngle(engine.Radians(10))))
		case tea.KeyLeft:
			m.body.SetVelocityVector(m.body.Velocity().Rotate(cp.ForAngle(engine.Radians(-10))))
		}
	case engine.TickMsg:
		// Step animation
		m.animation.Step(msg.Interval)
	}

	return nil
}
