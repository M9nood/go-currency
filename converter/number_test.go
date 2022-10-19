package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntegerAndDecimalDigits(t *testing.T) {
	numInt, numDec := getIntegerAndDecimalDigits(12345.101)
	assert.Equal(t, "12345", numInt)
	assert.Equal(t, "101", numDec)

	numInt, numDec = getIntegerAndDecimalDigits(12345.50)
	assert.Equal(t, "12345", numInt)
	assert.Equal(t, "5", numDec)

	numInt, numDec = getIntegerAndDecimalDigits(-10.50)
	assert.Equal(t, "-10", numInt)
	assert.Equal(t, "5", numDec)
}

func TestIsZero(t *testing.T) {
	assert.True(t, isZero(0))
	assert.False(t, isZero(1))
	assert.False(t, isZero(-1))
}

func TestReverse(t *testing.T) {
	str := "1523"
	assert.Equal(t, "3251", reverse(str))

	str = "100"
	assert.Equal(t, "001", reverse(str))
}
func TestGetMillionUnitCount(t *testing.T) {
	numStr := "10"
	assert.Equal(t, 0, getMillionUnitCount(numStr))
	numStr = "100000"
	assert.Equal(t, 0, getMillionUnitCount(numStr))
	numStr = "1000000"
	assert.Equal(t, 1, getMillionUnitCount(numStr))
	numStr = "1000000000000"
	assert.Equal(t, 2, getMillionUnitCount(numStr))
}

func TestGetUnitPosition(t *testing.T) {
	assert.Equal(t, 0, getUnitPosition(0))
	assert.Equal(t, 1, getUnitPosition(1))
	assert.Equal(t, 2, getUnitPosition(2))
	assert.Equal(t, 3, getUnitPosition(3))
	assert.Equal(t, 4, getUnitPosition(4))
	assert.Equal(t, 5, getUnitPosition(5))
	assert.Equal(t, 0, getUnitPosition(6))
	assert.Equal(t, 1, getUnitPosition(7))
	assert.Equal(t, 2, getUnitPosition(8))
	assert.Equal(t, 3, getUnitPosition(9))
	assert.Equal(t, 4, getUnitPosition(10))
	assert.Equal(t, 5, getUnitPosition(11))
	assert.Equal(t, 0, getUnitPosition(12))
	assert.Equal(t, 3, getUnitPosition(15))
}

func TestIsMillionPosition(t *testing.T) {
	assert.False(t, isMillionPosition(0))
	assert.False(t, isMillionPosition(2))
	assert.True(t, isMillionPosition(MILLION_POSITION))
	assert.False(t, isMillionPosition(7))
	assert.False(t, isMillionPosition(12))
}

func TestIsOverPrecisionDigit(t *testing.T) {
	assert.False(t, isOverPrecisionDigit("0"))
	assert.False(t, isOverPrecisionDigit("10"))
	assert.True(t, isOverPrecisionDigit("100"))
	assert.True(t, isOverPrecisionDigit("001"))
}

func TestIsEmptyDigit(t *testing.T) {
	assert.True(t, isEmptyDigit("0"))
	assert.True(t, isEmptyDigit(""))
	assert.False(t, isEmptyDigit("100"))
}
