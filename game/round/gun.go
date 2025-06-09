package round

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

func NewGun() *Gun {
	return &Gun{}
}

type Gun struct {
	path engine.Path2DPlayer
	engine.ImageDrawer
}

func (m *Gun) Init() tea.Cmd {
	// Path
	m.path.Path = engine.StaticPath2d{Position: engine.Vector2D{0, 0}}
	m.path.OnEnd = engine.PlayerOnEndPause

	// Drawer
	m.ImageDrawer.Pointer = engine.PointAdder{
		engine.Position2DPointer{&m.path},
		engine.Point{-6, -6},
	}
	m.ImageDrawer.Imager = assets.GunNormal

	return nil
}

func (m *Gun) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Step path
		m.path.Step(msg.Interval)
	case tea.MouseMsg:
		mouse := msg.Mouse()
		// Update path
		m.path.Path = engine.ElasticPath2D{
			m.path.Position(),
			engine.Vector2D{float64(mouse.X), float64(mouse.Y)},
			time.Second * 1,
			1, 0.25,
		}
		m.path.Rewind()
		m.path.Play()
		switch msg := msg.(type) {
		case tea.MouseClickMsg:
			switch msg.Button {
			case tea.MouseLeft:
				return GunShot(m.path.Position())
			}
		}
	}

	return nil
}

/************/
/* Messages */
/************/

func GunShot(position engine.Vector2D) tea.Cmd {
	return func() tea.Msg {
		return GunShotMsg(position)
	}
}

type GunShotMsg engine.Vector2D

func (msg GunShotMsg) Position() engine.Vector2D {
	return engine.Vector2D(msg)
}
