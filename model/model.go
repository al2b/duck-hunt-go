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
	intro engine.Model
	game  engine.Model
}

func (m *Model) Init() {
	m.game.Init()
}

func (m *Model) Update(msgs []tea.Msg) {
	m.game.Update(msgs)
}

func (m *Model) Bodies() (bodies engine.Bodies) {
	return m.game.Bodies()
}

func (m *Model) Sprites8() (sprites engine.Sprites8) {
	return m.game.Sprites8()
}

func (m *Model) Sprites24() (sprites engine.Sprites24) {
	return m.game.Sprites24()
}
