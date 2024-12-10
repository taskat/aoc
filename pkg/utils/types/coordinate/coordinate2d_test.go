package coordinate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskat/aoc/pkg/utils/types"
)

func TestNewCoordinate2D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		x        T
		y        T
		expected Coordinate2D[T]
	}
	testCases := []testCase[int]{
		{"Test 1", 1, 2, Coordinate2D[int]{X: 1, Y: 2}},
		{"Test 2", 3, 4, Coordinate2D[int]{X: 3, Y: 4}},
		{"Test negative", -1, -2, Coordinate2D[int]{X: -1, Y: -2}},
		{"Test zero", 0, 0, Coordinate2D[int]{X: 0, Y: 0}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := NewCoordinate2D(tc.x, tc.y)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestAdd(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c        Coordinate2D[T]
		other    Coordinate2D[T]
		expected Coordinate2D[T]
	}
	testCases := []testCase[int]{
		{"Test 1", Coordinate2D[int]{X: 1, Y: 2}, Coordinate2D[int]{X: 3, Y: 4}, Coordinate2D[int]{X: 4, Y: 6}},
		{"Test 2", Coordinate2D[int]{X: 3, Y: 4}, Coordinate2D[int]{X: 1, Y: 2}, Coordinate2D[int]{X: 4, Y: 6}},
		{"Test negative", Coordinate2D[int]{X: -1, Y: -2}, Coordinate2D[int]{X: 1, Y: 2}, Coordinate2D[int]{X: 0, Y: 0}},
		{"Test zero", Coordinate2D[int]{X: 0, Y: 0}, Coordinate2D[int]{X: 0, Y: 0}, Coordinate2D[int]{X: 0, Y: 0}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c.Add(tc.other)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGo(t *testing.T) {
	type testCase[T types.Real] struct {
		testName  string
		c         Coordinate2D[T]
		direction Direction
		expected  Coordinate2D[T]
	}
	testCases := []testCase[int]{
		{"Up", Coordinate2D[int]{X: 1, Y: 2}, Up(), Coordinate2D[int]{X: 1, Y: 1}},
		{"Up Right", Coordinate2D[int]{X: 1, Y: 2}, UpRight(), Coordinate2D[int]{X: 2, Y: 1}},
		{"Right", Coordinate2D[int]{X: 1, Y: 2}, Right(), Coordinate2D[int]{X: 2, Y: 2}},
		{"Down Right", Coordinate2D[int]{X: 1, Y: 2}, DownRight(), Coordinate2D[int]{X: 2, Y: 3}},
		{"Down", Coordinate2D[int]{X: 1, Y: 2}, Down(), Coordinate2D[int]{X: 1, Y: 3}},
		{"Down Left", Coordinate2D[int]{X: 1, Y: 2}, DownLeft(), Coordinate2D[int]{X: 0, Y: 3}},
		{"Left", Coordinate2D[int]{X: 1, Y: 2}, Left(), Coordinate2D[int]{X: 0, Y: 2}},
		{"Up Left", Coordinate2D[int]{X: 1, Y: 2}, UpLeft(), Coordinate2D[int]{X: 0, Y: 1}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c.Go(tc.direction)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestIn2DSlice(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c        Coordinate2D[T]
		width    T
		height   T
		expected bool
	}
	testCases := []testCase[int]{
		{"Test 1", Coordinate2D[int]{X: 1, Y: 2}, 10, 10, true},
		{"Test 2", Coordinate2D[int]{X: 3, Y: 4}, 10, 10, true},
		{"Test negative", Coordinate2D[int]{X: -1, Y: -2}, 10, 10, false},
		{"Test zero", Coordinate2D[int]{X: 0, Y: 0}, 10, 10, true},
		{"Test out of bounds", Coordinate2D[int]{X: 10, Y: 10}, 10, 10, false},
		{"Test out of bounds 2", Coordinate2D[int]{X: 10, Y: 9}, 10, 10, false},
		{"Test out of bounds 3", Coordinate2D[int]{X: 9, Y: 10}, 10, 10, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c.In2DSlice(tc.width, tc.height)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestInLimits(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c        Coordinate2D[T]
		minX     T
		minY     T
		maxX     T
		maxY     T
		expected bool
	}
	testCases := []testCase[int]{
		{"Test 1", Coordinate2D[int]{X: 1, Y: 2}, 0, 0, 10, 10, true},
		{"Test 2", Coordinate2D[int]{X: 3, Y: 4}, 0, 0, 10, 10, true},
		{"Test negative", Coordinate2D[int]{X: -1, Y: -2}, 0, 0, 10, 10, false},
		{"Test negative 2", Coordinate2D[int]{X: -1, Y: -2}, -10, -10, 10, 10, true},
		{"Test lower edge", Coordinate2D[int]{X: 0, Y: 0}, 0, 0, 10, 10, true},
		{"Test upper edge", Coordinate2D[int]{X: 10, Y: 10}, 0, 0, 10, 10, true},
		{"Test out of bounds", Coordinate2D[int]{X: 10, Y: 10}, 0, 0, 9, 9, false},
		{"Test out of bounds 2", Coordinate2D[int]{X: 10, Y: 9}, 0, 0, 9, 9, false},
		{"Test out of bounds 3", Coordinate2D[int]{X: 9, Y: 10}, 0, 0, 9, 9, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c.InLimits(tc.minX, tc.minY, tc.maxX, tc.maxY)
			assert.Equal(t, tc.expected, result)
		})
	}
}
