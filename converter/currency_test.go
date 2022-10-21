package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrencyPrefix(t *testing.T) {
	pf := GetCurrencyPrefix("thb")
	assert.Equal(t, CurrencyPrefix["thb"], pf)

	// not have in map
	pf = GetCurrencyPrefix("NO_KEY")
	assert.Equal(t, "", pf)
}

func TestFormatCurrency(t *testing.T) {
	var num float64
	num = 0
	assert.Equal(t, "0", FormatCurrency(num))
	assert.Equal(t, "0", FormatCurrency(-num))
	num = 1000
	assert.Equal(t, "1,000", FormatCurrency(num))
	assert.Equal(t, "-1,000", FormatCurrency(-num))
	num = 540000
	assert.Equal(t, "540,000", FormatCurrency(num))
	assert.Equal(t, "-540,000", FormatCurrency(-num))
	num = 32540000
	assert.Equal(t, "32,540,000", FormatCurrency(num))
	assert.Equal(t, "-32,540,000", FormatCurrency(-num))
	num = 0.59
	assert.Equal(t, "0.59", FormatCurrency(num))
	assert.Equal(t, "-0.59", FormatCurrency(-num))
	num = 1234550.50
	assert.Equal(t, "1,234,550.5", FormatCurrency(num))
	assert.Equal(t, "-1,234,550.5", FormatCurrency(-num))
}
