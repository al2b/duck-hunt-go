package game

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/config"
	"duck-hunt-go/game/menu"
	"duck-hunt-go/game/stage"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Game {
	return &Game{}
}

type Game struct {
	pause bool
	model engine.DrawModel
}

func (g *Game) Size(_ engine.Size) engine.Size {
	return engine.Size{256, 240}
}

func (g *Game) Init() tea.Cmd {
	// Set the initial model
	g.model = stage.New()

	return tea.Batch(
		tea.SetWindowTitle("Duck Hunt"),
		engine.StartTicker(config.TickInterval),
		g.model.Init(),
	)
}

func (g *Game) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Pause
		case "p":
			if !g.pause {
				g.pause = true
				return engine.StopTicker
			} else {
				g.pause = false
				return engine.StartTicker(config.TickInterval)
			}
		// Next tick
		case "n":
			if g.pause {
				return engine.StepTick(config.TickInterval)
			}
		// Switch the current model
		case "s":
			switch g.model.(type) {
			case *menu.Menu:
				g.model = stage.New()
			case *stage.Stage:
				g.model = menu.New()
			}
			return g.model.Init()
		// Re-init the current model
		case "i":
			switch g.model.(type) {
			case *menu.Menu:
				g.model = menu.New()
			case *stage.Stage:
				g.model = stage.New()
			}
			return g.model.Init()
		}
	}

	// Update the current model
	return g.model.Update(msg)
}

func (g *Game) Draw(dst *engine.Image) {
	// Draw the current model
	dst.Draw(g.model)
}
