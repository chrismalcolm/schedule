package client

import (
	"net/http"
	"time"

	"github.com/chrismalcolm/schedule/pkg/models"
)

type Client struct {
	client     *http.Client
	interval   int
	user       string
	password   string
	serverHost string
	status     models.Status
}

func New(serverHost string, interval int) (*Client, error) {
	// TODO Verify host and interval make sense
	return &Client{
		client:     &http.Client{},
		serverHost: serverHost,
		interval:   interval,
	}, nil
}

func (c *Client) SendAuthentication() error {
	// TODO Send request to set up basic auth with server
	return nil
}

func (c *Client) SendLocation() {
	// TODO Send request of location of the bus
}

func (c *Client) Terminated() bool {
	return c.status == models.FINISHED || c.status == models.DISCONTINUED
}

func (c *Client) Sleep() {
	time.Sleep(time.Duration(c.interval) * time.Second)
}
