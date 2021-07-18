package server

import (
	"net/http"
	"sync"

	"github.com/chrismalcolm/schedule/pkg/models"
)

type Server struct {
	router *http.ServeMux
	// TODO Add examples
	store []models.Bus
	mutex sync.Mutex
}

func New() *Server {
	s := &Server{}
	s.router = http.NewServeMux()
	s.routes()
	return s
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.router)
}
