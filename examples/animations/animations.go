package animations

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 60

var (
	//go:embed assets/*.apng assets/*.png assets/*.gif
	assets embed.FS

	// Animations
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
	return &Animations{
		animationPng:   engine.AnimationPlayer{Animation: animationPng, OnEnd: engine.PlayerOnEndLoop},
		animationGif:   engine.AnimationPlayer{Animation: animationGif},
		animationKirby: engine.AnimationPlayer{Animation: animationKirby, OnEnd: engine.PlayerOnEndLoop},
	}
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
	return engine.Size{70, 44}
}

func (s *Animations) Init() (cmd tea.Cmd) {
	s.animationPng.Play()
	s.animationGif.Play()
	s.animationKirby.Play()

	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Animations) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		s.animationPng.Step(msg.Interval)
		s.animationGif.Step(msg.Interval)
		s.animationKirby.Step(msg.Interval)
	}

	return nil
}

func (s *Animations) Draw(dst *engine.Image) {
	dst.Draw(
		engine.ImageDrawer{engine.Pt(0, 0), s.animationPng.Image()},
		engine.ImageDrawer{engine.Pt(35, 0), s.animationGif.Image()},
		engine.ImageDrawer{engine.Pt(0, 25), s.animationKirby.Image()},
	)
}
