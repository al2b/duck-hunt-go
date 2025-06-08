package text

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 60

func New() *Text {
	return &Text{}
}

type Text struct{}

func (s *Text) String() string {
	return "Text"
}

func (s *Text) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Text) Init() (cmd tea.Cmd) {
	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Text) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Text) Draw(dst *engine.Image) {
	dst.Draw(
		// 5x5
		engine.ImageDrawer{engine.Pt(0, 0), engine.Text{
			"The quick brown\nfox jumps over\nthe lazy dog.",
			engine.Font5x5, engine.ColorRed,
		}},

		// 6x6
		engine.ImageDrawer{engine.Pt(0, 16), engine.Text{
			"The quick\nbrown fox\njumps over\nthe lazy dog.",
			engine.Font6x6, engine.ColorGreen,
		}},
	)
}
