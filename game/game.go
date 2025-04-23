package game

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/mouse"
	"duck-hunt-go/game/round"
	"duck-hunt-go/game/state"
	"duck-hunt-go/game/title"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Game {
	return &Game{
		mouse: mouse.New(),
		title: title.New(),
		round: round.New(),
	}
}

type Game struct {
	mouse *mouse.Mouse
	title *title.Title
	round *round.Round
	state state.State
}

func (g *Game) Size(_ engine.Size) engine.Size {
	return engine.Size{256, 240}
}

func (g *Game) TPS() int {
	return 60
}

func (g *Game) Init() (cmd tea.Cmd) {
	g.state = state.Play

	switch g.state {
	case state.Title:
		cmd = g.title.Init()
	case state.Play:
		cmd = g.round.Init()
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
			case state.Title:
				g.state = state.Play
				return g.round.Init()
			case state.Play:
				g.state = state.Title
				return g.title.Init()
			}
		// Init current state
		case "i":
			switch g.state {
			case state.Title:
				return g.title.Init()
			case state.Play:
				return g.round.Init()
			}
		}
	}

	switch g.state {
	case state.Title:
		cmd = g.title.Update(msg)
	case state.Play:
		cmd = g.round.Update(msg)
	}

	return tea.Batch(
		g.mouse.Update(msg),
		cmd,
	)
}

func (g *Game) Draw(dst *engine.Image) {
	switch g.state {
	case state.Title:
		dst.Draw(g.title)
	case state.Play:
		dst.Draw(g.round)
	}

	// Mouse
	dst.Draw(g.mouse)
}
