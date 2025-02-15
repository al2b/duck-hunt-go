package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

type TickMsg struct {
	time.Time
	time.Duration
}

func tick(fps int) tea.Cmd {
	duration := time.Second / time.Duration(fps)
	return tea.Tick(
		duration,
		func(time time.Time) tea.Msg {
			return TickMsg{
				Time:     time,
				Duration: duration,
			}
		},
	)
}
