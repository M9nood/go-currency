package main

import (
	"fmt"
	"strconv"

	"github.com/M9nood/go-currency"
)

func main() {
	nums := []float64{
		124000000400000,
		124000000520000,
		21,
		10,
		310,
		501,
		12.10,
		500.107,
		32.01,
		0.1,
		0,
	}
	for _, n := range nums {
		printNum(n)
	}
}

func printNum(num float64) {
	fmt.Printf("num : %v, text:  %s\n----------------------\n", strconv.FormatFloat(num, 'f', -1, 64), currency.ThaiBahtText(num))
}
