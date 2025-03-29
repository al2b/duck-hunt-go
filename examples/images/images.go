package images

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*.png assets/*.gif
var assets embed.FS

var (
	imagePng = engine.Must(engine.LoadImage(assets, "assets/kirby.png"))
	imageGif = engine.Must(engine.LoadImage(assets, "assets/kirby.gif"))
)

func New() *Images {
	return &Images{}
}

type Images struct {
	imagePng *engine.Image
	imageGif *engine.Image
}

func (s *Images) String() string {
	return "Images"
}

func (s *Images) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 40, Height: 18}
}

func (s *Images) FPS() int {
	return 10
}

func (s *Images) Init() (cmd tea.Cmd) {
	s.imagePng = imagePng
	s.imageGif = imageGif
	return nil
}

func (s *Images) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Images) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(image.Pt(0, 0), s.imagePng),
		engine.DrawImage(image.Pt(20, 0), s.imageGif),
	)
}
