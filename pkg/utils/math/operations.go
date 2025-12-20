package math

import "github.com/taskat/aoc/pkg/utils/types"

// Abs returns the absolute value of the given number
func Abs[T types.Real](number T) T {
	if number < 0 {
		return -number
	}
	return number
}

// Ceil returns the smallest integer value greater than or equal to the given number
func Ceil[T types.Real](number T) int {
	i := int(number)
	if T(i) < number {
		return i + 1
	}
	return i
}

// Floor returns the largest integer value less than or equal to the given number
func Floor[T types.Real](number T) int {
	i := int(number)
	if T(i) > number {
		return i - 1
	}
	return i
}

// Round returns the nearest integer to the given number
func Round[T types.Real](number T) int {
	if number > 0 {
		return int(Floor(float64(number) + 0.5))
	}
	return int(Ceil(float64(number) - 0.5))
}
