package payping

import "context"

// refrence: https://docs.payping.ir/#operation/TransactionsReport
func (p *Payping) TransactionsList(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/TransactionsReportCount
func (p *Payping) TransactionsDetails(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/WithdrawTransactionsCount
func (p *Payping) TransactionsNumber(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/WithdrawTransactionsReport
func (p *Payping) SettelmentList(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}

// refrence: https://docs.payping.ir/#operation/WithdrawTransactionsCount
func (p *Payping) SettelmentListNumber(ctx context.Context, amount uint, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, error) {
	return &PaymentResponse{}, nil
}
