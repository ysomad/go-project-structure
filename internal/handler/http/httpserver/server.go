package httpserver

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"
)

const (
	defaultAddr            = ":80"
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 3 * time.Second
)

type server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(ctx context.Context, h http.Handler, opts ...Option) *server {
	srv := &http.Server{
		Addr:         defaultAddr,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Handler:      h,
	}

	s := &server{
		server:          srv,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start(ctx)
	return s
}

func (s *server) start(ctx context.Context) {
	go func() {
		slog.InfoContext(ctx, "starting http server at "+s.server.Addr)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *server) Notify() <-chan error {
	return s.notify
}

func (s *server) Shutdown(ctx context.Context) error {
	slog.InfoContext(ctx, "shutting down http server at "+s.server.Addr)

	ctx, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx) //nolint:wrapcheck // not needed here
}
