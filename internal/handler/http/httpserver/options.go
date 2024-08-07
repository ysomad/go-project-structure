package httpserver

import (
	"net"
	"strconv"
	"time"
)

type Option func(*server)

func WithPort(port int) Option {
	return func(s *server) {
		s.server.Addr = net.JoinHostPort("", strconv.Itoa(port))
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.WriteTimeout = timeout
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.shutdownTimeout = timeout
	}
}
