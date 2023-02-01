package payping

import (
	"net/http"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
)

const API_VERSION = "2"

const (
	PAYPING_HOST = "https://api.payping.ir/v2/pay"
)

const (
	PAYPING_REQUEST_API_ENDPOINT                = "/pg/v4/payment/request.json"
	PAYPING_VERIFY_API_ENDPOINT                 = "/pg/v4/payment/verify.json"
	PAYPING_UNVERIFIED_TRANSACTION_API_ENDPOINT = "/pg/v4/payment/unVerified.json"
)

// New create payping provider object for user factory request methods
func New(client client.Transporter, merchantID string, sandbox bool) (*Payping, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}

	payping := &Payping{
		client:             client,
		merchantID:         merchantID,
		baseUrl:            PAYPING_HOST,
		requestEndpoint:    PAYPING_REQUEST_API_ENDPOINT,
		verifyEndpoint:     PAYPING_VERIFY_API_ENDPOINT,
		unverifiedEndpoint: PAYPING_UNVERIFIED_TRANSACTION_API_ENDPOINT,
	}

	if err := client.GetValidator().Struct(payping); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return payping, nil
}
