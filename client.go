package currency

import (
	"github.com/M9nood/go-currency/converter"
)

type Currency interface {
	Text() string
}

func (c CurrencyModel) Text() string {
	switch c.CurrencyFormat {
	case CURRENCY_FORMAT_THB:
		return converter.ThaiBahtText(c.Number)
	default:
		return ""
	}

}

func THB(num float64) Currency {
	return CurrencyModel{
		Number:         num,
		CurrencyFormat: CURRENCY_FORMAT_THB,
	}
}
