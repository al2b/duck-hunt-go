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
		models: []engine.Model{
			layout.New(),
			duck.New(),
			gun.New(),
		},
	}
}

type Game struct {
	models []engine.Model
}

func (m *Game) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, model := range m.models {
		cmds = append(cmds, model.Init())
	}
	return tea.Batch(cmds...)
}

func (m *Game) Update(msg engine.Msg) tea.Cmd {
	switch msg := msg.Msg().(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Restart
		case "r":
			return m.Init()
		}
	}

	var cmds []tea.Cmd
	for _, model := range m.models {
		cmds = append(cmds, model.Update(msg))
	}
	return tea.Batch(cmds...)
}

func (m *Game) Sprites() (sprites engine.Sprites) {
	for _, model := range m.models {
		sprites = append(sprites, model.Sprites()...)
	}
	return sprites
}

func (m *Game) Bodies() (bodies engine.Bodies) {
	for _, model := range m.models {
		bodies = append(bodies, model.Bodies()...)
	}
	return bodies
}
