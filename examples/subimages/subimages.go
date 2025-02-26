package subimages

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

func New() *Subimages {
	return &Subimages{}
}

type Subimages struct{}

func (s *Subimages) String() string {
	return "Subimages"
}

func (s *Subimages) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 80, Height: 50}
}

func (s *Subimages) FPS() int {
	return 10
}

func (s *Subimages) Init() (cmd tea.Cmd) {
	return nil
}

func (s *Subimages) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Subimages) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawRectangle(image.Pt(0, 0), engine.Size{Width: 80, Height: 50}, engine.ColorRed),
	).SubImage(
		image.Pt(7, 3),
		engine.Size{Width: 50, Height: 40},
	).Draw(
		engine.DrawRectangle(image.Pt(0, 0), engine.Size{Width: 50, Height: 40}, engine.ColorGreen),
	).SubImage(
		image.Pt(3, 7),
		engine.Size{Width: 30, Height: 30},
	).Draw(
		engine.DrawRectangle(image.Pt(0, 0), engine.Size{Width: 30, Height: 30}, engine.ColorBlue),
		engine.DrawCircle(image.Pt(15, 15), 10, engine.ColorWhite),
	)
}
