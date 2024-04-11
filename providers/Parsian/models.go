package Parsian

import "github.com/GoFarsi/paygap/client"

type Parsian struct {
	client          client.Transporter
	LoginAccount    string `validate:"required"`
	baseUrl         string
	paymentEndpoint string
	verifyEndpoint  string
	reverseEndpoint string
}
type PaymentRequest struct {
	LoginAccount   string                 `json:"loginAccount" validate:"required"`
	Amount         uint                   `json:"amount" validate:"required,min=1000"`
	OrderId        int                    `json:"orderId" validate:"required"`
	CallBackURL    string                 `json:"callbackUrl" validate:"required,url"`
	Originator     string                 `json:"originator"`
	AdditionalData map[string]interface{} `json:"additionalData"`
}
type ReversalRequest struct {
	LoginAccount string `json:"loginAccount" validate:"required"`
	Token        string `json:"token" validate:"required"`
}
type VerifyRequest struct {
	LoginAccount string `json:"loginAccount" validate:"required"`
	Token        string `json:"token" validate:"required"`
}

type ReversalResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}
type PaymentResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

type VerifyResponse struct {
	Status           int    `json:"status"`
	RRN              int    `json:"RRN"`
	CardNumberMasked string `json:"cardNumberMasked"`
	Token            string `json:"token"`
}
