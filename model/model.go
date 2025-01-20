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

func (m *Model) Init() {
	m.state = StateGame
	m.models[m.state].Init()
}

func (m *Model) Update(msgs []tea.Msg) {
	// Messages
	for _, msg := range msgs {
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "s":
				switch m.state {
				case StateIntro:
					m.state = StateGame
				case StateGame:
					m.state = StateIntro
				}
				m.models[m.state].Init()
				return
			}
		}
	}

	m.models[m.state].Update(msgs)
}

func (m *Model) Bodies() (bodies engine.Bodies) {
	return m.models[m.state].Bodies()
}

func (m *Model) Sprites8() (sprites engine.Sprites8) {
	return m.models[m.state].Sprites8()
}

func (m *Model) Sprites24() (sprites engine.Sprites24) {
	return m.models[m.state].Sprites24()
}
