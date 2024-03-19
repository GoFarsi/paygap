package zibal

import "github.com/GoFarsi/paygap/client"

type Zibal struct {
	client   client.Transporter
	merchant string `validate:"required"`

	baseUrl         string
	requestEndpoint string
	verifyEndpoint  string
}

type paymentRequest struct {
	Merchant    string `json:"merchant" validate:"required"`
	Amount      uint   `json:"amount" validate:"required,min=1000"`
	CallbackURL string `json:"callbackUrl" validate:"required,url"`
	Description string `json:"description"`
}

type PaymentResponse struct {
	Result  int    `json:"result"`
	TrackID int    `json:"trackId"`
	Message string `json:"message"`
}

type VerificationRequest struct {
	Merchant string `json:"merchant" validate:"required"`
	TrackID  int    `json:"trackId" validate:"required"`
}

type VerificationResponse struct {
	PaidAt      string `json:"paidAt"`
	Amount      int    `json:"amount"`
	Result      int    `json:"result"`
	Status      int    `json:"status"`
	RefNumber   int    `json:"refNumber"`
	Description string `json:"description"`
	CardNumber  string `json:"cardNumber"`
	OrderID     string `json:"orderId"`
	Message     string `json:"message"`
}
