package zibal

import (
	"context"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"net/http"

	"google.golang.org/grpc/codes"
)

const zibalBaseURL = "https://gateway.zibal.ir"
const (
	zibalRequestEndpoint = "/v1/request"
	zibalVerifyEndpoint  = "/v1/verify"
)

// New creates a zibal provider object for user factory request methods
func New(client client.Transporter, merchant string) (*Zibal, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}

	zibal := &Zibal{
		client:          client,
		merchant:        merchant,
		baseUrl:         zibalBaseURL,
		requestEndpoint: zibalRequestEndpoint,
		verifyEndpoint:  zibalVerifyEndpoint,
	}
	if err := client.GetValidator().Struct(zibal); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return zibal, nil
}

// RequestPayment creates a payment request and returns a status code and authority.
func (z *Zibal) RequestPayment(ctx context.Context, amount uint, callBackUrl, description string) (*PaymentResponse, error) {
	req := &paymentRequest{
		Merchant:    z.merchant,
		Amount:      amount,
		CallbackURL: callBackUrl,
		Description: description,
	}
	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	response := &PaymentResponse{}
	resp, err := z.client.Post(ctx, &client.APIConfig{Host: z.baseUrl, Path: z.requestEndpoint, Headers: map[string]string{"Content-Type": "application/json"}}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return response, nil
}

// VerifyPayment verifies a payment and returns the payment details.
func (z *Zibal) VerifyPayment(ctx context.Context, trackID int) (*VerificationResponse, error) {
	req := &VerificationRequest{
		Merchant: z.merchant,
		TrackID:  trackID,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	response := &VerificationResponse{}
	resp, err := z.client.Post(ctx, &client.APIConfig{Host: z.baseUrl, Path: z.verifyEndpoint, Headers: map[string]string{"Content-Type": "application/json"}}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return response, nil
}
