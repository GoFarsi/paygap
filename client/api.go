package client

// APIConfig use for provider standard external api path address
type APIConfig struct {
	Host    string
	Path    string
	Query   map[string]string // Query use for http get request for set query params example.com/x?q1=v1&q2=v2
	Headers map[string]string // Headers set http header request
}
