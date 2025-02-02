package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
)

type Layout struct {
	engine.Coordinates
	engine.StaticImage
	engine.RectangleShape
}

func (m *Layout) Init() tea.Cmd {
	// Init coordinates
	m.Coordinates = engine.NewCoordinates(0, 0, 0)

	// Init image
	m.StaticImage = engine.NewStaticImage(imageLayout)

	// Init shape
	m.RectangleShape = engine.NewRectangleShape(
		0, 0,
		width-1, Ground-1,
	)

	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Layout) Sprites() engine.Sprites {
	return engine.Sprites{
		m,
		// Sky
		engine.NewCoordinatedSprite(
			m.Coordinates.SetZ(-math.MaxFloat64),
			imageSky,
		),
		// Tree
		engine.NewCoordinatedSprite(
			m.Coordinates.Add(6, 32, 10),
			imageTree,
		),
		// Shrub
		engine.NewCoordinatedSprite(
			m.Coordinates.Add(193, 122, 20),
			imageShrub,
		),
	}
}

func (m *Layout) Bodies() (bodies engine.Bodies) {
	return bodies.Append(
		m,
		// Tree
		engine.NewRectangleBody(
			m.Coordinates,
			engine.NewRectangleShape(
				0+6, 0+32,
				68+6, 150+32,
			),
		),
		// Shrub
		engine.NewPolygonBody(
			m.Coordinates,
			engine.NewPolygonShape(
				0+193, 60,
				0+193, 29+122,
				1+193, 22+122,
				3+193, 16+122,
				7+193, 11+122,
				8+193, 7+122,
				9+193, 4+122,
				16+193, 0+122,
				23+193, 2+122,
				25+193, 4+122,
				29+193, 15+122,
				30+193, 25+122,
				30+193, 60+122,
			),
		),
	)
}
