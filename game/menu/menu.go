package menu

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/state"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

func New() *Menu {
	return &Menu{
		choices: []choice{
			{engine.Point{64, 136},
				engine.Text{"GAME A   1 DUCK", assets.Font, color.RGBA{R: 255, G: 160, B: 0, A: 255}}.Image(),
				state.Mode1Duck,
			},
			{engine.Point{64, 152},
				engine.Text{"GAME B   2 DUCKS", assets.Font, color.RGBA{R: 255, G: 160, B: 0, A: 255}}.Image(),
				state.Mode2Ducks,
			},
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
	case tea.MouseMsg:
		mouse := msg.Mouse()
		for i, choice := range m.choices {
			choiceBounds := choice.Image.Bounds()
			choiceMax := choice.Point.Add(
				engine.Point{choiceBounds.Dx(), choiceBounds.Dy()},
			)
			if choice.Point.X <= mouse.X && mouse.X <= choiceMax.X && choice.Point.Y <= mouse.Y && mouse.Y <= choiceMax.Y {
				m.choice = i
				switch msg := msg.(type) {
				case tea.MouseClickMsg:
					switch msg.Button {
					case tea.MouseLeft:
						return state.SetMode(m.choices[m.choice].Mode)
					}
				}
			}
		}
	}

	return nil
}

func (m *Menu) Draw(dst *engine.Image) {
	dst.Draw(
		// Layout
		engine.ImageDrawer{engine.Point{0, 0}, assets.MenuLayout},
		// Top Score
		engine.ImageDrawer{engine.Point{56, 192}, engine.Text{
			fmt.Sprintf("TOP SCORE = %d", state.TopScore),
			assets.Font, color.RGBA{R: 112, G: 240, B: 64, A: 255},
		}},
		// Footer
		engine.ImageDrawer{engine.Point{40, 208}, engine.Text{
			"©1984 NINTENDO CO;LTD.",
			assets.Font, engine.ColorWhite,
		}},
	)

	// Choices
	for _, choice := range m.choices {
		dst.Draw(
			engine.ImageDrawer{choice.Point, choice.Image},
		)
	}

	// Cursor
	dst.Draw(
		engine.ImageDrawer{engine.Point{m.choices[m.choice].Point.X - 16, m.choices[m.choice].Point.Y}, assets.MenuCursor},
	)
}

type choice struct {
	Point engine.Point
	Image *engine.Image
	Mode  state.Mode
}
