package intro

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Intro {
	return &Intro{}
}

type Intro struct{}

func (m *Intro) Init() {}

func (m *Intro) Update(_ []tea.Msg) {}

func (m *Intro) Bodies() engine.Bodies {
	return nil
}

func (m *Intro) Sprites8() (sprites engine.Sprites8) {
	return sprites.Append(
		sprite8,
	)
}

func (m *Intro) Sprites24() (sprites engine.Sprites24) {
	return sprites.Append(
		sprite24,
	)
}
