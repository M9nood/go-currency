package currency

import (
	"fmt"

	"github.com/M9nood/go-currency/converter"
)

type Currency interface {
	Text() string
	Currency() string
	String() string
}

func (c CurrencyModel) Text() string {
	switch c.CurrencyFormat {
	case CURRENCY_FORMAT_THB:
		return converter.ThaiBahtText(c.Number)
	default:
		return ""
	}
}

func (c CurrencyModel) Currency() string {
	pf := converter.GetCurrencyPrefix(string(c.CurrencyFormat))
	if pf != "" {
		return fmt.Sprintf("%s %s", pf, converter.FormatCurrency(c.Number))
	}
	return converter.FormatCurrency(c.Number)
}

func (c CurrencyModel) String() string {
	return converter.FormatCurrency(c.Number)
}

func THB(num float64) Currency {
	return CurrencyModel{
		Number:         num,
		CurrencyFormat: CURRENCY_FORMAT_THB,
	}
}
