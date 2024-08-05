package slogx

import (
	"context"
	"log/slog"
)

var (
	Level              = &slog.LevelVar{} //nolint:gochecknoglobals // for updating log level in runtime
	_     slog.Handler = &LevelFilter{}
)

// LevelFilter is a slog handler middleware which is filtering by slog.Level.
type LevelFilter struct {
	next  slog.Handler
	level *slog.LevelVar
}

func NewLevelFilter(next slog.Handler, level string) *LevelFilter {
	Level.Set(ParseLevel(level))
	return &LevelFilter{next: next, level: Level}
}

func (h *LevelFilter) Handle(ctx context.Context, r slog.Record) error {
	if r.Level >= h.level.Level() {
		return h.next.Handle(ctx, r) //nolint:wrapcheck // not needed here
	}
	return nil
}

func (h *LevelFilter) Enabled(ctx context.Context, l slog.Level) bool { return h.next.Enabled(ctx, l) }
func (h *LevelFilter) WithAttrs(attrs []slog.Attr) slog.Handler       { return h.next.WithAttrs(attrs) }
func (h *LevelFilter) WithGroup(name string) slog.Handler             { return h.next.WithGroup(name) }
