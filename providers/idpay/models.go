package idpay

import "github.com/GoFarsi/paygap/client"

type IdPay struct {
	client  client.Transporter
	apiKey  string `validate:"required"`
	sandbox bool

	baseUrl              string
	paymentEndpoint      string
	verifyEndpoint       string
	inquiryEndpoint      string
	transactionsEndpoint string
}

type ErrorResponse struct {
	ErrorCode    any    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type PaymentRequest struct {
	OrderId  string `validate:"required,max=50" json:"order_id"`
	Amount   uint   `validate:"required,min=1000,max=5000000000" json:"amount"`
	Name     string `validate:"omitempty,max=255" json:"name"`
	Phone    string `validate:"omitempty,max=11" json:"phone"`
	Mail     string `validate:"email,max=255" json:"mail"`
	Desc     string `validate:"omitempty,max=255" json:"desc"`
	Callback string `validate:"required,max=2048,url" json:"callback"`
}

type PaymentResponse struct {
	Id   string `json:"id"`
	Link string `json:"link"`
}

type VerifyRequest struct {
	Id      string `validate:"required" json:"id"`
	OrderId string `validate:"required,max=50" json:"order_id"`
}

type VerifyResponse struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	TrackId string `json:"track_id"`
	OrderId string `json:"order_id"`
	Amount  string `json:"amount"`
	Date    string `json:"date"`
	Payment struct {
		TrackId      string `json:"track_id"`
		Amount       string `json:"amount"`
		CardNo       string `json:"card_no"`
		HashedCardNo string `json:"hashed_card_no"`
		Date         string `json:"date"`
	} `json:"payment"`
	Verify struct {
		Date int `json:"date"`
	} `json:"verify"`
}

type TransactionStatusRequest struct {
	Id      string `validate:"required" json:"id"`
	OrderId string `validate:"required,max=50" json:"order_id"`
}

type TransactionStatusResponse struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	TrackId string `json:"track_id"`
	OrderId string `json:"order_id"`
	Amount  string `json:"amount"`
	Wage    struct {
		By     string `json:"by"`
		Type   string `json:"type"`
		Amount string `json:"amount"`
	} `json:"wage"`
	Date  int `json:"date"`
	Payer struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Mail  string `json:"mail"`
		Desc  string `json:"desc"`
	} `json:"payer"`
	Payment struct {
		TrackId      string `json:"track_id"`
		Amount       string `json:"amount"`
		CardNo       string `json:"card_no"`
		HashedCardNo string `json:"hashed_card_no"`
		Date         string `json:"date"`
	} `json:"payment"`
	Verify struct {
		Date string `json:"date"`
	} `json:"verify"`
	Settlement struct {
		TrackId string `json:"track_id"`
		Amount  string `json:"amount"`
		Date    string `json:"date"`
	} `json:"settlement"`
}

type TransactionListRequest struct {
	Page     int `json:"-"`
	PageSize int `json:"-"`

	Id                  string   `json:"id"`
	OrderId             string   `validate:"omitempty,max=50" json:"order_id"`
	Amount              uint     `validate:"omitempty,max=5000000000,min=1000" json:"amount"`
	Status              []string `json:"status"`
	TrackId             string   `json:"track_id"`
	PaymentCardNo       string   `validate:"omitempty,max=16" json:"payment_card_no"`
	PaymentHashedCardNo string   `json:"payment_hashed_card_no"`
	PaymentDate         struct {
		Min uint `json:"min"`
		Max uint `json:"max"`
	} `json:"payment_date"`
	SettlementDate struct {
		Min uint `json:"min"`
		Max uint `json:"max"`
	} `json:"settlement_date"`
}

type TransactionListResponse struct {
	Attachment struct {
		Timestamp   int    `json:"timestamp"`
		TotalCount  int    `json:"total_count"`
		PageCount   int    `json:"page_count"`
		CurrentPage int    `json:"current_page"`
		TotalAmount string `json:"total_amount"`
		PageAmount  int    `json:"page_amount"`
	} `json:"attachment"`
	Records []*struct {
		Id      string `json:"id"`
		Status  string `json:"status"`
		TrackId string `json:"track_id"`
		OrderId string `json:"order_id"`
		Amount  string `json:"amount"`
		Wage    struct {
			By     string `json:"by"`
			Type   string `json:"type"`
			Amount string `json:"amount"`
		} `json:"wage"`
		Date  int `json:"date"`
		Payer struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
			Mail  string `json:"mail"`
			Desc  string `json:"desc"`
		} `json:"payer"`
		Payment struct {
			TrackId      string `json:"track_id"`
			Amount       string `json:"amount"`
			CardNo       string `json:"card_no"`
			HashedCardNo string `json:"hashed_card_no"`
			Date         string `json:"date"`
		} `json:"payment"`
		Verify struct {
			Date string `json:"date"`
		} `json:"verify"`
		Settlement struct {
			Account struct {
				Id     string `json:"id"`
				Iban   string `json:"iban"`
				Number string `json:"number"`
				Bank   struct {
					Id    string `json:"id"`
					Title string `json:"title"`
				} `json:"bank"`
			} `json:"account"`
			TrackId string `json:"track_id"`
			Amount  string `json:"amount"`
			Date    string `json:"date"`
		} `json:"settlement"`
	} `json:"records"`
}
