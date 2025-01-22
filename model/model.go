package model

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/model/game"
	"duck-hunt-go/model/intro"
	"duck-hunt-go/model/mouse"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Model {
	return &Model{
		mouse: mouse.New(),
		models: StateModels{
			StateIntro: intro.New(),
			StateGame:  game.New(),
		},
	}
}

type Model struct {
	mouse  engine.Model
	models StateModels
	state  State
}

func (m *Model) Init() tea.Cmd {
	m.state = StateGame

	return tea.Batch(
		m.mouse.Init(),
		m.models[m.state].Init(),
	)
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

	return tea.Batch(
		m.mouse.Update(msg),
		m.models[m.state].Update(msg),
	)
}

func (m *Model) Bodies() (bodies engine.Bodies) {
	return bodies.Appends(
		m.mouse.Bodies(),
		m.models[m.state].Bodies(),
	)
}

func (m *Model) Sprites() (sprites engine.Sprites) {
	return sprites.Appends(
		m.mouse.Sprites(),
		m.models[m.state].Sprites(),
	)
}
