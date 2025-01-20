package engine

import tea "github.com/charmbracelet/bubbletea/v2"

type Model interface {
	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	Bodies() Bodies
	Sprites8() Sprites8
	Sprites24() Sprites24
}

type ModelUpdatedMsg struct{}
type ModelIntersectedMsg struct{}
