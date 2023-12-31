package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func New(port string, handler http.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
