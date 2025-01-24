package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
)

const (
	width  = 256
	Ground = 184
)

func New() *Layout {
	return &Layout{}
}

type Layout struct {
	coordinates engine.Coordinates
}

func (m *Layout) Init() tea.Cmd {
	// Coordinates
	m.coordinates = engine.NewCoordinates(0, 0, 0)

	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Layout) Sprites() engine.Sprites {
	return engine.Sprites{
		engine.CoordinatedSprite{
			Coordinates: m.coordinates.SetZ(100),
			Image:       image,
		},
		engine.CoordinatedSprite{
			Coordinates: m.coordinates.SetZ(-math.MaxFloat64),
			Image:       imageSky,
		},
		engine.CoordinatedSprite{
			Coordinates: m.coordinates.Add(6, 32, 10),
			Image:       imageTree,
		},
		engine.CoordinatedSprite{
			Coordinates: m.coordinates.Add(193, 122, 20),
			Image:       imageShrub,
		},
	}
}

func (m *Layout) Bodies() engine.Bodies {
	return engine.Bodies{
		engine.NewBody(m.coordinates,
			engine.BodyShape{
				{0, 0},
				{width - 1, 0},
				{width - 1, Ground - 1},
				{0, Ground - 1},
			},
		),
		// Tree
		engine.NewBody(m.coordinates.Add(6, 32, 10),
			engine.BodyShape{
				{0, 0},
				{68, 0},
				{68, 150},
				{0, 150},
			},
		),
		// Shrub
		engine.NewBody(m.coordinates.Add(193, 122, 20),
			engine.BodyShape{
				{0, 60},
				{0, 29},
				{1, 22},
				{3, 16},
				{7, 11},
				{8, 7},
				{9, 4},
				{16, 0},
				{23, 2},
				{25, 4},
				{29, 15},
				{30, 25},
				{30, 60},
			},
		),
	}
}
