package dog

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Dog {
	return &Dog{}
}

type Dog struct {
	cinematic engine.Cinematic3DPlayer
	engine.OrderedDrawer
}

func (m *Dog) Init() tea.Cmd {
	// Cinematic
	m.cinematic.Cinematic = engine.SequenceCinematic3D{
		cinematicTrack,
		cinematicMock,
		cinematicRetrieve1,
		cinematicRetrieve2,
	}
	m.cinematic.OnEnd = engine.PlayerOnEndLoop
	m.cinematic.Play()

	// Drawer
	m.OrderedDrawer.Drawer = engine.ImageDrawer{
		engine.Position2DPointer{
			engine.Position3DProjector{&m.cinematic, engine.OrthographicProjector{}},
		},
		&m.cinematic,
	}
	m.OrderedDrawer.Orderer = engine.Position3DOrderer{&m.cinematic}

	return nil
}

func (m *Dog) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Cinematic
		m.cinematic.Step(msg.Interval)
	}

	return nil
}
