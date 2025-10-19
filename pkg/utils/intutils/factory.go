package intutils

import "github.com/taskat/aoc/pkg/utils/types"

// IsN returns a function that checks if a number is equal to n
func IsN[INT types.Integer](n INT) func(number INT) bool {
	return func(number INT) bool {
		return number == n
	}
}
