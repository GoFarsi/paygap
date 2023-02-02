package payping

import (
	"github.com/GoFarsi/paygap/client"
)

type Payping struct {
	client client.Transporter

	baseUrl                     string
	paymentEndpoint             string
	verifyEndpoint              string
	multiplePaymentEndpoint     string
	blockMoneyPaymentEndpoint   string
	unBlockMoneyPaymentEndpoint string
	posEndpoint                 string

	apiToken string
}
