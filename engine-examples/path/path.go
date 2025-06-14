package path

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 60

var (
	//go:embed assets/*.apng
	assets embed.FS

	// Animations
	animationKirbyBlow = engine.MustLoad(engine.AnimationLoader{assets, "assets/kirby.blow.apng"})
)

func New() *Path {
	return &Path{
		animationKirbyBlow: engine.AnimationPlayer{Animation: animationKirbyBlow, OnEnd: engine.PlayerOnEndLoop},
		pathLinear: engine.Path2DPlayer{
			Path:  engine.LinearPath2D{engine.Vector2D{0, 0}, engine.Vector2D{59, 27}, time.Second * 3},
			OnEnd: engine.PlayerOnEndPause,
		},
		pathElastic: engine.Path2DPlayer{
			Path:  engine.ElasticPath2D{engine.Vector2D{60, 0}, engine.Vector2D{0, 28}, time.Second * 3, 1, 0.25},
			OnEnd: engine.PlayerOnEndLoop,
		},
	}
}

type Path struct {
	animationKirbyBlow engine.AnimationPlayer
	pathLinear         engine.Path2DPlayer
	pathElastic        engine.Path2DPlayer
}

func (s *Path) String() string {
	return "Path"
}

func (s *Path) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Path) Init() (cmd tea.Cmd) {
	s.animationKirbyBlow.Play()
	s.pathLinear.Play()
	s.pathElastic.Play()

	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Path) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		s.animationKirbyBlow.Step(msg.Interval)
		s.pathLinear.Step(msg.Interval)
		s.pathElastic.Step(msg.Interval)
	}

	return nil
}

func (s *Path) Draw(dst *engine.Image) {
	dst.Draw(
		engine.ImageDrawer{
			engine.Position2DPointer{&s.pathLinear},
			&s.animationKirbyBlow,
		},
		engine.ImageDrawer{
			engine.Position2DPointer{&s.pathElastic},
			&s.animationKirbyBlow,
		},
	)
}
