package mouse

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	width  = 3
	height = 3
)

var coordinates = engine.NewCoordinates(0, 0, 1000)

func New() *Mouse {
	return &Mouse{}
}

type Mouse struct{}

func (m *Mouse) Init() tea.Cmd {
	return nil
}

func (m *Mouse) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		coordinates.
			SetX(float64(msg.X - (width / 2))).
			SetY(float64(msg.Y - (height / 2)))
	}

	return nil
}

func (m *Mouse) Bodies() engine.Bodies {
	return nil
}

func (m *Mouse) Sprites() (sprites engine.Sprites) {
	return sprites.Append(sprite)
}
