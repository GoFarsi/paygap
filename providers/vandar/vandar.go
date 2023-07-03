package vandar

import (
	"context"
	"net/http"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
)

const API_VERSION = "3"

const (
	VANDAR_HOST = "https://ipg.vandar.io"
)

const (
	VANDAR_REQUEST_API_ENDPOINT             = "/api/v3/send"
	VANDAR_REDIRECT_API_ENDPOINT            = "/v3/"
	VANDAR_VERIFY_API_ENDPOINT              = "/api/v3/verify"
	VANDAR_TRANSACTION_DETAIL_API_END_POINT = "/api/v3/transaction"
)

// New create vandar provider object for user factory request methods
func New(client client.Transporter, ApiKey string) (*Vandar, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}

	vandar := &Vandar{
		client:                    client,
		APIKey:                    ApiKey,
		baseUrl:                   VANDAR_HOST,
		requestEndpoint:           VANDAR_REQUEST_API_ENDPOINT,
		redirectEndpoint:          VANDAR_REDIRECT_API_ENDPOINT,
		verifyEndpoint:            VANDAR_VERIFY_API_ENDPOINT,
		transactionDetailEndpoint: VANDAR_TRANSACTION_DETAIL_API_END_POINT,
	}

	if err := client.GetValidator().Struct(vandar); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return vandar, nil
}

// RequestPayment create payment request and return status code and authority
// document of field https://vandarpay.github.io/docs/ipg/#step-1
func (v *Vandar) RequestPayment(ctx context.Context, amount int, callBackUrl, mobileNumber, factorNumber, nationalCode, validCardNumber, description string) (*PaymentResponse, error) {
	req := paymentRequest{
		ApiKey:          v.APIKey,
		Amount:          amount,
		CallBackURL:     callBackUrl,
		MobileNumber:    mobileNumber,
		FactorNumber:    factorNumber,
		NationalCode:    nationalCode,
		ValidCardNumber: validCardNumber,
		Description:     description,
	}

	if err := v.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	response := &PaymentResponse{}
	resp, err := v.client.Post(ctx, &client.APIConfig{Host: v.baseUrl, Path: v.requestEndpoint, Headers: headers}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return response, nil
}

// VerifyPayment transaction by merchant id, amount and authority to payment provider
// doc https://vandarpay.github.io/docs/ipg/#step-4
func (v *Vandar) VerifyPayment(ctx context.Context, token string) (*VerifyResponse, error) {
	req := &verifyRequest{
		APIKey: v.APIKey,
		Token:  token,
	}

	if err := v.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	response := &VerifyResponse{}
	resp, err := v.client.Post(ctx, &client.APIConfig{Host: v.baseUrl, Path: v.verifyEndpoint, Headers: headers}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return response, nil
}

// TransactionDetail Transaction Detail get Transaction status
// doc https://vandarpay.github.io/docs/ipg/#step-3
func (v *Vandar) TransactionDetail(ctx context.Context, token string) (*TransactionDetailResponse, error) {
	req := &transactionDetailRequest{
		APIKey: v.APIKey,
		Token:  token,
	}

	if err := v.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	response := &TransactionDetailResponse{}
	resp, err := v.client.Post(ctx, &client.APIConfig{Host: v.baseUrl, Path: v.transactionDetailEndpoint, Headers: headers}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return response, nil
}
