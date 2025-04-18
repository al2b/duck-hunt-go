package game

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/intro"
	"duck-hunt-go/game/mouse"
	"duck-hunt-go/game/round"
	"duck-hunt-go/game/state"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Game {
	return &Game{
		mouse: mouse.New(),
		intro: intro.New(),
		game:  round.New(),
	}
}

type Game struct {
	mouse *mouse.Mouse
	intro *intro.Intro
	game  *round.Game
	state state.State
}

func (g *Game) Size(_ engine.Size) engine.Size {
	return engine.Size{
		Width:  256,
		Height: 240,
	}
}

func (g *Game) FPS() int {
	return 60
}

func (g *Game) Init() (cmd tea.Cmd) {
	g.state = state.StateGame

	switch g.state {
	case state.StateIntro:
		cmd = g.intro.Init()
	case state.StateGame:
		cmd = g.game.Init()
	}

	return tea.Batch(
		tea.SetWindowTitle("Duck Hunt"),
		g.mouse.Init(),
		cmd,
	)
}

func (g *Game) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Switch state
		case "s":
			switch g.state {
			case state.StateIntro:
				g.state = state.StateGame
				return g.game.Init()
			case state.StateGame:
				g.state = state.StateIntro
				return g.intro.Init()
			}
		// Init current state
		case "i":
			switch g.state {
			case state.StateIntro:
				return g.intro.Init()
			case state.StateGame:
				return g.game.Init()
			}
		}
	}

	switch g.state {
	case state.StateIntro:
		cmd = g.intro.Update(msg)
	case state.StateGame:
		cmd = g.game.Update(msg)
	}

	return tea.Batch(
		g.mouse.Update(msg),
		cmd,
	)
}

func (g *Game) Draw(scene *engine.Image) {
	switch g.state {
	case state.StateIntro:
		scene.Draw(g.intro)
	case state.StateGame:
		scene.Draw(g.game)
	}
	scene.Draw(g.mouse)
}
