package engine

import tea "github.com/charmbracelet/bubbletea/v2"

type Model interface {
	Init()
	Update(msgs []tea.Msg)
	Bodies() Bodies
	Sprites8() Sprites8
	Sprites24() Sprites24
}

type ModelUpdatedMsg struct{}
type ModelIntersectedMsg struct{}
