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

func TestItoa(t *testing.T) {
	testCases := []struct {
		testName      string
		input         int
		expectedValue string
	}{
		{"Zero", 0, "0"},
		{"Positive", 1, "1"},
		{"Negative", -1, "-1"},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Itoa(tc.input)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
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

func TestStringToBool(t *testing.T) {
	testCases := []struct {
		testName      string
		input         string
		expectedValue bool
	}{
		{"True", "true", true},
		{"False", "false", false},
		{"0", "0", false},
		{"1", "1", true},
		{"t", "t", true},
		{"f", "f", false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := StringToBool(tc.input)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	t.Run("Invalid input", func(t *testing.T) {
		assert.Panics(t, func() {
			StringToBool("a")
		})
	})
}
