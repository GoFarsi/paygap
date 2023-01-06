package pay

import "github.com/GoFarsi/paygap/client"

type Pay struct {
	client client.Transporter
	apiKey string

	host            string
	paymentEndpoint string
	verifyEndpoint  string
}

type ErrorResponse struct {
	Status       int    `json:"status"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type Request struct {
	API string `json:"api"`
	*PaymentRequest
	*VerifyRequest
}

type PaymentRequest struct {
	Amount          uint   `json:"amount" validate:"required,min=10000"`
	Redirect        string `json:"redirect" validate:"required,url"`
	Mobile          string `json:"mobile,omitempty" validate:"omitempty,min=11"`
	FactorNumber    string `json:"factorNumber,omitempty"`
	Description     string `json:"description,omitempty" validate:"omitempty,max=255"`
	ValidCardNumber string `json:"validCardNumber,omitempty" validate:"omitempty,max=16"`
}

type PaymentResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

type VerifyRequest struct {
	Token string `json:"token" validate:"required"`
}

type VerifyResponse struct {
	Status       int    `json:"status"`
	Amount       string `json:"amount"`
	TransId      int    `json:"transId"`
	FactorNumber string `json:"factorNumber"`
	Mobile       string `json:"mobile"`
	Description  string `json:"description"`
	CardNumber   string `json:"cardNumber"`
	Message      string `json:"message"`
}
