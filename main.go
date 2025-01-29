package main

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/scene"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"io"
	"log"
	"os"
)

func main() {

	// Discard logs
	log.SetOutput(io.Discard)

	p := tea.NewProgram(
		engine.New(
			scene.New(),
		),
	)
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
