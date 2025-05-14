package round

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"duck-hunt-go/game/font"
	"duck-hunt-go/game/round/dog"
	"duck-hunt-go/game/round/duck"
	"duck-hunt-go/game/round/gun"
	"duck-hunt-go/game/round/layout"
	"duck-hunt-go/game/state"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

func New() *Round {
	m := &Round{
		debug: false,
		space: space.NewSpace().
			SetGravity(engine.Vec2D(0, 9.8)),
	}

	m.layout = layout.New(m.space)
	m.layoutTree = layout.NewTree(m.space)
	m.layoutShrub = layout.NewShrub(m.space)
	m.dog = dog.New()
	m.duck = duck.New(m.space)
	m.gun = gun.New(m.space)

	return m
}

type Round struct {
	debug       bool
	space       *space.Space
	layout      *layout.Layout
	layoutTree  *layout.Tree
	layoutShrub *layout.Shrub
	dog         *dog.Dog
	duck        *duck.Duck
	gun         *gun.Gun
}

func (m *Round) Init() tea.Cmd {
	// Init space
	m.space.Clear()

	return tea.Batch(
		m.space.Init(),
		m.layout.Init(),
		m.layoutTree.Init(),
		m.layoutShrub.Init(),
		m.dog.Init(),
		m.duck.Init(),
		m.gun.Init(),
	)
}

func (m *Round) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Debug
		case "d":
			m.debug = !m.debug
			return engine.ConsoleLog("Debug: %t", m.debug)
		}
	case engine.TickMsg:
		cmd := tea.Batch(
			m.space.Update(msg),
			m.layout.Update(msg),
			m.layoutTree.Update(msg),
			m.layoutShrub.Update(msg),
			m.dog.Update(msg),
			m.duck.Update(msg),
			m.gun.Update(msg),
		)

		// Update space
		m.space.Step(msg.Interval)

		return cmd
	}

	return tea.Batch(
		m.space.Update(msg),
		m.layout.Update(msg),
		m.layoutTree.Update(msg),
		m.layoutShrub.Update(msg),
		m.dog.Update(msg),
		m.duck.Update(msg),
		m.gun.Update(msg),
	)
}

func (m *Round) Draw(dst *engine.Image) {
	dst.
		// Sky
		Fill(color.NRGBA{R: 63, G: 191, B: 255, A: 255}).
		Draw(
			m.layoutTree,
			m.layoutShrub,
			m.duck,

			// Ordered drawers
			engine.OrderDrawers{
				// Layout
				m.layout,
				// Dog
				m.dog,
			},

			// Texts
			engine.TextDrawer{engine.Pt(24, 192),
				engine.Text{fmt.Sprintf("R=%d", state.Round), font.Font, color.RGBA{R: 131, G: 211, B: 19, A: 255}},
			},
			engine.TextDrawer{engine.Pt(64, 208),
				engine.Text{"HIT", font.Font, color.RGBA{R: 131, G: 211, B: 19, A: 255}},
			},
			engine.TextDrawer{engine.Pt(192, 208),
				engine.Text{fmt.Sprintf("%06d", state.Score), font.Font, engine.ColorWhite},
			},
			engine.TextDrawer{engine.Pt(200, 216),
				engine.Text{"SCORE", font.Font, engine.ColorWhite},
			},

			// Gun
			m.gun,
		)

	// Debug
	if m.debug {
		dst.Draw(m.space)
	}
}
