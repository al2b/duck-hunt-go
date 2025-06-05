package round

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	"duck-hunt-go/game/round/dog"
	"duck-hunt-go/game/round/duck"
	"duck-hunt-go/game/round/gun"
	"duck-hunt-go/game/round/layout"
	"duck-hunt-go/game/state"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
	"image/color"
)

func New(mode state.Mode, round int) *Round {
	// Model
	m := &Round{
		round: round,
	}

	// Space
	m.space = cp.NewSpace()

	// Layout
	m.layout = layout.New(m.space)
	m.layoutTree = layout.NewTree(m.space)
	m.layoutShrub = layout.NewShrub(m.space)

	// Dog
	m.dog = dog.New()

	// Duck(s)
	switch mode {
	case state.Mode1Duck:
		m.ducks = duck.NewDucks(m.space, 1)
	case state.Mode2Ducks:
		m.ducks = duck.NewDucks(m.space, 2)
	}

	// Gun
	m.gun = gun.New()

	return m
}

type Round struct {
	round       int
	ammos       int
	space       *cp.Space
	layout      *layout.Layout
	layoutTree  *layout.Tree
	layoutShrub *layout.Shrub
	dog         *dog.Dog
	ducks       duck.Ducks
	gun         *gun.Gun
}

func (m *Round) Init() tea.Cmd {
	// Ammos
	m.ammos = 3

	return tea.Batch(
		m.layout.Init(),
		m.layoutTree.Init(),
		m.layoutShrub.Init(),
		m.dog.Init(),
		m.ducks.Init(),
		m.gun.Init(),
	)
}

func (m *Round) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Update models
		cmd := tea.Batch(
			m.layout.Update(msg),
			m.layoutTree.Update(msg),
			m.layoutShrub.Update(msg),
			m.dog.Update(msg),
			m.ducks.Update(msg),
			m.gun.Update(msg),
		)
		// Step space
		m.space.Step(msg.Interval.Seconds())
		return cmd
	case tea.KeyPressMsg:
		switch msg.String() {
		// Debug
		case "d":
			config.Debug = !config.Debug
			return engine.ConsoleLog("Debug: %t", config.Debug)
		}
	case gun.ShotMsg:
		m.ammos = max(m.ammos-1, 0)
		nearest := m.space.PointQueryNearest(
			cp.Vector{msg.X, msg.Y},
			10,
			cp.NewShapeFilter(cp.NO_GROUP, cp.ALL_CATEGORIES, config.ShapeCategoryDuck),
		)
		if nearest.Shape != nil {
			return m.ducks.Update(
				duck.DiscriminatedShotMsg{
					Discriminator: nearest.Shape.Body().UserData,
				},
			)
		}
	}

	return tea.Batch(
		m.layout.Update(msg),
		m.layoutTree.Update(msg),
		m.layoutShrub.Update(msg),
		m.dog.Update(msg),
		m.ducks.Update(msg),
		m.gun.Update(msg),
	)
}

func (m *Round) Draw(dst *engine.Image) {
	dst.
		Fill(color.NRGBA{R: 63, G: 191, B: 255, A: 255}).
		Draw(
			m.layoutTree,
			m.layoutShrub,
			m.ducks,
			engine.OrderDrawers{
				m.layout,
				m.dog,
			},
			engine.TextDrawer{engine.Pt(24, 192),
				engine.Text{fmt.Sprintf("R=%d", m.round), assets.Font, color.RGBA{R: 131, G: 211, B: 19, A: 255}},
			},
			engine.TextDrawer{engine.Pt(64, 208),
				engine.Text{"HIT", assets.Font, color.RGBA{R: 131, G: 211, B: 19, A: 255}},
			},
			engine.TextDrawer{engine.Pt(192, 208),
				engine.Text{fmt.Sprintf("%06d", config.Score), assets.Font, engine.ColorWhite},
			},
			engine.TextDrawer{engine.Pt(200, 216),
				engine.Text{"SCORE", assets.Font, engine.ColorWhite},
			},
		)

	// Ammos
	for i := 0; i < m.ammos; i++ {
		dst.Draw(engine.ImageDrawer{
			engine.Pt(26+(i*8), 208),
			assets.LayoutAmmo,
		})
	}

	// Gun
	dst.Draw(m.gun)

	// Debug
	if config.Debug {
		dst.Draw(enginecp.SpaceDrawer{m.space})
	}
}
