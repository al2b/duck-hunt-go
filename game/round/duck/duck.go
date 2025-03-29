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
	m := &Duck{space: space}
	m.animation = engine.AnimationPlayer{Animation: NewAnimation(m), Loop: true}
	return m
}

type Duck struct {
	space *space.Space
	space.Body
	animation engine.AnimationPlayer
}

func (m *Duck) Init() tea.Cmd {
	// Init space body
	m.Body = m.space.AddNewBody(1.0).
		SetPosition(engine.Vec(
			85+math.Round(rand.Float64()*85),
			layout.Ground-120,
		)).
		SetVelocity(engine.Vector{}.
			FromAngle(235 + (rand.Float64() * 90)).
			Scale(1),
		)

	m.Body.AddNewCircle(18).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Duck) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		// Follow mouse position
		m.SetPosition(engine.Vec(
			float64(msg.X), float64(msg.Y),
		))
		return engine.ConsoleLog("Go!")
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			m.SetVelocity(m.Velocity().Rotate(10))
		case tea.KeyLeft:
			m.SetVelocity(m.Velocity().Rotate(-10))
		}
	case engine.TickMsg:
		// Step animation
		m.animation.Step(msg.Duration)
	}

	return nil
}

func (m *Duck) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawCenteredImage(m.Position().Point(), m.animation.Image()),
	)
}
