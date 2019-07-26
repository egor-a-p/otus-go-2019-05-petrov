package hw2

import (
	"strconv"
	"strings"
	"unicode"
)

const escape = '\\'
const zero = '0'

func Encode(str string) string {
	if len(str) == 0 {
		return ""
	}

	var (
		result   strings.Builder
		runes    = []rune(str)
		previous = runes[0]
		count    = 0
	)

	for index, current := range runes {
		if previous == current {
			count++
		} else {
			result.WriteString(encodeSymbol(previous, count))
			count = 1
			previous = current
		}

		if index == len(runes)-1 {
			result.WriteString(encodeSymbol(previous, count))
		}
	}

	return result.String()
}

func encodeSymbol(symbol rune, count int) string {
	var result strings.Builder

	if unicode.IsDigit(symbol) || symbol == escape {
		result.WriteRune(escape)
	}

	result.WriteRune(symbol)

	if count > 1 {
		result.WriteString(strconv.Itoa(count))
	}

	return result.String()
}

func Decode(str string) string {
	if len(str) == 0 {
		return ""
	}

	var (
		runes  = []rune(str)
		result strings.Builder
	)

	for index := 0; index < len(runes); {
		symbol, shift := decodeSymbol(runes[index:])
		index += shift
		count, shift := decodeCount(runes[index:])
		index += shift

		for j := 0; j < count; j++ {
			result.WriteRune(symbol)
		}
	}

	return result.String()
}

func decodeSymbol(runes []rune) (symbol rune, shift int) {
	if unicode.IsDigit(runes[0]) {
		panic("Illegal symbol")
	}

	shift = 1

	if runes[0] == escape {
		if len(runes) == 1 || (!unicode.IsDigit(runes[1]) && runes[1] != escape) {
			panic("Illegal symbol")
		}
		runes = runes[1:]
		shift++
	}

	return runes[0], shift
}

func decodeCount(runes []rune) (count int, shift int) {
	if len(runes) == 0 || !unicode.IsDigit(runes[0]) {
		return 1, 0
	}

	if runes[0] == zero {
		panic("Illegal symbol")
	}

	for shift < len(runes) && unicode.IsDigit(runes[shift]) {
		shift++
	}

	count, _ = strconv.Atoi(string(runes[0:shift]))

	return count, shift
}
