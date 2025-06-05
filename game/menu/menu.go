package menu

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/state"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

var (
	textColor = color.RGBA{R: 0xff, G: 0xa0, B: 0x00}
)

func New() *Menu {
	return &Menu{}
}

type Menu struct {
	cursor int
}

func (m *Menu) Init() tea.Cmd {
	m.cursor = 0

	return nil
}

func (m *Menu) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyUp:
			m.cursor = (m.cursor - 1 + 2) % 2
		case tea.KeyDown:
			m.cursor = (m.cursor + 1) % 2
		case tea.KeyEnter:
			switch m.cursor {
			case 0:
				return state.SetMode(state.Mode1Duck)
			case 1:
				return state.SetMode(state.Mode2Ducks)
			}
		}
	}

	return nil
}

func (m *Menu) Draw(dst *engine.Image) {
	dst.Draw(
		// Layout
		engine.ImageDrawer{engine.Pt(0, 0), assets.MenuLayout},

		// Choices
		engine.TextDrawer{engine.Pt(64, 136),
			engine.Text{"GAME A   1 DUCK", assets.Font, textColor},
		},
		engine.TextDrawer{engine.Pt(64, 152),
			engine.Text{"GAME B   2 DUCKS", assets.Font, textColor},
		},

		// Cursor
		engine.ImageDrawer{engine.Pt(48, 136+(m.cursor*16)), assets.MenuCursor},
	)
}
