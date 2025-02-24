package gun

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

//go:embed assets/*.png
var assets embed.FS

func New(space *space.Space) *Gun {
	return &Gun{
		space: space,
		Path:  engine.NewPath(),
		image: engine.MustLoadImage(engine.ImageFile(assets, "assets/gun.png")),
	}
}

type Gun struct {
	space *space.Space
	*engine.Path
	image *engine.Image
}

func (m *Gun) Init() tea.Cmd {
	// Init path
	m.Move(0, 0)

	// Init space body
	m.space.AddNewPositionableBody(m).
		AddNewPolygon(engine.Vectors{
			{-5.5, -18.5},
			{4.5, -18.5},
			{17.5, -5.5},
			{17.5, 4.5},
			{4.5, 17.5},
			{-5.5, 17.5},
			{-18.5, 4.5},
			{-18.5, -5.5},
		}, 0).
		SetElasticity(1).
		SetFriction(0)

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		// Update path
		m.To(engine.Vec(
			float64(msg.X), float64(msg.Y),
		), engine.ElasticEasing(1, 0.25), time.Second*1)
	case engine.TickMsg:
		// Step path
		m.Step(msg.Duration)
	}

	return nil
}

func (m *Gun) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawCenteredImage(m.Position().Point(), m.image),
	)
}
