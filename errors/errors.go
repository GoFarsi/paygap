package errors

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Error struct {
	ProviderStatusCode int
	GrpcStatusCode     codes.Code
	HttpStatusCode     int
	Message            string
}

// New create new pgp error
func New(providerStatusCode, httpStatusCode int, grpcCode codes.Code, message string, params ...interface{}) *Error {
	return &Error{
		ProviderStatusCode: providerStatusCode,
		HttpStatusCode:     httpStatusCode,
		GrpcStatusCode:     grpcCode,
		Message:            fmt.Sprintf(message, params...),
	}
}

// Error return type error
func (e *Error) Error() error {
	return errors.New(e.Message)
}

// GrpcStatusError create grpc error
func (e *Error) GrpcStatusError() *status.Status {
	return status.New(e.GrpcStatusCode, e.Message)
}

// HttpError create http error
func (e *Error) HttpError(w http.ResponseWriter) {
	http.Error(w, e.Message, e.HttpStatusCode)
}
