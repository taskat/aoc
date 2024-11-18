package stringutils

import "strconv"

// IsInteger returns true if the string is an integer and false otherwise
func IsInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
