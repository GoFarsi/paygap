package pay

import (
	"context"
	"errors"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
	"net/http"
	"reflect"
	"strconv"
)

const (
	PAY_HOST = "https://pay.ir"

	PAYMENT_ENDPOINT = "/pg/send"
	VERIFY_ENDPOINT  = "/pg/verify"
)

// CallBackFunc is a function for create call back action
type CallBackFunc func(ctx context.Context, status int, token string) error

// New create pay object for create and verify transaction in pay.ir service
func New(client client.Transporter, apiKey string, sandbox bool) (*Pay, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}

	pay := new(Pay)

	pay.client = client
	pay.host = PAY_HOST
	pay.apiKey = apiKey
	pay.paymentEndpoint = PAYMENT_ENDPOINT
	pay.verifyEndpoint = VERIFY_ENDPOINT

	if sandbox {
		pay.apiKey = "test"
	}

	if err := client.GetValidator().Struct(pay); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	return pay, nil
}

// CreateTransaction create new transaction and return status and token
func (p *Pay) CreateTransaction(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	gatewayReq := new(Request)
	gatewayReq.API = p.apiKey
	gatewayReq.PaymentRequest = req

	return request[*Request, *PaymentResponse](ctx, p, gatewayReq, p.paymentEndpoint)
}

// VerifyTransaction payment transaction via token
func (p *Pay) VerifyTransaction(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	if err := p.client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	gatewayReq := new(Request)
	gatewayReq.API = p.apiKey
	gatewayReq.VerifyRequest = req

	return request[*Request, *VerifyResponse](ctx, p, gatewayReq, p.verifyEndpoint)
}

// DefaultCallback is default callback for http server handler,
// example :
//
//	        http.Handle("/callback", pay.DefaultCallback(callback))
//		    http.ListenAndServe(":8080", nil)
func DefaultCallback(actionFunc CallBackFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, _ := strconv.Atoi(r.URL.Query().Get("status"))
		token := r.URL.Query().Get("token")

		if err := actionFunc(r.Context(), statusCode, token); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func request[RQ any, RS any](ctx context.Context, p *Pay, req RQ, endpoint string) (response RS, err error) {
	r, ok := reflect.New(reflect.TypeOf(response).Elem()).Interface().(RS)
	if !ok {
		return response, errors.New("response type is invalid")
	}

	errResp := new(ErrorResponse)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	resp, err := p.client.Post(ctx, &client.APIConfig{Host: p.host, Path: endpoint}, headers, req)
	if err != nil {
		return response, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if resp.GetHttpResponse().StatusCode != http.StatusOK {
		if err := resp.GetJSON(errResp); err != nil {
			return response, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
		}
		return response, status.New(errResp.ErrorCode, http.StatusFailedDependency, codes.OK, errResp.ErrorMessage)
	}

	if err := resp.GetJSON(r); err != nil {
		return response, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	return r, nil
}
