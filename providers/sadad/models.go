package sadad

import (
	"github.com/GoFarsi/paygap/client"
	"time"
)

type MultiplexingType int

const (
	Percentage MultiplexingType = iota
	Amount
)

type Sadad struct {
	Client      client.Transporter
	TerminalId  string `json:"terminal_id"`
	MerchantId  string `json:"merchant_id"`
	MerchantKey string `json:"merchant_key"`
}
type MultiplexingData struct {
	Type             MultiplexingType       `json:"type"`
	MultiplexingRows []MultiplexingDataItem `json:"multiplexingRows"`
}
type MultiplexingDataItem struct {
	IbanNumber int32 `json:"iban_number"`
	Value      int64 `json:"value"`
}

type PurchaseResult struct {
	OrderId          string           `json:"order_id"`
	Token            string           `json:"token"`
	ResCode          string           `json:"res_code"`
	VerifyResultData VerifyResultData `json:"verify_result_data"`
}
type PaymentRequest struct {
	TerminalId         string           `json:"terminal_id"`
	MerchantId         string           `json:"merchant_id"`
	Amount             int64            `json:"amount"`
	OrderId            string           `json:"order_id"`
	LocalDateTime      time.Time        `json:"local_date_time"`
	ReturnUrl          string           `json:"return_url"`
	SignData           string           `json:"sign_data"`
	EnableMultiplexing bool             `json:"enable_multiplexing"`
	MultiplexingData   MultiplexingData `json:"multiplexing_data"`
}
type PayResultData struct {
	ResCode     string `json:"res_code"`
	Token       string `json:"token"`
	Description string `json:"description"`
}
type VerifyRequest struct {
	token    string `json:"token"`
	SignData string `json:"signData"`
}
type VerifyResultData struct {
	Succeed       bool   `json:"succeed"`
	ResCode       string `json:"resCode"`
	Description   string `json:"description"`
	Amount        string `json:"amount"`
	RetrivalRefNo string `json:"retrivalRefNo"`
	SystemTraceNo string `json:"systemTraceNo"`
	OrderId       string `json:"orderId"`
}

func (data *MultiplexingData) IsValidated() bool {

	if data.Type != 0 && data.Type != 1 {
		return false
	}
	if len(data.MultiplexingRows) == 0 {
		return false
	}

	for i := 0; i < len(data.MultiplexingRows); i++ {
		if data.MultiplexingRows[i].Value < 0 {
			return false
		}
	}

	switch data.Type {
	case 0:
		var sum int64
		for i := 0; i < len(data.MultiplexingRows); i++ {
			sum = data.MultiplexingRows[i].Value + sum
		}
		if sum > 100 {
			return false
		}
		for i := 0; i < len(data.MultiplexingRows); i++ {
			if data.MultiplexingRows[i].Value > 99 {
				return false
			}
		}
		break
	case 1:
		break
	default:
		break
	}

	return true
}
