package main

import tea "github.com/charmbracelet/bubbletea"

type blockDuck struct {
}

func (b blockDuck) Init() tea.Cmd {
	return nil
}

func (b blockDuck) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return b, nil
}

func (b blockDuck) View() string {
	return "🦆"
}
