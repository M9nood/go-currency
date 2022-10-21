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
		printTHB(n)
		fmt.Println("----------------------")
	}
}

func printTHB(num float64) {
	fmt.Printf("THB string: %s\n", currency.THB(num).String())
	fmt.Printf("THB currency: %s\n", currency.THB(num).Currency())
	fmt.Printf("THB text: %s\n", currency.THB(num).Text())
}
