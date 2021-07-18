package main

import "github.com/chrismalcolm/schedule/pkg/server"

func main() {
	s := server.New()
	s.Serve()
}
