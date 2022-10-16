package client

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type Option func(client *Client)

// WithCustomClient use custom http client instead default client
func WithCustomClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

// WithRateLimit make rate limit for example every time 5 * time.Second for 50 request
func WithRateLimit(every time.Duration, requestPerTime int) Option {
	return func(client *Client) {
		client.rate = rate.NewLimiter(rate.Every(every), requestPerTime)
	}
}
