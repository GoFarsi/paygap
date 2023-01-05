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
	i, err := idpay.New(c, "YOUR_API_KEY", false)
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
		Callback: "http://example.com/callback",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
