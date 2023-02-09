# PayGap [![Go Reference](https://pkg.go.dev/badge/github.com/GoFarsi/paygap.svg)](https://pkg.go.dev/github.com/GoFarsi/paygap)
Payment gateway providers SDK Go for zarinpal, idpay, pay.ir, and other gateway providers.

![paygap](assets/banner.jpg)

## Install For ![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/GoFarsi/paygap/main)

```shell
go get -u github.com/GoFarsi/paygap
```

## Usage

example for zarinpal provider ([Other examples](_examples))

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

## Contributing

1. fork project in your GitHub account.
2. create new branch for new changes.
3. after change code, send Pull Request.

## TODO

- [x] zarinpal
- [x] idpay
- [x] pay.ir
- [x] payping
- [ ] rayanpay
- [ ] nextpay
- [ ] mellat
- [ ] parsian
- [ ] pasargad
- [x] sadad
