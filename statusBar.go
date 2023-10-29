package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type statusBar struct {
	deadDuckCount int
}

func (s statusBar) Init() tea.Cmd {
	return nil
}

func (s statusBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {

	case clickMsg:
		s.deadDuckCount += 1
	}

	return s, nil
}

func (s statusBar) View() string {
	return fmt.Sprintf("%d killed, psychopath!", s.deadDuckCount)
}
