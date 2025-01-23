package intro

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Intro {
	return &Intro{}
}

type Intro struct {
	coordinates engine.Coordinates
}

func (m *Intro) Init() tea.Cmd {
	// Coordinates
	m.coordinates = engine.NewCoordinates(0, 0, 1000)

	return nil
}

func (m *Intro) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Intro) Sprites() engine.Sprites {
	return engine.Sprites{
		engine.CoordinatedSprite{
			Coordinates: m.coordinates,
			Image:       image,
		},
	}
}

func (m *Intro) Bodies() engine.Bodies {
	return nil
}
