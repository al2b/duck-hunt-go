package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
)

func newBlockDuck(maxX int, maxY int) duck {
	return duck{X: rand.Intn(maxX), Y: rand.Intn(maxY)}
}

type duck struct {
	X      int
	Y      int
	isDead bool
}

func (d duck) Init() tea.Cmd {
	return nil
}

func (d duck) Update(msg tea.Msg) (duck, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch _msg := msg.(type) {

	case tickMsg:
		d.Y = rand.Intn(8)
		d.X = rand.Intn(8)
		d.isDead = false

	case tea.MouseMsg:

		switch _msg.Type {
		case tea.MouseLeft:
			if d.X == (_msg.X/2) && d.Y == _msg.Y {
				d.isDead = true
				cmds = append(cmds, func() tea.Msg { return clickMsg(true) })
			}
		}
	}

	return d, tea.Batch(cmds...)
}

func (d duck) View() string {
	if d.isDead {
		return "☠️"
	}
	return "🦆"
}
