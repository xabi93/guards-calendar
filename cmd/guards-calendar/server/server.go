package server

import (
	"net/http"

	"github.com/xabi93/guards-calendar/api"
)

func New() *Server {
	var s Server

	s.bootstrap()

	return &s
}

type Server struct {
	api http.Handler
}

func (s *Server) bootstrap() {
	s.api = api.New()
}

func (s Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.api.ServeHTTP(rw, r)
}
