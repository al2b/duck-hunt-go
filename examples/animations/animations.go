package animations

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*.apng assets/*.png assets/*.gif
var assets embed.FS

var (
	animationPng       = engine.Must(engine.LoadAnimation(assets, "assets/parrot.png"))
	animationGif       = engine.Must(engine.LoadAnimation(assets, "assets/parrot.gif"))
	animationKirbyWalk = engine.Must(engine.LoadAnimation(assets, "assets/kirby.walk.apng"))
	animationKirbyRun  = engine.Must(engine.LoadAnimation(assets, "assets/kirby.run.apng"))
	animationKirbyLand = engine.Must(engine.LoadAnimation(assets, "assets/kirby.land.apng"))
	animationKirby     = engine.SequenceAnimation{
		animationKirbyWalk, animationKirbyWalk,
		animationKirbyRun, animationKirbyRun, animationKirbyRun,
		animationKirbyLand, animationKirbyLand,
	}
)

func New() *Animations {
	return &Animations{}
}

type Animations struct {
	animationPng   engine.AnimationPlayer
	animationGif   engine.AnimationPlayer
	animationKirby engine.AnimationPlayer
}

func (s *Animations) String() string {
	return "Animations"
}

func (s *Animations) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 70, Height: 44}
}

func (s *Animations) FPS() int {
	return 60
}

func (s *Animations) Init() (cmd tea.Cmd) {
	s.animationPng = engine.AnimationPlayer{Animation: animationPng, Loop: true}
	s.animationGif = engine.AnimationPlayer{Animation: animationGif}
	s.animationKirby = engine.AnimationPlayer{Animation: animationKirby, Loop: true}
	return nil
}

func (s *Animations) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		s.animationPng.Step(msg.Duration)
		s.animationGif.Step(msg.Duration)
		s.animationKirby.Step(msg.Duration)
	}
	return nil
}

func (s *Animations) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(image.Pt(0, 0), s.animationPng.Image()),
		engine.DrawImage(image.Pt(35, 0), s.animationGif.Image()),
		engine.DrawImage(image.Pt(0, 25), s.animationKirby.Image()),
	)
}
