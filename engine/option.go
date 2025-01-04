package engine

import "log/slog"

type Option func(engine *Engine)

func WithLogHandler(handler slog.Handler) Option {
	return func(engine *Engine) {
		engine.logHandler = handler
	}
}
