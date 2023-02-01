package payping

import "context"

// RequestPayment create payment request and return status code and authority
// refrence: https://docs.payping.ir/#operation/CreateSinglePayment
func (p *Payping) RequestPayment(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

// VerifyPayment transaction by merchant id, amount and authority to payment provider
func (p *Payping) VerifyPayment(ctx context.Context, amount uint, authority string) (*VerifyResponse, error) {
	return &VerifyResponse{}, nil
}

func (p *Payping) RequestSharePayment(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

func (p *Payping) SharePaymentSettlement(ctx context.Context, amount uint, description, callbackUrl string, wages []*Wages, metaData map[string]interface{}) (*FloatingShareSettlementResponse, error) {
	return &FloatingShareSettlementResponse{}, nil
}

func (p *Payping) RequestBlockingPayment(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

func (p *Payping) ReleasingBlockingPayment(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

func (p *Payping) PaymentWithTracingId(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}
