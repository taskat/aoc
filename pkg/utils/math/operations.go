package math

import "github.com/taskat/aoc/pkg/utils/types"

// Abs returns the absolute value of the given number
func Abs[T types.Real](number T) T {
	if number < 0 {
		return -number
	}
	return number
}
