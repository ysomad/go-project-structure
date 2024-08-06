package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // for goose running migrations via pgx
	"github.com/pressly/goose/v3"
	"github.com/ysomad/go-project-structure/internal/slogx"
)

func runMigrations(ctx context.Context) error {
	goose.SetLogger(slogx.GooseLogger{})

	url, ok := os.LookupEnv("POSTGRES_URL")
	if !ok || url == "" {
		return errors.New("POSTGRES_URL env variable not declared") //nolint:err113 // not needed here
	}

	db, err := goose.OpenDBWithDriver("pgx", url)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			slog.Warn("db not closed: " + err.Error())
		}
	}()

	if err := goose.RunContext(ctx, "up", db, "./migrations"); err != nil {
		return fmt.Errorf("run: %w", err)
	}

	return nil
}
