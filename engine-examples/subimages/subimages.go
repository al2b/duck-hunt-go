package subimages

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 10

func New() *Subimages {
	return &Subimages{}
}

type Subimages struct{}

func (s *Subimages) String() string {
	return "Subimages"
}

func (s *Subimages) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Subimages) Init() (cmd tea.Cmd) {
	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Subimages) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Subimages) Draw(dst *engine.Image) {
	dst.
		Draw(
			engine.Rectangle{engine.Point{0, 0}, engine.Size{80, 50}, engine.ColorRed},
		).
		SubImage(
			engine.Point{7, 3},
			engine.Size{50, 40},
		).
		Draw(
			engine.Rectangle{engine.Point{0, 0}, engine.Size{50, 40}, engine.ColorGreen},
		).
		SubImage(
			engine.Point{3, 7},
			engine.Size{30, 30},
		).
		Draw(
			engine.Rectangle{engine.Point{0, 0}, engine.Size{30, 30}, engine.ColorBlue},
			engine.Circle{engine.Point{15, 15}, 10, engine.ColorWhite},
		)
}
