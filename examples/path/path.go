package path

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

var (
	//go:embed assets/*.apng
	assets embed.FS

	// Animations
	animationKirbyBlow = engine.Must(engine.LoadAnimation(assets, "assets/kirby.blow.apng"))
)

func New() *Path {
	return &Path{
		animationKirbyBlow: engine.AnimationPlayer{Animation: animationKirbyBlow, OnEnd: engine.PlayerOnEndLoop},
		pathLinear: engine.Path2DPlayer{
			Path:  engine.LinearPath2D{engine.Vec2D(0, 0), engine.Vec2D(59, 27), time.Second * 3},
			OnEnd: engine.PlayerOnEndPause,
		},
		pathElastic: engine.Path2DPlayer{
			Path:  engine.ElasticPath2D{engine.Vec2D(60, 0), engine.Vec2D(0, 28), time.Second * 3, 1, 0.25},
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

func (s *Path) TPS() int {
	return 60
}

func (s *Path) Init() (cmd tea.Cmd) {
	s.animationKirbyBlow.Play()
	s.pathLinear.Play()
	s.pathElastic.Play()

	return nil
}

func (s *Path) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		s.animationKirbyBlow.Step(msg.Duration)
		s.pathLinear.Step(msg.Duration)
		s.pathElastic.Step(msg.Duration)
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
