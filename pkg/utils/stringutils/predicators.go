package stringutils

import "strconv"

// IsDigit returns true if the string is a digit and false otherwise
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsEmpty returns true if the string is empty and false otherwise
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsInteger returns true if the string is an integer and false otherwise
func IsInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
