package mouse

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*.png
var assets embed.FS

func New() *Mouse {
	return &Mouse{
		image: engine.Must(engine.LoadImage(assets, "assets/mouse.png")),
	}
}

type Mouse struct {
	position image.Point
	image    *engine.Image
}

func (m *Mouse) Init() tea.Cmd {
	m.position = image.Pt(0, 0)
	return nil
}

func (m *Mouse) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		m.position = image.Pt(msg.X, msg.Y)
	}
	return nil
}

func (m *Mouse) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawCenteredImage(m.position, m.image),
	)
}
