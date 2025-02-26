package examples

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/examples/animations"
	"duck-hunt-go/examples/images"
	"duck-hunt-go/examples/primitives"
	"duck-hunt-go/examples/space"
	"duck-hunt-go/examples/subimages"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
)

type Example interface {
	engine.Scene
	fmt.Stringer
}

func New() *Examples {
	return &Examples{
		examples: []Example{
			primitives.New(),
			subimages.New(),
			images.New(),
			animations.New(),
			space.New(),
		},
	}
}

type Examples struct {
	current  int
	examples []Example
}

func (s *Examples) Size(windowSize engine.Size) engine.Size {
	return s.examples[s.current].
		Size(windowSize).
		Add(engine.Size{Width: 0, Height: 5})
}

func (s *Examples) FPS() int {
	return s.examples[s.current].FPS()
}

func (s *Examples) Init() (cmd tea.Cmd) {
	return s.examples[s.current].Init()
}

func (s *Examples) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			s.current = (s.current + 1) % len(s.examples)
			return s.examples[s.current].Init()
		case tea.KeyLeft:
			s.current = (s.current - 1 + len(s.examples)) % len(s.examples)
			return s.examples[s.current].Init()
		}
	}
	return s.examples[s.current].Update(msg)
}

func (s *Examples) Draw(scene *engine.Image) {
	// Title
	scene.Draw(
		engine.DrawText(image.Pt(0, 0), s.examples[s.current].String(), engine.Font5x5, engine.ColorWhite),
	)
	// Scene
	scene.SubImage(
		image.Pt(0, 5),
		scene.Size().Sub(engine.Size{Width: 0, Height: 5}),
	).Draw(
		s.examples[s.current],
	)
}
