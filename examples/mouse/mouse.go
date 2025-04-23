package mouse

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
)

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageMouse = engine.Must(engine.LoadImage(assets, "assets/mouse.png"))
)

func New() *Mouse {
	return &Mouse{}
}

type Mouse struct {
	point engine.Point
}

func (s *Mouse) String() string {
	return "Mouse"
}

func (s *Mouse) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Mouse) TPS() int {
	return 60
}

func (s *Mouse) Init() (cmd tea.Cmd) {
	s.point = engine.Pt(0, 0)

	return nil
}

func (s *Mouse) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMotionMsg:
		s.point = engine.Pt(msg.X, msg.Y-6)
	}

	return nil
}

func (s *Mouse) Draw(dst *engine.Image) {
	dst.Draw(
		engine.ImageDrawer{
			s.point,
			imageMouse,
		},
	)
}
