package main

import (
	"context"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/zibal"
	"log"
)

func main() {
	c := client.New()

	z, err := zibal.New(c, "zibal")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := z.RequestPayment(context.Background(), 5000, "https://example.com", "description")
	if err != nil {
		log.Fatal(err)
	}

	trackID := resp.TrackID

	_, err = z.VerifyPayment(context.Background(), trackID)
	if err != nil {
		log.Fatal(err)
	}
}
