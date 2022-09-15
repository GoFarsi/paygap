package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Ja7ad/pgp/errors"
	"google.golang.org/grpc/codes"
	"net/http"
	"sync"
)

type Client struct {
	mu     sync.Mutex
	client *http.Client
}

func (c *Client) Request(ctx context.Context, url, method, contentType string, payload interface{}) ([]byte, *errors.Error) {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return nil, errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, method, url, &buf)
	if err != nil {
		return nil, errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if len(contentType) == 0 {
		contentType = "application/json"
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

}
