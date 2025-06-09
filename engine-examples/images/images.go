package images

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 10

var (
	//go:embed assets/*.png assets/*.gif
	assets embed.FS

	// Images
	imagePng  = engine.MustLoad(engine.ImageLoader{assets, "assets/kirby.png"})
	imageGif  = engine.MustLoad(engine.ImageLoader{assets, "assets/kirby.gif"})
	imageStar = engine.MustLoad(engine.ImageLoader{assets, "assets/star.png"})
)

func New() *Images {
	return &Images{}
}

type Images struct{}

func (s *Images) String() string {
	return "Images"
}

func (s *Images) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Images) Init() (cmd tea.Cmd) {
	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Images) Update(_ tea.Msg) (cmd tea.Cmd) {
	return nil
}

func (s *Images) Draw(dst *engine.Image) {
	dst.Draw(
		engine.ImageDrawer{engine.Pt(0, 0), imagePng},
		engine.ImageDrawer{engine.Pt(20, 0), imageGif},

		engine.ImageDrawer{engine.Pt(2, 20), imageStar},
		engine.Dot{engine.Pt(2, 20), engine.ColorRed},

		engine.ImageDrawer{
			engine.PointAdder{
				engine.Pt(38, 31),
				engine.Pt(-11, -11),
			},
			imageStar,
		},
		engine.Dot{engine.Pt(38, 31), engine.ColorRed},

		engine.ImageDrawer{
			engine.PointAdder{
				engine.Pt(68, 38),
				engine.Pt(-16, -18),
			},
			imageStar,
		},
		engine.Dot{engine.Pt(68, 38), engine.ColorRed},
	)
}
