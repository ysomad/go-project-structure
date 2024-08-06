package slogx

import (
	"fmt"
	"log/slog"

	"github.com/pressly/goose/v3"
)

var _ goose.Logger = &GooseLogger{}

type GooseLogger struct{}

func (GooseLogger) Fatalf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v...))
}

func (GooseLogger) Printf(format string, v ...interface{}) {
	slog.Info(fmt.Sprintf(format, v...))
}
