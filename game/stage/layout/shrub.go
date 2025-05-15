package layout

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
	"duck-hunt-go/game/assets"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
)

func NewShrub(space *cp.Space) *Shrub {
	// Model
	m := &Shrub{}

	// Space body
	m.body = space.AddBody(cp.NewKinematicBody())
	m.body.SetPosition(cp.Vector{193, 122})

	bodyShapeVertices := []cp.Vector{
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

	bodyShape := space.AddShape(
		cp.NewPolyShapeRaw(m.body, len(bodyShapeVertices), bodyShapeVertices, 0),
	)
	bodyShape.SetElasticity(1)
	bodyShape.SetFriction(0)

	// Drawer
	m.ImageDrawer = engine.ImageDrawer{
		enginecp.PositionPointer{m.body},
		assets.LayoutShrub,
	}

	return m
}

type Shrub struct {
	body *cp.Body
	engine.ImageDrawer
}

func (m *Shrub) Init() tea.Cmd {
	return nil
}

func (m *Shrub) Update(_ tea.Msg) tea.Cmd {
	return nil
}
