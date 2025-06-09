package primitives

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 10

func New() *Primitives {
	return &Primitives{}
}

type Primitives struct{}

func (s *Primitives) String() string {
	return "Primitives"
}

func (s *Primitives) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Primitives) Init() (cmd tea.Cmd) {
	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Primitives) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Primitives) Draw(dst *engine.Image) {
	dst.Draw(
		// Dots
		engine.Dot{engine.Point{5, 10}, engine.ColorRed},
		engine.Dot{engine.Point{10, 10}, engine.ColorGreen},
		engine.Dot{engine.Point{15, 10}, engine.ColorBlue},

		// Segments
		engine.Segment{engine.Point{5, 20}, engine.Point{20, 30}, engine.ColorRed},
		engine.Segment{engine.Point{5, 25}, engine.Point{25, 20}, engine.ColorGreen},
		engine.Segment{engine.Point{5, 30}, engine.Point{30, 30}, engine.ColorBlue},

		// Rectangles
		engine.Rectangle{engine.Point{40, 0}, engine.Size{20, 20}, engine.ColorRed},
		engine.Rectangle{engine.Point{57, 7}, engine.Size{10, 20}, engine.ColorGreen},
		engine.Rectangle{engine.Point{45, 5}, engine.Size{10, 5}, engine.ColorBlue},

		// Circles
		engine.Circle{engine.Point{40, 30}, 5, engine.ColorRed},
		engine.Circle{engine.Point{50, 35}, 10, engine.ColorGreen},
		engine.Circle{engine.Point{60, 35}, 3, engine.ColorBlue},
	)
}
