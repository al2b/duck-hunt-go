package game

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/model/game/duck"
	"duck-hunt-go/model/game/gun"
	"duck-hunt-go/model/game/layout"
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

func (m *Game) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
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

func (m *Game) Bodies() (bodies engine.Bodies) {
	for _, model := range m.models {
		bodies = append(bodies, model.Bodies()...)
	}
	return bodies
}

func (m *Game) Sprites8() (sprites engine.Sprites8) {
	for _, model := range m.models {
		sprites = append(sprites, model.Sprites8()...)
	}
	return sprites
}

func (m *Game) Sprites24() (sprites engine.Sprites24) {
	for _, model := range m.models {
		sprites = append(sprites, model.Sprites24()...)
	}
	return sprites
}
