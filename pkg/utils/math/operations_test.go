package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected int
	}{
		{"Positive", 1, 1},
		{"Negative", -1, 1},
		{"Zero", 0, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Abs(tc.number)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		expected int
	}{
		{"Integer", 2.0, 2},
		{"Fractional", 2.3, 3},
		{"Negative Fractional", -2.3, -2},
		{"Negative Integer", -2.0, -2},
		{"Zero", 0.0, 0},
		{"Small positive", 0.000000000000001, 1},
		{"Large positive", 0.999999999999999, 1},
		{"Small negative", -0.000000000000001, 0},
		{"Large negative", -0.999999999999999, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Ceil(tc.number)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		expected int
	}{
		{"Integer", 2.0, 2},
		{"Fractional", 2.7, 2},
		{"Negative Fractional", -2.3, -3},
		{"Negative Integer", -2.0, -2},
		{"Zero", 0.0, 0},
		{"Small positive", 0.000000000000001, 0},
		{"Large positive", 0.999999999999999, 0},
		{"Small negative", -0.000000000000001, -1},
		{"Large negative", -0.999999999999999, -1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Floor(tc.number)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		expected int
	}{
		{"Positive Half", 2.5, 3},
		{"Positive Below Half", 2.4, 2},
		{"Positive Above Half", 2.6, 3},
		{"Negative Half", -2.5, -3},
		{"Negative Below Half", -2.6, -3},
		{"Negative Above Half", -2.4, -2},
		{"Zero", 0.0, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Round(tc.number)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
