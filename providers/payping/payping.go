package payping

import (
	"github.com/GoFarsi/paygap/client"
)

type Payping struct {
	client     client.Transporter
	merchantID string `validate:"required"`

	baseUrl            string
	requestEndpoint    string
	verifyEndpoint     string
	unverifiedEndpoint string
}
