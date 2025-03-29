package space

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
)

//go:embed assets/*.png
var assets embed.FS

var (
	imageKirby = engine.Must(engine.LoadImage(assets, "assets/kirby.png"))
)

func New() *Space {
	// Space
	space := space.NewSpace().
		SetGravity(engine.Vec(0, 9.8))

	return &Space{
		space:      space,
		imageKirby: imageKirby,
	}
}

type Space struct {
	space      *space.Space
	body       space.Body
	imageKirby *engine.Image
}

func (s *Space) String() string {
	return "Space"
}

func (s *Space) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 80, Height: 50}
}

func (s *Space) FPS() int {
	return 60
}

func (s *Space) Init() (cmd tea.Cmd) {
	// Space
	s.space.Clear()
	s.space.AddNewSegment(engine.Vec(0, 0), engine.Vec(79, 0), 0).
		SetElasticity(1).
		SetFriction(0)
	s.space.AddNewSegment(engine.Vec(79, 0), engine.Vec(79, 49), 0).
		SetElasticity(1).
		SetFriction(0)
	s.space.AddNewSegment(engine.Vec(79, 49), engine.Vec(0, 49), 0).
		SetElasticity(1).
		SetFriction(0)
	s.space.AddNewSegment(engine.Vec(0, 49), engine.Vec(0, 0), 0).
		SetElasticity(1).
		SetFriction(0)

	// Body
	s.body = s.space.AddNewBody(1.0).
		SetPosition(engine.Vec(10, 10)).
		SetVelocity(engine.Vec(30, 100))
	s.body.AddNewCircle(10).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (s *Space) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Update space
		s.space.Step(msg.Duration)
	}

	return nil
}

func (s *Space) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawCenteredImage(s.body.Position().Point(), s.imageKirby),
		s.space,
	)
}
