package payping

import (
	"context"
	"net/http"

	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
)

// refrence: https://docs.payping.ir/#operation/CreateSinglePayment
func (p *Payping) RequestPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	// return &PaymentResponse{}, nil
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	gatewayReq := new(Request)
	gatewayReq.API = p.apiKey
	gatewayReq.PaymentRequest = req

	return request[*PaymentRequest, *PaymentResponse](ctx, p, gatewayReq, p.paymentEndpoint)
}

// refrence: https://docs.payping.ir/#operation/VerifyPayment
func (p *Payping) VerifyPayment(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	return &VerifyResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/CreateMultiPayment
func (p *Payping) RequestSharePayment(ctx context.Context, req *SharePaymentRequest) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/CreateBlockPayment
func (p *Payping) RequestBlockingPayment(ctx context.Context, req *BlockedPaymentRequest) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/UnBlockPayment
func (p *Payping) ReleasingBlockedPayment(ctx context.Context, req *ReleasingBlockedPaymentRequest) error {
	return nil
}

// refrence: https://docs.payping.ir/#operation/CreateIdPosPayment
func (p *Payping) PaymentWithTracingId(ctx context.Context, req *PaymentWithTracerIdRequest) (*PaymentWithTracerIdResponse, error) {
	return &PaymentWithTracerIdResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/CancelPayment
func (p *Payping) PaymentSuspending(ctx context.Context, req *PaymentSuspedingRequest) error {
	return nil
}
