package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/table"
)

func newBoard() board {
	board := board{}
	for i := 0; i < 5; i++ {
		board.ducks[i] = newBlockDuck(8, 8)
	}

	return board
}

type board struct {
	gameBoard [][]string
	ducks     [5]duck
}

func (b board) Init() tea.Cmd {

	return nil
}

func (b board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch _msg := msg.(type) {

	case tickMsg:
		for i, v := range b.ducks {
			b.ducks[i], cmd = v.Update(msg)
			cmds = append(cmds, cmd)
		}

	case tea.MouseMsg:
		switch _msg.Type {
		case tea.MouseLeft:
			for i, d := range b.ducks {
				b.ducks[i], cmd = d.Update(msg)
				cmds = append(cmds, cmd)
			}
		}
	}

	return b, tea.Batch(cmds...)
}

func (b board) View() string {

	gameBoard := make([][]string, 8)
	for i := 0; i < 8; i++ {
		gameBoard[i] = make([]string, 8)
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			var found bool
			var duck duck
			for _, _duck := range b.ducks {
				if x == _duck.X && y == _duck.Y {
					found = true
					duck = _duck
					continue
				}
			}
			if found {
				gameBoard[y][x] = duck.View()
			} else {
				gameBoard[y][x] = "  "
			}
		}
	}

	t := table.New().
		BorderColumn(false).
		BorderRow(false).
		BorderTop(false).
		BorderRight(false).
		BorderBottom(false).
		BorderLeft(false).
		Rows(gameBoard...)

	return t.Render()
}
