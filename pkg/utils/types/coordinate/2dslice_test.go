package coordinate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy2DSlice(t *testing.T) {
	original := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	copied := Copy2DSlice(original)
	assert.Equal(t, original, copied)
	// Modify the copied slice and check that the original remains unchanged
	copied[0][0] = 99
	assert.NotEqual(t, original, copied)
	assert.Equal(t, 1, original[0][0])
	// Modify the original slice and check that the copied remains unchanged
	original[1][1] = 88
	assert.NotEqual(t, original, copied)
	assert.Equal(t, 5, copied[1][1])
}

func TestCountIn2DSlice(t *testing.T) {
	testCases := []struct {
		testName  string
		slice     [][]int
		predicate func(int) bool
		expected  int
	}{
		{
			testName: "Count even numbers",
			slice: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			predicate: func(n int) bool { return n%2 == 0 },
			expected:  4,
		},
		{
			testName: "Count numbers greater than 5",
			slice: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			predicate: func(n int) bool { return n > 5 },
			expected:  4,
		},
		{
			testName: "Count all numbers",
			slice: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			predicate: func(n int) bool { return true },
			expected:  9,
		},
		{
			testName: "Count no numbers",
			slice: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			predicate: func(n int) bool { return false },
			expected:  0,
		},
		{
			testName:  "Empty slice",
			slice:     [][]int{},
			predicate: func(n int) bool { return n == 1 },
			expected:  0,
		},
		{
			testName:  "nil slice",
			slice:     nil,
			predicate: func(n int) bool { return true },
			expected:  0,
		},
		{
			testName: "Slice with empty rows",
			slice: [][]int{
				{},
				{1, 2},
				{},
				{3},
			},
			predicate: func(n int) bool { return n > 1 },
			expected:  2,
		},
		{
			testName: "Slice with nil rows",
			slice: [][]int{
				nil,
				{1, 2},
				nil,
				{3},
			},
			predicate: func(n int) bool { return n > 1 },
			expected:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := CountIn2DSlice(tc.slice, tc.predicate)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFind2DSliceIndex(t *testing.T) {
	testCases := []struct {
		testName      string
		slice         [][]int
		predicate     func(int) bool
		expectedCoord Coordinate2D[int]
	}{
		{
			testName: "Element found",
			slice: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			predicate:     func(n int) bool { return n == 5 },
			expectedCoord: Coordinate2D[int]{X: 1, Y: -1},
		},
		{
			testName: "Element not found",
			slice: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			predicate:     func(n int) bool { return n == 10 },
			expectedCoord: Coordinate2D[int]{X: -1, Y: -1},
		},
		{
			testName:      "Empty slice",
			slice:         [][]int{},
			predicate:     func(n int) bool { return n == 1 },
			expectedCoord: Coordinate2D[int]{X: -1, Y: -1},
		},
		{
			testName: "Multiple matches, return first",
			slice: [][]int{
				{1, 2, 3},
				{4, 2, 6},
				{7, 8, 2},
			},
			predicate:     func(n int) bool { return n == 2 },
			expectedCoord: Coordinate2D[int]{X: 1, Y: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Find2DSliceIndex(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedCoord, result)
		})
	}
}
