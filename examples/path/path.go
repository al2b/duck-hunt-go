package path

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

//go:embed assets/*.apng
var assets embed.FS

var animationKirbyBlow = engine.Must(engine.LoadAnimation(assets, "assets/kirby.blow.apng"))

func New() *Path {
	return &Path{
		animationKirbyBlow: engine.AnimationPlayer{Animation: animationKirbyBlow},
		pathLinear: engine.PathPlayer{
			Path: engine.LinearPath{engine.Vec(0, 0), engine.Vec(60, 28), time.Second * 3},
		},
		pathElastic: engine.PathPlayer{
			Path: engine.ElasticPath{engine.Vec(60, 0), engine.Vec(0, 28), time.Second * 3, 1, 0.25},
		},
	}
}

type Path struct {
	animationKirbyBlow engine.AnimationPlayer
	pathLinear         engine.PathPlayer
	pathElastic        engine.PathPlayer
}

func (s *Path) String() string {
	return "Path"
}

func (s *Path) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 80, Height: 50}
}

func (s *Path) FPS() int {
	return 60
}

func (s *Path) Init() (cmd tea.Cmd) {
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

func (s *Path) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(s.pathLinear.Position().Point(), s.animationKirbyBlow.Image()),
		engine.DrawImage(s.pathElastic.Position().Point(), s.animationKirbyBlow.Image()),
	)
}
