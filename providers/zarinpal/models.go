package zarinpal

import "github.com/GoFarsi/paygap/client"

type Zarinpal struct {
	client     client.Transporter
	merchantID string `validate:"required"`

	baseUrl            string
	requestEndpoint    string
	verifyEndpoint     string
	unverifiedEndpoint string
}

type paymentRequest struct {
	MerchantID  string                 `json:"merchant_id" validate:"required"`
	Amount      int                    `json:"amount" validate:"gt=5000"`
	Currency    string                 `json:"currency"`
	CallBackURL string                 `json:"callback_url" validate:"url"`
	Description string                 `json:"description"`
	MetaData    map[string]interface{} `json:"metadata"`
}

type verifyRequest struct {
	MerchantID string `json:"merchant_id" validate:"required"`
	Amount     int    `json:"amount" validate:"gt=0"`
	Authority  string `json:"authority"`
}

type unverifiedTransactionsRequest struct {
	MerchantID string `json:"merchant_id" validate:"required"`
}

type floatingShareSettlementRequest struct {
	MerchantID  string                 `json:"merchant_id" validate:"required"`
	Amount      int                    `json:"amount" validate:"gt=5000"`
	CallBackURL string                 `json:"callback_url" validate:"url"`
	Description string                 `json:"description"`
	Wages       []*Wages               `json:"wages"`
	MetaData    map[string]interface{} `json:"metadata"`
}

type PaymentResponse struct {
	Data   *Data    `json:"data"`
	Errors []string `json:"errors"`
}

type VerifyResponse struct {
	Data   *Data    `json:"data"`
	Errors []string `json:"errors"`
}

type UnverifiedTransactionsResponse struct {
	Data *Data `json:"data"`
}

type FloatingShareSettlementResponse struct {
	Wages []*Wages `json:"wages"`
}

type VerifyFloatingShareSettlementResponse struct {
	Data   *Data    `json:"data"`
	Errors []string `json:"errors"`
}

type Data struct {
	Code        int      `json:"code"`
	Message     string   `json:"message"`
	Authority   string   `json:"authority,omitempty"`
	CardHash    string   `json:"card_hash,omitempty"`
	CardPan     string   `json:"card_pan,omitempty"`
	RefID       int      `json:"ref_id,omitempty"`
	FeeType     string   `json:"fee_type,omitempty"`
	Fee         int      `json:"fee,omitempty"`
	Wages       []*Wages `json:"wages,omitempty"`
	Authorities []struct {
		Authority   string `json:"authority"`
		Amount      int    `json:"amount"`
		CallBackURL string `json:"callback_url"`
		Referer     string `json:"referer"`
		Date        string `json:"date"`
	} `json:"authorities,omitempty"`
}

type Wages struct {
	Iban        string `json:"iban"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
