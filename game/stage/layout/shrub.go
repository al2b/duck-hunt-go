package layout

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
)

func NewShrub(space *cp.Space) *Shrub {
	// Model
	m := &Shrub{}

	// Space
	{
		m.body = space.AddBody(cp.NewKinematicBody())
		m.body.SetPosition(cp.Vector{193, 122})

		vertices := []cp.Vector{
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
		}

		shape := space.AddShape(
			cp.NewPolyShapeRaw(m.body, len(vertices), vertices, 0),
		)
		shape.SetFilter(cp.NewShapeFilter(cp.NO_GROUP, config.ShapeCategoryLayout, config.ShapeCategoryDuck))
		shape.SetElasticity(1)
		shape.SetFriction(0)
	}

	return m
}

type Shrub struct {
	body *cp.Body
	engine.ImageDrawer
}

func (m *Shrub) Init() tea.Cmd {
	// Drawer
	m.ImageDrawer.Pointer = enginecp.PositionPointer{m.body}
	m.ImageDrawer.Imager = assets.LayoutShrub

	return nil
}

func (m *Shrub) Update(_ tea.Msg) tea.Cmd {
	return nil
}
