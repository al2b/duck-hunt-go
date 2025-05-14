package space

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 60

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageKirby = engine.Must(engine.LoadImage(assets, "assets/kirby.png"))
)

func New() *Space {
	return &Space{
		space: space.NewSpace().
			SetGravity(engine.Vec2D(0, 9.8)),
	}
}

type Space struct {
	space *space.Space
	body  space.Body
}

func (s *Space) String() string {
	return "Space"
}

func (s *Space) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Space) Init() (cmd tea.Cmd) {
	// Space
	s.space.Clear()
	s.space.AddNewSegment(engine.Vec2D(0, 0), engine.Vec2D(79, 0), 0).
		SetElasticity(1).
		SetFriction(0)
	s.space.AddNewSegment(engine.Vec2D(79, 0), engine.Vec2D(79, 49), 0).
		SetElasticity(1).
		SetFriction(0)
	s.space.AddNewSegment(engine.Vec2D(79, 49), engine.Vec2D(0, 49), 0).
		SetElasticity(1).
		SetFriction(0)
	s.space.AddNewSegment(engine.Vec2D(0, 49), engine.Vec2D(0, 0), 0).
		SetElasticity(1).
		SetFriction(0)

	// Body
	s.body = s.space.AddNewBody(1.0).
		SetPosition(engine.Vec2D(10, 10)).
		SetVelocity(engine.Vec2D(30, 100))
	s.body.AddNewCircle(10).
		SetElasticity(1).
		SetFriction(0)

	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Space) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Update space
		s.space.Step(msg.Interval)
	}

	return nil
}

func (s *Space) Draw(dst *engine.Image) {
	dst.Draw(
		engine.ImageDrawer{
			engine.PointAdder{
				engine.Position2DPointer{s.body},
				engine.Pt(-9, -9),
			},
			imageKirby,
		},
		s.space,
	)
}
