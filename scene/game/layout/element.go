package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Element struct {
	engine.Coordinates
	engine.StaticImage
	engine.PolygonShape
}

func (m *Element) Init() tea.Cmd            { return nil }
func (m *Element) Update(_ tea.Msg) tea.Cmd { return nil }
