package client

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type Option func(client *Client)

// WithCustomClient use custom http httpClient instead default httpClient
func WithCustomClient(client *http.Client) Option {
	return func(c *Client) {
		if _, ok := client.Transport.(*transport); !ok {
			t := client.Transport
			if t != nil {
				client.Transport = &transport{Base: t}
			} else {
				client.Transport = &transport{Base: http.DefaultTransport}
			}
		}
		c.httpClient = client
	}
}

// WithRateLimit make rate limit for example every time 5 * time.Second for 50 request
func WithRateLimit(every time.Duration, requestPerTime int) Option {
	return func(client *Client) {
		client.rate = rate.NewLimiter(rate.Every(every), requestPerTime)
	}
}
