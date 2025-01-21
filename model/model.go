package model

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/model/game"
	"duck-hunt-go/model/intro"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Model {
	return &Model{
		models: StateModels{
			StateIntro: intro.New(),
			StateGame:  game.New(),
		},
	}
}

type Model struct {
	models StateModels
	state  State
}

func (m *Model) Init() tea.Cmd {
	m.state = StateGame
	return m.models[m.state].Init()
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Switch state
		case "s":
			switch m.state {
			case StateIntro:
				m.state = StateGame
			case StateGame:
				m.state = StateIntro
			}
			return m.models[m.state].Init()
		}
	}

	return m.models[m.state].Update(msg)
}

func (m *Model) Bodies() engine.Bodies {
	return m.models[m.state].Bodies()
}

func (m *Model) Sprites() engine.Sprites {
	return m.models[m.state].Sprites()
}
