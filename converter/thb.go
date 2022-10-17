package converter

import (
	"strconv"
	"strings"
)

var (
	THB_NUMBER_TEXTS   = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า", "สิบ"}
	THB_UNIT_TEXTS     = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}
	THB_PRIMARY_UNIT   = "บาท"
	THB_SECONDARY_UNIT = "สตางค์"
	THB_DOT            = "จุด"
)

func ConvertToThaiBahtText(num float64) string {
	var currencyText string
	if isZero(num) {
		return THB_NUMBER_TEXTS[0] + THB_PRIMARY_UNIT
	}
	intValue, decValue := getIntegerAndDecimalDigits(num)
	var intText, decText string
	if !isEmptyDigit(intValue) {
		intText = getIntegerText(intValue)
	}
	if !isEmptyDigit(decValue) {
		decText = getDecimalText(decValue)
	}
	currencyText += intText + decText
	return currencyText
}

func getIntegerText(intValueStr string) string {
	var intText string
	revIntStr := reverse(intValueStr)
	intStrRunes := []rune(revIntStr)
	millionCount := getMillionUnitCount(revIntStr)
	lengthOfDigit := len(intStrRunes)
	numTextByPos := make([]string, lengthOfDigit)
	for i := 0; i < lengthOfDigit; i++ {
		nextUnitValue := rune(-1)
		if i < lengthOfDigit-1 {
			nextUnitValue = intStrRunes[(i + 1)]
		}

		numTextByPos[(lengthOfDigit-1)-i] = convertToThaiBaht(intStrRunes[i], nextUnitValue, i, lengthOfDigit)
		if isMillionPosition(i) {
			numTextByPos[(lengthOfDigit-1)-i] += millionUnitName(millionCount)
		}
	}
	intText = strings.Join(numTextByPos[:], "") + THB_PRIMARY_UNIT
	return intText
}

func getDecimalText(decValueStr string) string {
	var decText string
	if !isOverPrecisionDigit(decValueStr) {
		if len(decValueStr) == 1 {
			decValueStr += "0"
		}
		revDecStr := reverse(decValueStr)
		dectrRunes := []rune(revDecStr)
		lengthOfDigit := len(dectrRunes)
		numTextByPos := make([]string, lengthOfDigit)
		for i := 0; i < lengthOfDigit; i++ {
			numTextByPos[(lengthOfDigit-1)-i] = convertToThaiBaht(dectrRunes[i], rune(-1), i, lengthOfDigit)
		}
		decText = strings.Join(numTextByPos[:], "") + THB_SECONDARY_UNIT
	} else {
		for _, n := range decValueStr {
			decText += convertToDecimal(string(n))
		}
		decText = THB_DOT + decText + THB_SECONDARY_UNIT
	}
	return decText
}

func convertToThaiBaht(unitValue, nextUnitValue rune, pos, lengthOfDigit int) string {
	intVar, _ := strconv.Atoi(string(unitValue))
	nextIntVar, _ := strconv.Atoi(string(nextUnitValue))
	if isZero(float64(intVar)) {
		return ""
	}
	numText := THB_NUMBER_TEXTS[intVar]
	unitPos := getUnitPosition(pos)
	if unitPos == TEN_POSITION && intVar == 2 {
		numText = "ยี่"
	}
	if unitPos == TEN_POSITION && intVar == 1 {
		numText = ""
	}
	if lengthOfDigit > 1 && unitPos == UNIT_POSITION && intVar == 1 && nextIntVar != 0 {
		numText = "เอ็ด"
	}

	unitText := THB_UNIT_TEXTS[getUnitPosition(pos)]

	return numText + unitText
}

func convertToDecimal(decimalValue string) string {
	decVar, _ := strconv.Atoi(decimalValue)
	return THB_NUMBER_TEXTS[decVar]
}

func millionUnitName(millionCount int) string {
	text := ""
	for i := 0; i < millionCount; i++ {
		text += THB_UNIT_TEXTS[MILLION_POSITION]
	}
	return text
}
