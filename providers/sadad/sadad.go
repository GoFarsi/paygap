package sadad

import (
	"context"
	"crypto/aes"
	b64 "encoding/base64"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"github.com/andreburgaud/crypt2go/ecb"
	"google.golang.org/grpc/codes"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	baseUrl           = "https://sadad.shaparak.ir"
	requestPaymentUrl = "/api/v0/Request/PaymentRequest"
	requestVerifyUrl  = "/api/v0/Advice/Verify"
)

func New(client client.Transporter, terminalId string, amount int64,
	merchantKey string, returnUrl string, merchantId string,
	purchasePage string, enableMultiplexing bool,
	multiplexingData *MultiplexingData) (*Sadad, error) {

	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}
	orderId := string(rand.Int())
	str := []string{terminalId, orderId, string(amount)}
	joinedString := strings.Join(str, ";")

	sadad := &Sadad{
		Client:             client,
		TerminalId:         terminalId,
		MerchantId:         merchantId,
		Amount:             amount,
		OrderId:            orderId,
		AdditionalData:     "",
		LocalDateTime:      time.Now(),
		ReturnUrl:          returnUrl,
		SignData:           "",
		EnableMultiplexing: enableMultiplexing,
		MultiplexingData:   *multiplexingData,
		MerchantKey:        merchantKey,
		PurchasePage:       purchasePage,
	}

	signedDataAsSadadWay, er := sadad.SigningData(joinedString)
	if er != nil {
		panic(er)
	}
	sadad.SignData = signedDataAsSadadWay

	return sadad, nil
}

func (s *Sadad) PayRequest(ctx context.Context) (*PayResultData, error) {

	req := &PaymentRequest{
		MerchantId:    s.MerchantId,
		Amount:        s.Amount,
		TerminalId:    s.TerminalId,
		ReturnUrl:     s.ReturnUrl,
		SignData:      s.SignData,
		OrderId:       s.OrderId,
		LocalDateTime: s.LocalDateTime,
	}
	if err := s.Client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}
	if !s.MultiplexingData.IsValidated() {
		return nil, status.New(1, http.StatusBadRequest, codes.FailedPrecondition, "خطا در دیتا ورودی برای تسهیم")
	}
	response := &PayResultData{}
	resp, err := s.Client.Post(ctx, &client.APIConfig{Host: baseUrl, Path: requestPaymentUrl}, req)
	if err != nil {
		return nil, status.New(2, http.StatusInternalServerError, codes.Internal, err.Error())
	}
	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(3, http.StatusInternalServerError, codes.Internal, err.Error())
	}
	if response.ResCode != "0" {
		return nil, status.New(4, http.StatusInternalServerError, codes.Canceled, "failed pay request")

	}
	return response, nil
}

func (s *Sadad) VerifyRequest(ctx context.Context, PayRes PayResultData) (*VerifyResultData, error) {

	signedDataAsSadadWay, er := s.SigningData(PayRes.Token)
	if er != nil {
		panic(er)
	}
	req := &VerifyRequest{
		token:    PayRes.Token,
		SignData: signedDataAsSadadWay,
	}

	if err := s.Client.GetValidator().Struct(req); err != nil {
		return nil, status.New(0, http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	resp, err := s.Client.Post(ctx, &client.APIConfig{Host: baseUrl, Path: requestVerifyUrl}, req)

	if err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}
	response := &VerifyResultData{}

	if err := resp.GetJSON(response); err != nil {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, err.Error())
	}

	if response.ResCode != "0" {
		return nil, status.New(0, http.StatusInternalServerError, codes.Internal, "verify failed")
	}

	return response, nil
}

func (s *Sadad) SigningData(plainText string) (string, error) {

	key := []byte(s.MerchantKey)
	plaintext := []byte(plainText)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	sb64Enc := b64.StdEncoding.EncodeToString([]byte(plaintext))
	return sb64Enc, nil
}
