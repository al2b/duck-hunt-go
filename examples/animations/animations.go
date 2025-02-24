package animations

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*.png assets/*.gif
var assets embed.FS

func New() *Animations {
	return &Animations{
		animationPng: engine.MustLoadAnimation(engine.AnimationFile(assets, "assets/parrot.png")),
		animationGif: engine.MustLoadAnimation(engine.AnimationFile(assets, "assets/parrot.gif")),
	}
}

type Animations struct {
	animationPng *engine.Animation
	animationGif *engine.Animation
}

func (a *Animations) String() string {
	return "Animations"
}

func (a *Animations) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 70, Height: 25}
}

func (a *Animations) FPS() int {
	return 60
}

func (a *Animations) Init() (cmd tea.Cmd) {
	return nil
}

func (a *Animations) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		a.animationPng.Step(msg.Duration)
		a.animationGif.Step(msg.Duration)
	}
	return nil
}

func (a *Animations) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(image.Pt(0, 0), a.animationPng.Image()),
		engine.DrawImage(image.Pt(35, 0), a.animationGif.Image()),
	)
}
