package client

//go:generate stringer -type=Method

// Method is http method for request
type Method int

const (
	POST Method = iota + 1
	GET
	PUT
	PATCH
	DELETE
)
