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

//go:embed assets/*.png
var assets embed.FS

func New(space *space.Space) *Layout {
	return &Layout{
		space: space,
		image: engine.MustLoadImage(engine.ImageFile(assets, "assets/layout.png")),
	}
}

type Layout struct {
	space *space.Space
	image *engine.Image
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
	scene.Draw(
		engine.DrawImage(image.Pt(0, 0), m.image),
	)
}
