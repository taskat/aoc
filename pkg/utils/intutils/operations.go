package intutils

import (
	"math"

	"github.com/taskat/aoc/pkg/utils/types"
)

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

// Length returns the number of digits in the number
func Length[INT types.SignedInteger](number INT) INT {
	if number == 0 {
		return 1
	}
	return INT(math.Log10(float64(Abs(number)))) + 1
}

// Power returns the power of the number
// as Power(a, b) = a^b. If b is negatice, it panics
func Power[INT types.Integer](a, b INT) INT {
	if b < 0 {
		panic("negative exponent is not supported")
	}
	return INT(math.Pow(float64(a), float64(b)))
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
