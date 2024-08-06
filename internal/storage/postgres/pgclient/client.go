package pgclient

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxConns     = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type Client struct {
	maxConns     int32
	connAttempts uint8
	connTimeout  time.Duration
	tracer       pgx.QueryTracer

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(ctx context.Context, url string, opts ...Option) (*Client, error) {
	cli := &Client{
		maxConns:     defaultMaxConns,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
		Builder:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	for _, opt := range opts {
		opt(cli)
	}

	conf, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if cli.tracer != nil {
		conf.ConnConfig.Tracer = cli.tracer
	}

	conf.MaxConns = cli.maxConns

	cli.Pool, err = pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("pool create: %w", err)
	}

	for cli.connAttempts > 0 {
		if err = cli.Pool.Ping(ctx); err == nil {
			break
		}

		slog.InfoContext(ctx, "connecting to postgres, attempts left", "attempts", cli.connAttempts)
		time.Sleep(cli.connTimeout)
		cli.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return cli, nil
}

func (p *Client) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
