package payping

import (
	"errors"
	"log"
	"net/http"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
)

const (
	PAYPING_HOST = "https://api.payping.ir"
)

const (
	PAYPING_REQUEST_API_ENDPOINT               = "/v2/pay"
	PAYPING_VERIFY_API_ENDPOINT                = "/v2/pay/verify"
	PAYPING_MULTI_PAYMENT_API_ENDPOINT         = "/v2/pay/multi"
	PAYPING_BLOCK_MONEY_PAYMENT_API_ENDPOINT   = "/v2/pay/BlockMoney"
	PAYPING_UNBLOCK_MONEY_PAYMENT_API_ENDPOINT = "/v2/pay/UnBlockMoney"
	PAYPING_POS_PAYMENT_API_ENDPOINT           = "/v1/pos"
)

// New create payping provider object
func New(client client.Transporter, apiToken string) (*Payping, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}

	payping := &Payping{
		client:   client,
		apiToken: apiToken,

		baseUrl:                     PAYPING_HOST,
		paymentEndpoint:             PAYPING_REQUEST_API_ENDPOINT,
		verifyEndpoint:              PAYPING_VERIFY_API_ENDPOINT,
		multiplePaymentEndpoint:     PAYPING_MULTI_PAYMENT_API_ENDPOINT,
		blockMoneyPaymentEndpoint:   PAYPING_BLOCK_MONEY_PAYMENT_API_ENDPOINT,
		unBlockMoneyPaymentEndpoint: PAYPING_UNBLOCK_MONEY_PAYMENT_API_ENDPOINT,
		posEndpoint:                 PAYPING_POS_PAYMENT_API_ENDPOINT,
	}

	if payping.apiToken == "" || len(payping.apiToken) < 10 {
		log.Fatal("jwt token is invalid")
		return nil, errors.New("jwt token is invalid")
	}

	if err := client.GetValidator().Struct(payping); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return payping, nil
}
