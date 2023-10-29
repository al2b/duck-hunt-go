package main

import (
	_ "embed"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type clickMsg bool

//go:embed resources/duck-start.txt
var duckStart string

//go:embed resources/backgroundPlay.txt
var backgroundPlay string

func newScene() scene {
	return scene{
		duck:      blockDuck{},
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
	duck      tea.Model
}

func (s scene) Init() tea.Cmd {
	return tea.Batch(
		s.duck.Init(),
		s.statusBar.Init(),
	)
}

func (s scene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	s.duck, cmd = s.duck.Update(msg)
	cmds = append(cmds, cmd)

	switch _msg := msg.(type) {

	case tea.MouseMsg:
		s.X, s.Y = _msg.X, _msg.Y

		switch _msg.Type {
		case tea.MouseLeft:
			if !s.state {
				s.state = true
			} else {
				s.statusBar, cmd = s.statusBar.Update(clickMsg(true))
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
	return lipgloss.Place(s.width, s.height, lipgloss.Right, lipgloss.Bottom, s.statusBar.View())
}

func main() {
	p := tea.NewProgram(newScene(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Goodbye Duck")
}
