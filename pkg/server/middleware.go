package server

import "net/http"

func (s *Server) basicAuth(h http.HandlerFunc) http.HandlerFunc {
	// TODO Implement middleware or basic auth
	return h
}

func (s *Server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	// TODO Implement middleware or basic auth for admin only
	return h
}
