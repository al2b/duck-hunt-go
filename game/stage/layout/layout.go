package layout

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
)

func New(space *cp.Space) *Layout {
	// Model
	m := &Layout{}

	// Space
	{
		borders := []cp.Vector{
			{0, 0}, {255, 0},
			{255, 0}, {255, config.Ground},
			{255, config.Ground}, {0, config.Ground},
			{0, config.Ground}, {0, 0},
		}

		for i := 0; i < len(borders)-1; i += 2 {
			shape := space.AddShape(cp.NewSegment(space.StaticBody, borders[i], borders[i+1], 0))
			shape.SetElasticity(1)
			shape.SetFriction(0)
		}
	}

	return m
}

type Layout struct {
	engine.OrderedDrawer
}

func (m *Layout) Init() tea.Cmd {
	// Drawer
	m.OrderedDrawer.Drawer = engine.ImageDrawer{
		engine.Pt(0, 0),
		assets.Layout,
	}
	m.OrderedDrawer.Orderer = engine.Order(0)

	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}
