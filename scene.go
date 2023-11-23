package main

import (
	_ "embed"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

type clickMsg bool

type tickMsg time.Time

//go:embed resources/duck-start.txt
var duckStart string

//go:embed resources/backgroundPlay.txt
var backgroundPlay string

func newScene() scene {
	return scene{
		board:     newBoard(),
		statusBar: statusBar{},
	}
}

type scene struct {
	width     int
	height    int
	X         int
	Y         int
	state     bool
	statusBar tea.Model

	board tea.Model
}

func (s scene) Init() tea.Cmd {
	return tea.Batch(
		s.statusBar.Init(),
		s.board.Init(),
		doTick(),
	)
}

func (s scene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch _msg := msg.(type) {

	case tickMsg:
		if s.state {
			s.board, cmd = s.board.Update(_msg)
			cmds = append(cmds, cmd)
			cmds = append(cmds, doTick())
		}
	case clickMsg:
		s.statusBar, cmd = s.statusBar.Update(clickMsg(true))
		//s.board, cmd = s.board.Update(clickMsg(true))
	case tea.MouseMsg:
		s.X, s.Y = _msg.X, _msg.Y

		switch _msg.Type {
		case tea.MouseLeft:
			if !s.state {
				s.state = true
			} else {
				s.board, cmd = s.board.Update(_msg)

				cmds = append(cmds, cmd)
			}
		}

	case tea.WindowSizeMsg:
		s.height, s.width = _msg.Height, _msg.Width

	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch _msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return s, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return s, tea.Batch(cmds...)
}

func (s scene) View() string {
	//return fmt.Sprintf( "Hello Duck! \nWidth: %d, Height: %d and Mouse X: %d, Mouse Y: %d", s.width, s.height, s.X, s.Y)
	if !s.state {
		return duckStart
	}

	// Place a paragraph in the bottom right corner of a 30x80 cell space.
	return lipgloss.JoinVertical(lipgloss.Left, s.board.View(), s.statusBar.View())
}

func doTick() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
