package intro

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image/color"
)

//go:embed assets/*
var assets embed.FS

var (
	textColor = color.RGBA{R: 0xff, G: 0xa0, B: 0x00}
)

func New() *Intro {
	return &Intro{
		AbsolutePosition: engine.NewAbsolutePosition(0, 0),
		StaticImage: engine.NewStaticImage(
			engine.MustLoadImage(assets, "assets/layout.png"),
		),
	}
}

type Intro struct {
	*engine.AbsolutePosition
	*engine.StaticImage
}

func (m *Intro) Init() tea.Cmd {
	return nil
}

func (m *Intro) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (m *Intro) Draw(scene *engine.Image) {
	scene.DrawImage(m.Position(), m.Image())

	// Menu
	scene.DrawImage(m.Position().Add(engine.Position{
		X: 64,
		Y: 136,
	}), engine.NewText8x8("GAME A   1 DUCK", textColor).Image())
	scene.DrawImage(m.Position().Add(engine.Position{
		X: 64,
		Y: 152,
	}), engine.NewText8x8("GAME B   2 DUCKS", textColor).Image())
	scene.DrawImage(m.Position().Add(engine.Position{
		X: 64,
		Y: 168,
	}), engine.NewText8x8("GAME C   CLAY SHOOTING", textColor).Image())
}
