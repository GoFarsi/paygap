# PayGap
Payment gateway providers SDK Go for zarinpal, idpay, pay.ir, and other gateway providers.

## Install

```shell
go get -u github.com/GoFarsi/paygap
```

## Usage

example for zarinpal provider

```go
package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/pgp/client"
	zarinpal "github.com/GoFarsi/pgp/providers/zarinpal/v4"
	"log"
)

func main() {
	c := client.New()
	z, err := zarinpal.New(c, "YOUR_MERCHANT_ID", false)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := z.RequestPayment(context.Background(), 1000, "YOUR_CALL_BACK", "YOUR_CURRENCY", "description", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(resp)
}
```

