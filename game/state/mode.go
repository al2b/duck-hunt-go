package state

import (
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Mode int

const (
	Mode1Duck Mode = iota
	Mode2Ducks
)

func SetMode(mode Mode) tea.Cmd {
	return func() tea.Msg {
		return SetModeMsg(mode)
	}
}

type SetModeMsg Mode

func (msg SetModeMsg) Mode() Mode {
	return Mode(msg)
}
