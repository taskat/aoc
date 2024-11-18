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
