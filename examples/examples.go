package examples

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/examples/animations"
	"duck-hunt-go/examples/images"
	"duck-hunt-go/examples/mouse"
	"duck-hunt-go/examples/path"
	"duck-hunt-go/examples/primitives"
	"duck-hunt-go/examples/space"
	"duck-hunt-go/examples/subimages"
	"duck-hunt-go/examples/text"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Example interface {
	engine.Scene
	fmt.Stringer
}

func New(start int) *Examples {
	var examples = []Example{
		primitives.New(),
		subimages.New(),
		images.New(),
		animations.New(),
		mouse.New(),
		text.New(),
		path.New(),
		space.New(),
	}

	return &Examples{
		examples: examples,
		current:  (start - 1) % len(examples),
	}
}

type Examples struct {
	current  int
	examples []Example
}

func (s *Examples) Size(windowSize engine.Size) engine.Size {
	return s.examples[s.current].
		Size(windowSize).
		Add(engine.Size{0, 6})
}

func (s *Examples) Init() (cmd tea.Cmd) {
	return s.examples[s.current].Init()
}

func (s *Examples) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		msg.Y -= 6
		return s.examples[s.current].Update(msg)
	case tea.MouseReleaseMsg:
		msg.Y -= 6
		return s.examples[s.current].Update(msg)
	case tea.MouseWheelMsg:
		msg.Y -= 6
		return s.examples[s.current].Update(msg)
	case tea.MouseMotionMsg:
		msg.Y -= 6
		return s.examples[s.current].Update(msg)
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

func (s *Examples) Draw(dst *engine.Image) {
	// Title
	dst.Draw(
		engine.TextDrawer{engine.Pt(0, 0),
			engine.Text{s.examples[s.current].String(), engine.Font6x6, engine.ColorWhite},
		},
	)

	// Scene
	dst.SubImage(
		engine.Pt(0, 6),
		dst.Size().Sub(engine.Size{0, 6}),
	).Draw(
		s.examples[s.current],
	)
}
