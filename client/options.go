package client

import (
	"net/http"
)

type Option func(client *Client)

func WithCustomClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}
