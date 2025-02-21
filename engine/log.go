package engine

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"log/slog"
	"os"
	"time"
)

func NewLogFileHandler(path string) (*slog.TextHandler, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %v", path, err)
	}

	return slog.NewTextHandler(file, &slog.HandlerOptions{}), nil
}

func MustNewLogFileHandler(path string) *slog.TextHandler {
	handler, err := NewLogFileHandler(path)
	if err != nil {
		panic(err)
	}
	return handler
}

type LogMsg struct {
	slog.Record
}

func LogInfo(msg string, args ...any) tea.Cmd {
	t := time.Now()
	return func() tea.Msg {
		r := slog.NewRecord(t,
			slog.LevelInfo,
			msg,
			0,
		)
		r.Add(args...)
		return LogMsg{r}
	}
}
