package main

import (
	"flag"

	"github.com/chrismalcolm/schedule/pkg/client"
)

func main() {

	var (
		serverHost string
		interval   int
	)

	flag.StringVar(&serverHost, "serverHost", "", "hostname of the server")
	flag.IntVar(&interval, "interval", 5, "frequency of the location updates in seconds")
	flag.Parse()

	c, err := client.New(serverHost, interval)
	if err != nil {
		panic(err)
	}

	err = c.SendAuthentication()
	if err != nil {
		panic(err)
	}

	for !c.Terminated() {
		c.SendLocation()
		c.Sleep()
	}
}
