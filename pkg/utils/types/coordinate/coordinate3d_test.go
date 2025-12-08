package coordinate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskat/aoc/pkg/utils/types"
)

func TestNewCoordinate3D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		x, y, z  T
		expected Coordinate3D[T]
	}
	testCases := []testCase[int]{
		{"All positive", 1, 2, 3, Coordinate3D[int]{1, 2, 3}},
		{"Negative values", -1, -2, -3, Coordinate3D[int]{-1, -2, -3}},
		{"Mixed values", -1, 2, -3, Coordinate3D[int]{-1, 2, -3}},
		{"Zero values", 0, 0, 0, Coordinate3D[int]{0, 0, 0}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := NewCoordinate3D(tc.x, tc.y, tc.z)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestAdd3D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c1, c2   Coordinate3D[T]
		expected Coordinate3D[T]
	}
	testCases := []testCase[int]{
		{"All positive", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{4, 5, 6}, Coordinate3D[int]{5, 7, 9}},
		{"Negative values", Coordinate3D[int]{-1, -2, -3}, Coordinate3D[int]{-4, -5, -6}, Coordinate3D[int]{-5, -7, -9}},
		{"Mixed values", Coordinate3D[int]{-1, 2, -3}, Coordinate3D[int]{4, -5, 6}, Coordinate3D[int]{3, -3, 3}},
		{"Zero values", Coordinate3D[int]{0, 0, 0}, Coordinate3D[int]{0, 0, 0}, Coordinate3D[int]{0, 0, 0}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c1.Add(tc.c2)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestDistance(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c1, c2   Coordinate3D[T]
		expected float64
	}
	testCases := []testCase[int]{
		{"Same point", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{1, 2, 3}, 0},
		{"Unit distance", Coordinate3D[int]{0, 0, 0}, Coordinate3D[int]{1, 0, 0}, 1},
		{"Pythagorean triple", Coordinate3D[int]{0, 0, 0}, Coordinate3D[int]{3, 4, 0}, 5},
		{"3D distance", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{4, 6, 8}, 7.0710678118654755},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c1.Distance(tc.c2)
			assert.InDelta(t, tc.expected, result, 1e-9)
		})
	}
}

func TestEqual3D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c1, c2   Coordinate3D[T]
		expected bool
	}
	testCases := []testCase[int]{
		{"Equal coordinates", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{1, 2, 3}, true},
		{"Different X", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{4, 2, 3}, false},
		{"Different Y", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{1, 5, 3}, false},
		{"Different Z", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{1, 2, 6}, false},
		{"All different", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{4, 5, 6}, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c1.Equal(tc.c2)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestIn3DSlice(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c        Coordinate3D[T]
		width    T
		height   T
		depth    T
		expected bool
	}
	testCases := []testCase[int]{
		{"Inside slice", Coordinate3D[int]{1, 1, 1}, 3, 3, 3, true},
		{"On boundary", Coordinate3D[int]{0, 0, 0}, 3, 3, 3, true},
		{"On opposite boundary", Coordinate3D[int]{2, 2, 2}, 3, 3, 3, true},
		{"Outside negative", Coordinate3D[int]{-1, 1, 1}, 3, 3, 3, false},
		{"Outside positive", Coordinate3D[int]{3, 1, 1}, 3, 3, 3, false},
		{"Outside height", Coordinate3D[int]{1, 3, 1}, 3, 3, 3, false},
		{"Outside depth", Coordinate3D[int]{1, 1, 3}, 3, 3, 3, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c.In3DSlice(tc.width, tc.height, tc.depth)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestInLimits3D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName         string
		c                Coordinate3D[T]
		minX, minY, minZ T
		maxX, maxY, maxZ T
		expected         bool
	}
	testCases := []testCase[int]{
		{"Inside limits", Coordinate3D[int]{2, 2, 2}, 1, 1, 1, 3, 3, 3, true},
		{"On min boundary", Coordinate3D[int]{1, 1, 1}, 1, 1, 1, 3, 3, 3, true},
		{"On max boundary", Coordinate3D[int]{3, 3, 3}, 1, 1, 1, 3, 3, 3, true},
		{"Outside min X", Coordinate3D[int]{0, 2, 2}, 1, 1, 1, 3, 3, 3, false},
		{"Outside max X", Coordinate3D[int]{4, 2, 2}, 1, 1, 1, 3, 3, 3, false},
		{"Outside min Y", Coordinate3D[int]{2, 0, 2}, 1, 1, 1, 3, 3, 3, false},
		{"Outside max Y", Coordinate3D[int]{2, 4, 2}, 1, 1, 1, 3, 3, 3, false},
		{"Outside min Z", Coordinate3D[int]{2, 2, 0}, 1, 1, 1, 3, 3, 3, false},
		{"Outside max Z", Coordinate3D[int]{2, 2, 4}, 1, 1, 1, 3, 3, 3, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c.InLimits(tc.minX, tc.minY, tc.minZ, tc.maxX, tc.maxY, tc.maxZ)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestManhattanDistance3D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c1, c2   Coordinate3D[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Same point", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{1, 2, 3}, 0},
		{"Unit distance", Coordinate3D[int]{0, 0, 0}, Coordinate3D[int]{1, 0, 0}, 1},
		{"Mixed coordinates", Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{4, 6, 8}, 12},
		{"Negative coordinates", Coordinate3D[int]{-1, -2, -3}, Coordinate3D[int]{1, 2, 3}, 12},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c1.ManhattanDistance(tc.c2)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestCoordinateString3D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c        Coordinate3D[T]
		expected string
	}
	testCases := []testCase[int]{
		{"Positive coordinates", Coordinate3D[int]{1, 2, 3}, "(1, 2, 3)"},
		{"Negative coordinates", Coordinate3D[int]{-1, -2, -3}, "(-1, -2, -3)"},
		{"Mixed coordinates", Coordinate3D[int]{-1, 2, -3}, "(-1, 2, -3)"},
		{"Zero coordinates", Coordinate3D[int]{0, 0, 0}, "(0, 0, 0)"},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c.String()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSub3D(t *testing.T) {
	type testCase[T types.Real] struct {
		testName string
		c1, c2   Coordinate3D[T]
		expected Coordinate3D[T]
	}
	testCases := []testCase[int]{
		{"All positive", Coordinate3D[int]{4, 5, 6}, Coordinate3D[int]{1, 2, 3}, Coordinate3D[int]{3, 3, 3}},
		{"Negative values", Coordinate3D[int]{-1, -2, -3}, Coordinate3D[int]{-4, -5, -6}, Coordinate3D[int]{3, 3, 3}},
		{"Mixed values", Coordinate3D[int]{4, -5, 6}, Coordinate3D[int]{-1, 2, -3}, Coordinate3D[int]{5, -7, 9}},
		{"Zero values", Coordinate3D[int]{0, 0, 0}, Coordinate3D[int]{0, 0, 0}, Coordinate3D[int]{0, 0, 0}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.c1.Sub(tc.c2)
			assert.Equal(t, tc.expected, result)
		})
	}
}
