package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Ja7ad/pgp/status"
	"github.com/go-playground/validator/v10"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/codes"
	"io/ioutil"
	"net/http"
	"sync"
)

var _ Transporter = (*Client)(nil)

type Client struct {
	mu        *sync.Mutex
	client    *http.Client
	validator *validator.Validate
	rate      *rate.Limiter
}

type Transporter interface {
	Request(ctx context.Context, url, contentType string, method Method, headers map[string]string, request interface{}, response interface{}) *status.Status
	GetClient() *http.Client
	GetValidator() *validator.Validate
}

// New create client constructor
func New(opts ...Option) Transporter {
	client := &Client{
		mu:        &sync.Mutex{},
		validator: validator.New(),
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

// Request to providers endpoint with http client
func (c *Client) Request(ctx context.Context, url, contentType string, method Method, headers map[string]string, request interface{}, response interface{}) *status.Status {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, method.String(), url, &buf)
	if err != nil {
		return status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if len(contentType) == 0 {
		contentType = "application/json"
	}

	req.Header.Set("Content-Type", contentType)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if c.rate != nil {
		if !c.rate.Allow() {
			return &status.Status{ProviderStatusCode: -1, HttpStatusCode: 429, GrpcStatusCode: 8, Message: "too many requests"}
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}
	resp.Close = true
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return &status.Status{GrpcStatusCode: 0, HttpStatusCode: resp.StatusCode}
}

func (c *Client) GetClient() *http.Client {
	return c.client
}

func (c *Client) GetValidator() *validator.Validate {
	return c.validator
}

func (c *Client) GetMutex() *sync.Mutex {
	return c.mu
}
