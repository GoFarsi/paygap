package status

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Status struct {
	ProviderStatusCode any        `json:"provider_status_code" xml:"provider_status_code"`
	GrpcStatusCode     codes.Code `json:"grpc_status_code" xml:"grpc_status_code"`
	HttpStatusCode     int        `json:"http_status_code" xml:"http_status_code"`
	Message            string     `json:"message" xml:"message"`
}

// New create new pgp status
func New(providerStatusCode any, httpStatusCode int, grpcCode codes.Code, message string, params ...interface{}) *Status {
	return &Status{
		ProviderStatusCode: providerStatusCode,
		HttpStatusCode:     httpStatusCode,
		GrpcStatusCode:     grpcCode,
		Message:            fmt.Sprintf(message, params...),
	}
}

// Status return type status
func (e *Status) Error() string {
	return e.Message
}

// GrpcStatus create grpc status
func (e *Status) GrpcStatus() *status.Status {
	return status.New(e.GrpcStatusCode, e.Message)
}

// HttpError create http status
func (e *Status) HttpError(w http.ResponseWriter) {
	http.Error(w, e.Message, e.HttpStatusCode)
}
