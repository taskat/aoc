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

func TestEquals(t *testing.T) {
	testCases := []struct {
		testName      string
		a             int
		b             int
		expectedValue bool
	}{
		{"Positive numbers", 5, 3, false},
		{"Negative numbers", -5, -3, false},
		{"Positive and negative numbers", 5, -3, false},
		{"Negative and positive numbers", -5, 3, false},
		{"Zero and positive numbers", 0, 3, false},
		{"Positive and zero numbers", 5, 0, false},
		{"Zero and negative numbers", 0, -3, false},
		{"Negative and zero numbers", -5, 0, false},
		{"Equal numbers", 5, 5, true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Equals(tc.a, tc.b)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestDiff(t *testing.T) {
	testCases := []struct {
		testName      string
		a             int
		b             int
		expectedValue int
	}{
		{"Positive numbers", 5, 3, 2},
		{"Negative numbers", -5, -3, -2},
		{"Positive and negative numbers", 5, -3, 8},
		{"Negative and positive numbers", -5, 3, -8},
		{"Zero and positive numbers", 0, 3, -3},
		{"Positive and zero numbers", 5, 0, 5},
		{"Zero and negative numbers", 0, -3, 3},
		{"Negative and zero numbers", -5, 0, -5},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Diff(tc.a, tc.b)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestLength(t *testing.T) {
	testCases := []struct {
		testName      string
		number        int
		expectedValue int
	}{
		{"Positive number", 5, 1},
		{"Negative number", -5, 1},
		{"Zero", 0, 1},
		{"Positive number with multiple digits", 123, 3},
		{"Negative number with multiple digits", -123, 3},
		{"Long positive number", 1234567890, 10},
		{"Long negative number", -1234567890, 10},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Length(tc.number)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestPower(t *testing.T) {
	testCases := []struct {
		testName      string
		a             int
		b             int
		expectedValue int
	}{
		{"Positive numbers", 5, 3, 125},
		{"Negative numbers", -5, 3, -125},
		{"Zero and positive numbers", 0, 3, 0},
		{"Positive and zero numbers", 5, 0, 1},
		{"Negative and zero numbers", -5, 0, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Power(tc.a, tc.b)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestPowerPanic(t *testing.T) {
	testCases := []struct {
		testName string
		a        int
		b        int
	}{
		{"Negative exponent", 5, -3},
		{"Negative base and negative exponent", -5, -3},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() {
				Power(tc.a, tc.b)
			})
		})
	}
}

func TestProduct(t *testing.T) {
	testCases := []struct {
		testName      string
		a             int
		b             int
		expectedValue int
	}{
		{"Positive numbers", 5, 3, 15},
		{"Negative numbers", -5, -3, 15},
		{"Positive and negative numbers", 5, -3, -15},
		{"Negative and positive numbers", -5, 3, -15},
		{"Zero and positive numbers", 0, 3, 0},
		{"Positive and zero numbers", 5, 0, 0},
		{"Zero and negative numbers", 0, -3, 0},
		{"Negative and zero numbers", -5, 0, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Product(tc.a, tc.b)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestQuotient(t *testing.T) {
	testCases := []struct {
		testName      string
		a             int
		b             int
		expectedValue int
	}{
		{"Positive numbers", 5, 3, 1},
		{"Negative numbers", -5, -3, 1},
		{"Positive and negative numbers", 5, -3, -1},
		{"Negative and positive numbers", -5, 3, -1},
		{"Zero and positive numbers", 0, 3, 0},
		{"Zero and negative numbers", 0, -3, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Quotient(tc.a, tc.b)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		testName      string
		a             int
		b             int
		expectedValue int
	}{
		{"Positive numbers", 5, 3, 8},
		{"Negative numbers", -5, -3, -8},
		{"Positive and negative numbers", 5, -3, 2},
		{"Negative and positive numbers", -5, 3, -2},
		{"Zero and positive numbers", 0, 3, 3},
		{"Positive and zero numbers", 5, 0, 5},
		{"Zero and negative numbers", 0, -3, -3},
		{"Negative and zero numbers", -5, 0, -5},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Sum(tc.a, tc.b)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}
