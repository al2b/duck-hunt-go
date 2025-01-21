package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
)

const (
	width  = engine.Width
	Ground = 184
)

var (
	coordinates      = engine.NewCoordinates(0, 0, 0)
	skyCoordinates   = engine.NewRelativeCoordinates(coordinates, 0, 0, -math.MaxFloat64)
	treeCoordinates  = engine.NewRelativeCoordinates(coordinates, 6, 32, 10)
	shrubCoordinates = engine.NewRelativeCoordinates(coordinates, 193, 122, 20)
)

func New() *Layout {
	return &Layout{}
}

type Layout struct{}

func (m *Layout) Init() tea.Cmd {
	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Layout) Bodies() engine.Bodies {
	return bodies
}

func (m *Layout) Sprites() engine.Sprites {
	return sprites
}
