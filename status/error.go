package status

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

var (
	// ERR_CLIENT_IS_NIL client is nil object, before use provider need initiate client
	ERR_CLIENT_IS_NIL = New(-1, http.StatusInternalServerError, codes.Internal, "client is nil")
)
