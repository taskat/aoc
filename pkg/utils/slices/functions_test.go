package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		testName      string
		slice         []int
		item          int
		expectedValue bool
	}{
		{"Nil slice", nil, 1, false},
		{"Empty slice", []int{}, 1, false},
		{"Item not in slice", []int{2, 3, 4}, 1, false},
		{"Item in slice", []int{2, 3, 4}, 3, true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Contains(tc.slice, tc.item)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestFilter(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		predicate     func(T) bool
		expectedValue []T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i int) bool { return true }, []int{}},
		{"Empty slice", []int{}, func(i int) bool { return true }, []int{}},
		{"Filter even numbers", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }, []int{2, 4}},
		{"Filter odd numbers", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 1 }, []int{1, 3}},
		{"Filter all numbers", []int{1, 2, 3, 4}, func(i int) bool { return true }, []int{1, 2, 3, 4}},
		{"Filter no numbers", []int{1, 2, 3, 4}, func(i int) bool { return false }, []int{}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Filter(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestFind(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		predicate     func(T) bool
		expectedValue T
		expectedFound bool
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i int) bool { return true }, 0, false},
		{"Empty slice", []int{}, func(i int) bool { return true }, 0, false},
		{"Find even number", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }, 2, true},
		{"Find odd number", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 1 }, 1, true},
		{"Find first number", []int{1, 2, 3, 4}, func(i int) bool { return true }, 1, true},
		{"Find no number", []int{1, 2, 3, 4}, func(i int) bool { return false }, 0, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, found := Find(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedValue, result)
			assert.Equal(t, tc.expectedFound, found)
		})
	}
}

func TestMap(t *testing.T) {
	type testCase[T, U any] struct {
		testName      string
		slice         []T
		f             func(T) U
		expectedValue []U
	}
	testCases := []testCase[int, string]{
		{"Nil slice", nil, func(i int) string { return "a" }, []string{}},
		{"Empty slice", []int{}, func(i int) string { return "a" }, []string{}},
		{"Map from int to string", []int{1, 2, 3}, func(i int) string { return fmt.Sprintf("%d", i) }, []string{"1", "2", "3"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Map(tc.slice, tc.f)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestMap_i(t *testing.T) {
	type testCase[T, U any] struct {
		testName      string
		slice         []T
		f             func(T, int) U
		expectedValue []U
	}
	testCases := []testCase[int, string]{
		{"Nil slice", nil, func(i int, j int) string { return "a" }, []string{}},
		{"Empty slice", []int{}, func(i int, j int) string { return "a" }, []string{}},
		{"Map from int to string", []int{1, 2, 3}, func(i int, j int) string { return fmt.Sprintf("%d", i+j) }, []string{"1", "3", "5"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Map_i(tc.slice, tc.f)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		testName      string
		slice         []int
		expectedValue int
	}{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
		{"Sum of positive numbers", []int{1, 2, 3}, 6},
		{"Sum of negative numbers", []int{-1, -2, -3}, -6},
		{"Sum of mixed numbers", []int{1, -2, 3}, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Sum(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}
