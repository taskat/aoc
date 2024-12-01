package intutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbc(t *testing.T) {
	testCases := []struct {
		testName      string
		number        int
		expectedValue int
	}{
		{"Positive number", 5, 5},
		{"Negative number", -5, 5},
		{"Zero", 0, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Abs(tc.number)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}
