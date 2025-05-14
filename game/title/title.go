package title

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

var (
	textColor = color.RGBA{R: 0xff, G: 0xa0, B: 0x00}
)

func New() *Title {
	return &Title{}
}

type Title struct {
	cursor int
}

func (m *Title) Init() tea.Cmd {
	m.cursor = 0

	return nil
}

func (m *Title) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyUp:
			m.cursor = (m.cursor - 1 + 3) % 3
		case tea.KeyDown:
			m.cursor = (m.cursor + 1) % 3
		}
	}

	return nil
}

func (m *Title) Draw(dst *engine.Image) {
	dst.Draw(
		// Layout
		engine.ImageDrawer{engine.Pt(0, 0), assets.TitleLayout},

		// Menu
		engine.TextDrawer{engine.Pt(64, 136),
			engine.Text{"GAME A   1 DUCK", assets.Font, textColor},
		},
		engine.TextDrawer{engine.Pt(64, 152),
			engine.Text{"GAME B   2 DUCKS", assets.Font, textColor},
		},
		engine.TextDrawer{engine.Pt(64, 168),
			engine.Text{"GAME C   CLAY SHOOTING", assets.Font, textColor},
		},

		// Cursor
		engine.ImageDrawer{engine.Pt(48, 136+(m.cursor*16)), assets.TitleCursor},
	)
}
