package primitives

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

func New() *Scene {
	return &Scene{}
}

type Scene struct{}

func (s *Scene) String() string {
	return "Primitives"
}

func (s *Scene) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 80, Height: 50}
}

func (s *Scene) FPS() int {
	return 10
}

func (s *Scene) Init() (cmd tea.Cmd) {
	return nil
}

func (s *Scene) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Scene) Draw(scene *engine.Image) {
	scene.Draw(
		// Dots
		engine.DrawDot(image.Pt(5, 10), engine.ColorRed),
		engine.DrawDot(image.Pt(10, 10), engine.ColorGreen),
		engine.DrawDot(image.Pt(15, 10), engine.ColorBlue),
		// Segments
		engine.DrawSegment(image.Pt(5, 20), image.Pt(20, 30), engine.ColorRed),
		engine.DrawSegment(image.Pt(5, 25), image.Pt(25, 20), engine.ColorGreen),
		engine.DrawSegment(image.Pt(5, 30), image.Pt(30, 30), engine.ColorBlue),
		// Circles
		engine.DrawCircle(image.Pt(40, 30), 5, engine.ColorRed),
		engine.DrawCircle(image.Pt(50, 35), 10, engine.ColorGreen),
		engine.DrawCircle(image.Pt(60, 35), 3, engine.ColorBlue),
	)
}
