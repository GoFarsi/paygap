package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	_everyRateLimitDuration = 10 * time.Second
	_preRequest             = 50
)

type Message struct {
	Status  int
	Message string
}

func TestClient_Request(t *testing.T) {
	t.Parallel()
	ctx := context.TODO()
	server := serverTest()

	c := client(server.Client())

	t.Run("GET", func(t *testing.T) {
		resp := &Message{}
		s := c.Request(ctx, server.URL, "", GET, map[string]string{}, nil, resp)
		if s.HttpStatusCode != http.StatusOK {
			t.Fatalf("want status code ok, but got %d and err %v", s.HttpStatusCode, s.Error())
		}
		t.Log(s, resp)
	})

	t.Run("POST", func(t *testing.T) {
		req := &Message{Status: http.StatusOK, Message: "test"}
		resp := &Message{}

		s := c.Request(ctx, server.URL, "", POST, map[string]string{}, req, resp)
		if s.HttpStatusCode != http.StatusOK {
			t.Fatalf("want status code ok, but got %d and err %v", s.HttpStatusCode, s.Error())
		}
		t.Log(s, resp)
	})

}

func TestClientWithRateLimit_Request(t *testing.T) {
	ctx := context.TODO()
	server := serverTest()
	request := 50

	c := clientWithRateLimit(server.Client())

	for i := 0; i < request; i++ {
		s := c.Request(ctx, server.URL, "", GET, map[string]string{}, nil, nil)
		if s.HttpStatusCode != http.StatusOK {
			t.Fatalf("want status code ok, but got %d and err %v", s.HttpStatusCode, s.Error())
		}
		t.Log(s)
		i++
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
