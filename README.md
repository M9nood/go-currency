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
    printTHBString(num)
    // OUTPUT : 124,000,000,000,000
    printTHBCurrency(num)
    // OUTPUT : ฿ 124,000,000,000,000
    printTHBText(num)
    // OUTPUT : หนึ่งร้อยยี่สิบสี่ล้านล้านบาท

    num = 21
    printTHBString(num)
    // OUTPUT : 21
    printTHBCurrency(num)
    // OUTPUT : ฿ 21
    printTHBText(num)
    // OUTPUT : ยี่สิบเอ็ดบาท

    num = 501
    printTHBString(num)
    // OUTPUT : 501
    printTHBCurrency(num)
    // OUTPUT : ฿ 501
    printTHBText(num)
    // OUTPUT : ห้าร้อยหนึ่งบาท

    num = 32.01
    printTHBString(num)
    // OUTPUT : 32.01
    printTHBCurrency(num)
    // OUTPUT : ฿ 32.01
    printTHBText(num)
    // OUTPUT : สามสิบสองบาทหนึ่งสตางค์

    num = 0.1
    printTHBString(num)
    // OUTPUT : 0.1
    printTHBCurrency(num)
    // OUTPUT : ฿ 0.1
    printTHBText(num)
    // OUTPUT : สิบสตางค์

}


func printTHBText(num float64) {
	fmt.Printf("THB text: %s\n", currency.THB(num).Text())
}

func printTHBCurrency(num float64) {
	fmt.Printf("THB : %s\n", currency.THB(num).Currency())
}

func printTHBString(num float64) {
	fmt.Printf("THB value: %s\n", currency.THB(num).String())
}

```


# License
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/M9nood/go-currency/blob/master/LICENSE)
