package gun

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"duck-hunt-go/game/assets"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

func New(space *space.Space) *Gun {
	m := &Gun{
		space: space,
		path: engine.Path2DPlayer{
			OnEnd: engine.PlayerOnEndPause,
		},
	}

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
	space *space.Space
	path  engine.Path2DPlayer
	engine.ImageDrawer
}

func (m *Gun) Init() tea.Cmd {
	// Path
	m.path.Path = engine.StaticPath2d{Position: engine.Vec2D(0, 0)}

	// Init space body
	m.space.AddNewPositionableBody(&m.path).
		AddNewPolygon(engine.Vectors2D{
			{-5.5, -18.5},
			{4.5, -18.5},
			{17.5, -5.5},
			{17.5, 4.5},
			{4.5, 17.5},
			{-5.5, 17.5},
			{-18.5, 4.5},
			{-18.5, -5.5},
		}, 0).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		mouse := msg.Mouse()
		// Path
		m.path.Path = engine.ElasticPath2D{
			m.path.Position(),
			engine.Vec2D(float64(mouse.X), float64(mouse.Y)),
			time.Second * 1,
			1, 0.25,
		}
		m.path.Rewind()
		m.path.Play()
	case engine.TickMsg:
		// Path
		m.path.Step(msg.Interval)
	}

	return nil
}
