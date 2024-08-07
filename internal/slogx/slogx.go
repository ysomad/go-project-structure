package slogx

import (
	"context"
	"log/slog"
)

func TraceContext(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelTrace.Level(), msg, args...)
}

func Trace(msg string, args ...any) {
	slog.Log(context.Background(), LevelTrace.Level(), msg, args...)
}

func FatalContext(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelFatal.Level(), msg, args...)
}

func Fatal(msg string, args ...any) {
	slog.Log(context.Background(), LevelFatal.Level(), msg, args...)
}
