package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*
var assets embed.FS

func New(space *space.Space) *Duck {
	m := &Duck{space: space}
	m.Animator = NewAnimator(m)
	return m
}

type Duck struct {
	space *space.Space
	space.Body
	Animator
}

func (m *Duck) Init() tea.Cmd {
	// Init space body
	m.Body = m.space.AddNewBody(1.0).
		SetPosition(engine.Vec(
			128,
			1,
		)).
		SetVelocity(engine.Vector{}.
			FromAngle(270).
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
		m.Animation().Step(msg.Duration)
	}

	return nil
}

func (m *Duck) Draw(scene *engine.Image) {
	position := m.Position()
	scene.DrawCenteredImage(image.Pt(
		int(position.X),
		int(position.Y),
	), m.Animation().Image())
}
