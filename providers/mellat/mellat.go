package mellat

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"google.golang.org/grpc/codes"
)

const mellatURL = "https://bpm.shaparak.ir/pgwchannel/services/pgw?wsdl"

// New create mellat object for create new request
func New(client client.Transporter, username, password string) (*Mellat, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}
	mellat := &Mellat{
		client:   client,
		username: username,
		password: password,
		url:      mellatURL,
	}
	if err := client.GetValidator().Struct(mellat); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	return mellat, nil
}

func (m *Mellat) CreateTransaction(ctx context.Context, req *paymentRequest) (*PaymentResponse, error) {
	payload, err := req.raw(m.username, m.password)
	if err != nil {
		return nil, err
	}
	return request[*PaymentResponse](m.client, http.MethodPost, m.url, payload)
}

func (m *Mellat) VerifyTransaction(ctx context.Context, req *verifyRequest) (*VerifyResponse, error) {
	payload, err := req.raw(m.username, m.password)
	if err != nil {
		return nil, err
	}
	return request[*VerifyResponse](m.client, http.MethodPost, m.url, payload)
}
func request[Rs response](
	clientTransporter client.Transporter,
	method, url string,
	body []byte,
) (
	response Rs,
	err error,
) {
	responsePtr := new(Rs)

	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return response, err
	}
	request.Header.Set("Content-Type", "text/xml")
	request.Header.Set("charset", "utf-8")

	resp, err := clientTransporter.GetClient().Do(request)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if err := xml.Unmarshal(responseBody, responsePtr); err != nil {
		return response, fmt.Errorf("error raw response: %v", string(responseBody))
	}

	response = *responsePtr
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return response, status.New(
			response.ResponseCode(), http.StatusFailedDependency, codes.OK,
			fmt.Sprintf("response code %v", response.ResponseCode()),
		)
	}
	if err := response.modifyResponse(); err != nil {
		return response, err
	}
	return response, nil
}
