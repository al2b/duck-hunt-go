package gun

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
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

	// Space body
	body := space.AddBody(cp.NewKinematicBody())
	body.SetPositionUpdateFunc(enginecp.BodyPositioner2DFunc(&m.path))

	bodyShapeVertices := []cp.Vector{
		{-5.5, -18.5},
		{4.5, -18.5},
		{17.5, -5.5},
		{17.5, 4.5},
		{4.5, 17.5},
		{-5.5, 17.5},
		{-18.5, 4.5},
		{-18.5, -5.5},
	}

	bodyShape := space.AddShape(
		cp.NewPolyShapeRaw(body, len(bodyShapeVertices), bodyShapeVertices, 0),
	)
	bodyShape.SetElasticity(1)
	bodyShape.SetFriction(0)

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
		// Step path
		m.path.Step(msg.Interval)
	}

	return nil
}
