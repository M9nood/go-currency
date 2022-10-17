package main

import (
	"fmt"

	"github.com/M9nood/go-currency"
)

func main() {
	nums := []float64{
		124000000000000,
		124000000520000,
		21,
		10,
		310,
		501,
	}
	for _, n := range nums {
		printNum(n)
	}
}

func printNum(num float64) {
	fmt.Printf("num : %v, text:  %s\n----------------------\n", num, currency.ThaiBahtText(num))
}
