package dog

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Dog {
	// Model
	m := &Dog{
		cinematic: engine.Cinematic3DPlayer{
			OnEnd: engine.PlayerOnEndLoop,
		},
	}

	// Drawer
	m.OrderedDrawer = engine.OrderedDrawer{
		engine.ImageDrawer{
			engine.Position2DPointer{
				engine.Position3DProjector{&m.cinematic, engine.OrthographicProjector{}},
			},
			&m.cinematic,
		},
		engine.Position3DOrderer{&m.cinematic},
	}

	return m
}

type Dog struct {
	cinematic engine.Cinematic3DPlayer
	engine.OrderedDrawer
}

func (m *Dog) Init() tea.Cmd {
	m.cinematic.Cinematic = engine.SequenceCinematic3D{
		cinematicTrack,
		cinematicMock,
		cinematicRetrieve1,
		cinematicRetrieve2,
	}
	m.cinematic.Play()

	return nil
}

func (m *Dog) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		m.cinematic.Step(msg.Interval)
	}

	return nil
}
