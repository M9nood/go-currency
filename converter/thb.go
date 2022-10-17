package converter

import (
	"strconv"
	"strings"
)

func ConvertToThaiBahtText(num float64) string {
	var currencyText string
	if isZero(num) {
		return ""
	}
	intValue, _ := getIntegerAndDecimalDigits(num)

	intText := getIntegerText(intValue)
	currencyText += intText
	return currencyText
}

func getIntegerText(intValueStr string) string {
	var intText string
	revIntStr := reverse(intValueStr)
	runes := []rune(revIntStr)
	millionCount := getMillionUnitCount(revIntStr)
	lengthOfDigit := len(runes)
	numTextByPos := make([]string, lengthOfDigit)
	for i := 0; i < lengthOfDigit; i++ {
		numTextByPos[(lengthOfDigit-1)-i] = convertToThaiBaht(runes[i], i, lengthOfDigit)
		if isMillionPosition(i) {
			numTextByPos[(lengthOfDigit-1)-i] += millionUnitName(millionCount)
		}
	}
	intText = strings.Join(numTextByPos[:], "")
	return intText
}

func convertToThaiBaht(numChar rune, pos, lengthOfDigit int) string {
	intVar, _ := strconv.Atoi(string(numChar))
	if isZero(float64(intVar)) {
		return ""
	}
	numText := THAI_NUMBER_TEXTS[intVar]
	unitPos := getUnitPosition(pos)
	if unitPos == TEN_POSITION && intVar == 2 {
		numText = "ยี่"
	}
	if unitPos == TEN_POSITION && intVar == 1 {
		numText = ""
	}
	if lengthOfDigit > 1 && unitPos == UNIT_POSITION && intVar == 1 {
		numText = "เอ็ด"
	}

	unitText := THAI_UNIT_TEXTS[getUnitPosition(pos)]

	return numText + unitText
}

func millionUnitName(millionCount int) string {
	text := ""
	for i := 0; i < millionCount; i++ {
		text += THAI_UNIT_TEXTS[MILLION_POSITION]
	}
	return text
}
