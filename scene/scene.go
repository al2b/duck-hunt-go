package scene

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/scene/game"
	"duck-hunt-go/scene/intro"
	"duck-hunt-go/scene/mouse"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Scene {
	return &Scene{
		mouse: mouse.New(),
		intro: intro.New(),
		game:  game.New(),
	}
}

type Scene struct {
	mouse *mouse.Mouse
	intro *intro.Intro
	game  *game.Game
	state State
}

func (s *Scene) Size(_ engine.Size) engine.Size {
	return engine.Size{
		Width:  256,
		Height: 240,
	}
}

func (s *Scene) FPS() int {
	return 60
}

func (s *Scene) Init() (cmd tea.Cmd) {
	s.state = StateGame

	switch s.state {
	case StateIntro:
		cmd = s.intro.Init()
	case StateGame:
		cmd = s.game.Init()
	}

	return tea.Batch(
		tea.SetWindowTitle("Duck Hunt"),
		s.mouse.Init(),
		cmd,
	)
}

func (s *Scene) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Switch state
		case "s":
			switch s.state {
			case StateIntro:
				s.state = StateGame
				return s.game.Init()
			case StateGame:
				s.state = StateIntro
				return s.intro.Init()
			}
		// Init current state
		case "i":
			switch s.state {
			case StateIntro:
				return s.intro.Init()
			case StateGame:
				return s.game.Init()
			}
		}
	}

	switch s.state {
	case StateIntro:
		cmd = s.intro.Update(msg)
	case StateGame:
		cmd = s.game.Update(msg)
	}

	return tea.Batch(
		s.mouse.Update(msg),
		cmd,
	)
}

func (s *Scene) Draw(scene *engine.Image) {
	switch s.state {
	case StateIntro:
		scene.Draw(s.intro)
	case StateGame:
		scene.Draw(s.game)
	}
	scene.Draw(s.mouse)
}
