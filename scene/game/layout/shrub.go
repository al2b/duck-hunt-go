package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Shrub struct {
	engine.Coordinates
	engine.StaticImage
	engine.PolygonShape
}

func (m *Shrub) Init() tea.Cmd {
	// Init coordinates
	m.Coordinates = engine.NewCoordinates(193, 122, 20)

	// Init image
	m.StaticImage = engine.NewStaticImage(imageShrub)

	// Init shape
	m.PolygonShape = engine.NewPolygonShape(
		0, 60,
		0, 29,
		1, 22,
		3, 16,
		7, 11,
		8, 7,
		9, 4,
		16, 0,
		23, 2,
		25, 4,
		29, 15,
		30, 25,
		30, 60,
	)

	return nil
}

func (m *Shrub) Update(_ tea.Msg) tea.Cmd {
	return nil
}
