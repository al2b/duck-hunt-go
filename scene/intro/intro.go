package intro

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
	"image/color"
)

//go:embed assets/*
var assets embed.FS

var (
	textColor = color.RGBA{R: 0xff, G: 0xa0, B: 0x00}
)

func New() *Intro {
	return &Intro{
		layoutImage: engine.MustLoadImage(engine.ImageFile(assets, "assets/layout.png")),
		cursorImage: engine.MustLoadImage(engine.ImageFile(assets, "assets/cursor.png")),
	}
}

type Intro struct {
	layoutImage    *engine.Image
	cursorImage    *engine.Image
	cursorPosition int
}

func (m *Intro) Init() tea.Cmd {
	m.cursorPosition = 0
	return nil
}

func (m *Intro) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyUp:
			m.cursorPosition = (m.cursorPosition - 1 + 3) % 3
		case tea.KeyDown:
			m.cursorPosition = (m.cursorPosition + 1) % 3
		}
	}
	return nil
}

func (m *Intro) Draw(scene *engine.Image) {
	scene.Draw(
		// Layout
		engine.DrawImage(image.Pt(0, 0), m.layoutImage),
		// Menu
		engine.DrawImage(image.Pt(64, 136), engine.NewText8x8("GAME A   1 DUCK", textColor).Image()),
		engine.DrawImage(image.Pt(64, 152), engine.NewText8x8("GAME B   2 DUCKS", textColor).Image()),
		engine.DrawImage(image.Pt(64, 168), engine.NewText8x8("GAME C   CLAY SHOOTING", textColor).Image()),
		// Cursor
		engine.DrawImage(image.Pt(48, 136+(m.cursorPosition*16)), m.cursorImage),
	)
}
