package animations

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*
var assets embed.FS

func New() *Scene {
	return &Scene{
		animationPng: engine.MustLoadAnimation(engine.AnimationFile(assets, "assets/parrot.png")),
		animationGif: engine.MustLoadAnimation(engine.AnimationFile(assets, "assets/parrot.gif")),
	}
}

type Scene struct {
	animationPng *engine.Animation
	animationGif *engine.Animation
}

func (s *Scene) String() string {
	return "Animations"
}

func (s *Scene) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 70, Height: 25}
}

func (s *Scene) FPS() int {
	return 60
}

func (s *Scene) Init() (cmd tea.Cmd) {
	return nil
}

func (s *Scene) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		s.animationPng.Step(msg.Duration)
		s.animationGif.Step(msg.Duration)
	}
	return nil
}

func (s *Scene) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(image.Pt(0, 0), s.animationPng.Image()),
		engine.DrawImage(image.Pt(35, 0), s.animationGif.Image()),
	)
}
