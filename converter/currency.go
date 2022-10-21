package converter

import "fmt"

func GetCurrencyPrefix(c string) string {
	return CurrencyPrefix[c]
}

func FormatCurrency(num float64) string {
	var prefix = ""
	var format string
	intDigit, decDigit := getIntegerAndDecimalDigits(num)
	if num == 0 {
		return "0"
	}
	if num < 0 {
		prefix += "-"
		intDigit = intDigit[1:]
	}
	intDigitRv := reverse(intDigit)
	for i, n := range intDigitRv {
		if i != 0 && i%COMMA_POSITION == 0 {
			format = fmt.Sprintf(",%s", format)
		}
		format = fmt.Sprintf("%s%s", string(n), format)
	}

	if decDigit != "" {
		decDigit = fmt.Sprintf(".%s", decDigit)
	}

	return fmt.Sprintf("%s%s%s", prefix, format, decDigit)
}
