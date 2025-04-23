package title

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/font"
	tea "github.com/charmbracelet/bubbletea/v2"
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
		engine.ImageDrawer{engine.Pt(0, 0), imageLayout},

		// Menu
		engine.TextDrawer{engine.Pt(64, 136),
			engine.Text{"GAME A   1 DUCK", font.Font, colorText},
		},
		engine.TextDrawer{engine.Pt(64, 152),
			engine.Text{"GAME B   2 DUCKS", font.Font, colorText},
		},
		engine.TextDrawer{engine.Pt(64, 168),
			engine.Text{"GAME C   CLAY SHOOTING", font.Font, colorText},
		},

		// Cursor
		engine.ImageDrawer{engine.Pt(48, 136+(m.cursor*16)), imageCursor},
	)
}
