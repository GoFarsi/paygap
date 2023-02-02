package payping

import (
	"context"
	"net/http"

	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
)

// refrence: https://docs.payping.ir/#operation/CreateSinglePayment
func (p *Payping) RequestPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return request[*PaymentRequest, *PaymentResponse](ctx, p, req, p.baseUrl, p.paymentEndpoint, nil)
}

// refrence: https://docs.payping.ir/#operation/VerifyPayment
func (p *Payping) VerifyPayment(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	return request[*VerifyRequest, *VerifyResponse](ctx, p, req, p.baseUrl, p.paymentEndpoint, nil)
}

// refrence: https://docs.payping.ir/#operation/CreateMultiPayment
func (p *Payping) RequestSharePayment(ctx context.Context, req *SharePaymentRequest) (*PaymentResponse, error) {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	return request[*SharePaymentRequest, *PaymentResponse](ctx, p, req, p.baseUrl, p.paymentEndpoint, nil)
}

// refrence: https://docs.payping.ir/#operation/CreateBlockPayment
func (p *Payping) RequestBlockingPayment(ctx context.Context, req *BlockedPaymentRequest) (*PaymentResponse, error) {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	return request[*BlockedPaymentRequest, *PaymentResponse](ctx, p, req, p.baseUrl, p.paymentEndpoint, nil)
}

// refrence: https://docs.payping.ir/#operation/UnBlockPayment
func (p *Payping) ReleasingBlockedPayment(ctx context.Context, req *ReleasingBlockedPaymentRequest) error {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	_, err := request[*ReleasingBlockedPaymentRequest, error](ctx, p, req, p.baseUrl, p.paymentEndpoint, nil)
	return err
}

// refrence: https://docs.payping.ir/#operation/CreateIdPosPayment
func (p *Payping) PaymentWithTracingId(ctx context.Context, req *PaymentWithTracerIdRequest) (*PaymentWithTracerIdResponse, error) {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	return request[*PaymentWithTracerIdRequest, *PaymentWithTracerIdResponse](ctx, p, req, p.baseUrl, p.paymentEndpoint, nil)
}

// refrence: https://docs.payping.ir/#operation/CancelPayment
func (p *Payping) PaymentSuspending(ctx context.Context, req *PaymentSuspedingRequest) error {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	_, err := request[*PaymentSuspedingRequest, error](ctx, p, req, p.baseUrl, p.paymentEndpoint, nil)
	return err
}
