package mouse

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
)

//go:embed assets/*
var assets embed.FS

func New() *Mouse {
	return &Mouse{
		StaticImage: engine.NewStaticImage(
			engine.MustLoadImage(engine.ImagePngFile(assets, "assets/mouse.png")),
		),
	}
}

type Mouse struct {
	engine.AbsolutePosition
	engine.StaticImage
}

func (m *Mouse) Init() tea.Cmd {
	// Position
	m.SetPosition(engine.Vec(0, 0))

	return nil
}

func (m *Mouse) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		// Position
		m.SetPosition(engine.Vec(
			float64(msg.X), float64(msg.Y),
		))
	}
	return nil
}

func (m *Mouse) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawCenteredImage(m.Position().Point(), m.Image()),
	)
}
