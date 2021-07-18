package server

import "net/http"

// TODO Finished implementing handles

func (s *Server) handleLocation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(""))
			return
		}
	}
}

func (s *Server) handleAddBuses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(""))
			return
		}
	}
}

func (s *Server) handleUpdateBuses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(""))
			return
		}
	}
}

func (s *Server) handleDeleteBuses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(""))
			return
		}
	}
}
