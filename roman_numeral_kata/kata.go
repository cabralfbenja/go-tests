package romannumeralkata

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > 3999 {
		return "", errors.New("Cannot convert greater than 3999")
	}
	var res strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			res.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return res.String(), nil
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return arabic
}
