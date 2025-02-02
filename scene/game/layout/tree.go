package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Tree struct {
	engine.Coordinates
	engine.StaticImage
	engine.RectangleShape
}

func (m *Tree) Init() tea.Cmd {
	// Init coordinates
	m.Coordinates = engine.NewCoordinates(6, 32, 10)

	// Init image
	m.StaticImage = engine.NewStaticImage(imageTree)

	// Init shape
	m.RectangleShape = engine.NewRectangleShape(
		0, 0,
		68, 150,
	)

	return nil
}

func (m *Tree) Update(_ tea.Msg) tea.Cmd {
	return nil
}
