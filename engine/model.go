package engine

import tea "github.com/charmbracelet/bubbletea/v2"

type Model interface {
	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
}

type DrawModel interface {
	Model
	Drawer
}
