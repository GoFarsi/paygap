package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/sadad"
	"log"
)

func main() {

	c := client.New()
	s, err := sadad.New(c, "9001", 1500, ":)", "back_url", "5888789", "yourPurchasePage",
		false, nil)
	if err != nil {
		log.Fatal(err)
	}
	//pay
	resp, err := s.PayRequest(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	//verify
	v_res, v_err := s.VerifyRequest(context.Background(), *resp)
	if v_err != nil {
		log.Fatal(v_err)
	}

	fmt.Println(v_res)
}
