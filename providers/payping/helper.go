package payping

import (
	"context"
	"errors"
	"net/http"
	"reflect"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
)

func request[RQ any, RS any](ctx context.Context, payping *Payping, req RQ, baseUrl string, endpoint string, queryParams map[string]string) (response RS, err error) {
	r, ok := reflect.New(reflect.TypeOf(response).Elem()).Interface().(RS)
	if !ok {
		return response, errors.New("response type is invalid")
	}

	headers := make(map[string]string)
	headers["X-API-KEY"] = payping.apiKey
	headers["Content-Type"] = "application/json"

	// TODO: can review if SANDBOX was available
	// if i.sandbox {
	// 	headers["X-SANDBOX"] = "1"
	// }

	errResp := &ErrorResponse{}
	resp, err := payping.client.Post(ctx, &client.APIConfig{Host: baseUrl, Path: endpoint, Headers: headers, Query: queryParams}, req)
	if err != nil {
		return response, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if resp.GetHttpResponse().StatusCode != http.StatusOK|http.StatusCreated {
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
