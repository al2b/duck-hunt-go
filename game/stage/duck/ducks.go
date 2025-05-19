package duck

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
)

type id int

func NewDucks(space *cp.Space, n int) Ducks {
	ducks := make(Ducks, n)
	for i := 0; i < n; i++ {
		ducks[i] = New(space, id(i))
	}
	return ducks
}

type Ducks []*Duck

func (ducks Ducks) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, duck := range ducks {
		cmds = append(cmds, duck.Init())
	}
	return tea.Batch(cmds...)
}

func (ducks Ducks) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case DiscriminatedShotMsg:
		if id, ok := msg.Discriminator.(id); ok {
			if id >= 0 && int(id) < len(ducks) {
				return ducks[id].Update(msg.ShotMsg)
			}
		}
		return nil
	}

	for _, duck := range ducks {
		cmds = append(cmds, duck.Update(msg))
	}
	return tea.Batch(cmds...)
}

func (ducks Ducks) Draw(dst *engine.Image) {
	for _, duck := range ducks {
		dst.Draw(duck)
	}
}
