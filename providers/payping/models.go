package payping

type PaymentRequest struct {
	Amount        int32  `json:"amount" validate:"required,min=100,max=50000000"`
	PayerIdentity string `json:"payerIdentity"`
	PayerName     string `json:"payerName" validate:"required"`
	Description   string `json:"description"`
	ReturnUrl     string `json:"returnUrl" validate:"required,url"`
	ClientRefId   string `json:"clientRefId"`
}

type PaymentResponse struct {
	Code string `json:"code"`
}

type VerifyRequest struct {
	RefId  string `json:"refId" validate:"required"`
	Amount int32  `json:"Amount" validate:"required,min=100,max=50000000"`
}

type VerifyResponse struct {
	Amount      int32  `json:"amount"`
	CardNumber  string `json:"cardNumber"`
	CardHashPan string `json:"cardHashPan"`
}

type SharePaymentRequest struct {
	Pairs       []pairs `json:"pairs" validate:"required"`
	PayerName   string  `json:"payerName" validate:"required"`
	ReturnUrl   string  `json:"returnUrl" validate:"required,url"`
	ClientRefId string  `json:"clientRefId"`
}

type pairs struct {
	Amount       int32  `json:"amount" validate:"required"`
	Name         string `json:"name"`
	UserIdentity string `json:"userIdentity" validate:"required"`
	Description  string `json:"description"`
}

type BlockedPaymentRequest struct {
	Pairs       []pairs `json:"pairs" validate:"required"`
	PayerName   string  `json:"payerName" validate:"required"`
	ReturnUrl   string  `json:"returnUrl" validate:"required,url"`
	ClientRefId string  `json:"clientRefId"`
}

type ReleasingBlockedPaymentRequest struct {
	Code     string `json:"code"`
	ClientId string `json:"clientId"`
}

type PaymentWithTracerIdRequest struct {
	Amount        int32  `json:"authority" validate:"required,min=100,max=50000000"`
	PayerIdentity string `json:"payerIdentity"`
	PayerName     string `json:"payerName"`
	Description   string `json:"description"`
	ReturnUrl     string `json:"returnUrl" validate:"required,url"`
	ClientRefId   string `json:"clientRefId"`
}

type PaymentWithTracerIdResponse struct {
	Itd  int32  `json:"itd"`
	Code string `json:"code"`
}

type PaymentSuspedingRequest struct {
	Amount        int32  `json:"authority" validate:"required,min=100,max=50000000"`
	PayerIdentity string `json:"payerIdentity"`
	PayerName     string `json:"payerName"`
	Description   string `json:"description"`
	ReturnUrl     string `json:"returnUrl" validate:"required,url"`
	ClientRefId   string `json:"clientRefId"`
}

type ErrorResponse struct {
	ErrorCode    any    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
