package intutils

import "github.com/taskat/aoc/pkg/utils/types"

// Abs returns the absolute value of the given number
func Abs[INT types.SignedInteger](number INT) INT {
	if number < 0 {
		return -number
	}
	return number
}

// Equals returns true if the two numbers are equal and false otherwise
func Equals[INT types.Integer](a, b INT) bool {
	return a == b
}

// Diff returns the difference between the two numbers
// as Diff(a, b) = a - b
func Diff[INT types.SignedInteger](a, b INT) INT {
	return a - b
}

// Product returns the product of the two numbers
// as Product(a, b) = a * b
func Product[INT types.SignedInteger](a, b INT) INT {
	return a * b
}

// Quotient returns the quotient of the two numbers
// as Quotient(a, b) = a / b as integer
func Quotient[INT types.SignedInteger](a, b INT) INT {
	return a / b
}

// Sum returns the sum of the two numbers
// as Sum(a, b) = a + b
func Sum[INT types.SignedInteger](a, b INT) INT {
	return a + b
}
