package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func MustNewFileHandler(path string) *slog.TextHandler {
	handler, err := NewFileHandler(path)
	if err != nil {
		panic(err)
	}
	return handler
}

func NewFileHandler(path string) (*slog.TextHandler, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %v", path, err)
	}

	return slog.NewTextHandler(file, &slog.HandlerOptions{}), nil
}

// DiscardHandler discards all log output.
// DiscardHandler.Enabled returns false for all Levels.
var DiscardHandler slog.Handler = discardHandler{}

type discardHandler struct{}

func (dh discardHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (dh discardHandler) Handle(context.Context, slog.Record) error { return nil }
func (dh discardHandler) WithAttrs(attrs []slog.Attr) slog.Handler  { return dh }
func (dh discardHandler) WithGroup(name string) slog.Handler        { return dh }
