package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultAddr         = ":80"
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
)

type Server struct {
	notify       chan error
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration

	srv *http.Server
}

func New(router http.Handler, opts ...Option) *Server {
	s := &Server{
		address:      _defaultAddr,
		readTimeout:  _defaultReadTimeout,
		writeTimeout: _defaultWriteTimeout,
		notify:       make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         s.address,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
	}

	s.srv = srv

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.srv.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
