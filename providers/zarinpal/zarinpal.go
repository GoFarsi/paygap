package zarinpal

import (
	"context"
	"github.com/Ja7ad/pgp/client"
	"github.com/Ja7ad/pgp/status"
	"google.golang.org/grpc/codes"
	"net/http"
)

const (
	ZARINPAL_REQUEST_API_ENDPOINT                = "https://api.zarinpal.com/pg/v4/payment/request.json"
	ZARINPAL_VERIFY_API_ENDPOINT                 = "https://api.zarinpal.com/pg/v4/payment/verify.json"
	ZARINPAL_UNVERIFIED_TRANSACTION_API_ENDPOINT = "https://api.zarinpal.com/pg/v4/payment/unVerified.json"

	ZARINPAL_REQUEST_SANDBOX_API_ENDPOINT = "https://sandbox.zarinpal.com/pg/v4/payment/request.json"
	ZARINPAL_VERIFY_SANDBOX_API_ENDPOINT  = "https://sandbox.zarinpal.com/pg/v4/payment/verify.json"
)

var _ Interface = (*Zarinpal)(nil)

type Interface interface {
	RequestPayment(ctx context.Context, amount int, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, *status.Status)
	VerifyPayment(ctx context.Context, amount int, authority string) (*VerifyResponse, *status.Status)
	UnverifiedTransactions(ctx context.Context) (*UnverifiedTransactionsResponse, *status.Status)
	FloatingShareSettlement(ctx context.Context, amount int, description, callbackUrl string, wages []*Wages, metaData map[string]interface{}) (*FloatingShareSettlementResponse, *status.Status)
	VerifyFloatingShareSettlement(ctx context.Context, amount int, authority string) (*VerifyFloatingShareSettlementResponse, *status.Status)
}

// New create zarinpal provider object for user factory request methods
func New(client *client.Client, merchantID string, sandbox bool) (Interface, *status.Status) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}

	client.GetMutex().Lock()
	defer client.GetMutex().Unlock()

	zarinpal := &Zarinpal{
		client:             client,
		merchantID:         merchantID,
		requestEndpoint:    ZARINPAL_REQUEST_API_ENDPOINT,
		verifyEndpoint:     ZARINPAL_VERIFY_API_ENDPOINT,
		unverifiedEndpoint: ZARINPAL_UNVERIFIED_TRANSACTION_API_ENDPOINT,
	}

	if sandbox {
		zarinpal.requestEndpoint = ZARINPAL_REQUEST_SANDBOX_API_ENDPOINT
		zarinpal.verifyEndpoint = ZARINPAL_VERIFY_SANDBOX_API_ENDPOINT
	}

	if err := client.GetValidator().Struct(zarinpal); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return zarinpal, nil
}

// RequestPayment create payment request and return status code and authority
func (z *Zarinpal) RequestPayment(ctx context.Context, amount int, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, *status.Status) {
	req := &paymentRequest{
		MerchantID:  z.merchantID,
		Amount:      amount,
		Currency:    currency,
		CallBackURL: callBackUrl,
		Description: description,
		MetaData:    metaData,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &PaymentResponse{}
	s := z.client.Request(ctx, z.requestEndpoint, "POST", client.POST, map[string]string{}, req, resp)
	if s.HttpStatusCode == http.StatusInternalServerError {
		return nil, s
	}

	s.ProviderStatusCode = resp.Data.Code

	return resp, s
}

// VerifyPayment transaction by merchant id, amount and authority to payment provider
func (z *Zarinpal) VerifyPayment(ctx context.Context, amount int, authority string) (*VerifyResponse, *status.Status) {
	req := &verifyRequest{
		MerchantID: z.merchantID,
		Amount:     amount,
		Authority:  authority,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &VerifyResponse{}
	s := z.client.Request(ctx, z.requestEndpoint, "POST", client.POST, map[string]string{}, req, resp)
	if s.HttpStatusCode == http.StatusInternalServerError {
		return nil, s
	}

	s.ProviderStatusCode = resp.Data.Code

	return resp, s
}

// UnverifiedTransactions get unverified transactions from provider
func (z *Zarinpal) UnverifiedTransactions(ctx context.Context) (*UnverifiedTransactionsResponse, *status.Status) {
	req := &unverifiedTransactionsRequest{
		MerchantID: z.merchantID,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &UnverifiedTransactionsResponse{}
	s := z.client.Request(ctx, z.requestEndpoint, "POST", client.POST, map[string]string{}, req, resp)
	if s.HttpStatusCode == http.StatusInternalServerError {
		return nil, s
	}

	s.ProviderStatusCode = resp.Data.Code

	return resp, s
}

// FloatingShareSettlement a special method is used for sellers who benefit from an incoming amount,
// more information in https://docs.zarinpal.com/paymentGateway/setshare.html#%D8%AA%D8%B3%D9%88%DB%8C%D9%87-%D8%A7%D8%B4%D8%AA%D8%B1%D8%A7%DA%A9%DB%8C-%D8%B4%D9%86%D8%A7%D9%88%D8%B1
func (z *Zarinpal) FloatingShareSettlement(ctx context.Context, amount int, description, callbackUrl string, wages []*Wages, metaData map[string]interface{}) (*FloatingShareSettlementResponse, *status.Status) {
	req := &floatingShareSettlementRequest{
		MerchantID:  z.merchantID,
		Amount:      amount,
		CallBackURL: callbackUrl,
		Description: description,
		MetaData:    metaData,
		Wages:       wages,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &FloatingShareSettlementResponse{}
	s := z.client.Request(ctx, z.requestEndpoint, "POST", client.POST, map[string]string{}, req, resp)
	if s.HttpStatusCode == http.StatusInternalServerError {
		return nil, s
	}

	return resp, s
}

// VerifyFloatingShareSettlement verify floating share settlement
func (z *Zarinpal) VerifyFloatingShareSettlement(ctx context.Context, amount int, authority string) (*VerifyFloatingShareSettlementResponse, *status.Status) {
	req := &verifyRequest{
		MerchantID: z.merchantID,
		Amount:     amount,
		Authority:  authority,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &VerifyFloatingShareSettlementResponse{}
	s := z.client.Request(ctx, z.requestEndpoint, "POST", client.POST, map[string]string{}, req, resp)
	if s.HttpStatusCode == http.StatusInternalServerError {
		return nil, s
	}

	s.ProviderStatusCode = resp.Data.Code

	return resp, s
}
