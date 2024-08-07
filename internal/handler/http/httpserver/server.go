package httpserver

import (
	"context"
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

	s.start()
	return s
}

func (s *server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *server) Notify() <-chan error {
	return s.notify
}

func (s *server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx) //nolint:wrapcheck // not needed here
}
