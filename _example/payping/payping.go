package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/payping"
)

func main() {
	p, err := payping.New(client.New(), "YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	request := &payping.PaymentRequest{
		Amount:        11000,
		PayerIdentity: "124500",
		PayerName:     "Ali Hesami",
		ClientRefId:   "example-arbitary-code",
		ReturnUrl:     "http://example.com/callback",
		Description:   "desc test",
	}
	resp, err := p.RequestPayment(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	verify := &payping.VerifyRequest{
		Amount: request.Amount,
		RefId:  resp.Code,
	}
	verifyResp, err := p.VerifyPayment(context.Background(), verify)

	fmt.Println(verifyResp)
}
