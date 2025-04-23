package layout

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewShrub(space *space.Space) *Shrub {
	m := &Shrub{
		space: space,
	}

	m.ImageDrawer = engine.ImageDrawer{
		engine.Position2DPointer{&m.position},
		imageShrub,
	}

	return m
}

type Shrub struct {
	space    *space.Space
	position engine.Vector2D
	engine.ImageDrawer
}

func (m *Shrub) Init() tea.Cmd {
	// Position
	m.position = engine.Vec2D(193, 122)

	// Space
	m.space.AddNewPositionableBody(&m.position).
		AddNewPolygon(engine.Vectors2D{
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
