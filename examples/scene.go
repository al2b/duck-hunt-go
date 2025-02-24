package examples

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/examples/animations"
	"duck-hunt-go/examples/images"
	"duck-hunt-go/examples/primitives"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

type StringScene interface {
	engine.Scene
	fmt.Stringer
}

func New() *Scene {
	return &Scene{
		scenes: []StringScene{
			primitives.New(),
			images.New(),
			animations.New(),
		},
	}
}

type Scene struct {
	current int
	scenes  []StringScene
}

func (s *Scene) Size(windowSize engine.Size) engine.Size {
	return s.scenes[s.current].
		Size(windowSize).
		Add(engine.Size{Width: 0, Height: 5})
}

func (s *Scene) FPS() int {
	return s.scenes[s.current].FPS()
}

func (s *Scene) Init() (cmd tea.Cmd) {
	return s.scenes[s.current].Init()
}

func (s *Scene) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			s.current = (s.current + 1) % len(s.scenes)
			return s.scenes[s.current].Init()
		case tea.KeyLeft:
			s.current = (s.current - 1 + len(s.scenes)) % len(s.scenes)
			return s.scenes[s.current].Init()
		}
	}
	return s.scenes[s.current].Update(msg)
}

func (s *Scene) Draw(scene *engine.Image) {
	// Title
	scene.Draw(
		engine.DrawText(image.Pt(0, 0), s.scenes[s.current].String(), engine.Font5x5, engine.ColorWhite),
	)

	// Scene
	scene.SubImage(
		image.Pt(0, 5),
		scene.Size().Sub(engine.Size{Width: 0, Height: 5}),
	).Draw(
		s.scenes[s.current],
	)
}
