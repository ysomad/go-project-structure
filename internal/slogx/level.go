package slogx

import (
	"context"
	"log/slog"
	"strings"
)

func init() { //nolint:gochecknoinits // to init default log level
	LevelVar.Set(slog.LevelInfo)
}

var (
	// LevelVar is global logging level which allows to change it in runtime.
	LevelVar customLevelVar = customLevelVar{&slog.LevelVar{}} //nolint:gochecknoglobals // for updating log level in runtime
	_        slog.Handler   = &levelFilter{}
)

// customLevelVar is a wrapper for slog.customLevelVar.String method to return custom log level strings.
type customLevelVar struct {
	*slog.LevelVar
}

// String rewrites slog.LevelVar.String method.
func (v customLevelVar) String() string {
	return customLevel(v.Level()).String()
}

type customLevel slog.Level

func (l customLevel) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelFatal:
		return "FATAL"
	default:
		return slog.Level(l).String()
	}
}

// Level returns receiver as slog.Level, implements slog.Leveler.
func (l customLevel) Level() slog.Level {
	return slog.Level(l)
}

const (
	// LevelTrace and LevelFatal implements opentelemetry logging data model.
	// Opentelemetry by default adding offset +9 to slog.Level thats why LevelTrace=-8 not 1
	// https://opentelemetry.io/docs/specs/otel/logs/data-model/#field-severitynumber
	LevelTrace customLevel = -8
	LevelFatal customLevel = 12
)

func ParseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error", "err":
		return slog.LevelError
	case "trace":
		return slog.Level(LevelTrace)
	case "fatal":
		return slog.Level(LevelFatal)
	default:
		return slog.LevelInfo
	}
}

// levelFilter is a slog handler middleware which is filtering by slog.Level.
type levelFilter struct {
	next slog.Handler
}

func NewLevelFilter(next slog.Handler) *levelFilter {
	return &levelFilter{next: next}
}

func (h *levelFilter) Handle(ctx context.Context, r slog.Record) error {
	if r.Level >= LevelVar.Level() {
		return h.next.Handle(ctx, r) //nolint:wrapcheck // not needed here
	}
	return nil
}

func (h *levelFilter) Enabled(ctx context.Context, l slog.Level) bool { return h.next.Enabled(ctx, l) }
func (h *levelFilter) WithAttrs(attrs []slog.Attr) slog.Handler       { return h.next.WithAttrs(attrs) }
func (h *levelFilter) WithGroup(name string) slog.Handler             { return h.next.WithGroup(name) }
