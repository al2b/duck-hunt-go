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

type Example interface {
	engine.Scene
	fmt.Stringer
}

func New() *Examples {
	return &Examples{
		examples: []Example{
			primitives.New(),
			images.New(),
			animations.New(),
		},
	}
}

type Examples struct {
	current  int
	examples []Example
}

func (e *Examples) Size(windowSize engine.Size) engine.Size {
	return e.examples[e.current].
		Size(windowSize).
		Add(engine.Size{Width: 0, Height: 5})
}

func (e *Examples) FPS() int {
	return e.examples[e.current].FPS()
}

func (e *Examples) Init() (cmd tea.Cmd) {
	return e.examples[e.current].Init()
}

func (e *Examples) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyRight:
			e.current = (e.current + 1) % len(e.examples)
			return e.examples[e.current].Init()
		case tea.KeyLeft:
			e.current = (e.current - 1 + len(e.examples)) % len(e.examples)
			return e.examples[e.current].Init()
		}
	}
	return e.examples[e.current].Update(msg)
}

func (e *Examples) Draw(scene *engine.Image) {
	// Title
	scene.Draw(
		engine.DrawText(image.Pt(0, 0), e.examples[e.current].String(), engine.Font5x5, engine.ColorWhite),
	)
	// Scene
	scene.SubImage(
		image.Pt(0, 5),
		scene.Size().Sub(engine.Size{Width: 0, Height: 5}),
	).Draw(
		e.examples[e.current],
	)
}
