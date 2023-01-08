package client

import (
	"fmt"
	"net/http"
)

const userAgent = "Go-http-httpClient/1.1"

type transport struct {
	Base http.RoundTripper
}

// RoundTrip appends userAgent existing User-Agent header and performs the request
// via t.Base.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req)
	ua := req.Header.Get("User-Agent")
	if ua == "" {
		ua = userAgent
	} else {
		ua = fmt.Sprintf("%s;%s", ua, userAgent)
	}
	req.Header.Set("User-Agent", ua)
	return t.Base.RoundTrip(req)
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Headers map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Headers
	r2.Header = make(http.Header)
	for k, s := range r.Header {
		r2.Header[k] = s
	}
	return r2
}
