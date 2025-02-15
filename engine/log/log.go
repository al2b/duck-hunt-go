package log

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"log/slog"
	"time"
)

type Msg struct {
	slog.Record
}

func Info(msg string, args ...any) tea.Cmd {
	t := time.Now()
	return func() tea.Msg {
		r := slog.NewRecord(t,
			slog.LevelInfo,
			msg,
			0,
		)
		r.Add(args...)
		return Msg{r}
	}
}
