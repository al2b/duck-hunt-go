package layout

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const Ground = 183

func New(space *space.Space) *Layout {
	return &Layout{
		space: space,
		OrderedDrawer: engine.OrderedDrawer{
			engine.ImageDrawer{
				engine.Pt(0, 0),
				imageLayout,
			},
			engine.Order(0),
		},
	}
}

type Layout struct {
	space *space.Space
	engine.OrderedDrawer
}

func (m *Layout) Init() tea.Cmd {
	// Space
	m.space.AddNewSegment(engine.Vec2D(0, 0), engine.Vec2D(255, 0), 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Vec2D(255, 0), engine.Vec2D(255, Ground), 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Vec2D(255, Ground), engine.Vec2D(0, Ground), 0).
		SetElasticity(1).
		SetFriction(0)
	m.space.AddNewSegment(engine.Vec2D(0, Ground), engine.Vec2D(0, 0), 0).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}
