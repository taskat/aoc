package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	testCases := []struct {
		testName      string
		input         rune
		expectedValue bool
	}{
		{"Digit 0", '0', true},
		{"Digit 1", '1', true},
		{"Digit 9", '9', true},
		{"Not a digit", 'a', false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := IsDigit(tc.input)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

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
