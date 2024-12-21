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
