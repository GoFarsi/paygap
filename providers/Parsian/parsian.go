package Parsian

import (
	"context"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
	"net/http"
)

const API_VERSION = "1"

const (
	PARSIAN_HOST                         = "https://pec.shaparak.ir"
	PARSIAN_PAYMENT_REQUEST_API_ENDPOINT = "/NewIPGServices/Sale/SaleService.asmx"
	PARSIAN_VERIFY_API_ENDPOINT          = "/NewIPGServices/Confirm/ConfirmService.asmx"
	PARSIAN_REVERSAL_API_ENDPOINT        = "/NewIPGServices/Reverse/ReversalService.asmx"
)

func New(client client.Transporter, LoginAccount string) (*Parsian, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}

	parsian := &Parsian{
		client:          client,
		LoginAccount:    LoginAccount,
		baseUrl:         PARSIAN_HOST,
		paymentEndpoint: PARSIAN_PAYMENT_REQUEST_API_ENDPOINT,
		verifyEndpoint:  PARSIAN_VERIFY_API_ENDPOINT,
		reverseEndpoint: PARSIAN_REVERSAL_API_ENDPOINT,
	}

	if err := client.GetValidator().Struct(parsian); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return parsian, nil
}

// RequestPayment create payment request
func (p *Parsian) RequestPayment(ctx context.Context, amount uint, callBackUrl string, OrderId int, Originator string, AdditionalData map[string]interface{}) (*PaymentResponse, error) {
	req := &PaymentRequest{
		LoginAccount:   p.LoginAccount,
		Amount:         amount,
		OrderId:        OrderId,
		CallBackURL:    callBackUrl,
		Originator:     Originator,
		AdditionalData: AdditionalData,
	}

	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	response := &PaymentResponse{}
	resp, err := p.client.Post(ctx, &client.APIConfig{Host: p.baseUrl, Path: p.paymentEndpoint}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return response, nil
}

// VerifyPayment transaction
func (p *Parsian) VerifyPayment(ctx context.Context, token string) (*VerifyResponse, error) {
	req := &VerifyRequest{
		LoginAccount: p.LoginAccount,
		Token:        token,
	}

	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	response := &VerifyResponse{}
	resp, err := p.client.Post(ctx, &client.APIConfig{Host: p.baseUrl, Path: p.verifyEndpoint}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}
	if response.RRN <= 0 || response.Status != 0 {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, "verify failed")
	}
	return response, nil
}

func (p *Parsian) ReversePayment(ctx context.Context, token string) (*ReversalResponse, error) {
	req := &ReversalRequest{
		LoginAccount: p.LoginAccount,
		Token:        token,
	}

	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	response := &ReversalResponse{}
	resp, err := p.client.Post(ctx, &client.APIConfig{Host: p.baseUrl, Path: p.reverseEndpoint}, req)
	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return response, nil
}
