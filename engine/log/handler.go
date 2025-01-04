package log

import (
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
