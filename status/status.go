package status

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Status struct {
	ProviderStatusCode int
	GrpcStatusCode     codes.Code
	HttpStatusCode     int
	Message            string
}

// New create new pgp status
func New(providerStatusCode, httpStatusCode int, grpcCode codes.Code, message string, params ...interface{}) *Status {
	return &Status{
		ProviderStatusCode: providerStatusCode,
		HttpStatusCode:     httpStatusCode,
		GrpcStatusCode:     grpcCode,
		Message:            fmt.Sprintf(message, params...),
	}
}

// Status return type status
func (e *Status) Error() error {
	return errors.New(e.Message)
}

// GrpcStatusError create grpc status
func (e *Status) GrpcStatusError() *status.Status {
	return status.New(e.GrpcStatusCode, e.Message)
}

// HttpError create http status
func (e *Status) HttpError(w http.ResponseWriter) {
	http.Error(w, e.Message, e.HttpStatusCode)
}
