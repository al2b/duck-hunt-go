package primitives

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

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

func (s *Primitives) TPS() int {
	return 10
}

func (s *Primitives) Init() (cmd tea.Cmd) {
	return nil
}

func (s *Primitives) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Primitives) Draw(dst *engine.Image) {
	dst.Draw(
		// Dots
		engine.Dot{engine.Pt(5, 10), engine.ColorRed},
		engine.Dot{engine.Pt(10, 10), engine.ColorGreen},
		engine.Dot{engine.Pt(15, 10), engine.ColorBlue},

		// Segments
		engine.Segment{engine.Pt(5, 20), engine.Pt(20, 30), engine.ColorRed},
		engine.Segment{engine.Pt(5, 25), engine.Pt(25, 20), engine.ColorGreen},
		engine.Segment{engine.Pt(5, 30), engine.Pt(30, 30), engine.ColorBlue},

		// Rectangles
		engine.Rectangle{engine.Pt(40, 0), engine.Size{20, 20}, engine.ColorRed},
		engine.Rectangle{engine.Pt(57, 7), engine.Size{10, 20}, engine.ColorGreen},
		engine.Rectangle{engine.Pt(45, 5), engine.Size{10, 5}, engine.ColorBlue},

		// Circles
		engine.Circle{engine.Pt(40, 30), 5, engine.ColorRed},
		engine.Circle{engine.Pt(50, 35), 10, engine.ColorGreen},
		engine.Circle{engine.Pt(60, 35), 3, engine.ColorBlue},
	)
}
