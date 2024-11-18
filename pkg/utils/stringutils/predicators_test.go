package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInteger(t *testing.T) {
	testCases := []struct {
		testName      string
		input         string
		expectedValue bool
	}{
		{"Empty input", "", false},
		{"Valid input", "1", true},
		{"Negative input", "-1", true},
		{"Big input", "123456", true},
		{"Invalid input", "a", false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := IsInteger(tc.input)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}
