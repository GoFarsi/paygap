package sadad

import (
	"crypto/aes"
	b64 "encoding/base64"
	"errors"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/status"
	"github.com/andreburgaud/crypt2go/ecb"
	"math/rand"
	"strings"
	"time"
)

const (
	requestPaymentUrl = "/api/v0/Request/PaymentRequest"
)

func New(client client.Transporter, terminalId string, amount int64, merchantKey string, verifyUrl string, merchantId string, purchasePage string, enableMultiplexing bool, multiplexingData MultiplexingData) (*Sadad, error) {
	if client == nil {
		return nil, status.ERR_CLIENT_IS_NIL
	}
	orderId := string(rand.Int())
	str := []string{terminalId, orderId, string(amount)}
	joinedString := strings.Join(str, ";")
	key := []byte(merchantKey)
	plaintext := []byte(joinedString)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	sb64Enc := b64.StdEncoding.EncodeToString([]byte(plaintext))

	sadad := &Sadad{
		Client:             client,
		TerminalId:         terminalId,
		MerchantId:         merchantId,
		Amount:             amount,
		OrderId:            orderId,
		AdditionalData:     "",
		LocalDateTime:      time.Now(),
		ReturnUrl:          verifyUrl,
		SignData:           sb64Enc,
		EnableMultiplexing: enableMultiplexing,
		MultiplexingData:   multiplexingData,
		MerchantKey:        merchantKey,
		PurchasePage:       purchasePage,
	}

	if sadad.EnableMultiplexing {
		if !sadad.MultiplexingData.IsValidated() {
			return nil, errors.New("multiplexing data is not valid")
		}
	}
	return sadad, nil
}
