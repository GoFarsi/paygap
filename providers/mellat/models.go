package mellat

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"github.com/GoFarsi/paygap/client"
)

type Mellat struct {
	client   client.Transporter
	username string `validate:"required"`
	password string `validate:"required"`

	url string
}

type response interface {
	ResponseCode() int
	modifyResponse() error
}

type paymentRequest struct {
	XMLName      xml.Name `xml:"ns1:bpPayRequest"`
	TerminalId   string   `xml:"terminalId"`
	UserName     string   `xml:"userName"`
	UserPassword string   `xml:"userPassword"`
	OrderId      string   `xml:"orderId"`
	Amount       uint64   `xml:"amount"`
	LocalDate    string   `xml:"localDate"`
	LocalTime    string   `xml:"localTime"`
	CallBackUrl  string   `xml:"callBackUrl"`
	PayerId      string   `xml:"payerId"`
}

func NewPaymentRequest(
	orderId string,
	amount uint64,
	callBackUrl string,
	payerId string,
) *paymentRequest {
	return &paymentRequest{
		OrderId:     orderId,
		Amount:      amount,
		LocalDate:   time.Now().Format("20060402"),
		LocalTime:   time.Now().Format("150405"),
		CallBackUrl: callBackUrl,
		PayerId:     payerId,
	}
}

func (pr *paymentRequest) raw(
	userId string,
	password string,
) ([]byte, error) {
	pr.TerminalId = userId
	pr.UserName = userId
	pr.UserPassword = password
	root := newSoapRoot()
	root.Body.Request = pr
	return root.Marshal()
}

type PaymentResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		BpPay   struct {
			XMLName xml.Name `xml:"bpPayRequestResponse"`
			Return  string   `xml:"return"`
		}
	}
	responseCode int    `xml:"-"`
	refId        string `xml:"-"`
}

func (pr *PaymentResponse) ResponseCode() int {
	if pr == nil {
		return 0
	}
	return pr.responseCode
}

func (pr *PaymentResponse) RefId() string {
	return pr.refId
}

func (pr *PaymentResponse) modifyResponse() error {
	params := strings.Split(pr.Body.BpPay.Return, ",")
	if len(params) > 0 {
		if params[0] == "0" {
			pr.responseCode = 0
		} else {
			code, err := strconv.Atoi(params[0])
			if err != nil {
				return err
			}
			pr.responseCode = code
		}
		if len(params) > 1 {
			pr.refId = params[1]
		}
	}
	return nil
}

type verifyRequest struct {
	XMLName         xml.Name `xml:"ns1:bpVerifyRequest"`
	TerminalId      string   `xml:"terminalId"`
	UserName        string   `xml:"userName"`
	Password        string   `xml:"userPassword"`
	OrderId         string   `xml:"orderId"`
	SaleOrderId     string   `xml:"saleOrderId"`
	SaleReferenceId string   `xml:"saleReferenceId"`
}

func NewVerifyRequest(
	orderId string,
	saleOrderId string,
	saleReferenceId string,
) *verifyRequest {
	return &verifyRequest{
		OrderId:         orderId,
		SaleOrderId:     saleOrderId,
		SaleReferenceId: saleReferenceId,
	}
}

func (vr *verifyRequest) raw(
	userId string,
	password string,
) ([]byte, error) {
	vr.TerminalId = userId
	vr.UserName = userId
	vr.Password = password
	root := newSoapRoot()
	root.Body.Request = vr
	return root.Marshal()
}

type VerifyResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		BpPay   struct {
			XMLName xml.Name `xml:"bpVerifyRequestResponse"`
			Return  string   `xml:"return"`
		}
	}
	responseCode int `xml:"-"`
}

func (vr *VerifyResponse) ResponseCode() int {
	return vr.responseCode
}

func (vr *VerifyResponse) modifyResponse() error {
	if vr.Body.BpPay.Return == "0" {
		vr.responseCode = -1
		return nil
	}
	code, err := strconv.Atoi(vr.Body.BpPay.Return)
	if err != nil {
		return err
	}
	vr.responseCode = code
	return nil
}

type soapRoot struct {
	XMLName xml.Name `xml:"x:Envelope"`
	X       string   `xml:"xmlns:x,attr"`
	Ns1     string   `xml:"xmlns:ns1,attr"`
	Body    soapBody
}

func (r *soapRoot) Marshal() ([]byte, error) {
	return xml.MarshalIndent(r, "", "  ")
}

type soapBody struct {
	XMLName xml.Name `xml:"x:Body"`
	Request interface{}
}

func newSoapRoot() *soapRoot {
	return &soapRoot{
		X:   "http://schemas.xmlsoap.org/soap/envelope/",
		Ns1: "http://interfaces.core.sw.bps.com/",
	}
}
