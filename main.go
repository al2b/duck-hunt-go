package main

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/model"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"os"
)

func main() {

	p := tea.NewProgram(
		engine.New(
			model.New(),
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
