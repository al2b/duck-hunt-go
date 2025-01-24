package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

type TickMsg time.Time

func tick(fps int) tea.Cmd {
	return tea.Tick(time.Second/time.Duration(fps), func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
