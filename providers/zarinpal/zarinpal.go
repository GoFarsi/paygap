package zarinpal

import (
	"context"
	"github.com/Ja7ad/pgp/client"
	"github.com/Ja7ad/pgp/errors"
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

var _ FactoryRequest = (*Zarinpal)(nil)

type FactoryRequest interface {
	RequestPayment(ctx context.Context, amount int, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, *errors.Error)
	VerifyPayment(ctx context.Context, amount int, authority string) (*VerifyResponse, *errors.Error)
	UnverifiedTransactions(ctx context.Context) (*UnverifiedTransactionsResponse, *errors.Error)
	FloatingShareSettlement(ctx context.Context, amount int, description, callbackUrl string, wages []*Wages, metaData map[string]interface{}) (*FloatingShareSettlementResponse, *errors.Error)
	VerifyFloatingShareSettlement(ctx context.Context, amount int, authority string) (*VerifyFloatingShareSettlementResponse, *errors.Error)
}

// New create zarinpal provider object for user factory request methods
func New(client *client.Client, merchantID string, sandbox bool) (FactoryRequest, *errors.Error) {
	if client == nil {
		return nil, errors.New(0, http.StatusInternalServerError, codes.Internal, "client is nil")
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
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return zarinpal, nil
}

// RequestPayment create payment request and return status code and authority
func (z *Zarinpal) RequestPayment(ctx context.Context, amount int, callBackUrl, currency, description string, metaData map[string]interface{}) (*PaymentResponse, *errors.Error) {
	errCh := make(chan *errors.Error)
	req := &paymentRequest{
		MerchantID:  z.merchantID,
		Amount:      amount,
		Currency:    currency,
		CallBackURL: callBackUrl,
		Description: description,
		MetaData:    metaData,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &PaymentResponse{}
	go z.client.Request(ctx, z.requestEndpoint, "POST", "", map[string]string{}, req, resp, errCh)

	select {
	case err := <-errCh:
		close(errCh)
		return resp, err
	case <-ctx.Done():
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, ctx.Err().Error())
	}
}

// VerifyPayment transaction by merchant id, amount and authority to payment provider
func (z *Zarinpal) VerifyPayment(ctx context.Context, amount int, authority string) (*VerifyResponse, *errors.Error) {
	errCh := make(chan *errors.Error)
	req := &verifyRequest{
		MerchantID: z.merchantID,
		Amount:     amount,
		Authority:  authority,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &VerifyResponse{}
	go z.client.Request(ctx, z.requestEndpoint, "POST", "", map[string]string{}, req, resp, errCh)

	select {
	case err := <-errCh:
		close(errCh)
		return resp, err
	case <-ctx.Done():
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, ctx.Err().Error())
	}
}

// UnverifiedTransactions get unverified transactions from provider
func (z *Zarinpal) UnverifiedTransactions(ctx context.Context) (*UnverifiedTransactionsResponse, *errors.Error) {
	errCh := make(chan *errors.Error)
	req := &unverifiedTransactionsRequest{
		MerchantID: z.merchantID,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &UnverifiedTransactionsResponse{}
	go z.client.Request(ctx, z.requestEndpoint, "POST", "", map[string]string{}, req, resp, errCh)

	select {
	case err := <-errCh:
		close(errCh)
		return resp, err
	case <-ctx.Done():
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, ctx.Err().Error())
	}
}

// FloatingShareSettlement a special method is used for sellers who benefit from an incoming amount,
// more information in https://docs.zarinpal.com/paymentGateway/setshare.html#%D8%AA%D8%B3%D9%88%DB%8C%D9%87-%D8%A7%D8%B4%D8%AA%D8%B1%D8%A7%DA%A9%DB%8C-%D8%B4%D9%86%D8%A7%D9%88%D8%B1
func (z *Zarinpal) FloatingShareSettlement(ctx context.Context, amount int, description, callbackUrl string, wages []*Wages, metaData map[string]interface{}) (*FloatingShareSettlementResponse, *errors.Error) {
	errCh := make(chan *errors.Error)
	req := &floatingShareSettlementRequest{
		MerchantID:  z.merchantID,
		Amount:      amount,
		CallBackURL: callbackUrl,
		Description: description,
		MetaData:    metaData,
		Wages:       wages,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &FloatingShareSettlementResponse{}
	go z.client.Request(ctx, z.requestEndpoint, "POST", "", map[string]string{}, req, resp, errCh)

	select {
	case err := <-errCh:
		close(errCh)
		return resp, err
	case <-ctx.Done():
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, ctx.Err().Error())
	}
}

// VerifyFloatingShareSettlement verify floating share settlement
func (z *Zarinpal) VerifyFloatingShareSettlement(ctx context.Context, amount int, authority string) (*VerifyFloatingShareSettlementResponse, *errors.Error) {
	errCh := make(chan *errors.Error)
	req := &verifyRequest{
		MerchantID: z.merchantID,
		Amount:     amount,
		Authority:  authority,
	}

	if err := z.client.GetValidator().Struct(req); err != nil {
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp := &VerifyFloatingShareSettlementResponse{}
	go z.client.Request(ctx, z.requestEndpoint, "POST", "", map[string]string{}, req, resp, errCh)

	select {
	case err := <-errCh:
		close(errCh)
		return resp, err
	case <-ctx.Done():
		return nil, errors.New(0, http.StatusBadRequest, codes.InvalidArgument, ctx.Err().Error())
	}
}
