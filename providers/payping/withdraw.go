package payping

import "context"

func (p *Payping) SettlementRequest(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

func (p *Payping) SettlementDetails(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}
