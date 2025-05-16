package gun

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
	"time"
)

func New(space *cp.Space) *Gun {
	// Model
	m := &Gun{
		path: engine.Path2DPlayer{
			OnEnd: engine.PlayerOnEndPause,
		},
	}

	// Drawer
	m.ImageDrawer = engine.ImageDrawer{
		engine.PointAdder{
			engine.Position2DPointer{&m.path},
			engine.Pt(-18, -18),
		},
		assets.Gun,
	}

	return m
}

type Gun struct {
	path engine.Path2DPlayer
	engine.ImageDrawer
}

func (m *Gun) Init() tea.Cmd {
	// Path
	m.path.Path = engine.StaticPath2d{Position: engine.Vec2D(0, 0)}

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Step path
		m.path.Step(msg.Interval)
	case tea.MouseMsg:
		mouse := msg.Mouse()
		// Update path
		m.path.Path = engine.ElasticPath2D{
			m.path.Position(),
			engine.Vec2D(float64(mouse.X), float64(mouse.Y)),
			time.Second * 1,
			1, 0.25,
		}
		m.path.Rewind()
		m.path.Play()
		switch msg := msg.(type) {
		case tea.MouseClickMsg:
			switch msg.Button {
			case tea.MouseLeft:
				return func() tea.Msg { return ShotMsg(m.path.Position()) }
			}
		}
	}

	return nil
}
