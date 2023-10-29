package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type scene struct {
	width  int
	height int
	X      int
	Y      int
}

func (s scene) Init() tea.Cmd {
	return nil
}

func (s scene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch _msg := msg.(type) {

	case tea.MouseMsg:
		s.X, s.Y = _msg.X, _msg.Y

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
	return s, nil
}

func (s scene) View() string {
	return fmt.Sprintf("Hello Duck! Width: %d, Height: %d and Mouse X: %d, Mouse Y: %d", s.width, s.height, s.X, s.Y)
}

func main() {
	p := tea.NewProgram(scene{}, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Goodbye Duck")
}
