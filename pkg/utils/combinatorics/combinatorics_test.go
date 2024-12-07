package combinatorics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartesianProduct(t *testing.T) {
	type testCase[T any] struct {
		testName       string
		element        []T
		length         int
		expectedResult [][]T
	}
	testCases := []testCase[int]{
		{"Nil element", nil, 0, [][]int{{}}},
		{"Empty element", []int{}, 0, [][]int{{}}},
		{"One element", []int{1}, 1, [][]int{{1}}},
		{"One element 3 length", []int{1}, 3, [][]int{{1, 1, 1}}},
		{"Two elements", []int{1, 2}, 1, [][]int{{1}, {2}}},
		{"Two elements 2 length", []int{1, 2}, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{"Two elements 3 length", []int{1, 2}, 3, [][]int{{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2}, {2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2}}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := CartesianProduct(tc.element, tc.length)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
