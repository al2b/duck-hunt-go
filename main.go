package main

import (
	"duck-hunt-go/engine"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func main() {

	p := tea.NewProgram(
		engine.New(
			NewGame(),
		),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
