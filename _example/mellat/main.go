package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/mellat"
)

func main() {
	c := client.New()
	m, err := mellat.New(c, "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := m.CreateTransaction(
		context.Background(),
		mellat.NewPaymentRequest(
			"1",                            // order id
			1_000_000,                      //amount
			"https://example.com/callback", // callback
			"1",                            // payer id (user id)
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	verifyResp, err := m.VerifyTransaction(
		context.Background(),
		mellat.NewVerifyRequest(
			"1",               // order id
			"SaleOrderId",     // SaleOrderId PostForm parameter in callback
			"SaleReferenceId", // SaleReferenceId PostForm parameter in callback
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(verifyResp)
}
