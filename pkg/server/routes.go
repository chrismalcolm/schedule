package server

func (s *Server) routes() {
	s.router.HandleFunc("/location", s.handleLocation())
	s.router.HandleFunc("/add", s.basicAuth(s.handleAddBuses()))
	s.router.HandleFunc("/update", s.adminOnly(s.handleUpdateBuses()))
	s.router.HandleFunc("/delete", s.adminOnly(s.handleDeleteBuses()))
}
