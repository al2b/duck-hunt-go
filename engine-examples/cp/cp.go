package cp

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
	"time"
)

const TickInterval = time.Second / 60

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageKirby = engine.Must(engine.LoadImage(assets, "assets/kirby.png"))
)

func New() *Cp {
	return &Cp{}
}

type Cp struct {
	space *cp.Space
	body  *cp.Body
}

func (s *Cp) String() string {
	return "Cp"
}

func (s *Cp) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Cp) Init() (cmd tea.Cmd) {
	// Space
	s.space = cp.NewSpace()
	s.space.SetGravity(cp.Vector{0, 9.8})

	// Space walls
	walls := []cp.Vector{
		{0, 0}, {79, 0},
		{79, 0}, {79, 49},
		{79, 49}, {0, 49},
		{0, 49}, {0, 0},
	}
	for i := 0; i < len(walls)-1; i += 2 {
		shape := s.space.AddShape(cp.NewSegment(s.space.StaticBody, walls[i], walls[i+1], 0))
		shape.SetElasticity(1)
		shape.SetFriction(0)
	}

	// Space body
	bodyMass := 1.0
	bodyShapeRadius := 10.0
	s.body = s.space.AddBody(cp.NewBody(bodyMass, cp.MomentForCircle(bodyMass, 0, bodyShapeRadius, cp.Vector{})))
	s.body.SetPosition(cp.Vector{10, 10})
	s.body.SetVelocityVector(cp.Vector{30, 100})
	bodyShape := s.space.AddShape(cp.NewCircle(s.body, bodyShapeRadius, cp.Vector{}))
	bodyShape.SetElasticity(1)
	bodyShape.SetFriction(0)

	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Cp) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Step space
		s.space.Step(msg.Interval.Seconds())
	}

	return nil
}

func (s *Cp) Draw(dst *engine.Image) {
	dst.Draw(
		engine.ImageDrawer{
			engine.PointAdder{
				enginecp.PositionPointer{s.body},
				engine.Pt(-9, -9),
			},
			imageKirby,
		},
		enginecp.SpaceDrawer{s.space},
	)
}
