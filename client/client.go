package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
)

var _ Transporter = (*Client)(nil)

type Client struct {
	httpClient *http.Client
	validator  *validator.Validate
	rate       *rate.Limiter
}

type Response struct {
	resp *http.Response
}

type Transporter interface {
	GetClient() *http.Client
	GetValidator() *validator.Validate
	Get(ctx context.Context, apiConfig *APIConfig) (*Response, error)
	Post(ctx context.Context, apiConfig *APIConfig, apiRequest any) (*Response, error)
}

// New create httpClient constructor
func New(opts ...Option) Transporter {
	client := &Client{
		validator: validator.New(),
	}
	for _, opt := range opts {
		opt(client)
	}

	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
	}

	return client
}

func (c *Client) GetClient() *http.Client {
	return c.httpClient
}

func (c *Client) GetValidator() *validator.Validate {
	return c.validator
}

// Get do get request and return response
func (c *Client) Get(ctx context.Context, apiConfig *APIConfig) (*Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(GET.String(), apiConfig.Host+apiConfig.Path, nil)
	if err != nil {
		return nil, err
	}

	if len(apiConfig.Headers) != 0 {
		for k, v := range apiConfig.Headers {
			req.Header.Add(k, v)
		}
	}

	if len(apiConfig.Query) != 0 {
		req.URL.RawQuery = c.queryBuilder(apiConfig.Query)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}

	return &Response{resp}, nil
}

// Post do post request and return response
func (c *Client) Post(ctx context.Context, apiConfig *APIConfig, apiRequest any) (*Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	body, err := json.Marshal(apiRequest)
	if err != nil {
		return nil, err
	}

	url, err := url.JoinPath(apiConfig.Host, apiConfig.Path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(POST.String(), url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if len(apiConfig.Headers) != 0 {
		for k, v := range apiConfig.Headers {
			req.Header.Add(k, v)
		}
	}

	if len(apiConfig.Query) != 0 {
		req.URL.RawQuery = c.queryBuilder(apiConfig.Query)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}

	return &Response{resp}, nil
}

// GetJSON decode response body to your response
func (r *Response) GetJSON(response any) error {
	defer r.resp.Body.Close()
	if err := json.NewDecoder(r.resp.Body).Decode(response); err != nil {
		return err
	}
	return nil
}

// GetHttpResponse return http response
func (r *Response) GetHttpResponse() *http.Response {
	return r.resp
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}

	return client.Do(req.WithContext(ctx))
}

func (c *Client) awaitRateLimiter(ctx context.Context) error {
	if c.rate == nil {
		return nil
	}
	return c.rate.Wait(ctx)
}

func (c *Client) queryBuilder(params map[string]string) string {
	query := url.Values{}
	for k, v := range params {
		query[k] = []string{v}
	}
	return query.Encode()
}
