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
	animationKirby     = engine.NewAnimationSequence(
		animationKirbyWalk, animationKirbyWalk,
		animationKirbyRun, animationKirbyRun, animationKirbyRun,
		animationKirbyLand, animationKirbyLand,
	)
)

func New() *Animations {
	return &Animations{
		animationPng:   engine.NewAnimationPlayer(animationPng),
		animationGif:   engine.NewAnimationPlayer(animationGif),
		animationKirby: engine.NewAnimationPlayer(animationKirby),
	}
}

type Animations struct {
	animationPng   *engine.AnimationPlayer
	animationGif   *engine.AnimationPlayer
	animationKirby *engine.AnimationPlayer
}

func (a *Animations) String() string {
	return "Animations"
}

func (a *Animations) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 70, Height: 44}
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
		a.animationKirby.Step(msg.Duration)
	}
	return nil
}

func (a *Animations) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(image.Pt(0, 0), a.animationPng.Image()),
		engine.DrawImage(image.Pt(35, 0), a.animationGif.Image()),
		engine.DrawImage(image.Pt(0, 25), a.animationKirby.Image()),
	)
}
