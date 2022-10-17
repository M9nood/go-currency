package currency

import (
	"github.com/M9nood/go-currency/converter"
)

func ThaiBahtText(num float64) string {
	return converter.ConvertToThaiBahtText(num)
}
