package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/pay"
	"log"
)

func main() {
	ctx := context.Background()

	p, err := pay.New(client.New(), "YOUR_API_KEY", true)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := p.CreateTransaction(ctx, &pay.PaymentRequest{
		Amount:          11000,
		Redirect:        "http://example.com/callback",
		Mobile:          "09151234567",
		FactorNumber:    "xxxxx",
		Description:     "desc test",
		ValidCardNumber: "1234123412341234",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	verifyResp, err := p.VerifyTransaction(ctx, &pay.VerifyRequest{
		Token: resp.Token,
	})

	fmt.Println(verifyResp)
}
