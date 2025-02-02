package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
)

func New() *Layout {
	return &Layout{
		tree:  &Tree,
		shrub: &Shrub,
	}
}

type Layout struct {
	engine.Coordinates
	engine.StaticImage
	engine.RectangleShape
	tree  *Element
	shrub *Element
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

	return tea.Batch(
		m.tree.Init(),
		m.shrub.Init(),
	)
}

func (m *Layout) Update(msg tea.Msg) tea.Cmd {
	return tea.Batch(
		m.tree.Update(msg),
		m.shrub.Update(msg),
	)
}

func (m *Layout) Sprites() (sprites engine.Sprites) {
	return sprites.Append(
		m, m.tree, m.shrub,
		// Sky
		engine.NewCoordinatedSprite(
			m.Coordinates.SetZ(-math.MaxFloat64),
			imageSky,
		),
	)
}

func (m *Layout) Bodies() (bodies engine.Bodies) {
	return bodies.Append(
		m, m.tree, m.shrub,
	)
}
