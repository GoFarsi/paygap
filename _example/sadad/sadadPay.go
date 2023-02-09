package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/sadad"
	"log"
	"math/rand"
)

func main() {

	c := client.New()
	s, err := sadad.New(c, "9001", "RecivedMerchenantkey", "returnUrl", "1565879", "purchagePage", false)
	if err != nil {
		log.Fatal(err)
	}
	//pay
	//اگر پرداخت به صورت تسهیمی است باید آبجکت مولتی پلکسینگ نیز مقدار دهی شود
	orderId := string(rand.Int())
	resp, err := s.PaymentRequest(context.Background(), 50000, orderId, nil)
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
