package main

import (
	"context"
	"log"

	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/vandar"
)

func main() {
	c := client.New()
	v, err := vandar.New(c, "4924fc5aeca14bfdaa0d44162baf4fb7ee19d706")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := v.RequestPayment(context.Background(), 10000, "http://vandagateway.local:8081/callback",
		"09353917307", "", "", "", "")
	if err != nil {
		log.Printf("%+v\n", err)
	}

	//make a break point and Go to payment page :=https://ipg.vandar.io/v3/{token} and token is resp.Token

	//get transaction detail
	detailResp, detailErr := v.TransactionDetail(context.Background(), resp.Token)
	if detailErr != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", detailResp)

	//and verify Transction
	verifyResp, verifyErr := v.VerifyPayment(context.Background(), resp.Token)
	if verifyErr != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", verifyResp)

}
