package rangetype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testCase struct {
		testName string
		start    int
		end      int
		expected Range
	}

	testCases := []testCase{
		{"Normal range", 1, 5, Range{Start: 1, End: 5}},
		{"Zero range", 0, 1, Range{Start: 0, End: 1}},
		{"Negative range", -5, -1, Range{Start: -5, End: -1}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := New(tc.start, tc.end)
			assert.Equal(t, tc.expected, result)
		})
	}
	// Test panic case
	t.Run("Test for panic", func(t *testing.T) {
		assert.Panics(t, func() {
			New(5, 1)
		})
	})
}

func TestNewAllInclusive(t *testing.T) {
	type testCase struct {
		testName string
		start    int
		end      int
		expected Range
	}

	testCases := []testCase{
		{"Normal range", 1, 5, Range{Start: 1, End: 6}},
		{"Zero range", 0, 0, Range{Start: 0, End: 1}},
		{"Negative range", -5, -1, Range{Start: -5, End: 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := NewAllInclusive(tc.start, tc.end)
			assert.Equal(t, tc.expected, result)
		})
	}
	// Test panic case
	t.Run("Test for panic", func(t *testing.T) {
		assert.Panics(t, func() {
			NewAllInclusive(5, 1)
		})
	})
}

func TestNewAllExclusive(t *testing.T) {
	type testCase struct {
		testName string
		start    int
		end      int
		expected Range
	}

	testCases := []testCase{
		{"Normal range", 1, 5, Range{Start: 2, End: 5}},
		{"Zero range", -1, 1, Range{Start: 0, End: 1}},
		{"Negative range", -5, -1, Range{Start: -4, End: -1}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := NewAllExclusive(tc.start, tc.end)
			assert.Equal(t, tc.expected, result)
		})
	}
	// Test panic case
	t.Run("Test for panic", func(t *testing.T) {
		assert.Panics(t, func() {
			NewAllExclusive(5, 1)
		})
	})
}

func TestContains(t *testing.T) {
	testCases := []struct {
		testName      string
		r             Range
		value         int
		expectedValue bool
	}{
		{"Value inside range", New(1, 5), 3, true},
		{"Value at start", New(1, 5), 1, true},
		{"Value at end", New(1, 5), 5, false},
		{"Value below range", New(1, 5), 0, false},
		{"Value above range", New(1, 5), 6, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.r.Contains(tc.value)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestLength(t *testing.T) {
	testCases := []struct {
		testName      string
		r             Range
		expectedValue int
	}{
		{"Normal range", New(1, 5), 4},
		{"Negative range", New(-5, -1), 4},
		{"Single element range", New(3, 4), 1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.r.Length()
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestOverlaps(t *testing.T) {
	testCases := []struct {
		testName      string
		r1            Range
		r2            Range
		expectedValue bool
	}{
		{"Overlapping ranges", New(1, 5), New(4, 8), true},
		{"Non-overlapping ranges", New(1, 3), New(4, 6), false},
		{"Touching ranges", New(1, 5), New(5, 10), false},
		{"One range inside another", New(1, 10), New(3, 7), true},
		{"Identical ranges", New(2, 6), New(2, 6), true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.r1.Overlaps(tc.r2)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		testName      string
		r             Range
		expectedValue string
	}{
		{"Normal range", New(1, 5), "[1, 5)"},
		{"Negative range", New(-5, -1), "[-5, -1)"},
		{"Single element range", New(3, 4), "[3, 4)"},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.r.String()
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}
