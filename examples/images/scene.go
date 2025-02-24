package images

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*
var assets embed.FS

func New() *Scene {
	return &Scene{
		imagePng: engine.MustLoadImage(engine.ImageFile(assets, "assets/kirby.png")),
		imageGif: engine.MustLoadImage(engine.ImageFile(assets, "assets/kirby.gif")),
	}
}

type Scene struct {
	imagePng *engine.Image
	imageGif *engine.Image
}

func (s *Scene) String() string {
	return "Images"
}

func (s *Scene) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 40, Height: 18}
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
		engine.DrawImage(image.Pt(0, 0), s.imagePng),
		engine.DrawImage(image.Pt(20, 0), s.imageGif),
	)
}
