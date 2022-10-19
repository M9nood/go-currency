package converter

import (
	"strconv"
	"strings"
)

func getIntegerAndDecimalDigits(num float64) (string, string) {
	numStrSplit := strings.Split(strconv.FormatFloat(num, 'f', -1, 64), ".")
	if len(numStrSplit) < 2 {
		return numStrSplit[0], ""
	}
	return numStrSplit[0], numStrSplit[1]
}

func isZero(num float64) bool {
	return num == 0
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func getUnitPosition(pos int) int {
	if pos != 0 && pos%MILLION_POSITION == 0 {
		return MILLION_POSITION
	}
	return pos % MILLION_POSITION
}

func isMillionPosition(pos int) bool {
	return pos == MILLION_POSITION
}

func isOverPrecisionDigit(digit string) bool {
	return len(digit) > DEFAULT_PRECISION_DIGIT
}

func isEmptyDigit(num string) bool {
	return num == "0" || num == ""
}
