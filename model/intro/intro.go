package intro

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

var coordinates = engine.NewCoordinates(0, 0, 0)

func New() *Intro {
	return &Intro{}
}

type Intro struct{}

func (m *Intro) Init() tea.Cmd {
	return nil
}

func (m *Intro) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Intro) Bodies() engine.Bodies {
	return nil
}

func (m *Intro) Sprites() engine.Sprites {
	return sprites
}
