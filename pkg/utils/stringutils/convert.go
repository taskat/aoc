package stringutils

import "strconv"

// Atoi converts a string to an integer
// If the string cannot be converted, it panics
func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

// RuneToInt converts a rune to an integer
// If the rune is not a digit, it panics
func RuneToInt(r rune) int {
	if !IsDigit(r) {
		panic("Rune is not a digit")
	}
	return int(r - '0')
}
