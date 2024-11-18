package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	testCases := []struct {
		testName      string
		input         string
		expectedValue int
	}{
		{"Valid input", "1", 1},
		{"Negative input", "-1", -1},
		{"Big input", "123456", 123456},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Atoi(tc.input)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	t.Run("Invalid input", func(t *testing.T) {
		assert.Panics(t, func() {
			Atoi("a")
		})
	})
}

func TestRuneToInt(t *testing.T) {
	testCases := []struct {
		testName      string
		input         rune
		expectedValue int
	}{
		{"Digit 0", '0', 0},
		{"Digit 1", '1', 1},
		{"Digit 9", '9', 9},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := RuneToInt(tc.input)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	t.Run("Invalid input", func(t *testing.T) {
		assert.Panics(t, func() {
			RuneToInt('a')
		})
	})
}
