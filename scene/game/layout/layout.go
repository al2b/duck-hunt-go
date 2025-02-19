package layout

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

const (
	Ground = 183
)

//go:embed assets/*
var assets embed.FS

func New(space *space.Space) *Layout {
	return &Layout{
		space: space,
		StaticImage: engine.NewStaticImage(
			engine.MustLoadImage(engine.ImagePngFile(assets, "assets/layout.png")),
		),
	}
}

type Layout struct {
	space *space.Space
	engine.AbsolutePosition
	engine.StaticImage
}

func (m *Layout) Init() tea.Cmd {
	// Space
	m.space.AddNewSegment(engine.Vec(0, 0), engine.Vec(255, 0), 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Vec(255, 0), engine.Vec(255, Ground), 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Vec(255, Ground), engine.Vec(0, Ground), 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Vec(0, Ground), engine.Vec(0, 0), 0).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Layout) Draw(scene *engine.Image) {
	position := m.Position()
	scene.DrawImage(image.Pt(
		int(position.X),
		int(position.Y),
	), m.Image())
}
