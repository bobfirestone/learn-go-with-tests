package romannumerals

import "strings"

// RomanNumeral struct mapping the arabic numerals to their roman
type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals is the mapping of arabic numbers to roman numeral symbols
var RomanNumerals = []RomanNumeral{
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

// ConvertToRoman takes an integer and returns a string of roman numerals
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range RomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

// ConvertToArabic takes a roman numeral and converts it to a digit
func ConvertToArabic(roman string) int {
	total := 0
	for range roman {
		total++
	}
	return total
}
