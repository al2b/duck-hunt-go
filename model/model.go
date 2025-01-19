package model

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/model/game"
	"duck-hunt-go/model/intro"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Model {
	return &Model{
		intro: intro.New(),
		game:  game.New(),
	}
}

type Model struct {
	model engine.Model
	intro engine.Model
	game  engine.Model
}

func (m *Model) Init() {
	m.intro.Init()
	m.game.Init()

	m.model = m.game
}

func (m *Model) Update(msgs []tea.Msg) {
	m.model.Update(msgs)
}

func (m *Model) Bodies() (bodies engine.Bodies) {
	return m.model.Bodies()
}

func (m *Model) Sprites8() (sprites engine.Sprites8) {
	return m.model.Sprites8()
}

func (m *Model) Sprites24() (sprites engine.Sprites24) {
	return m.model.Sprites24()
}
