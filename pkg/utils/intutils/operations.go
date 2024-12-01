package intutils

import "github.com/taskat/aoc/pkg/utils/types"

// Abs returns the absolute value of the given number
func Abs[INT types.SignedInteger](number INT) INT {
	if number < 0 {
		return -number
	}
	return number
}
