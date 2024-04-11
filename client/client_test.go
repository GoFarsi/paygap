package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	_everyRateLimitDuration = 3 * time.Second
	_preRequest             = 30
)

type Message struct {
	Status  int
	Message string
}

func Test_ClientRequest(t *testing.T) {
	t.Parallel()
	ctx := context.TODO()
	server := serverTest()
	c := client(server.Client())

	t.Run("GET", func(t *testing.T) {
		response := &Message{}
		query := make(map[string]string)
		query["foo"] = "bar"

		resp, err := c.Get(ctx, &APIConfig{Host: server.URL, Query: query})

		assert.Nil(t, err)
		assert.Nil(t, resp.GetJSON(response))
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, query["foo"], resp.GetHttpResponse().Request.URL.Query().Get("foo"))
	})

	t.Run("POST", func(t *testing.T) {
		req := &Message{Status: http.StatusOK, Message: "test"}
		response := &Message{}
		headers := make(map[string]string)
		headers["foo"] = "bar"

		resp, err := c.Post(ctx, &APIConfig{Host: server.URL, Headers: headers}, req)

		assert.Nil(t, err)
		assert.Nil(t, resp.GetJSON(response))
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, headers["foo"], resp.GetHttpResponse().Request.Header.Get("foo"))
	})

}

func Test_ClientWithRateLimitRequest(t *testing.T) {
	ctx := context.TODO()
	server := serverTest()
	request := 50

	c := clientWithRateLimit(server.Client())

	for i := 0; i < request; i++ {
		t.Run("GET", func(t *testing.T) {
			response := &Message{}
			query := make(map[string]string)
			query["foo"] = "bar"

			resp, err := c.Get(ctx, &APIConfig{Host: server.URL, Query: query})

			assert.Nil(t, err)
			assert.Nil(t, resp.GetJSON(response))
			assert.Equal(t, http.StatusOK, response.Status)
			assert.Equal(t, query["foo"], resp.GetHttpResponse().Request.URL.Query().Get("foo"))
		})

		t.Run("POST", func(t *testing.T) {
			req := &Message{Status: http.StatusOK, Message: "test"}
			response := &Message{}
			headers := make(map[string]string)
			headers["foo"] = "bar"

			resp, err := c.Post(ctx, &APIConfig{Host: server.URL, Headers: headers}, req)

			assert.Nil(t, err)
			assert.Nil(t, resp.GetJSON(response))
			assert.Equal(t, http.StatusOK, response.Status)
			assert.Equal(t, headers["foo"], resp.GetHttpResponse().Request.Header.Get("foo"))
		})
	}
}

func serverTest() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case POST.String():
			m := &Message{}
			if err := json.NewDecoder(req.Body).Decode(&m); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
			}
			b, err := json.Marshal(m)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
			}
			rw.Write(b)
		case GET.String():
			b, err := json.Marshal(&Message{
				Status:  http.StatusOK,
				Message: "OK",
			})
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
			}
			rw.Write(b)
		}
	}))
}

func client(client *http.Client) Transporter {
	return New(WithCustomClient(client))
}

func clientWithRateLimit(client *http.Client) Transporter {
	return New(WithCustomClient(client), WithRateLimit(_everyRateLimitDuration, _preRequest))
}
