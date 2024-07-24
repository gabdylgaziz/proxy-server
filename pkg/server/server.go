package server

import (
	"context"
	"net"
	"net/http"
)

type Server struct {
	http     *http.Server
	listener net.Listener
}

type Configuration func(r *Server) error

func New(configs ...Configuration) (r *Server, err error) {
	// Create the Server
	r = &Server{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(r); err != nil {
			return
		}
	}
	return
}

func (s *Server) Run() (err error) {
	if s.http != nil {
		go func() {
			if err = s.http.ListenAndServe(); err != nil {
				//TODO: log add
				return
			}
		}()
	}

	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	if s.http != nil {
		if err = s.http.Shutdown(ctx); err != nil {
			return
		}
	}

	return
}

func WithHTTPServer(handler http.Handler, port string) Configuration {
	return func(s *Server) (err error) {
		s.http = &http.Server{
			Handler: handler,
			Addr:    ":" + port,
		}
		return
	}
}
