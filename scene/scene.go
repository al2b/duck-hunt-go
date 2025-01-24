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
		models: StateModels{
			StateIntro: intro.New(),
			StateGame:  game.New(),
		},
	}
}

type Scene struct {
	mouse  engine.Model
	models StateModels
	state  State
}

func (s *Scene) Width() int {
	return 256
}

func (s *Scene) Height() int {
	return 240
}

func (s *Scene) FPS() int {
	return 60
}

func (s *Scene) Init() tea.Cmd {
	s.state = StateGame

	return tea.Batch(
		s.mouse.Init(),
		s.models[s.state].Init(),
	)
}

func (s *Scene) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		// Switch state
		case "s":
			switch s.state {
			case StateIntro:
				s.state = StateGame
			case StateGame:
				s.state = StateIntro
			}
			return s.models[s.state].Init()
		}
	}

	return tea.Batch(
		s.mouse.Update(msg),
		s.models[s.state].Update(msg),
	)
}

func (s *Scene) Sprites() (sprites engine.Sprites) {
	sprites = append(sprites, s.mouse.Sprites()...)
	sprites = append(sprites, s.models[s.state].Sprites()...)
	return sprites
}

func (s *Scene) Bodies() (bodies engine.Bodies) {
	return bodies.Appends(
		s.mouse.Bodies(),
		s.models[s.state].Bodies(),
	)
}
