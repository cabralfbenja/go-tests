package romannumeralkata

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, err := ConvertToRoman(test.Arabic)

			if err != nil {
				t.Errorf("got %q, should not have had error", err)
			}

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}

		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:1] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

func TestPropertiesOfConsecutiveSymbols(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman, err := ConvertToRoman(arabic)
		if err != nil {
			return true
		}
		var count = 0
		var prev rune
		for _, char := range roman {
			if count == 4 {
				return false
			}
			if char == prev {
				count++
			} else {
				count = 0
			}
			prev = char
		}
		return true
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
func TestPropertiesOfSubstractors(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman, err := ConvertToRoman(arabic)
		if err != nil {
			return true
		}
		for index, char := range roman {
			if index == len(roman)-1 {
				return true
			}
			next := rune(roman[index+1])
			if !isSubstractor(char) {
				if next == getNext(char) {
					return false
				}
			}
		}
		return true
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

func isSubstractor(char rune) bool {
	return char == 'I' || char == 'X' || char == 'C'
}

func getNext(char rune) rune {
	switch char {
	case 'V':
		return 'X'
	case 'L':
		return 'C'
	case 'D':
		return 'M'
	default:
		return 0
	}
}
