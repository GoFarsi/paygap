package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/Parsian"
	"log"
)

func main() {
	c := client.New()
	p, err := Parsian.New(c, "YOUR_MERCHANT_ID")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := p.RequestPayment(context.Background(), 1000, "YOUR_CALL_BACK", 1, "description", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
