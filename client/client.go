package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Ja7ad/pgp/errors"
	"github.com/go-playground/validator/v10"
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
}

type Transporter interface {
	Request(ctx context.Context, url, method, contentType string, headers map[string]string, request interface{}, response interface{}, errCh chan<- *errors.Error)
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
func (c *Client) Request(ctx context.Context, url, method, contentType string, headers map[string]string, request interface{}, response interface{}, errCh chan<- *errors.Error) {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		errCh <- errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, method, url, &buf)
	if err != nil {
		errCh <- errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if len(contentType) == 0 {
		contentType = "application/json"
	}

	req.Header.Set("Content-Type", contentType)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		errCh <- errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}
	resp.Close = true
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errCh <- errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := json.Unmarshal(data, response); err != nil {
		errCh <- errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	errCh <- nil
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
