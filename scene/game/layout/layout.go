package layout

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	Ground = 183
)

//go:embed assets/*
var assets embed.FS

func New(space *space.Space) *Layout {
	return &Layout{
		space:            space,
		AbsolutePosition: engine.NewAbsolutePosition(0, 0),
		StaticImage: engine.NewStaticImage(
			engine.MustLoadImage(assets, "assets/layout.png"),
		),
	}
}

type Layout struct {
	space *space.Space
	*engine.AbsolutePosition
	*engine.StaticImage
}

func (m *Layout) Init() tea.Cmd {
	// Init space segments
	m.space.AddNewSegment(engine.Position{0, 0}, engine.Position{255, 0}, 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Position{255, 0}, engine.Position{255, Ground}, 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Position{255, Ground}, engine.Position{0, Ground}, 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Position{0, Ground}, engine.Position{0, 0}, 0).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Layout) Draw(scene *engine.Image) {
	scene.DrawImage(m.Position(), m.Image())
}
