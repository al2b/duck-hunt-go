package engine

import tea "github.com/charmbracelet/bubbletea/v2"

type Model interface {
	Init() tea.Cmd
	Update(msg Msg) tea.Cmd
	Sprites() Sprites
	Bodies() Bodies
}

type ModelUpdatedMsg struct{}
type ModelIntersectedMsg struct{}
