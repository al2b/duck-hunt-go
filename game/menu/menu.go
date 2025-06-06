package menu

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/state"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

func New() *Menu {
	return &Menu{
		choices: []choice{
			{engine.Text{"GAME A   1 DUCK", assets.Font}, state.Mode1Duck},
			{engine.Text{"GAME B   2 DUCKS", assets.Font}, state.Mode2Ducks},
		},
	}
}

type Menu struct {
	choices []choice
	choice  int
}

func (m *Menu) Init() tea.Cmd {
	m.choice = 0

	return nil
}

func (m *Menu) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyUp:
			m.choice = (m.choice - 1 + len(m.choices)) % len(m.choices)
		case tea.KeyDown:
			m.choice = (m.choice + 1) % len(m.choices)
		case tea.KeyEnter:
			return state.SetMode(m.choices[m.choice].Mode)
		}
	}

	return nil
}

func (m *Menu) Draw(dst *engine.Image) {
	// Layout
	dst.Draw(
		engine.ImageDrawer{engine.Pt(0, 0), assets.MenuLayout},
	)

	choiceHeight := assets.Font.Size().Height * 2

	// Choices
	for i, choice := range m.choices {
		dst.Draw(
			engine.TextDrawer{engine.Pt(64, 136+(i*choiceHeight)),
				choice.Text,
				color.RGBA{R: 0xff, G: 0xa0, B: 0x00},
			},
		)
	}

	// Cursor
	dst.Draw(
		engine.ImageDrawer{engine.Pt(48, 136+(m.choice*choiceHeight)), assets.MenuCursor},
	)
}

type choice struct {
	Text engine.Text
	Mode state.Mode
}
