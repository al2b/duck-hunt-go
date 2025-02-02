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
		mouse: &mouse.Mouse{},
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

func (s *Scene) Size() (int, int) {
	return 256, 240
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

func (s *Scene) Sprites() (sprites engine.Sprites) {
	switch s.state {
	case StateIntro:
		sprites = sprites.Appends(s.intro.Sprites())
	case StateGame:
		sprites = sprites.Appends(s.game.Sprites())
	}

	return sprites.Append(s.mouse)
}

func (s *Scene) Bodies() (bodies engine.Bodies) {
	switch s.state {
	case StateGame:
		bodies = bodies.Appends(s.game.Bodies())
	}

	return bodies
}
