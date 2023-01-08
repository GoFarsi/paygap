package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/idpay"
	"log"
)

func main() {
	c := client.New()
	i, err := idpay.New(c, "key", true)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := i.CreateTransaction(context.Background(), &idpay.PaymentRequest{
		OrderId:  "xxxxxxxx-xxxxxx-xxxx",
		Amount:   5000,
		Name:     "transaction A",
		Phone:    "09151234567",
		Mail:     "idpay@gmail.com",
		Desc:     "example description",
		Callback: "https://example.com/callback",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	verifyResp, err := i.VerifyTransaction(context.Background(), &idpay.VerifyRequest{
		Id:      resp.Id,
		OrderId: "xxxxxxxx-xxxxxx-xxxx",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(verifyResp)
}
