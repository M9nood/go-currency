# go-currency
Go library for convert number to text currency

## Currency Support
 - Thailand currency (THB)

### Installation
```bash
go get github.com/M9nood/go-currency
```

### Usage
#### Thai Baht (THB)

```golang
package main

import (
	"fmt"
	"strconv"

	"github.com/M9nood/go-currency"
)

func main() {
    var num float64 = 124000000000000
    printNumText(num)
    // OUTPUT : หนึ่งร้อยยี่สิบสี่ล้านล้านบาท

    num = 21
    printNumText(num)
    // OUTPUT : ยี่สิบเอ็ดบาท

    num = 501
    printNumText(num)
    // OUTPUT : ห้าร้อยหนึ่งบาท

    num = 32.01
    printNumText(num)
    // OUTPUT : สามสิบสองบาทหนึ่งสตางค์

    num = 0.1
    printNumText(num)
    // OUTPUT : สิบสตางค์

}

func printNumText(num float64) {
	fmt.Printf("%s\n", currency.ThaiBahtText(num))
}

```


# License
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/M9nood/go-iterror/blob/master/LICENSE)
