package round

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/engine/space"
	"duck-hunt-go/game/round/duck"
	"duck-hunt-go/game/round/gun"
	"duck-hunt-go/game/round/layout"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

func New() *Game {
	// Space
	space := space.NewSpace().
		SetGravity(engine.Vec(0, 9.8))

	return &Game{
		debug:       false,
		space:       space,
		layout:      layout.New(space),
		layoutTree:  layout.NewTree(space),
		layoutShrub: layout.NewShrub(space),
		duck:        duck.New(space),
		gun:         gun.New(space),
	}
}

type Game struct {
	debug       bool
	space       *space.Space
	layout      *layout.Layout
	layoutTree  *layout.Tree
	layoutShrub *layout.Shrub
	duck        *duck.Duck
	gun         *gun.Gun
}

func (m *Game) Init() tea.Cmd {
	// Init space
	m.space.Clear()

	return tea.Batch(
		m.space.Init(),
		m.layout.Init(),
		m.layoutTree.Init(),
		m.layoutShrub.Init(),
		m.duck.Init(),
		m.gun.Init(),
	)
}

func (m *Game) Update(msg tea.Msg) tea.Cmd {
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
			m.duck.Update(msg),
			m.gun.Update(msg),
		)

		// Update space
		m.space.Step(msg.Duration)

		return cmd
	}

	return tea.Batch(
		m.space.Update(msg),
		m.layout.Update(msg),
		m.layoutTree.Update(msg),
		m.layoutShrub.Update(msg),
		m.duck.Update(msg),
		m.gun.Update(msg),
	)
}

func (m *Game) Draw(scene *engine.Image) {
	scene.
		// Sky
		Fill(color.NRGBA{R: 143, G: 192, B: 255, A: 255}).
		Draw(
			m.layoutTree,
			m.layoutShrub,
			m.duck,
			// Layout
			m.layout,
			// Gun
			m.gun,
		)

	// Debug
	if m.debug {
		scene.Draw(m.space)
	}
}
