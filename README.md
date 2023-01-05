# PayGap [![Go Reference](https://pkg.go.dev/badge/github.com/gofarsi/paygap.svg)](https://pkg.go.dev/github.com/gofarsi/paygap)
Payment gateway providers SDK Go for zarinpal, idpay, pay.ir, and other gateway providers.

## Install

```shell
go get -u github.com/gofarsi/paygap
```

## Usage

example for zarinpal provider

```go
package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/paygap/client"
	"github.com/GoFarsi/paygap/providers/zarinpal"
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

## TODO

- [x] zarinpal
- [x] idpay
- [ ] pay.ir
- [ ] yekpay
- [ ] payping
- [ ] rayanpay
- [ ] nextpay
- [ ] mellat
- [ ] parsian
- [ ] pasargad
- [ ] sadad