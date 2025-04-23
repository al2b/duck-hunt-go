package mouse

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Mouse {
	m := &Mouse{}

	m.ImageDrawer = engine.ImageDrawer{
		engine.PointAdder{&m.point, engine.Pt(-1, -1)},
		imageMouse,
	}

	return m
}

type Mouse struct {
	point engine.Point
	engine.ImageDrawer
}

func (m *Mouse) Init() tea.Cmd {
	m.point = engine.Pt(0, 0)

	return nil
}

func (m *Mouse) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		m.point = engine.Pt(msg.X, msg.Y)
	}

	return nil
}
