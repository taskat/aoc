package intutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsN(t *testing.T) {
	testCases := []struct {
		testName string
		number   int
		n        int
		expected bool
	}{
		{"Number is equal to n", 5, 5, true},
		{"Number is not equal to n", 5, 3, false},
		{"Negative number equal to n", -5, -5, true},
		{"Negative number not equal to n", -5, -3, false},
		{"Zero equal to n", 0, 0, true},
		{"Zero not equal to n", 0, 1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			isN := IsN(tc.n)
			result := isN(tc.number)
			assert.Equal(t, tc.expected, result)
		})
	}
}
