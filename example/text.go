package main

import (
	"fmt"
	"strconv"

	"github.com/M9nood/go-currency"
)

func main() {
	nums := []float64{
		124000000400000,
		124005674340400000,
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
		fmt.Printf("num : %v\n", strconv.FormatFloat(n, 'f', -1, 64))
		printTHBCurrency(n)
		printTHBText(n)
		fmt.Println("----------------------")
	}
}

func printTHBText(num float64) {
	fmt.Printf("THB text: %s\n", currency.THB(num).Text())
}

func printTHBCurrency(num float64) {
	fmt.Printf("THB : %s\n", currency.THB(num).Currency())
}
