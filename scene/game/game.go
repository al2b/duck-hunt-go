package game

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/scene/game/duck"
	"duck-hunt-go/scene/game/gun"
	"duck-hunt-go/scene/game/layout"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Game {
	return &Game{
		layout: &layout.Layout{},
		duck:   &duck.Duck{},
		gun:    &gun.Gun{},
	}
}

type Game struct {
	layout *layout.Layout
	duck   *duck.Duck
	gun    *gun.Gun
}

func (m *Game) Init() tea.Cmd {
	return tea.Batch(
		m.layout.Init(),
		m.duck.Init(),
		m.gun.Init(),
	)
}

func (m *Game) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Restart
		case "r":
			return m.Init()
		}
	}

	return tea.Batch(
		m.layout.Update(msg),
		m.duck.Update(msg),
		m.gun.Update(msg),
	)
}

func (m *Game) Sprites() (sprites engine.Sprites) {
	return sprites.
		Appends(m.layout.Sprites()).
		Append(m.duck, m.gun)
}

func (m *Game) Bodies() (bodies engine.Bodies) {
	return bodies.
		Appends(m.layout.Bodies()).
		Append(m.duck, m.gun)
}
