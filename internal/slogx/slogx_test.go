package slogx

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	t.Parallel()
	type args struct {
		level string
	}
	tests := map[string]struct {
		args args
		want slog.Level
	}{
		"empty": {
			args: args{level: ""},
			want: slog.LevelInfo,
		},
		"invalid": {
			args: args{level: "INVALID_LOG_LEVEL"},
			want: slog.LevelInfo,
		},
		"info1": {
			args: args{level: "info"},
			want: slog.LevelInfo,
		},
		"info2": {
			args: args{level: "InFo"},
			want: slog.LevelInfo,
		},
		"warn1": {
			args: args{level: "warn"},
			want: slog.LevelWarn,
		},
		"warn2": {
			args: args{level: "wARN"},
			want: slog.LevelWarn,
		},
		"warn3": {
			args: args{level: "warning"},
			want: slog.LevelWarn,
		},
		"warn4": {
			args: args{level: "WarNiNG"},
			want: slog.LevelWarn,
		},
		"error1": {
			args: args{level: "error"},
			want: slog.LevelError,
		},
		"error2": {
			args: args{level: "ErRoR"},
			want: slog.LevelError,
		},
		"error3": {
			args: args{level: "err"},
			want: slog.LevelError,
		},
		"error4": {
			args: args{level: "ERr"},
			want: slog.LevelError,
		},
		"debug1": {
			args: args{level: "debug"},
			want: slog.LevelDebug,
		},
		"debug2": {
			args: args{level: "DebUG"},
			want: slog.LevelDebug,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ParseLevel(tt.args.level))
		})
	}
}
