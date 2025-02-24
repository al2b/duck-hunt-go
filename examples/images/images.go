package images

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

//go:embed assets/*.png assets/*.gif
var assets embed.FS

func New() *Images {
	return &Images{
		imagePng: engine.MustLoadImage(engine.ImageFile(assets, "assets/kirby.png")),
		imageGif: engine.MustLoadImage(engine.ImageFile(assets, "assets/kirby.gif")),
	}
}

type Images struct {
	imagePng *engine.Image
	imageGif *engine.Image
}

func (i *Images) String() string {
	return "Images"
}

func (i *Images) Size(_ engine.Size) engine.Size {
	return engine.Size{Width: 40, Height: 18}
}

func (i *Images) FPS() int {
	return 10
}

func (i *Images) Init() (cmd tea.Cmd) {
	return nil
}

func (i *Images) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (i *Images) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(image.Pt(0, 0), i.imagePng),
		engine.DrawImage(image.Pt(20, 0), i.imageGif),
	)
}
