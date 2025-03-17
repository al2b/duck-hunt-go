package dog

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*.apng
var assets embed.FS

var (
	animation = engine.Must(engine.LoadAnimation(assets, "assets/dog.apng"))
)

func New() *Dog {
	return &Dog{
		animation: engine.AnimationPlayer{Animation: animation},
	}
}

type Dog struct {
	animation engine.AnimationPlayer
}

func (m *Dog) Init() tea.Cmd {
	return nil
}

func (m *Dog) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		m.animation.Step(msg.Duration)
	}
	return nil
}

func (m *Dog) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(image.Pt(2, 146), m.animation.Image()),
	)
}
