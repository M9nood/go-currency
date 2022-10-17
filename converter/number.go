package converter

import (
	"strconv"
	"strings"
)

// func getIntgerDigits(num float64) string {
// 	numStrSplit := strings.Split(fmt.Sprintf("%v", num), ".")
// 	return numStrSplit[0]
// }

// func getDecimalDigits(num float64) string {
// 	numStrSplit := strings.Split(fmt.Sprintf("%v", num), ".")
// 	return numStrSplit[1]
// }

func getIntegerAndDecimalDigits(num float64) (string, string) {
	numStrSplit := strings.Split(strconv.FormatFloat(num, 'f', 6, 64), ".")
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

func getMillionUnitCount(s string) int {
	return len(s) / MILLION_POSITION
}

func getUnitPosition(pos int) int {
	return pos % MILLION_POSITION
}

func isMillionPosition(pos int) bool {
	return pos == MILLION_POSITION
}
