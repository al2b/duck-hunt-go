package layout

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

func NewShrub(space *space.Space) *Shrub {
	return &Shrub{
		space: space,
		StaticImage: engine.NewStaticImage(
			engine.MustLoadImage(engine.ImagePngFile(assets, "assets/shrub.png")),
		),
	}
}

type Shrub struct {
	space *space.Space
	engine.AbsolutePosition
	engine.StaticImage
}

func (m *Shrub) Init() tea.Cmd {
	// Position
	m.SetPosition(engine.Vec(193, 122))
	// Space
	m.space.AddNewPositionableBody(m).
		AddNewPolygon(engine.Vectors{
			{0, 61},
			{0, 29},
			{1, 22},
			{3, 16},
			{7, 11},
			{8, 7},
			{9, 4},
			{16, 0},
			{23, 2},
			{25, 4},
			{29, 15},
			{30, 25},
			{30, 61},
		}, 0).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Shrub) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Shrub) Draw(scene *engine.Image) {
	position := m.Position()
	scene.DrawImage(image.Pt(
		int(position.X),
		int(position.Y),
	), m.Image())
}
