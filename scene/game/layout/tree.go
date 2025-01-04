package layout

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

func NewTree(space *space.Space) *Tree {
	return &Tree{
		space: space,
		StaticImage: engine.NewStaticImage(
			engine.MustLoadImage(assets, "assets/tree.png"),
		),
	}
}

type Tree struct {
	space *space.Space
	engine.AbsolutePosition
	engine.StaticImage
}

func (m *Tree) Init() tea.Cmd {
	// Position
	m.SetPosition(engine.Vec(6, 32))
	// Space
	body := m.space.AddNewPositionableBody(m)
	body.AddNewPolygon(engine.Vectors{
		{22, 0},
		{26, 0},
		{33, 6},
		{36, 11},
		{35, 17},
		{33, 20},
		{29, 22},
		{15, 22},
		{11, 20},
		{7, 15},
		{7, 10},
		{8, 7},
		{13, 2},
	}, 0).
		SetElasticity(1).
		SetFriction(0)
	body.AddNewPolygon(engine.Vectors{
		{42, 24},
		{51, 25},
		{57, 30},
		{57, 34},
		{56, 39},
		{52, 40},
		{51, 42},
		{46, 43},
		{39, 41},
		{35, 39},
		{32, 36},
		{32, 31},
		{33, 28},
		{37, 25},
	}, 0).
		SetElasticity(1).
		SetFriction(0)
	body.AddNewPolygon(engine.Vectors{
		{31, 42},
		{35, 42},
		{38, 43},
		{42, 46},
		{43, 51},
		{40, 55},
		{32, 56},
		{25, 54},
		{24, 50},
		{26, 45},
	}, 0).
		SetElasticity(1).
		SetFriction(0)
	body.AddNewPolygon(engine.Vectors{
		{8, 48},
		{14, 48},
		{18, 49},
		{22, 55},
		{22, 60},
		{19, 65},
		{11, 67},
		{7, 65},
		{3, 63},
		{0, 60},
		{0, 55},
		{1, 52},
		{4, 50},
	}, 0).
		SetElasticity(1).
		SetFriction(0)
	body.AddNewPolygon(engine.Vectors{
		{54, 56},
		{58, 56},
		{64, 61},
		{67, 66},
		{68, 68},
		{67, 73},
		{65, 76},
		{59, 79},
		{47, 78},
		{43, 76},
		{39, 71},
		{39, 66},
		{40, 64},
		{42, 61},
		{45, 58},
	}, 0).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Tree) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Tree) Draw(scene *engine.Image) {
	position := m.Position()
	scene.DrawImage(image.Pt(
		int(position.X),
		int(position.Y),
	), m.Image())
}
