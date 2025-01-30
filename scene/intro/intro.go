package intro

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

var textColor = color.RGBA{R: 0xff, G: 0xa0, B: 0x00}

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

func (m *Intro) Update(_ engine.Msg) tea.Cmd {
	return nil
}

func (m *Intro) Sprites() engine.Sprites {
	return engine.Sprites{
		engine.NewCoordinatedSprite(
			m.coordinates,
			imageLayout,
		),
		engine.NewCoordinatedSprite(
			m.coordinates.SetXY(64, 136),
			engine.NewText8x8("GAME A   1 DUCK", textColor).Image(),
		),
		engine.NewCoordinatedSprite(
			m.coordinates.SetXY(64, 152),
			engine.NewText8x8("GAME B   2 DUCKS", textColor).Image(),
		),
		engine.NewCoordinatedSprite(
			m.coordinates.SetXY(64, 168),
			engine.NewText8x8("GAME C   CLAY SHOOTING", textColor).Image(),
		),
	}
}

func (m *Intro) Bodies() engine.Bodies {
	return nil
}
